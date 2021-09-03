package elisp

import (
	"fmt"
	"strconv"
	"strings"
)

type Parser struct {
	
}

type readContext struct {
	source string
	runes  []rune
	i      int
}

func (ctx *readContext) read() (rune, error) {
	if ctx.i == len(ctx.runes) {
		return 0, fmt.Errorf("eof")
	}

	r := ctx.runes[ctx.i]
	ctx.i++
	return r, nil
}

func (ctx *readContext) unread() {
	if ctx.i == 0 {
		panic("nothing to unread")
	}

	ctx.i--
}

func stringToNumber(s string) (lispObject, error) {
	nInt, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return makeInteger(nInt), nil
	}

	nFloat, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return makeFloat(nFloat), nil
	}

	return nil, fmt.Errorf("unknown number format")
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
	var err error
	quoted := false
	buf := []rune{}

	for {
		if c == '\\' {
			c, err = ctx.read()
			if err != nil {
				return nil, err
			}

			quoted = true
		}

		buf = append(buf, c)

		c, err = ctx.read()
		if err != nil {
			return nil, err
		}

		special := strings.Contains("\"';()[]#`,", string(c))
		if c <= 040 || c == nbsp || special {
			break
		}
	}

	ctx.unread()
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
		c, err = ctx.read()
		if err != nil {
			return nil, 0, err
		}

		switch {
		case c == '(':
			// List
			return p.read1Result(p.readList(ctx))
		case c == ')':
			return nil, ')', nil
		case c <= 040 || c == nbsp:
		case c == ';':
			for {
				c, err = ctx.read()
				if err != nil {
					return nil, 0, err
				}

				if c == '\n' {
					break
				}
			}
		case c == '?':
			c, err = ctx.read()
			if err != nil {
				return nil, 0, err
			}

			if c == ' ' || c == '\t' {
				return makeInteger(int64(c)), 0, nil
			}

			// TODO: update lread--unescaped-character-literals

			if c == '\\' {
				
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
		case c == '.':
			next, err := ctx.read()
			if err != nil {
				return nil, 0, err
			}

			ctx.unread()

			special := strings.Contains("\"';([#?`,", string(next))
			if next <= 040 || special {
				return nil, c, nil
			}
			return p.read1Result(p.readAtom(c, ctx))
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
