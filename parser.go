package main

import (
	"fmt"
	"regexp"
	"strconv"
	"unicode/utf8"
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
	value int
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

func makeSymbol(name string) *lispSymbol {
	return &lispSymbol{
		name: name,
	}
}

func makeInteger(value int) *lispInteger {
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

func makeList(objs ...lispObject) *lispCons {
	if len(objs) == 0 {
		return nil
	}

	tmp := makeCons(objs[0], lispNil)
	val := tmp

	for _, obj := range objs[1:] {
		tmp.cdr = makeCons(obj, lispNil)
		tmp = tmp.cdr.(*lispCons)
	}

	return val
}

var lispNil = makeSymbol("nil")
var lispQuote = makeSymbol("quote")

type readContext struct {
	source string
	runes  []rune
	i      int
}

func (ctx *readContext) currentRune() (rune, error) {
	if ctx.i >= len(ctx.runes) {
		return 0, fmt.Errorf("eof")
	}

	return ctx.runes[ctx.i], nil
}

func (ctx *readContext) currentString() (string, error) {
	if ctx.i >= len(ctx.runes) {
		return "", fmt.Errorf("eof")
	}

	return string(ctx.runes[ctx.i:]), nil
}

func (ctx *readContext) advanceN(count int) {
	ctx.i += count
}

func (ctx *readContext) advance() {
	ctx.advanceN(1)
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
			default:
				return nil, fmt.Errorf("invalid list ending: %v", c)
			}
		}

		tmp := makeCons(elt, lispNil)

		if tail != lispNil {
			tail.(*lispCons).cdr = tmp
		} else {
			val = tmp
		}

		tail = tmp
	}
}

func readAtom(ctx *readContext) (lispObject, error) {
	val, err := ctx.currentString()
	if err != nil {
		return nil, err
	}

	integerRegexp := regexp.MustCompile(`[\d]+`)
	loc := integerRegexp.FindStringIndex(val)
	if loc != nil {
		match := val[loc[0]:loc[1]]
		value, err := strconv.Atoi(match)
		if err != nil {
			return nil, err
		}

		ctx.advanceN(utf8.RuneCountInString(match))
		return makeInteger(value), nil
	}

	return nil, fmt.Errorf("unknown atom format")
}

func readResult(obj lispObject, err error) (lispObject, rune, error) {
	return obj, 0, err
}

func read1(ctx *readContext) (lispObject, rune, error) {
	var err error
	var c rune

	for {
		c, err = ctx.currentRune()
		if err != nil {
			return nil, 0, err
		}

		switch {
		case c == '(':
			// List
			ctx.advance()
			return readResult(readList(ctx))
		case c == ')':
			ctx.advance()
			return nil, ')', nil
		case c <= 040 || c == nbsp:
			ctx.advance()
		case c == ';':
			for {
				c, err = ctx.currentRune()
				if err != nil {
					return nil, 0, err
				}
				ctx.advance()

				if c == '\n' {
					break
				}
			}
		case c == '\'':
			ctx.advance()
			obj, err := read0(ctx)
			if err != nil {
				return nil, 0, err
			}

			list := makeList(lispQuote, obj)
			return list, 0, nil
		default:
			return readResult(readAtom(ctx))
		}
	}
}

func read0(ctx *readContext) (lispObject, error) {
	obj, c, err := read1(ctx)
	if err != nil {
		return nil, err
	}

	if c != 0 {
		return nil, fmt.Errorf("unexpected character: %v", string(c))
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
