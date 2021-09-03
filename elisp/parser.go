package elisp

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	eightBitCodeOffset rune = 0x3fff00
	max5ByteChar       rune = 0x3fff7f
	eofRune            rune = -1
)

type Parser struct {
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

func stringToNumber(s string) (lispObject, error) {
	nInt, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return makeInteger(lispInt(nInt)), nil
	}

	nFloat, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return makeFloat(nFloat), nil
	}

	return nil, fmt.Errorf("unknown number format")
}

func byte8toChar(c rune) rune {
	return c + eightBitCodeOffset
}

func charByte8(c rune) bool {
	return c > max5ByteChar
}

func charToByte8(c rune) rune {
	if charByte8(c) {
		return c - eightBitCodeOffset
	}
	return c & 0xff
}

func readEscape(ctx *readContext, stringp bool) (rune, error) {
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
		return 0, fmt.Errorf("eof on read escape")
	case ' ':
		if stringp {
			return 0, fmt.Errorf("eof on read escape")
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

		for count := 0; count < 3; count++ {
			c := ctx.read()

			if c >= '0' && c <= '7' {
				num *= 8
				num += c - '0'
			} else {
				ctx.unread(c)
				break
			}

			if num >= 0x80 && num < 0x100 {
				num = byte8toChar(num)
			}

			return num, nil
		}
	case 'x':
	case 'U':
	case 'u':
	case 'N':
	default:
		return c, nil
	}

	return 0, fmt.Errorf("unimplemented escape code: '%v", string(c))
}

func (p *Parser) readList(ctx *readContext) (lispObject, error) {
	var val lispObject = lispNil
	var tail lispObject = lispNil

	for {
		elt, c, err := p.read1(ctx)
		if err != nil {
			return nil, err
		}

		if c != 0 {
			switch c {
			case ')':
				return val, nil
			case '.':
				if tail != lispNil {
					obj, err := p.read0(ctx)
					if err != nil {
						return nil, err
					}
					setCdr(tail, obj)
				} else {
					val, err = p.read0(ctx)
				}

				_, c, _ = p.read1(ctx)
				if c == ')' {
					return val, nil
				}

				return nil, fmt.Errorf("'.' in wrong context")
			default:
				return nil, fmt.Errorf("invalid list ending: '%v'", string(c))
			}
		}

		tmp := makeCons(elt, lispNil)

		if tail != lispNil {
			setCdr(tail, tmp)
		} else {
			val = tmp
		}

		tail = tmp
	}
}

func (p *Parser) readAtom(c rune, ctx *readContext) (lispObject, error) {
	quoted := false
	buf := []rune{}

	for {
		if c == '\\' {
			c := ctx.read()
			if c == eofRune {
				return nil, fmt.Errorf("eof reading atom")
			}

			quoted = true
		}

		buf = append(buf, c)
		c = ctx.read()

		special := strings.Contains("\"';()[]#`,", string(c))
		if c <= 040 || c == nbsp || special {
			break
		}
	}

	ctx.unread(c)
	s := string(buf)

	if !quoted {
		num, err := stringToNumber(s)
		if err == nil {
			return num, nil
		}
	}

	// interned
	if s == "nil" {
		return lispNil, nil
	}

	// uninterned
	return makeSymbol(s), nil
}

func (p *Parser) read1Result(obj lispObject, err error) (lispObject, rune, error) {
	return obj, 0, err
}

func (p *Parser) read1(ctx *readContext) (lispObject, rune, error) {
	var err error
	var c rune

	for {
		c = ctx.read()
		if c == eofRune {
			return nil, 0, fmt.Errorf("eof on read1")
		}

		switch {
		case c == '(':
			return p.read1Result(p.readList(ctx))
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
			obj, err := p.read0(ctx)
			if err != nil {
				return nil, 0, err
			}

			q := lispQuote
			if c == '`' {
				q = lispBackQuote
			}

			list := makeList(q, obj)
			return list, 0, nil
		case c == ',':
			return nil, 0, fmt.Errorf("unimplented read: '%v'", string(c))
		case c == '?':
			c = ctx.read()
			if c == eofRune {
				return nil, 0, fmt.Errorf("eof reading character")
			}

			if c == ' ' || c == '\t' {
				return makeInteger(lispInt(c)), 0, nil
			}

			if c == '\\' {
				c, err = readEscape(ctx, false)
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
				return makeInteger(lispInt(c)), 0, nil
			}

			return nil, 0, fmt.Errorf("invalid syntax for '?'")
		case c == '"':
			return nil, 0, fmt.Errorf("unimplented read: '%v'", string(c))
		case c == '.':
			next := ctx.read()
			ctx.unread(next)

			special := strings.Contains("\"';([#?`,", string(next))
			if next <= 040 || special {
				return nil, c, nil
			}
			return p.read1Result(p.readAtom(c, ctx))
		case c <= 040 || c == nbsp:
		default:
			return p.read1Result(p.readAtom(c, ctx))
		}
	}
}

func (p *Parser) read0(ctx *readContext) (lispObject, error) {
	obj, c, err := p.read1(ctx)
	if err != nil {
		return nil, err
	}

	if c != 0 {
		return nil, fmt.Errorf("unexpected character: '%v'", string(c))
	}

	return obj, nil
}

func (p *Parser) ReadString(source string) (lispObject, error) {
	ctx := readContext{
		source: source,
		runes:  []rune(source),
		i:      0,
	}

	return p.read0(&ctx)
}
