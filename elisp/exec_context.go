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
	entryLet stackEntryTag = iota
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
	inp     *interpreter
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
	// TODO: revise
	return fmt.Sprintf("stack jump: tag: %v", jmp.tag)
}

func (jmp *stackJumpSignal) Error() string {
	// TODO: revise
	return fmt.Sprintf("stack jump: signal: %v", jmp.errorSymbol)
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
		panic("nothing to unread")
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

	panic("unimplemented escape code")
}

func (ec *execContext) defSubr(name string, sub *subroutine) {
	if sub.maxArgs >= 0 && sub.minArgs > sub.maxArgs {
		panic("min args must be smaller or equal to max args")
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
		panic("different symbol with that name already interned")
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

				return nil, fmt.Errorf("'.' in wrong context")
			default:
				return nil, fmt.Errorf("invalid list ending: '%v'", string(c))
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

	if form.getType() == lispTypeSymbol {
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
			return nil, fmt.Errorf("void-variable '%v'", sym.name)
		}

		return sym.value, nil
	} else if form.getType() != lispTypeCons {
		return form, nil
	}

	car := xCar(form)
	cdr := xCdr(form)
	original := cdr

	if car.getType() != lispTypeSymbol {
		return nil, fmt.Errorf("function is not a symbol: %v", car.getType())
	} else if car == ec.nil_ {
		return nil, fmt.Errorf("void-function")
	}

	sym := xSymbol(car)
	fn := sym.function

	if fn.getType() != lispTypeVectorLike {
		return nil, fmt.Errorf("function must be vector-like")
	}

	vec := xVectorLike(fn)
	if vec.vecType != vectorLikeTypeSubroutine {
		return nil, fmt.Errorf("vector must be a subroutine")
	}

	sub := vec.value.(*subroutine)

	args := []lispObject{}
	nArgs := 0

	for {
		if cdr != ec.nil_ && cdr.getType() != lispTypeCons {
			return nil, fmt.Errorf("wrong type argument: '%v'", cdr.getType())
		}

		if cdr == ec.nil_ {
			break
		}

		nArgs++

		arg := xCar(cdr)
		var processed lispObject

		if sub.maxArgs != argsUnevalled {
			processed, err = ec.evalSub(arg)
			if err != nil {
				return nil, err
			}

			args = append(args, processed)
		}

		cdr = xCdr(cdr)
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
	case argsUnevalled:
		result, err = sub.callabe1(original)
	default:
		panic("unknown maxargs value")
	}

	if err != nil {
		return nil, err
	}

	if sub.noReturn {
		panic("subroutine with noreturn returned value")
	}

	if count != ec.stackSize() {
		panic("one or more bindings not undone after call")
	}

	return result, nil
}

func (ec *execContext) apply(fn lispObject, args ...lispObject) (lispObject, error) {
	return nil, nil
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

func (ec *execContext) unbind(target int) {
	size := len(ec.stack)
	if target < 0 || size <= target {
		panic(fmt.Sprintf("unable to unbind back to %v, size is %v", target, size))
	}

	for len(ec.stack) > target {
		current := ec.stack[len(ec.stack)-1]

		switch current.tag() {
		case entryLet:
			let := current.(*stackEntryLet)
			let.symbol.value = let.oldVal
		case entryCatch:
			// Nothing to do.
		default:
			panic("unknown specBinding tag")
		}

		ec.stack = ec.stack[:len(ec.stack)-1]
	}
}
