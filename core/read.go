package core

import (
	"bufio"
	"fmt"
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
	line        int
	path        string
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

func newReadContextFile(f *os.File) readContextFile {
	return readContextFile{
		reader: bufio.NewReader(f),
		path:   f.Name(),
		line:   1,
	}
}

func (ctx *readContextFile) read() rune {
	var r rune
	if len(ctx.unreadRunes) > 0 {
		r = ctx.unreadRunes[len(ctx.unreadRunes)-1]
		ctx.unreadRunes = ctx.unreadRunes[:len(ctx.unreadRunes)-1]
		ctx.i++
	} else {
		var err error
		r, _, err = ctx.reader.ReadRune()
		if err != nil {
			r = eofRune
		} else {
			ctx.i++
		}
	}

	if r == '\n' {
		ctx.line++
	}

	return r
}

func (ctx *readContextFile) unread(c rune) {
	if c == eofRune {
		return
	}

	if ctx.i == 0 {
		terminate("unbalanced read/unread call with rune: '%v'", c)
	}

	if c == '\n' {
		ctx.line--
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

func (ec *execContext) stringToNumber(s string, radix int) (lispObject, error) {
	nInt, err := strconv.ParseInt(s, radix, 64)
	if err == nil {
		return newInteger(lispInt(nInt)), nil
	}

	nFloat, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return newFloat(lispFp(nFloat)), nil
	}

	return ec.pimacsUnimplemented(ec.s.read, "unknown number format: '%v'", s)
}

func (ec *execContext) readEscape(ctx readContext, nextChar rune) (rune, error) {
	var chr, modifiers rune
	nControl := 0
again:
	c := nextChar

	switch c {
	case eofRune:
		return 0, xErrOnly(ec.signalN(ec.s.endOfFile))
	case 'a':
		chr = '\a'
	case 'b':
		chr = '\b'
	case 'd':
		chr = 0177
	case 'e':
		chr = 033
	case 'f':
		chr = '\f'
	case 'n':
		chr = '\n'
	case 'r':
		chr = '\r'
	case 't':
		chr = '\t'
	case 'v':
		chr = '\v'
	case '\n':
		return 0, xErrOnly(ec.signalError("Invalid escape char syntax: \\<newline>"))
	case 'M', 'S', 'H', 'A', 's':
		var mod rune

		switch c {
		case 'M':
			mod = charMeta
		case 'S':
			mod = charShift
		case 'H':
			mod = charHyper
		case 'A':
			mod = charAlt
		case 's':
			mod = charSuper
		}

		c1 := ctx.read()
		if c1 != '-' {
			if c == 's' {
				ctx.unread(c1)
				chr = ' '
				break
			} else {
				return 0, xErrOnly(ec.signalError("Invalid escape char syntax: \\%c not followed by -", c))
			}
		}

		modifiers |= mod
		c1 = ctx.read()
		if c1 == '\\' {
			nextChar = ctx.read()
			goto again
		}
		chr = c1
	case 'C':
		c1 := ctx.read()
		if c1 != '-' {
			return 0, xErrOnly(ec.signalError("Invalid escape char syntax: \\%c not followed by -", c))
		}

		fallthrough
	case '^':
		nControl++
		c1 := ctx.read()
		if c1 == '\\' {
			nextChar = ctx.read()
			goto again
		}
		chr = c1
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

		chr = num
	case 'x':
		count := 0
		var i lispInt
		for {
			c = ctx.read()
			digit := lispInt(ec.digitToNumber(c, 16))
			if digit < 0 {
				ctx.unread(c)
				break
			}

			i = (i << 4) + digit

			if i > runeToLispInt(charMeta|(charMeta-1)) {
				return 0, xErrOnly(ec.signalError("Hex character out of range: \\x%x", i))
			}

			count++
		}

		if count == 0 {
			return 0, xErrOnly(ec.signalError("Invalid escape char syntax: \\x not followed by hex digit"))
		}

		if count < 3 && i >= 0x80 {
			i = runeToLispInt(byte8toChar(lispIntToRune(i)))
		}

		modifiers |= lispIntToRune(i) & charModMask
		chr = lispIntToRune(i) & ^charModMask
	case 'U', 'u', 'N':
		return 0, xErrOnly(ec.pimacsUnimplemented(ec.s.read, "unimplemented escape sequence: \\%c", c))
	default:
		chr = c
	}

	for nControl > 0 {
		if (chr >= '@' && chr <= '_') || (chr >= 'a' && chr <= 'z') {
			chr &= 0x1f
		} else if chr == '?' {
			chr = 127
		} else {
			modifiers |= charCtrl
		}
		nControl--
	}

	return chr | modifiers, nil
}

func (ec *execContext) digitToNumber(c rune, radix int) int {
	digit := 0

	if '0' <= c && c <= '9' {
		digit = int(c - '0')
	} else if 'a' <= c && c <= 'z' {
		digit = int(c - 'a' + 10)
	} else if 'A' <= c && c <= 'Z' {
		digit = int(c - 'A' + 10)
	} else {
		return -2
	}

	if digit < radix {
		return digit
	}
	return -1
}

func (ec *execContext) readInteger(ctx readContext, radix int) (lispObject, error) {
	valid := -1 // 1 -> valid, 0 -> invalid, -1 -> incomplete
	p := []rune{}

	c := ctx.read()
	if c == '-' || c == '+' {
		p = append(p, c)
		c = ctx.read()
	}

	if c == '0' {
		p = append(p, c)
		valid = 1

		for ok := true; ok; ok = (c == '0') {
			c = ctx.read()
		}
	}

	digit := ec.digitToNumber(c, radix)
	for {
		if digit < -1 {
			break
		} else if digit == -1 {
			valid = 0
		}
		if valid < 0 {
			valid = 1
		}

		p = append(p, c)
		c = ctx.read()
		digit = ec.digitToNumber(c, radix)
	}

	ctx.unread(c)
	if valid != 1 {
		return ec.invalidReadSyntax("invalid radix %v integer: %v", radix, string(p))
	}

	return ec.stringToNumber(string(p), radix)
}

func (ec *execContext) readStringLiteral(ctx readContext) (lispObject, error) {
	builder := strings.Builder{}
	multibyte := false
	c := ctx.read()

	for ; c != eofRune && c != '"'; c = ctx.read() {
		if c == '\\' {
			c = ctx.read()
			switch c {
			case 's':
				c = ' '
			case ' ':
				fallthrough
			case '\n':
				continue
			default:
				var err error
				c, err = ec.readEscape(ctx, c)
				if err != nil {
					return nil, err
				}
			}
		}

		multibyte = multibyte || !asciiCharp(c)
		builder.WriteRune(c)
	}

	if c == eofRune {
		return ec.signalN(ec.s.endOfFile)
	}

	return newString(builder.String(), multibyte), nil
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
		c, err = ec.readEscape(ctx, ctx.read())
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
	multibyte := false

	for {
		if c == '\\' {
			c = ctx.read()
			if c == eofRune {
				return ec.signalN(ec.s.endOfFile)
			}

			quoted = true
		}

		multibyte = multibyte || !asciiCharp(c)
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
		num, err := ec.stringToNumber(s, 10)
		if err == nil {
			return num, nil
		}
	}

	return ec.intern(newString(s, multibyte), ec.nil_)
}

func (ec *execContext) read0(ctx readContext) (lispObject, error) {
	stack := readStack{}
	var obj lispObject
	var err error
	var c rune

read_obj:
	c = ctx.read()
	switch {
	case c == eofRune:
		return ec.signalN(ec.s.endOfFile)
	case c == '(':
		stack.push(readStackElem{elementType: readStackListStart})
		goto read_obj
	case c == '[':
		stack.push(readStackElem{elementType: readStackVector})
		goto read_obj
	case c == ')':
		if stack.isEmpty() {
			return ec.invalidReadSyntax(")")
		}

		switch stack.peek().elementType {
		case readStackListStart:
			stack.pop()
			obj = ec.nil_
		case readStackList:
			obj = stack.pop().listHead
		default:
			return ec.invalidReadSyntax(")")
		}
	case c == ']':
		if stack.isEmpty() {
			return ec.invalidReadSyntax("]")
		}
		switch stack.peek().elementType {
		case readStackVector:
			obj = newVector(stack.pop().elems)
		default:
			return ec.invalidReadSyntax("]")
		}
	case c == '#':
		c = ctx.read()

		switch c {
		case '#':
			var err error
			obj, err = ec.intern(newString("", false), ec.nil_)
			if err != nil {
				return nil, err
			}
		case '\'':
			stack.push(readStackElem{
				elementType: readStackSpecial,
				special:     ec.s.function,
			})
			goto read_obj
		case 'o', 'O':
			var err error
			obj, err = ec.readInteger(ctx, 8)
			if err != nil {
				return nil, err
			}
		case 'x', 'X':
			var err error
			obj, err = ec.readInteger(ctx, 16)
			if err != nil {
				return nil, err
			}
		case 'b', 'B':
			var err error
			obj, err = ec.readInteger(ctx, 2)
			if err != nil {
				return nil, err
			}
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

			return ec.invalidReadSyntax(".")
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
				return ec.invalidReadSyntax("expected )")
			}

			stack.pop()

			if top.listHead == nil {
				// See Emacs bug #62020
				return ec.invalidReadSyntax(".")
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
			return ec.signalN(ec.s.read, newString("unknown read stack entry type", false))
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

	if ec.gl.obarray.containsSymbol(xStringValue(str)) {
		return ec.internInternal(str, &ec.gl.obarray)
	}

	return ec.internInternal(str, &ec.obarray)
}

func (ec *execContext) xIntern(str string) lispObject {
	return xEnsure(ec.intern(newString(str, true), ec.nil_))
}

func (ec *execContext) internInternal(str lispObject, obarray *obarrayType) (lispObject, error) {
	sym := ec.makeSymbol(xStringValue(str), true)
	sym, existed := obarray.loadOrStoreSymbol(sym)

	if !existed && strings.HasPrefix(sym.name, ":") {
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

	return ec.bool(ec.obarray.removeSymbol(nameStr))
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
		return ec.funcall(ec.s.readFromMinibuffer, newString("Lisp expression: ", false))
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

	ctx := newReadContextFile(f)

	if ec.lexicallyBoundp(&ctx) {
		if err := ec.stackPushLet(ec.gl.lexicalBinding.sym, ec.t); err != nil {
			return nil, err
		}
	}

	if err := ec.readEvalLoop(&ctx); err != nil {
		if signal, ok := err.(*stackJumpSignal); ok {
			signal.location = append(signal.location, fmt.Sprintf("%v:%v", ctx.path, ctx.line))
		}

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
	ec.defSym(&ec.s.commaAt, ",@")
	ec.defSym(&ec.s.variableDocumentation, "variable-documentation")

	ec.defSubr2(nil, "intern", (*execContext).intern, 1)
	ec.defSubr2(nil, "unintern", (*execContext).unintern, 1)
	ec.defSubr3(nil, "read-from-string", (*execContext).readFromString, 1)
	ec.defSubr1(&ec.s.read, "read", (*execContext).read, 0)
	ec.defSubr5(nil, "load", (*execContext).load, 1)
}
