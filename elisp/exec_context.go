package elisp

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	eofRune rune = -1
	nbsp         = '\u00A0'
)

type stackEntryTag int

const (
	entryLet stackEntryTag = iota + 1
	entryCatch
)

type stackJumpTag struct {
	tag   lispObject
	value lispObject
}

type stackJumpSignal struct {
	errorSymbol lispObject
	data        lispObject
}

type stackEntry interface {
	tag() stackEntryTag
}

type execContext struct {
	stack   []stackEntry
	nil_    lispObject
	t       lispObject
	g       globals
	obarray map[string]*lispSymbol
}

type stackEntryLet struct {
	symbol *lispSymbol
	oldVal lispObject
}

type stackEntryCatch struct {
	catchTag lispObject
}

type readContext struct {
	source string
	runes  []rune
	i      int
}

func (jmp *stackJumpTag) Error() string {
	format := "stack jump: tag: '%+v'"
	if symbolp(jmp.tag) {
		return fmt.Sprintf(format, xSymbol(jmp.tag).name)
	}

	return fmt.Sprintf(format, jmp.tag)
}

func (jmp *stackJumpSignal) Error() string {
	format := "stack jump: signal: '%+v'"
	if symbolp(jmp.errorSymbol) {
		return fmt.Sprintf(format, xSymbol(jmp.errorSymbol).name)
	}

	return fmt.Sprintf(format, jmp.errorSymbol)
}

func (sel *stackEntryLet) tag() stackEntryTag {
	return entryLet
}

func (sec *stackEntryCatch) tag() stackEntryTag {
	return entryCatch
}

func (ctx *readContext) read() rune {
	if ctx.i == len(ctx.runes) {
		return eofRune
	}

	r := ctx.runes[ctx.i]
	ctx.i++
	return r
}

func (ctx *readContext) unread(c rune) {
	if c == eofRune {
		return
	}

	if ctx.i == 0 {
		terminate("unbalanced read/unread call with rune: '%v'", c)
	}

	ctx.i--
}

func (ec *execContext) stringToNumber(s string) (lispObject, error) {
	nInt, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return ec.makeInteger(lispInt(nInt)), nil
	}

	nFloat, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return ec.makeFloat(lispFp(nFloat)), nil
	}

	return ec.pimacsUnimplemented(ec.g.read, "unknown number format: '%v'", s)
}

func (ec *execContext) makeSymbolBase(name string) *lispSymbol {
	return &lispSymbol{
		name: name,
	}
}

func (ec *execContext) makeSymbol(name string) *lispSymbol {
	base := ec.makeSymbolBase(name)
	base.value = ec.g.unbound
	base.function = ec.nil_
	base.plist = ec.nil_
	return base
}

func (ec *execContext) makeInteger(value lispInt) *lispInteger {
	return &lispInteger{
		value: value,
	}
}

func (ec *execContext) makeCons(car, cdr lispObject) *lispCons {
	return &lispCons{
		car: car,
		cdr: cdr,
	}
}

func (ec *execContext) makeFloat(value lispFp) *lispFloat {
	return &lispFloat{
		value: value,
	}
}

func (ec *execContext) makeString(value string) *lispString {
	return &lispString{
		value: value,
	}
}

func (ec *execContext) makeList(objs ...lispObject) lispObject {
	if len(objs) == 0 {
		return ec.nil_
	}

	tmp := ec.makeCons(objs[0], ec.nil_)
	val := tmp

	for _, obj := range objs[1:] {
		tmp.cdr = ec.makeCons(obj, ec.nil_)
		tmp = xCons(tmp.cdr)
	}

	return val
}

func (ec *execContext) makeVectorLike(vecType vectorLikeType, value vectorLikeValue) *lispVectorLike {
	return &lispVectorLike{
		vecType: vecType,
		value:   value,
	}
}

