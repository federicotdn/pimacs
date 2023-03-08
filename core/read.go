package core

import (
	"strconv"
	"strings"
)

type readStackElementType int

const (
	readStackListStart = iota + 1
	readStackList
	readStackListDot
	readStackSpecial
)

type readContext interface {
	read() rune
	unread(rune)
	position() int
}

type readContextString struct {
	source string
	runes  []rune
	i      int
	end    int
}

type readStackElem struct {
	elementType readStackElementType
	special     lispObject
	listHead    lispObject
	listTail    lispObject
}

type readStack struct {
	elements []readStackElem
}

func (ctx *readContextString) read() rune {
	if ctx.i == ctx.end {
		return eofRune
	}

	r := ctx.runes[ctx.i]
	ctx.i++
	return r
}

func (ctx *readContextString) unread(c rune) {
	if c == eofRune {
		return
	}

	if ctx.i == 0 {
		terminate("unbalanced read/unread call with rune: '%v'", c)
	}

	ctx.i--
}

func (ctx *readContextString) position() int {
	return ctx.i
}

func (stack *readStack) push(elem readStackElem) {
	stack.elements = append(stack.elements, elem)
}

func (stack *readStack) peek() *readStackElem {
	if stack.isEmpty() {
		terminate("can't peek: read stack is empty")
	}
	return &stack.elements[len(stack.elements)-1]
}

func (stack *readStack) pop() readStackElem {
	if stack.isEmpty() {
		terminate("can't pop: read stack is empty")
	}
	elem := stack.elements[len(stack.elements)-1]
	stack.elements = stack.elements[:len(stack.elements)-1]
	return elem
}

func (stack *readStack) isEmpty() bool {
	return len(stack.elements) == 0
}

