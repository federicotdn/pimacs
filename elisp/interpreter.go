package elisp

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	eightBitCodeOffset rune = 0x3fff00
	max5ByteChar            = 0x3fff7f
	eofRune                 = -1
	nbsp                    = '\u00A0'
)

type Interpreter interface {
	Print(obj lispObject, ec *execContext) (string, error)
	ReadString(source string, ec *execContext) (lispObject, error)
	ReadEvalPrint(source string, ec *execContext) (string, error)
}

type interpreter struct {
}

type specBinding struct {
}

type execContext struct {
	specpdl []specBinding
	inp     *interpreter
	nil_    lispObject
	obarray map[string]*lispSymbol
}

type readContext struct {
	source string
	runes  []rune
	i      int
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
		panic("nothing to unread")
	}

	ctx.i--
}

func (ec *execContext) xCar(obj lispObject) lispObject {
	return obj.(*lispCons).car
}

func (ec *execContext) xCdr(obj lispObject) lispObject {
	return obj.(*lispCons).cdr
}

func (ec *execContext) xSetCdr(obj, newCdr lispObject) lispObject {
	obj.(*lispCons).cdr = newCdr
	return newCdr
}

func (inp *interpreter) Print(obj lispObject, ec *execContext) (string, error) {
	if ec == nil {
		ec = NewExecContext()
	}

	str, err := ec.prin1(obj)
	if err != nil {
		return "", err
	}

	return str.(*lispString).value, nil
}

func (ec *execContext) prin1(obj lispObject) (lispObject, error) {
	type_ := obj.getType()
	var s string

	switch type_ {
	case symbol:
		s = obj.(*lispSymbol).name
	case integer:
		s = fmt.Sprint(obj.(*lispInteger).value)
	case string_:
		s = "\"" + obj.(*lispString).value + "\""
	case vector:
		return nil, fmt.Errorf("prin1 unimplemented")
	case cons:
		// TODO: Clean up when string functions are avaiable (don't type-assert)
		lc := obj.(*lispCons)
		current := lc

		carStr, err := ec.prin1(lc.car)
		if err != nil {
			return nil, err
		}
		s = "(" + carStr.(*lispString).value

		for {
			next, ok := current.cdr.(*lispCons)
			if ok {
				nextStr, err := ec.prin1(next.car)
				if err != nil {
					return nil, err
				}

				s += " " + nextStr.(*lispString).value
				current = next
			} else {
				if current.cdr != ec.nil_ {
					cdrStr, err := ec.prin1(current.cdr)
					if err != nil {
						return nil, err
					}
					s += " . " + cdrStr.(*lispString).value
				}

				break
			}
		}

		s += ")"
	case float:
		s = fmt.Sprint(obj.(*lispFloat).value)
	default:
		return nil, fmt.Errorf("prin1 unimplemented for '%v'", type_)
	}

	return ec.makeString(s), nil
}

func (ec *execContext) stringToNumber(s string) (lispObject, error) {
	nInt, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return ec.makeInteger(lispInt(nInt)), nil
	}

	nFloat, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return ec.makeFloat(nFloat), nil
	}

	return nil, fmt.Errorf("unknown number format")
}

func (ec *execContext) makeSymbol(name string) *lispSymbol {
	return &lispSymbol{
		name: name,
	}
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

func (ec *execContext) makeFloat(value float64) *lispFloat {
	return &lispFloat{
		value: value,
	}
}

func (ec *execContext) makeString(value string) *lispString {
	return &lispString{
		value: value,
	}
}

func (ec *execContext) makeList(obj lispObject, objs ...lispObject) *lispCons {
	tmp := ec.makeCons(obj, ec.nil_)
	val := tmp

	for _, obj := range objs {
		tmp.cdr = ec.makeCons(obj, ec.nil_)
		tmp = tmp.cdr.(*lispCons)
	}

	return val
}

func (ec *execContext) readEscape(ctx *readContext, stringp bool) (rune, error) {
	c := ctx.read()
	if c == eofRune {
		return 0, fmt.Errorf("eof")
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

	return 0, fmt.Errorf("unimplemented escape code: '%v", string(c))
}

func (ec *execContext) defun0(name string, fn lispFn0) *lispSymbol {
	sym := ec.intern(name)
	sym.function = &builtInFunction{callabe0: fn}
	return sym
}

func (ec *execContext) defun1(name string, fn lispFn1, minArgs int) *lispSymbol {
	sym := ec.intern(name)
	sym.function = &builtInFunction{callabe1: fn, minArgs: minArgs, maxArgs: 1}
	return sym
}

func (ec *execContext) defun2(name string, fn lispFn2, minArgs int) *lispSymbol {
	sym := ec.intern(name)
	sym.function = &builtInFunction{callabe2: fn, minArgs: minArgs, maxArgs: 2}
	return sym
}

func (ec *execContext) defunM(name string, fn lispFnM, minArgs int) *lispSymbol {
	sym := ec.intern(name)
	sym.function = &builtInFunction{callabem: fn, minArgs: minArgs, maxArgs: argsMany}
	return sym
}

func (ec *execContext) defunU(name string, fn lispFn1, minArgs int) *lispSymbol {
	sym := ec.intern(name)
	sym.function = &builtInFunction{callabe1: fn, minArgs: minArgs, maxArgs: argsUnevalled}
	return sym
}

func (ec *execContext) initialDefs() {
	nil_ := ec.intern("nil")
	nil_.value = nil_

	t := ec.intern("t")
	t.value = t
}

func NewInterpreter() Interpreter {
	return &interpreter{}
}

func NewExecContext() *execContext {
	ec := execContext{}
	ec.obarray = make(map[string]*lispSymbol)

	ec.initialDefs()          // interpreter.go
	ec.initialDefsFunctions() // functions.go

	ec.nil_ = ec.intern("nil")

	return &ec
}

func (ec *execContext) intern(name string) *lispSymbol {
	sym, ok := ec.obarray[name]
	if !ok {
		sym = ec.makeSymbol(name)
		ec.obarray[name] = sym
	}

	return sym
}

func (ec *execContext) setInternal(symbol, newVal lispObject) (lispObject, error) {
	sym, ok := symbol.(*lispSymbol)
	if !ok {
		return nil, fmt.Errorf("not a symbol")
	}

	sym.value = newVal
	return newVal, nil
}

func (ec *execContext) readList(ctx *readContext) (lispObject, error) {
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
					ec.xSetCdr(tail, obj)
				} else {
					val, err = ec.read0(ctx)
				}

				_, c, _ = ec.read1(ctx)
				if c == ')' {
					return val, nil
				}

				return nil, fmt.Errorf("'.' in wrong context")
			default:
				return nil, fmt.Errorf("invalid list ending: '%v'", string(c))
			}
		}

		tmp := ec.makeCons(elt, ec.nil_)

		if tail != ec.nil_ {
			ec.xSetCdr(tail, tmp)
		} else {
			val = tmp
		}

		tail = tmp
	}
}

