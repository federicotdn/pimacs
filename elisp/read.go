package elisp

import (
	"strconv"
	"strings"
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

func (ec *execContext) readList(ctx readContext) (lispObject, error) {
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

func (ec *execContext) readAtom(c rune, ctx readContext) (lispObject, error) {
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

func (ec *execContext) read1Result(obj lispObject, err error) (lispObject, rune, error) {
	return obj, 0, err
}

func (ec *execContext) read1(ctx readContext) (lispObject, rune, error) {
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
				return ec.read1Result(ec.intern(ec.emptyString(), ec.nil_))
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

func (ec *execContext) read0(ctx readContext) (lispObject, error) {
	obj, c, err := ec.read1(ctx)
	if err != nil {
		return nil, err
	}

	if c != 0 {
		return ec.invalidReadSyntax("unexpected character: '%v'", c)
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

func (ec *execContext) intern(str, obarray lispObject) (lispObject, error) {
	if !stringp(str) {
		return ec.wrongTypeArgument(ec.g.stringp, str)
	}

	name := xStringValue(str)

	sym, ok := ec.obarray[name]
	if !ok {
		sym = ec.makeSymbol(name)
		ec.internNewSymbol(sym)
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
	ec.defSubr2("intern", ec.intern, 1)
	ec.defSubr3("read-from-string", ec.readFromString, 1)
	ec.defSubr1("read", ec.read, 0)
	ec.defSubr5("load", ec.load, 1)

	ec.defVarStatic(ec.g.standardInput, ec.t)
	ec.defVarStatic(ec.g.lexicalBinding, ec.nil_)
}
