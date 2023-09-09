package core

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode/utf8"
)

type readStackElementType int

const (
	readStackListStart = iota + 1
	readStackList
	readStackListDot
	readStackSpecial
	readStackVector
)

type readContext interface {
	read() rune
	unread(rune)
	position() int
}

type readContextString struct {
	runes []rune
	i     int
	end   int
}

type readContextFile struct {
	reader      *bufio.Reader
	unreadRunes []rune
	i           int
}

type readStackElem struct {
	elementType readStackElementType
	special     lispObject
	listHead    lispObject
	listTail    lispObject
	elems       []lispObject
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

func (ctx *readContextFile) read() rune {
	if len(ctx.unreadRunes) > 0 {
		r := ctx.unreadRunes[len(ctx.unreadRunes)-1]
		ctx.unreadRunes = ctx.unreadRunes[:len(ctx.unreadRunes)-1]
		ctx.i++
		return r
	}

	r, _, err := ctx.reader.ReadRune()
	if err != nil {
		return eofRune
	}

	ctx.i++
	return r
}

func (ctx *readContextFile) unread(c rune) {
	if c == eofRune {
		return
	}

	if ctx.i == 0 {
		terminate("unbalanced read/unread call with rune: '%v'", c)
	}

	ctx.unreadRunes = append(ctx.unreadRunes, c)
	ctx.i--
}

func (ctx *readContextFile) position() int {
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
			return xErrOnly(ec.signalN(ec.s.endOfFile))
		}
	}

	ctx.unread(c)
	return nil
}

func (ec *execContext) stringToNumber(s string) (lispObject, error) {
	nInt, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return newInteger(lispInt(nInt)), nil
	}

	nFloat, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return newFloat(lispFp(nFloat)), nil
	}

	return ec.pimacsUnimplemented(ec.s.read, "unknown number format: '%v'", s)
}

func (ec *execContext) readEscape(ctx readContext, stringp bool) (rune, error) {
	c := ctx.read()
	if c == eofRune {
		return 0, xErrOnly(ec.signalN(ec.s.endOfFile))
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

	return 0, xErrOnly(ec.pimacsUnimplemented(ec.s.read, "unknown escape code: '\\%v'", c))
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
		return ec.signalN(ec.s.endOfFile)
	}

	return newString(builder.String()), nil
}

func (ec *execContext) readCharLiteral(ctx readContext) (lispObject, error) {
	c := ctx.read()
	if c == eofRune {
		return ec.signalN(ec.s.endOfFile)
	}

	if c == ' ' || c == '\t' {
		return newInteger(lispInt(c)), nil
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
		return newInteger(lispInt(c)), nil
	}

	return ec.invalidReadSyntax("invalid syntax for '?'")
}