func (ec *execContext) readAtom(c rune, ctx *readContext) (lispObject, error) {
	quoted := false
	builder := strings.Builder{}

	for {
		if c == '\\' {
			c = ctx.read()
			if c == eofRune {
				return nil, fmt.Errorf("eof reading atom")
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
	var err error
	var c rune

	for {
		c = ctx.read()
		if c == eofRune {
			return nil, 0, fmt.Errorf("eof on read1")
		}

		switch {
		case c == '(':
			return ec.read1Result(ec.readList(ctx))
		case c == '[':
			return nil, 0, fmt.Errorf("unimplented read: '%v'", string(c))
		case c == ')' || c == ']':
			return nil, c, nil
		case c == '#':
			return nil, 0, fmt.Errorf("unimplented read: '%v'", string(c))
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

			name := "quote"
			if c == '`' {
				name = "`"
			}

			list := ec.makeList(ec.intern(name), obj)
			return list, 0, nil
		case c == ',':
			return nil, 0, fmt.Errorf("unimplented read: '%v'", string(c))
		case c == '?':
			c = ctx.read()
			if c == eofRune {
				return nil, 0, fmt.Errorf("eof reading character")
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

			return nil, 0, fmt.Errorf("invalid syntax for '?'")
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
				return nil, c, fmt.Errorf("eof reading string")
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
		return nil, fmt.Errorf("unexpected character: '%v'", string(c))
	}

	return obj, nil
}

func (inp *interpreter) ReadString(source string, ec *execContext) (lispObject, error) {
	ctx := readContext{
		source: source,
		runes:  []rune(source),
		i:      0,
	}

	if ec == nil {
		ec = NewExecContext()
	}

	return ec.read0(&ctx)
}

func (inp *interpreter) ReadEvalPrint(source string, ec *execContext) (string, error) {
	if ec == nil {
		ec = NewExecContext()
	}

	obj, err := inp.ReadString(source, ec)
	if err != nil {
		return "", err
	}

	result, err := ec.evalSub(obj)
	if err != nil {
		return "", err
	}

	return inp.Print(result, ec)
}

func (ec *execContext) evalSub(form lispObject) (lispObject, error) {
	if form.getType() == symbol {
		// TODO: Use environments!
		sym := form.(*lispSymbol)
		if sym.value == nil {
			return nil, fmt.Errorf("void-variable '%v'", sym.name)
		}

		return sym.value, nil
	} else if form.getType() != cons {
		return form, nil
	}

	car := ec.xCar(form)
	cdr := ec.xCdr(form)
	original := cdr

	sym := car.(*lispSymbol)
	fn := sym.function

	if fn == nil {
		return nil, fmt.Errorf("void function '%v'", sym.name)
	}

	args := []lispObject{}
	nArgs := 0

	for {
		if cdr != ec.nil_ && cdr.getType() != cons {
			return nil, fmt.Errorf("wrong type argument: '%v'", cdr.getType())
		}

		if cdr == ec.nil_ {
			break
		}

		nArgs++

		arg := ec.xCar(cdr)
		var processed lispObject

		if fn.maxArgs != argsUnevalled {
			var err error
			processed, err = ec.evalSub(arg)
			if err != nil {
				return nil, err
			}

			args = append(args, processed)
		}

		cdr = ec.xCdr(cdr)
	}

	if nArgs < fn.minArgs {
		return nil, fmt.Errorf("expected at least %v arguments but got %v", fn.minArgs, len(args))
	} else if fn.maxArgs >= 0 && nArgs > fn.maxArgs {
		return nil, fmt.Errorf("expected at most %v arguments but got %v", fn.maxArgs, len(args))
	}

	if fn.maxArgs >= 0 {
		for nArgs < fn.maxArgs {
			args = append(args, ec.nil_)
			nArgs++
		}
	}

	switch fn.maxArgs {
	case 0:
		return fn.callabe0()
	case 1:
		return fn.callabe1(args[0])
	case 2:
		return fn.callabe2(args[0], args[1])
	case argsMany:
		return fn.callabem(args...)
	case argsUnevalled:
		return fn.callabe1(original)
	default:
		return nil, fmt.Errorf("unknown max args value: '%v'", fn.maxArgs)
	}
}

func (ec *execContext) apply(fn lispObject, args ...lispObject) (lispObject, error) {
	return nil, nil
}