func (ec *execContext) readEscape(ctx *readContext, stringp bool) (rune, error) {
	// TAGS: incomplete
	c := ctx.read()
	if c == eofRune {
		return 0, xErrOnly(ec.signalN(ec.g.endOfFile))
	}

	switch c {
	case 'a':
		return '\a', nil
	case 'b':
		return '\b', nil
	case 'd':
		return 0177, nil
	case 'e':
		return 033, nil
	case 'f':
		return '\f', nil
	case 'n':
		return '\n', nil
	case 'r':
		return '\r', nil
	case 't':
		return '\t', nil
	case 'v':
		return '\v', nil
	case '\n':
		return -1, nil
	case ' ':
		if stringp {
			return -1, nil
		}
		return ' ', nil
	case 'M':
	case 'S':
	case 'H':
	case 'A':
	case 's':
	case 'C':
	case '^':
	case '0', '1', '2', '3', '4', '5', '6', '7':
		num := c - '0'

		for count := 0; count < 2; count++ {
			c = ctx.read()

			if c >= '0' && c <= '7' {
				num *= 8
				num += c - '0'
			} else {
				ctx.unread(c)
				break
			}
		}

		if num >= 0x80 && num < 0x100 {
			num = byte8toChar(num)
		}

		return num, nil
	case 'x':
	case 'U':
	case 'u':
	case 'N':
	default:
		return c, nil
	}

	return 0, xErrOnly(ec.pimacsUnimplemented(ec.g.read, "unknown escape code: '\\%v'", c))
}

func (ec *execContext) defSubr(name string, sub *subroutine) {
	if sub.maxArgs >= 0 && sub.minArgs > sub.maxArgs {
		ec.terminate("min args must be smaller or equal to max args (subroutine: '%+v')", sub)
	}
	sym := ec.intern(name)
	vec := ec.makeVectorLike(vectorLikeTypeSubroutine, sub)
	sym.function = vec
}

func (ec *execContext) defSubr0(name string, fn lispFn0) *subroutine {
	sub := &subroutine{
		callabe0: fn,
	}
	ec.defSubr(name, sub)
	return sub
}

func (ec *execContext) defSubr1(name string, fn lispFn1, minArgs int) *subroutine {
	sub := &subroutine{
		callabe1: fn,
		minArgs:  minArgs,
		maxArgs:  1,
	}
	ec.defSubr(name, sub)
	return sub
}

func (ec *execContext) defSubr2(name string, fn lispFn2, minArgs int) *subroutine {
	sub := &subroutine{
		callabe2: fn,
		minArgs:  minArgs,
		maxArgs:  2,
	}
	ec.defSubr(name, sub)
	return sub
}

func (ec *execContext) defSubr3(name string, fn lispFn3, minArgs int) *subroutine {
	sub := &subroutine{
		callabe3: fn,
		minArgs:  minArgs,
		maxArgs:  3,
	}
	ec.defSubr(name, sub)
	return sub
}

func (ec *execContext) defSubrM(name string, fn lispFnM, minArgs int) *subroutine {
	sub := &subroutine{
		callabem: fn,
		minArgs:  minArgs,
		maxArgs:  argsMany,
	}
	ec.defSubr(name, sub)
	return sub
}

func (ec *execContext) defSubrU(name string, fn lispFn1, minArgs int) *subroutine {
	sub := &subroutine{
		callabe1: fn,
		minArgs:  minArgs,
		maxArgs:  argsUnevalled,
	}
	ec.defSubr(name, sub)
	return sub
}

func newExecContext() *execContext {
	ec := execContext{}

	ec.obarray = make(map[string]*lispSymbol)
	ec.stack = []stackEntry{}

	ec.initialDefsSymbols()   // symbols.go
	ec.initialDefsErrors()    // errors.gmo
	ec.initialDefsVariables() // variables.go
	ec.initialDefsFunctions() // functions.go

	return &ec
}