func (ec *execContext) skipSpaceAndComments(ctx readContext) error {
	var c rune

	// Imitate do...while loop
	for ok := true; ok; ok = (c <= 040 || c == nbsp) {
		c = ctx.read()

		if c == ';' {
			// Imitate do...while loop (again)
			for ok2 := true; ok2; ok2 = (c >= 0 && c != '\n') {
			}
		}

		if c < 0 {
			return xErrOnly(ec.signalN(ec.g.endOfFile))
		}
	}

	ctx.unread(c)
	return nil
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

func (ec *execContext) readEscape(ctx readContext, stringp bool) (rune, error) {
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

func (ec *execContext) readStringLiteral(ctx readContext) (lispObject, error) {
	builder := strings.Builder{}
	var c rune

	for {
		c = ctx.read()
		if c == eofRune || c == '"' {
			break
		}

		if c == '\\' {
			var err error
			c, err = ec.readEscape(ctx, true)
			if err != nil {
				return nil, err
			}

			if c == -1 {
				continue
			}
		}

		builder.WriteRune(c)
	}

	if c == eofRune {
		return ec.signalN(ec.g.endOfFile)
	}

	return ec.makeString(builder.String()), nil
}

func (ec *execContext) readCharLiteral(ctx readContext) (lispObject, error) {
	c := ctx.read()
	if c == eofRune {
		return ec.signalN(ec.g.endOfFile)
	}

	if c == ' ' || c == '\t' {
		return ec.makeInteger(lispInt(c)), nil
	}

	if c == '\\' {
		var err error
		c, err = ec.readEscape(ctx, false)
		if err != nil {
			return nil, err
		}
	}

	if charByte8(c) {
		c = charToByte8(c)
	}

	next := ctx.read()
	ctx.unread(next)

	special := strings.Contains("\"';()[]#?`,.", string(next))
	if next <= 040 || special {
		return ec.makeInteger(lispInt(c)), nil
	}

	return ec.invalidReadSyntax("invalid syntax for '?'")
}

func (ec *execContext) readSymbol(c rune, ctx readContext) (lispObject, error) {
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

	return ec.intern(ec.makeString(s), ec.nil_)
}

func (ec *execContext) read0(ctx readContext) (lispObject, error) {
	// TAGS: incomplete
	stack := readStack{}
	var obj lispObject
	var err error
	var c rune

read_obj:
	c = ctx.read()
	if c == eofRune {
		return ec.signalN(ec.g.endOfFile)
	}

	switch {
	case c == '(':
		stack.push(readStackElem{elementType: readStackListStart})
		goto read_obj
	case c == '[':
		return ec.pimacsUnimplemented(ec.g.read, "unknown token: '['")
	case c == ')':
		if stack.isEmpty() {
			return ec.signalN(ec.g.invalidReadSyntax, ec.makeString(")"))
		}

		switch stack.peek().elementType {
		case readStackListStart:
			stack.pop()
			obj = ec.nil_
		case readStackList:
			obj = stack.pop().listHead
		default:
			return ec.signalN(ec.g.invalidReadSyntax, ec.makeString(")"))
		}
	case c == ']':
		return ec.pimacsUnimplemented(ec.g.read, "unknown token: ']'")
	case c == '#':
		c = ctx.read()

		switch {
		case c == '#':
			obj = ec.g.emptySymbol
		case c == '\'':
			stack.push(readStackElem{
				elementType: readStackSpecial,
				special:     ec.g.function,
			})
			goto read_obj
		default:
			return ec.pimacsUnimplemented(ec.g.read, "unknown token: '#%c'", c)
		}
	case c == ';':
		for {
			c = ctx.read()
			if c == eofRune || c == '\n' {
				break
			}
		}
		goto read_obj
	case c == '\'':
		stack.push(readStackElem{elementType: readStackSpecial, special: ec.g.quote})
		goto read_obj
	case c == '`':
		stack.push(readStackElem{elementType: readStackSpecial, special: ec.g.backquote})
		goto read_obj
	case c == ',':
		return ec.pimacsUnimplemented(ec.g.read, "unknown token: ','")
	case c == '?':
		obj, err = ec.readCharLiteral(ctx)
		if err != nil {
			return nil, err
		}
	case c == '"':
		obj, err = ec.readStringLiteral(ctx)
		if err != nil {
			return nil, err
		}
	case c == '.':
		next := ctx.read()
		ctx.unread(next)

		special := strings.Contains("\"';([#?`,", string(next))
		if next <= 040 || c == nbsp || special {
			if !stack.isEmpty() && (stack.peek().elementType == readStackList ||
				stack.peek().elementType == readStackListStart) {
				stack.peek().elementType = readStackListDot
				goto read_obj
			}

			println("next", string([]rune{next}), next)
			return ec.signalN(ec.g.invalidReadSyntax, ec.makeString("."))
		}

		obj, err = ec.readSymbol(c, ctx)
		if err != nil {
			return nil, err
		}
	case c <= 040 || c == nbsp:
		goto read_obj
	default:
		obj, err = ec.readSymbol(c, ctx)
		if err != nil {
			return nil, err
		}
	}

	for !stack.isEmpty() {
		top := stack.peek()

		switch top.elementType {
		case readStackListStart:
			top.elementType = readStackList
			top.listHead = ec.makeCons(obj, ec.nil_)
			top.listTail = top.listHead
			goto read_obj

		case readStackList:
			tail := ec.makeCons(obj, ec.nil_)
			xSetCdr(top.listTail, tail)
			top.listTail = tail
			goto read_obj

		case readStackListDot:
			ec.skipSpaceAndComments(ctx)
			ch := ctx.read()
			if ch != ')' {
				return ec.signalN(ec.g.invalidReadSyntax, ec.makeString("expected )"))
			}

			stack.pop()

			if top.listHead == nil {
				// We never got to handle readStackListStart, since we got
				// (. x) as input. Use obj (set previously) as the return value.
				break
			}

			xSetCdr(top.listTail, obj)
			obj = top.listHead

		case readStackSpecial:
			stack.pop()
			obj = ec.makeList(top.special, obj)
		default:
			return ec.signalN(ec.g.read, ec.makeString("unknown read stack entry type"))
		}
	}

	if obj == nil {
		ec.terminate("read0 failed to read object")
	}

	return obj, nil

}

func (ec *execContext) readInternalStart(stream, start, end lispObject) (lispObject, readContext, error) {
	var ctx readContext

	switch stream.getType() {
	case lispTypeString:
		length, err := ec.length(stream)
		if err != nil {
			return nil, nil, err
		}

		lengthInt := int(xIntegerValue(length))
		startInt := 0
		endInt := lengthInt

		if integerp(start) {
			startInt = int(xIntegerValue(start))

			if startInt < 0 {
				startInt += lengthInt
			}
		} else if start != ec.nil_ {
			return nil, nil, xErrOnly(ec.wrongTypeArgument(ec.g.integerp, start))
		}

		if integerp(end) {
			endInt = int(xIntegerValue(end))

			if endInt < 0 {
				endInt += lengthInt
			}

		} else if end != ec.nil_ {
			return nil, nil, xErrOnly(ec.wrongTypeArgument(ec.g.integerp, end))
		}

		if !(0 <= startInt && startInt <= endInt && endInt <= lengthInt) {
			return nil, nil, xErrOnly(ec.signalN(ec.g.argsOutOfRange, stream, start, end))
		}

		str := xStringValue(stream)

		ctx = &readContextString{
			source: str,
			runes:  []rune(str),
			i:      startInt,
			end:    endInt,
		}
	default:
		return nil, nil, xErrOnly(ec.pimacsUnimplemented(ec.g.read, "unknown source type"))
	}

	result, err := ec.read0(ctx)
	return result, ctx, err
}

func (ec *execContext) intern(str, _ lispObject) (lispObject, error) {
	if !stringp(str) {
		return ec.wrongTypeArgument(ec.g.stringp, str)
	}

	name := xStringValue(str)

	sym, ok := ec.obarray[name]
	if !ok {
		sym = ec.makeSymbol(name)
		ec.internNewSymbol(sym, true)
	}

	return sym, nil
}

func (ec *execContext) read(stream lispObject) (lispObject, error) {
	if stream == ec.nil_ {
		stream = xSymbolValue(ec.g.standardInput)
	}

	if stream == ec.t {
		stream = ec.g.readChar
	}

	if stream == ec.g.readChar {
		return ec.funcall(ec.g.readFromMinibuffer, ec.makeString("Lisp expression: "))
	}

	result, _, err := ec.readInternalStart(stream, ec.nil_, ec.nil_)
	return result, err
}

func (ec *execContext) readFromString(str, start, end lispObject) (lispObject, error) {
	if !stringp(str) {
		return ec.wrongTypeArgument(ec.g.stringp, str)
	}

	result, ctx, err := ec.readInternalStart(str, start, end)
	if err != nil {
		return nil, err
	}

	pos := ec.makeInteger(lispInt(ctx.position()))
	return ec.cons(result, pos)
}

func (ec *execContext) load(file, noError, noMessage, noSuffix, mustSuffix lispObject) (lispObject, error) {
	if !stringp(file) {
		return ec.wrongTypeArgument(ec.g.stringp, file)
	}

	return nil, nil
}

func (ec *execContext) symbolsOfRead() {
	ec.defVar(&ec.g.standardInput, "standard-input", ec.t)
	ec.defVar(&ec.g.lexicalBinding, "lexical-binding", ec.nil_)
	ec.defVar(&ec.g.loadPath, "load-path", ec.nil_)
	ec.defVarUninterned(&ec.g.readChar, "read-char", ec.g.unbound)

	ec.defSubr2(nil, "intern", ec.intern, 1)
	ec.defSubr3(nil, "read-from-string", ec.readFromString, 1)
	ec.defSubr1(&ec.g.read, "read", ec.read, 0)
	ec.defSubr5(nil, "load", ec.load, 1)
}