func (ec *execContext) readSymbol(c rune, ctx readContext) (lispObject, error) {
	quoted := false
	builder := strings.Builder{}

	for {
		if c == '\\' {
			c = ctx.read()
			if c == eofRune {
				return ec.signalN(ec.s.endOfFile)
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

	return ec.intern(newString(s), ec.nil_)
}

func (ec *execContext) read0(ctx readContext) (lispObject, error) {
	stack := readStack{}
	var obj lispObject
	var err error
	var c rune

read_obj:
	c = ctx.read()
	if c == eofRune {
		return ec.signalN(ec.s.endOfFile)
	}

	switch {
	case c == '(':
		stack.push(readStackElem{elementType: readStackListStart})
		goto read_obj
	case c == '[':
		stack.push(readStackElem{elementType: readStackVector})
		goto read_obj
	case c == ')':
		if stack.isEmpty() {
			return ec.signalN(ec.s.invalidReadSyntax, newString(")"))
		}

		switch stack.peek().elementType {
		case readStackListStart:
			stack.pop()
			obj = ec.nil_
		case readStackList:
			obj = stack.pop().listHead
		default:
			return ec.signalN(ec.s.invalidReadSyntax, newString(")"))
		}
	case c == ']':
		if stack.isEmpty() {
			return ec.signalN(ec.s.invalidReadSyntax, newString("]"))
		}
		switch stack.peek().elementType {
		case readStackVector:
			obj = newVector(stack.pop().elems)
		default:
			return ec.signalN(ec.s.invalidReadSyntax, newString("]"))
		}
	case c == '#':
		c = ctx.read()

		switch {
		case c == '#':
			obj = ec.s.emptySymbol
		case c == '\'':
			stack.push(readStackElem{
				elementType: readStackSpecial,
				special:     ec.s.function,
			})
			goto read_obj
		default:
			return ec.pimacsUnimplemented(ec.s.read, "unknown token: '#%c'", c)
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
		stack.push(readStackElem{elementType: readStackSpecial, special: ec.s.quote})
		goto read_obj
	case c == '`':
		stack.push(readStackElem{elementType: readStackSpecial, special: ec.s.backquote})
		goto read_obj
	case c == ',':
		next := ctx.read()
		var sym lispObject

		if next == '@' {
			sym = ec.s.commaAt
		} else {
			if next >= 0 {
				ctx.unread(next)
			}
			sym = ec.s.comma
		}
		stack.push(readStackElem{elementType: readStackSpecial, special: sym})
		goto read_obj
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

			return ec.signalN(ec.s.invalidReadSyntax, newString("."))
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
			top.listHead = newCons(obj, ec.nil_)
			top.listTail = top.listHead
			goto read_obj

		case readStackList:
			tail := newCons(obj, ec.nil_)
			xSetCdr(top.listTail, tail)
			top.listTail = tail
			goto read_obj

		case readStackListDot:
			err := ec.skipSpaceAndComments(ctx)
			if err != nil {
				return nil, err
			}
			ch := ctx.read()
			if ch != ')' {
				return ec.signalN(ec.s.invalidReadSyntax, newString("expected )"))
			}

			stack.pop()

			if top.listHead == nil {
				// See Emacs bug #62020
				return ec.signalN(ec.s.invalidReadSyntax, newString("."))
			}

			xSetCdr(top.listTail, obj)
			obj = top.listHead
		case readStackVector:
			top.elems = append(top.elems, obj)
			goto read_obj
		case readStackSpecial:
			stack.pop()
			obj = ec.makeList(top.special, obj)
		default:
			return ec.signalN(ec.s.read, newString("unknown read stack entry type"))
		}
	}

	if obj == nil {
		ec.terminate("read0 failed to read object")
	}

	return obj, nil

}

func (ec *execContext) streamReadContext(stream, start, end lispObject) (readContext, error) {
	var ctx readContext

	switch stream.getType() {
	case lispTypeString:
		length, err := ec.length(stream)
		if err != nil {
			return nil, err
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
			return nil, xErrOnly(ec.wrongTypeArgument(ec.s.integerp, start))
		}

		if integerp(end) {
			endInt = int(xIntegerValue(end))

			if endInt < 0 {
				endInt += lengthInt
			}

		} else if end != ec.nil_ {
			return nil, xErrOnly(ec.wrongTypeArgument(ec.s.integerp, end))
		}

		if !(0 <= startInt && startInt <= endInt && endInt <= lengthInt) {
			return nil, xErrOnly(ec.signalN(ec.s.argsOutOfRange, stream, start, end))
		}

		str := xStringValue(stream)

		ctx = &readContextString{
			runes: []rune(str),
			i:     startInt,
			end:   endInt,
		}
	default:
		return nil, xErrOnly(ec.pimacsUnimplemented(ec.s.read, "unknown source type"))
	}

	return ctx, nil
}

func (ec *execContext) intern(str, _ lispObject) (lispObject, error) {
	if !stringp(str) {
		return ec.wrongTypeArgument(ec.s.stringp, str)
	}

	name := xStringValue(str)
	sym, existed := ec.internInternal(name)

	if !existed && strings.HasPrefix(name, ":") {
		_, err := ec.set(sym, sym)
		if err != nil {
			return nil, err
		}
		sym.special = true
	}

	return sym, nil
}

func (ec *execContext) unintern(name, _ lispObject) (lispObject, error) {
	nameStr := ""
	if symbolp(name) {
		nameStr = xSymbolName(name)
	} else if stringp(name) {
		nameStr = xStringValue(name)
	} else {
		return ec.wrongTypeArgument(ec.s.stringp, name)
	}

	return ec.bool(ec.uninternInternal(nameStr))
}

func (ec *execContext) lexicallyBoundp(ctx readContext) bool {
	c := ctx.read()

	if c == '#' {
		c = ctx.read()
		if c != '!' {
			ctx.unread(c)
			ctx.unread('#')
			return false
		}

		for c != '\n' && c != eofRune {
			c = ctx.read()
		}

		if c == '\n' {
			c = ctx.read()
		}
	}

	if c != ';' {
		ctx.unread(c)
		return false
	}

	builder := strings.Builder{}

	for {
		c = ctx.read()
		if c == '\n' || c == eofRune {
			break
		}

		builder.WriteRune(c)
	}

	line := strings.TrimRight(builder.String(), " \t")
	line, found := strings.CutSuffix(line, "-*-")
	if !found {
		return false
	}

	parts := strings.Split(line, "-*-")
	if len(parts) == 1 {
		return false
	}
	line = parts[1]

	strCtx := readContextString{runes: []rune(line), end: utf8.RuneCountInString(line)}

	for {
		variable, err := ec.read0(&strCtx)
		if err != nil || !symbolp(variable) || !strings.HasSuffix(xSymbolName(variable), ":") {
			return false
		}

		value, err := ec.read0(&strCtx)
		if err != nil {
			return false
		}

		// The reader will include the ":" in the symbol name, so use it here too
		if xSymbolName(variable) == "lexical-binding:" {
			return (value != ec.nil_)
		}

		for {
			c = strCtx.read()
			if c != ' ' && c != '\t' && c != ';' {
				strCtx.unread(c)
				break
			}
		}

		if c == eofRune || c == '\n' {
			break
		}
	}

	return false
}

func (ec *execContext) setupInternalInterpreterEnv() error {
	lex := ec.gl.lexicalBinding.val
	if lex == ec.nil_ || lex == ec.s.unbound {
		return ec.stackPushLet(ec.gl.internalInterpreterEnv.sym, ec.nil_)
	} else {
		return ec.stackPushLet(ec.gl.internalInterpreterEnv.sym, ec.makeList(ec.t))
	}
}

func (ec *execContext) readEvalLoop(ctx readContext) error {
	defer ec.unwind()()

	err := ec.setupInternalInterpreterEnv()
	if err != nil {
		return err
	}

	for {
		c := ctx.read()

		if c == ';' {
			for {
				c = ctx.read()
				if c == eofRune || c == '\n' {
					break
				}
			}
			continue
		}

		if c == eofRune {
			break
		}

		if c == ' ' || c == '\t' || c == '\n' || c == '\f' || c == '\r' || c == nbsp {
			continue
		}

		ctx.unread(c)

		val, err := ec.read0(ctx)
		if err != nil {
			return err
		}

		_, err = ec.evalSub(val)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ec *execContext) read(stream lispObject) (lispObject, error) {
	if stream == ec.nil_ {
		stream = ec.v.standardInput.val
	}

	if stream == ec.t {
		stream = ec.s.readChar
	}

	if stream == ec.s.readChar {
		return ec.funcall(ec.s.readFromMinibuffer, newString("Lisp expression: "))
	}

	ctx, err := ec.streamReadContext(stream, ec.nil_, ec.nil_)
	if err != nil {
		return nil, err
	}
	return ec.read0(ctx)
}

func (ec *execContext) readFromString(str, start, end lispObject) (lispObject, error) {
	if !stringp(str) {
		return ec.wrongTypeArgument(ec.s.stringp, str)
	}

	ctx, err := ec.streamReadContext(str, ec.nil_, ec.nil_)
	if err != nil {
		return nil, err
	}
	result, err := ec.read0(ctx)
	if err != nil {
		return nil, err
	}

	pos := newInteger(lispInt(ctx.position()))
	return ec.cons(result, pos)
}

func (ec *execContext) load(file, noError, noMessage, noSuffix, mustSuffix lispObject) (lispObject, error) {
	if !stringp(file) {
		return ec.wrongTypeArgument(ec.s.stringp, file)
	}

	loadPath := ec.v.loadPath.val
	iter := ec.iterate(loadPath)
	var f *os.File

	for ; iter.hasNext(); loadPath = iter.nextCons() {
		loadPath = xCar(loadPath)

		if !stringp(loadPath) {
			return ec.wrongTypeArgument(ec.s.stringp, loadPath)
		}

		base := xStringValue(loadPath)
		fullPath := filepath.Join(base, xStringValue(file))

		var err error
		f, err = os.Open(fullPath)
		if err == nil {
			defer func() {
				_ = f.Close()
			}()
			break
		}
	}

	if iter.hasError() {
		return iter.error()
	}

	if f == nil {
		if noError == ec.nil_ {
			return ec.signalN(ec.s.fileMissing, file)
		}
		return ec.nil_, nil
	}

	defer ec.unwind()()
	// Load is dynamic by default
	if err := ec.stackPushLet(ec.gl.lexicalBinding.sym, ec.nil_); err != nil {
		return nil, err
	}

	ctx := readContextFile{reader: bufio.NewReader(f)}

	if ec.lexicallyBoundp(&ctx) {
		if err := ec.stackPushLet(ec.gl.lexicalBinding.sym, ec.t); err != nil {
			return nil, err
		}
	}

	if err := ec.readEvalLoop(&ctx); err != nil {
		return nil, err
	}

	return ec.t, nil
}

func (ec *execContext) symbolsOfRead() {
	ec.defVarLisp(&ec.v.standardInput, "standard-input", ec.t)
	ec.defVarLisp(&ec.v.loadPath, "load-path", ec.nil_)
	ec.defSym(&ec.s.readChar, "read-char")
	ec.defSym(&ec.s.backquote, "`")
	ec.defSym(&ec.s.comma, ",")
	ec.defSym(&ec.s.commaAt, "@,")

	ec.defSubr2(nil, "intern", (*execContext).intern, 1)
	ec.defSubr2(nil, "unintern", (*execContext).unintern, 1)
	ec.defSubr3(nil, "read-from-string", (*execContext).readFromString, 1)
	ec.defSubr1(&ec.s.read, "read", (*execContext).read, 0)
	ec.defSubr5(nil, "load", (*execContext).load, 1)
}