func (ec *execContext) internSymbol(symbol *lispSymbol) {
	prev, ok := ec.obarray[symbol.name]
	if ok && prev != symbol {
		ec.terminate("different symbol with that name already interned, name: '%v'", symbol.name)
	}
	ec.obarray[symbol.name] = symbol
}

func (ec *execContext) intern(name string) *lispSymbol {
	sym, ok := ec.obarray[name]
	if !ok {
		sym = ec.makeSymbol(name)
		ec.internSymbol(sym)
	}

	return sym
}

func (ec *execContext) readList(ctx *readContext) (lispObject, error) {
	// TAGS: incomplete
	var val lispObject = ec.nil_
	var tail lispObject = ec.nil_

	for {
		elt, c, err := ec.read1(ctx)
		if err != nil {
			return nil, err
		}

		if c != 0 {
			switch c {
			case ')':
				return val, nil
			case '.':
				if tail != ec.nil_ {
					obj, err := ec.read0(ctx)
					if err != nil {
						return nil, err
					}
					xSetCdr(tail, obj)
				} else {
					val, err = ec.read0(ctx)
				}

				_, c, _ = ec.read1(ctx)
				if c == ')' {
					return val, nil
				}

				return ec.invalidReadSyntax("'.' in wrong context")
			default:
				return ec.invalidReadSyntax("invalid list ending: '%v'", string(c))
			}
		}

		tmp := ec.makeCons(elt, ec.nil_)

		if tail != ec.nil_ {
			xSetCdr(tail, tmp)
		} else {
			val = tmp
		}

		tail = tmp
	}
}

func (ec *execContext) readAtom(c rune, ctx *readContext) (lispObject, error) {
	// TAGS: incomplete
	quoted := false
	builder := strings.Builder{}

	for {
		if c == '\\' {
			c = ctx.read()
			if c == eofRune {
				return ec.signalN(ec.g.endOfFile)
			}

			quoted = true
		}

		builder.WriteRune(c)
		c = ctx.read()

		special := strings.Contains("\"';()[]#`,", string(c))
		if c <= 040 || c == nbsp || special {
			break
		}
	}

	ctx.unread(c)
	s := builder.String()

	if !quoted {
		num, err := ec.stringToNumber(s)
		if err == nil {
			return num, nil
		}
	}

	return ec.intern(s), nil
}

func (ec *execContext) read1Result(obj lispObject, err error) (lispObject, rune, error) {
	return obj, 0, err
}

