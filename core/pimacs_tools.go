package core

import "fmt"

const (
	debugReprMaxDepth            = 3
	debugReprLispStackMaxLineLen = 80
	debugReprLongStringLen       = 25
)

func debugRepr(objs ...lispObject) string {
	switch len(objs) {
	case 0:
		return "[]"
	case 1:
		return debugReprInternal(objs[0], debugReprMaxDepth)
	default:
		s := "[]lispObject{"
		for i, obj := range objs {
			s += debugReprInternal(obj, debugReprMaxDepth)
			if i < len(objs)-1 {
				s += ", "
			}
		}
		s += "}"
		return s
	}
}

func debugReprInternal(obj lispObject, depth int) string {
	if depth <= 0 {
		return "?"
	}
	depth--

	if obj == nil {
		return "null"
	}

	switch obj.getType() {
	case lispTypeCons:
		list := true
		elems := 0
		tail := obj
		s := "("

		for ; consp(tail); tail = xCdr(tail) {
			if tail != obj {
				s += " "
			}
			elems++
			s += debugReprInternal(xCar(tail), depth)
		}

		if symbolp(tail) && xSymbolName(tail) == "nil" {
			s += ")"
		} else {
			list = false
			s += fmt.Sprintf(" . %v)", debugReprInternal(tail, depth))
		}

		if list && elems == 2 {
			first := xCar(obj)
			val := xCar(xCdr(obj))
			if symbolp(first) && xSymbolName(first) == "quote" {
				return fmt.Sprintf("'%v", debugReprInternal(val, depth))
			}
		}

		return s
	case lispTypeFloat:
		return fmt.Sprintf("%v", xFloatValue(obj))
	case lispTypeString:
		val := xStringValue(obj)
		if len(val) > debugReprLongStringLen {
			val = val[:debugReprLongStringLen] + "..."
		}
		return fmt.Sprintf("\"%v\"", val)
	case lispTypeInteger:
		return fmt.Sprintf("%v", xIntegerValue(obj))
	case lispTypeSymbol:
		sym := xSymbol(obj)
		if sym.val == sym || sym.function == sym {
			return sym.name
		}

		s := fmt.Sprintf("%v", xSymbolName(sym))

		val := "?"
		switch sym.redirect {
		case symbolRedirectPlain:
			val = debugReprInternal(sym.val, depth)
		case symbolRedirectFwd:
			val = "FWD"
		}
		if val == "unbound" {
			val = ""
		}

		fn := debugReprInternal(sym.function, depth)
		if fn == "nil" {
			fn = ""
		}

		if val != "" || fn != "" {
			s += "{"

			if val != "" {
				s += fmt.Sprintf("=%v", val)
			}

			if fn != "" {
				if val != "" {
					s += ","
				}
				s += fmt.Sprintf("f=%v", fn)
			}

			s += "}"
		}

		return s
	case lispTypeVector:
		s := "["
		val := xVector(obj).val
		for i, elem := range val {
			s += debugReprInternal(elem, depth)
			if i < len(val)-1 {
				s += ", "
			}
		}
		return s + "]"
	case lispTypeBuffer:
		buf := xBuffer(obj)
		return fmt.Sprintf("buf(name=%v, live=%v)", buf.name, buf.live)
	case lispTypeSubroutine:
		subr := xSubroutine(obj)
		return fmt.Sprintf("subr(%v)", subr.name)
	case lispTypeCharTable:
		return fmt.Sprintf("chartab(subtype=%v)", xCharTable(obj).subtype)
	case lispTypeChannel:
		return fmt.Sprintf("channel(%v)", xChannel(obj).val)
	case lispTypeHashTable:
		table := xHashTable(obj)
		s := "hashtable{"
		i := 0
		for k, v := range table.val {
			s += fmt.Sprintf("%v: %v", k, v)
			if i < len(table.val)-1 {
				s += ", "
			}
			i++
		}

		return s + "}"
	default:
		return "<unknown object>"
	}
}

func debugReprLispStack(stack []stackEntry) string {
	lispStack := ""
	for i := len(stack) - 1; i >= 0; i-- {

		switch elem := stack[i].(type) {
		case *stackEntryBacktrace:
			functionName := "<unknown function>"
			if symbolp(elem.function) {
				functionName = xSymbolName(elem.function)
			}

			lispStack += fmt.Sprintf("  - bt: %v(", functionName)

			for j, arg := range elem.args {
				printed := debugRepr(arg)
				if len(printed) > debugReprLispStackMaxLineLen {
					printed = printed[:debugReprLispStackMaxLineLen] + "[...]"
				}
				lispStack += printed

				if j < len(elem.args)-1 {
					lispStack += " "
				}
			}

			lispStack += ")"
		case *stackEntryLet:
			lispStack += fmt.Sprintf("  - let: %v = %v", debugRepr(elem.symbol), debugRepr(elem.oldVal))
		case *stackEntryLetForwarded:
			lispStack += fmt.Sprintf("  - letfwd: %v = %v", debugRepr(elem.symbol), debugRepr(elem.oldVal))
		default:
			lispStack += "  - other"
		}

		if i > 0 {
			lispStack += "\n"
		}

	}
	return lispStack
}

func (ec *execContext) pimacsSymbolDebug(symbol lispObject) (lispObject, error) {
	if !symbolp(symbol) {
		return ec.wrongTypeArgument(ec.s.symbolp, symbol)
	}

	sym := xSymbol(symbol)
	val := xEnsure(ec.findSymbolValue(symbol))
	fwd := ec.nil_
	if sym.fwd != nil {
		fwd = sym.fwd.value(ec)
	}

	return ec.makeKwPlist(map[string]lispObject{
		"value":    val,
		"function": sym.function,
		"name":     newString(sym.name, true),
		"special":  xEnsure(ec.bool(sym.special)),
		"plist":    sym.plist,
		"redirect": newInteger(lispInt(sym.redirect)),
		"fwd":      fwd,
	})
}

func (ec *execContext) pimacsDebugRepr(objs ...lispObject) (lispObject, error) {
	return newString(debugRepr(objs...), true), nil
}

func (ec *execContext) symbolsOfPimacsTools() {
	ec.defVarLisp(
		&ec.v.pimacsRepo,
		"pimacs--repo",
		newString("https://github.com/federicotdn/pimacs", false),
	)

	ec.defSubr1(nil, "pimacs--symbol-debug", (*execContext).pimacsSymbolDebug, 1)
	ec.defSubrM(nil, "pimacs--debug-repr", (*execContext).pimacsDebugRepr, 0)
	ec.defSubrM(nil, "dr", (*execContext).pimacsDebugRepr, 0)
}
