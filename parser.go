package main

import (
	"fmt"
	"strconv"
	"strings"
)

type lispType int

const (
	symbol lispType = iota
	integer
	str
	vector
	cons
	float
)

const nbsp = '\u00A0'

type lispSymbol struct {
	name string
}

type lispCons struct {
	car lispObject
	cdr lispObject
}

type lispInteger struct {
	value int64
}

type lispFloat struct {
	value float64
}

type lispObject interface {
	getType() lispType
	printObj() string
}

func (ls *lispSymbol) getType() lispType {
	return symbol
}

func (ls *lispSymbol) printObj() string {
	return ls.name
}

func (lc *lispCons) getType() lispType {
	return cons
}

func (lc *lispCons) printObj() string {
	result := "(" + lc.car.printObj()
	current := lc

	for {
		next, ok := current.cdr.(*lispCons)
		if ok {
			result += " " + next.car.printObj()
			current = next
		} else {
			if current.cdr != lispNil {
				result += " . " + current.cdr.printObj()
			}

			break
		}
	}

	return result + ")"
}

func (li *lispInteger) getType() lispType {
	return integer
}

func (li *lispInteger) printObj() string {
	return fmt.Sprint(li.value)
}

func (lf *lispFloat) getType() lispType {
	return float
}

func (lf *lispFloat) printObj() string {
	return fmt.Sprint(lf.value)
}

func makeSymbol(name string) *lispSymbol {
	return &lispSymbol{
		name: name,
	}
}

func makeInteger(value int64) *lispInteger {
	return &lispInteger{
		value: value,
	}
}

func makeCons(car lispObject, cdr lispObject) *lispCons {
	return &lispCons{
		car: car,
		cdr: cdr,
	}
}

func setCdr(c lispObject, cdr lispObject) {
	c.(*lispCons).cdr = cdr
}

func makeFloat(value float64) *lispFloat {
	return &lispFloat{
		value: value,
	}
}

func makeList(obj lispObject, objs ...lispObject) *lispCons {
	tmp := makeCons(obj, lispNil)
	val := tmp

	for _, obj := range objs {
		tmp.cdr = makeCons(obj, lispNil)
		tmp = tmp.cdr.(*lispCons)
	}

	return val
}

var lispNil = makeSymbol("nil")
var lispQuote = makeSymbol("quote")
var lispBackQuote = makeSymbol("`")

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

func readList(ctx *readContext) (lispObject, error) {
	var val lispObject = lispNil
	var tail lispObject = lispNil

	for {
		elt, c, err := read1(ctx)
		if err != nil {
			return nil, err
		}

		if c != 0 {
			switch c {
			case ')':
				return val, nil
			case '.':
				if tail != lispNil {
					obj, err := read0(ctx)
					if err != nil {
						return nil, err
					}
					setCdr(tail, obj)
				} else {
					val, err = read0(ctx)
				}

				_, c, _ = read1(ctx)
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

func readAtom(c rune, ctx *readContext) (lispObject, error) {
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

func readResult(obj lispObject, err error) (lispObject, rune, error) {
	return obj, 0, err
}

func read1(ctx *readContext) (lispObject, rune, error) {
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
			return readResult(readList(ctx))
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
		case c == '\'' || c == '`':
			obj, err := read0(ctx)
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
			return readResult(readAtom(c, ctx))
		default:
			return readResult(readAtom(c, ctx))
		}
	}
}

func read0(ctx *readContext) (lispObject, error) {
	obj, c, err := read1(ctx)
	if err != nil {
		return nil, err
	}

	if c != 0 {
		return nil, fmt.Errorf("unexpected character: '%v'", string(c))
	}

	return obj, nil
}

func readString(source string) (lispObject, error) {
	ctx := readContext{
		source: source,
		runes:  []rune(source),
		i:      0,
	}

	return read0(&ctx)
}