func (ec *execContext) read1(ctx *readContext) (lispObject, rune, error) {
	// TAGS: incomplete
	var err error
	var c rune

	for {
		c = ctx.read()
		if c == eofRune {
			return ec.read1Result(ec.signalN(ec.g.endOfFile))
		}

		switch {
		case c == '(':
			return ec.read1Result(ec.readList(ctx))
		case c == '[':
			return ec.read1Result(ec.pimacsUnimplemented(ec.g.read, "unknown token: '%v'", c))
		case c == ')' || c == ']':
			return nil, c, nil
		case c == '#':
			c = ctx.read()

			switch {
			case c == '#':
				return ec.intern(""), 0, nil
			case c == '\'':
				body, err := ec.read0(ctx)
				if err != nil {
					return nil, 0, err
				}
				return ec.makeList(ec.g.function, body), 0, nil
			default:
				ctx.unread(c)
				return ec.read1Result(ec.signalN(ec.g.invalidReadSyntax, ec.makeString(string(c))))
			}
		case c == ';':
			for {
				c = ctx.read()
				if c == eofRune || c == '\n' {
					break
				}
			}
		case c == '\'' || c == '`':
			obj, err := ec.read0(ctx)
			if err != nil {
				return nil, 0, err
			}

			sym := ec.g.quote
			if c == '`' {
				sym = ec.g.backquote
			}

			list := ec.makeList(sym, obj)
			return list, 0, nil
		case c == ',':
			return ec.read1Result(ec.pimacsUnimplemented(ec.g.read, "unknown token: '%v'", c))
		case c == '?':
			c = ctx.read()
			if c == eofRune {
				return ec.read1Result(ec.signalN(ec.g.endOfFile))
			}

			if c == ' ' || c == '\t' {
				return ec.makeInteger(lispInt(c)), 0, nil
			}

			if c == '\\' {
				c, err = ec.readEscape(ctx, false)
				if err != nil {
					return nil, 0, err
				}
			}

			if charByte8(c) {
				c = charToByte8(c)
			}

			next := ctx.read()
			ctx.unread(next)

			special := strings.Contains("\"';()[]#?`,.", string(next))
			if next <= 040 || special {
				return ec.makeInteger(lispInt(c)), 0, nil
			}

			return ec.read1Result(ec.invalidReadSyntax("invalid syntax for '?'"))
		case c == '"':
			builder := strings.Builder{}

			for {
				c = ctx.read()
				if c == eofRune || c == '"' {
					break
				}

				if c == '\\' {
					c, err = ec.readEscape(ctx, true)
					if err != nil {
						return nil, 0, err
					}

					if c == -1 {
						continue
					}
				}

				builder.WriteRune(c)
			}

			if c == eofRune {
				return ec.read1Result(ec.signalN(ec.g.endOfFile))
			}

			return ec.makeString(builder.String()), 0, nil
		case c == '.':
			next := ctx.read()
			ctx.unread(next)

			special := strings.Contains("\"';([#?`,", string(next))
			if next <= 040 || special {
				return nil, c, nil
			}
			return ec.read1Result(ec.readAtom(c, ctx))
		case c <= 040 || c == nbsp:
		default:
			return ec.read1Result(ec.readAtom(c, ctx))
		}
	}
}

func (ec *execContext) read0(ctx *readContext) (lispObject, error) {
	obj, c, err := ec.read1(ctx)
	if err != nil {
		return nil, err
	}

	if c != 0 {
		return ec.invalidReadSyntax("unexpected character: '%v'", c)
	}

	return obj, nil
}

func (ec *execContext) readFromString(source string) (lispObject, error) {
	ctx := readContext{
		source: source,
		runes:  []rune(source),
		i:      0,
	}

	return ec.read0(&ctx)
}

func (ec *execContext) listToSlice(list lispObject) ([]lispObject, error) {
	result := []lispObject{}
	tail := list

	for ; consp(tail); tail = xCdr(tail) {
		result = append(result, xCar(tail))
	}

	if tail != ec.nil_ {
		return nil, xErrOnly(ec.wrongTypeArgument(ec.g.listp, list))
	}

	return result, nil
}

func (ec *execContext) evalSub(form lispObject) (lispObject, error) {
	// TAGS: incomplete
	var err error

	if symbolp(form) {
		var lex lispObject

		envVal := xSymbolValue(ec.g.internalInterpreterEnv)
		if envVal != ec.nil_ {
			lex, err = ec.assq(form, envVal)
			if err != nil {
				return nil, err
			}
		} else {
			lex = ec.nil_
		}

		if lex != ec.nil_ {
			return xCdr(lex), nil
		}

		sym := xSymbol(form)
		if sym.value == ec.g.unbound {
			return ec.signalN(ec.g.voidVariable, form)
		}

		return sym.value, nil
	} else if !consp(form) {
		return form, nil
	}

	originalFn := xCar(form)
	originalArgs := xCdr(form)
	fn := originalFn

	if !symbolp(fn) {
		fn, err = ec.function(ec.makeList(fn))
		if err != nil {
			return nil, err
		}
	} else {
		fn = xSymbol(fn).function
	}

	if subroutinep(fn) {
		return ec.applySubroutine(fn, originalArgs)
	} else if fn == ec.nil_ {
		return ec.signalN(ec.g.voidFunction, originalFn)
	} else if !consp(fn) {
		return ec.signalN(ec.g.invalidFunction, originalFn)
	}

	fnCar := xCar(fn)
	if !symbolp(fnCar) {
		return ec.signalN(ec.g.invalidFunction, originalFn)
	}

	if fnCar == ec.g.macro {
		val := ec.t
		if xSymbolValue(ec.g.internalInterpreterEnv) == ec.nil_ {
			val = ec.nil_
		}

		count := ec.stackSize()
		ec.stackPushLet(ec.g.lexicalBinding, val)

		args := []lispObject{xCdr(fn), originalArgs}
		exp, err := ec.funcall(args...)

		ec.stackPopTo(count)

		if err != nil {
			return nil, err
		}

		return ec.evalSub(exp)
	} else if fnCar == ec.g.lambda || fnCar == ec.g.closure {
		return ec.applyLambda(fn, originalArgs)
	} else {
		return ec.signalN(ec.g.invalidFunction, originalFn)
	}
}

