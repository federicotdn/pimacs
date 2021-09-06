package elisp

import (
	"fmt"
)

type lispType int
type lispInt int64

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
	value lispInt
}

type lispFloat struct {
	value float64
}

type lispString struct {
	value string
}

type lispObject interface {
	getType() lispType
	PrintObj() string
}

func (ls *lispSymbol) getType() lispType {
	return symbol
}

func (ls *lispSymbol) PrintObj() string {
	return ls.name
}

func (lc *lispCons) getType() lispType {
	return cons
}

func (lc *lispCons) PrintObj() string {
	result := "(" + lc.car.PrintObj()
	current := lc

	for {
		next, ok := current.cdr.(*lispCons)
		if ok {
			result += " " + next.car.PrintObj()
			current = next
		} else {
			if current.cdr != lispNil {
				result += " . " + current.cdr.PrintObj()
			}

			break
		}
	}

	return result + ")"
}

func (li *lispInteger) getType() lispType {
	return integer
}

func (li *lispInteger) PrintObj() string {
	return fmt.Sprint(li.value)
}

func (lf *lispFloat) getType() lispType {
	return float
}

func (lf *lispFloat) PrintObj() string {
	return fmt.Sprint(lf.value)
}

func (ls *lispString) getType() lispType {
	return str
}

func (ls *lispString) PrintObj() string {
	return "\"" + ls.value + "\""
}

func makeSymbol(name string) *lispSymbol {
	return &lispSymbol{
		name: name,
	}
}

func makeInteger(value lispInt) *lispInteger {
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

func makeString(value string) *lispString {
	return &lispString{
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
