package main

import (
	"fmt"
	"os"
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
	return "(" + lc.car.printObj() + " . " + lc.cdr.printObj() + ")"
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

var lispNil = makeSymbol("nil")

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
	println("advance:", count)
	println("progress:", ctx.i, "/", len(ctx.runes))
}

func (ctx *readContext) advance() {
	ctx.advanceN(1)
}

func readList(ctx *readContext) (lispObject, error) {
	var val lispObject = lispNil
	var tail lispObject = lispNil

	for {
		elt, c, err := read(ctx)
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

func read(ctx *readContext) (lispObject, rune, error) {
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
		default:
			return readResult(readAtom(ctx))
		}
	}
}

func readFile(filename string) (lispObject, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	text := string(data)

	ctx := readContext{
		source: text,
		runes:  []rune(text),
		i:      0,
	}

	obj, c, err := read(&ctx)
	if err != nil {
		return nil, err
	}

	if c != 0 {
		return nil, fmt.Errorf("unexpected character: %v", string(c))
	}

	return obj, nil
}