func (ec *execContext) applySubroutine(fn, originalArgs lispObject) (lispObject, error) {
	sub := xVectorLike(fn).value.(*subroutine)

	args, err := ec.listToSlice(originalArgs)
	if err != nil {
		return nil, err
	}

	if len(args) < sub.minArgs || (sub.maxArgs >= 0 && len(args) > sub.maxArgs) {
		return ec.wrongNumberOfArguments(fn, lispInt(len(args)))
	}

	if sub.maxArgs != argsUnevalled {
		for i := 0; i < len(args); i++ {
			evalled, err := ec.evalSub(args[i])
			if err != nil {
				return nil, err
			}
			args[i] = evalled
		}

		return ec.funcallSubroutine(fn, args...)
	}

	// Handle case: maxArgs == argsUnevalled
	result, err := sub.callabe1(originalArgs)
	if err != nil {
		return nil, err
	}

	if sub.noReturn {
		ec.terminate("subroutine with noreturn returned value")
	}

	return result, nil
}

func (ec *execContext) funcallSubroutine(fn lispObject, args ...lispObject) (lispObject, error) {
	sub := xVectorLike(fn).value.(*subroutine)

	if sub.maxArgs == argsUnevalled {
		return ec.signalN(ec.g.invalidFunction, fn)
	}

	if len(args) < sub.minArgs || (sub.maxArgs >= 0 && len(args) > sub.maxArgs) {
		return ec.wrongNumberOfArguments(fn, lispInt(len(args)))
	}

	missing := sub.maxArgs - len(args)
	for i := 0; i < missing && sub.maxArgs >= 0; i++ {
		args = append(args, ec.nil_)
	}

	var err error
	var result lispObject
	count := ec.stackSize()

	switch sub.maxArgs {
	case 0:
		result, err = sub.callabe0()
	case 1:
		result, err = sub.callabe1(args[0])
	case 2:
		result, err = sub.callabe2(args[0], args[1])
	case 3:
		result, err = sub.callabe3(args[0], args[1], args[2])
	case argsMany:
		result, err = sub.callabem(args...)
	default:
		ec.terminate("unknown subroutine maxargs value")
	}

	if err != nil {
		return nil, err
	}

	if sub.noReturn {
		ec.terminate("subroutine with noreturn returned value")
	}

	if count != ec.stackSize() {
		ec.terminate("subroutine did not pop one or more stack items")
	}

	return result, nil
}

func (ec *execContext) applyLambda(fn, originalArgs lispObject) (lispObject, error) {
	args, err := ec.listToSlice(originalArgs)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(args); i++ {
		evalled, err := ec.evalSub(args[i])
		if err != nil {
			return nil, err
		}
		args[i] = evalled
	}

	return ec.funcallLambda(fn, args...)
}

