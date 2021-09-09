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


type specBindingTag int

const (
	specLet specBindingTag = iota
)

type specBinding interface {
	tag() specBindingTag
}

type execContext struct {
	bindings []specBinding
	inp      *interpreter
	nil_     lispObject
	t        lispObject
	unbound  lispObject
	obarray  map[string]*lispSymbol
	env      *lispSymbol
}

type specBindingLet struct {
	symbol *lispSymbol
	oldVal lispObject
}

func (sbl *specBindingLet) tag() specBindingTag {
	return specLet
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

func (ec *execContext) makeSymbolBase(name string) *lispSymbol {
	return &lispSymbol{
		name: name,
	}
}

func (ec *execContext) makeSymbol(name string) *lispSymbol {
	if ec.nil_ == nil || ec.unbound == nil {
		panic("context not initialized")
	}

	base := ec.makeSymbolBase(name)
	base.value = ec.unbound
	base.function = ec.nil_
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

func (ec *execContext) makeVectorLike(vecType vectorLikeType, value vectorLikeValue) *lispVectorLike {
	return &lispVectorLike{
		vecType: vecType,
		value:   value,
	}
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

func (ec *execContext) defSubr(name string, sub *subroutine) *lispSymbol {
	sym := ec.intern(name)
	vec := ec.makeVectorLike(vectorLikeTypeSubroutine, sub)
	sym.function = vec
	return sym
}

func (ec *execContext) defSubr0(name string, fn lispFn0) *lispSymbol {
	return ec.defSubr(name, &subroutine{
		callabe0: fn,
	})
}

func (ec *execContext) defSubr1(name string, fn lispFn1, minArgs int) *lispSymbol {
	return ec.defSubr(name, &subroutine{
		callabe1: fn,
		minArgs:  minArgs,
		maxArgs:  1,
	})
}

func (ec *execContext) defSubr2(name string, fn lispFn2, minArgs int) *lispSymbol {
	return ec.defSubr(name, &subroutine{
		callabe2: fn,
		minArgs:  minArgs,
		maxArgs:  2,
	})
}

func (ec *execContext) defSubrM(name string, fn lispFnM, minArgs int) *lispSymbol {
	return ec.defSubr(name, &subroutine{
		callabem: fn,
		minArgs:  minArgs,
		maxArgs:  argsMany,
	})
}

func (ec *execContext) defSubrU(name string, fn lispFn1, minArgs int) *lispSymbol {
	return ec.defSubr(name, &subroutine{
		callabe1: fn,
		minArgs:  minArgs,
		maxArgs:  argsUnevalled,
	})
}

func (ec *execContext) initialDefs() {
	nil_ := ec.intern("nil")
	nil_.value = nil_

	t := ec.intern("t")
	t.value = t
}

func newExecContext() *execContext {
	ec := execContext{}
	ec.obarray = make(map[string]*lispSymbol)
	ec.bindings = []specBinding{}

	nil_ := ec.makeSymbolBase("nil")
	nil_.value = nil_
	nil_.function = nil_
	ec.obarray["nil"] = nil_
	ec.nil_ = nil_

	ec.unbound = ec.makeSymbolBase("unbound")

	t := ec.intern("t")
	t.value = t
	ec.t = t

	env := ec.makeSymbol("internal-interpreter-environment")
	env.value = ec.nil_
	ec.env = env

	ec.initialDefs()          // exec_context.go
	ec.initialDefsFunctions() // functions.go

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

func (ec *execContext) readFromString(source string) (lispObject, error) {
	ctx := readContext{
		source: source,
		runes:  []rune(source),
		i:      0,
	}

	return ec.read0(&ctx)
}

func (ec *execContext) evalSub(form lispObject) (lispObject, error) {
	var err error

	if form.getType() == symbol {
		var lex lispObject

		if ec.env.value != ec.nil_ {
			lex, err = ec.assq(form, ec.env.value)
			if err != nil {
				return nil, err
			}
		} else {
			lex = ec.nil_
		}

		if lex != ec.nil_ {
			return ec.xCdr(lex), nil
		}

		sym := form.(*lispSymbol)
		if sym.value == ec.unbound {
			return nil, fmt.Errorf("void-variable '%v'", sym.name)
		}

		return sym.value, nil
	} else if form.getType() != cons {
		return form, nil
	}

	car := ec.xCar(form)
	cdr := ec.xCdr(form)
	original := cdr

	if car.getType() != symbol {
		return nil, fmt.Errorf("function is not a symbol")
	} else if car == ec.nil_ {
		return nil, fmt.Errorf("void-function")
	}

	sym := car.(*lispSymbol)
	fn := sym.function

	if fn.getType() != vectorLike {
		return nil, fmt.Errorf("function must be vector-like")
	}

	vec := fn.(*lispVectorLike)
	if vec.vecType != vectorLikeTypeSubroutine {
		return nil, fmt.Errorf("vector must be a subroutine")
	}

	sub := vec.value.(*subroutine)

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

		if sub.maxArgs != argsUnevalled {
			processed, err = ec.evalSub(arg)
			if err != nil {
				return nil, err
			}

			args = append(args, processed)
		}

		cdr = ec.xCdr(cdr)
	}

	if nArgs < sub.minArgs {
		return nil, fmt.Errorf("expected at least %v arguments but got %v", sub.minArgs, len(args))
	} else if sub.maxArgs >= 0 && nArgs > sub.maxArgs {
		return nil, fmt.Errorf("expected at most %v arguments but got %v", sub.maxArgs, len(args))
	}

	if sub.maxArgs >= 0 {
		for nArgs < sub.maxArgs {
			args = append(args, ec.nil_)
			nArgs++
		}
	}

	switch sub.maxArgs {
	case 0:
		return sub.callabe0()
	case 1:
		return sub.callabe1(args[0])
	case 2:
		return sub.callabe2(args[0], args[1])
	case argsMany:
		return sub.callabem(args...)
	case argsUnevalled:
		return sub.callabe1(original)
	default:
		return nil, fmt.Errorf("unknown max args value: '%v'", sub.maxArgs)
	}
}

func (ec *execContext) apply(fn lispObject, args ...lispObject) (lispObject, error) {
	return nil, nil
}

func (ec *execContext) specBind(symbol *lispSymbol, value lispObject) {
	ec.bindings = append(ec.bindings, &specBindingLet{
		symbol: symbol,
		oldVal: symbol.value,
	})

	symbol.value = value
}

func (ec *execContext) bindingsSize() int {
	return len(ec.bindings)
}

func (ec *execContext) unbindTo(target int, obj lispObject, err error) (lispObject, error) {
	if err != nil {
		return nil, err
	}

	// TODO: Check if defer can be used for this
	size := len(ec.bindings)
	if size < target {
		panic(fmt.Sprintf("unable to unbind back to %v, size is %v", target, size))
	}

	for len(ec.bindings) > target {
		current := ec.bindings[len(ec.bindings)-1]

		switch current.tag() {
		case specLet:
			let := current.(*specBindingLet)
			let.symbol.value = let.oldVal
		default:
			panic("unknown specBinding tag")
		}

		ec.bindings = ec.bindings[:len(ec.bindings)-1]
	}

	return obj, nil
}