func (ec *execContext) funcallLambda(fn lispObject, args ...lispObject) (lispObject, error) {
	symsLeft := ec.nil_
	lexEnv := ec.nil_
	size := ec.stackSize()
	defer ec.stackPopTo(size)

	if consp(fn) {
		if xCar(fn) == ec.g.closure {
			cdr := xCdr(fn) // Drop 'closure'
			if !consp(cdr) {
				return ec.signalN(ec.g.invalidFunction, fn)
			}

			fn = cdr
			lexEnv = xCar(fn)
		}

		symsLeft = xCdr(fn)
		if !consp(symsLeft) {
			return ec.signalN(ec.g.invalidFunction, fn)
		}

		symsLeft = xCar(symsLeft)
	} else {
		ec.terminate("lambda is not a cons cell")
	}

	rest, optional := false, false
	i := 0

	for ; consp(symsLeft); symsLeft = xCdr(symsLeft) {
		next := xCar(symsLeft)

		if !symbolp(next) {
			return ec.signalN(ec.g.invalidFunction, fn)
		}

		if next == ec.g.andRest {
			if rest {
				return ec.signalN(ec.g.invalidFunction, fn)
			}
			rest = true
		} else if next == ec.g.andOptional {
			if optional {
				return ec.signalN(ec.g.invalidFunction, fn)
			}
			optional = true
		} else {
			var arg lispObject

			if rest {
				arg = ec.makeList(args[i:]...)
				i = len(args)
			} else if i < len(args) {
				arg = args[i]
				i++
			} else if !optional {
				return ec.wrongNumberOfArguments(fn, lispInt(len(args)))
			} else {
				arg = ec.nil_
			}

			if lexEnv != ec.nil_ {
				lexEnv = ec.makeCons(ec.makeCons(next, arg), lexEnv)
			} else {
				ec.stackPushLet(next, arg)
			}
		}
	}

	if symsLeft != ec.nil_ {
		return ec.signalN(ec.g.invalidFunction, fn)
	} else if i < len(args) {
		return ec.wrongNumberOfArguments(fn, lispInt(len(args)))
	}

	if lexEnv != xSymbolValue(ec.g.internalInterpreterEnv) {
		ec.stackPushLet(ec.g.internalInterpreterEnv, lexEnv)
	}

	if consp(fn) {
		return ec.progn(xCdr(xCdr(fn)))
	}

	return ec.pimacsUnimplemented(ec.g.eval, "unknown lambda body type")
}

func (ec *execContext) progIgnore(body lispObject) error {
	_, err := ec.progn(body)
	return err
}

func (ec *execContext) stackPushLet(symbol lispObject, value lispObject) {
	sym := xSymbol(symbol)
	ec.stack = append(ec.stack, &stackEntryLet{
		symbol: sym,
		oldVal: sym.value,
	})

	sym.value = value
}

func (ec *execContext) stackPushCatch(tag lispObject) {
	ec.stack = append(ec.stack, &stackEntryCatch{
		catchTag: tag,
	})
}

func (ec *execContext) catchInStack(tag lispObject) bool {
	for _, binding := range ec.stack {
		if binding.tag() != entryCatch {
			continue
		}

		catchTag := binding.(*stackEntryCatch).catchTag
		if catchTag == tag {
			return true
		}
	}

	return false
}

func (ec *execContext) stackSize() int {
	return len(ec.stack)
}

func (ec *execContext) stackPopTo(target int) {
	size := len(ec.stack)
	if target < 0 || size < target {
		ec.terminate("unable to pop back to '%v', size is '%v'", target, size)
	}

	for len(ec.stack) > target {
		current := ec.stack[len(ec.stack)-1]

		switch current.tag() {
		case entryLet:
			let := current.(*stackEntryLet)
			let.symbol.value = let.oldVal
		case entryCatch:
		default:
			ec.terminate("unknown stack item tag: '%v'", current.tag())
		}

		ec.stack = ec.stack[:len(ec.stack)-1]
	}
}

func (ec *execContext) bool(b bool) (lispObject, error) {
	if b {
		return ec.t, nil
	}
	return ec.g.nil_, nil
}

func (ec *execContext) terminate(format string, v ...interface{}) {
	// TAGS: incomplete
	terminate(format, v...)
}
