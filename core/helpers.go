package core

import (
	"errors"
	"reflect"
)

// General helpers //

func xEnsure(obj lispObject, err error) lispObject {
	if err != nil {
		terminate("cannot return value due to error: '%+v'", err)
	}

	return obj
}

func xErrOnly(obj lispObject, err error) error {
	if obj != nil {
		terminate("was handed a non-nil Lisp object: '%+v'", obj)
	}

	return err
}

func xCast[T lispObject](obj lispObject, typeName string) T {
	val, ok := obj.(T)
	if !ok {
		terminate("object is not a %v: '%+v'", typeName, obj)
	}

	return val
}

// Symbol helpers //

func symbolp(obj lispObject) bool {
	return obj.getType() == lispTypeSymbol
}

func xSymbol(obj lispObject) *lispSymbol {
	return xCast[*lispSymbol](obj, "symbol")
}

func xSymbolName(obj lispObject) string {
	return xStringValue(xSymbol(obj).name)
}

// Cons helpers //

func consp(obj lispObject) bool {
	return obj.getType() == lispTypeCons
}

func xCons(obj lispObject) *lispCons {
	return xCast[*lispCons](obj, "cons")
}

func xCarCdr(obj lispObject) (lispObject, lispObject) {
	cons := xCons(obj)
	return cons.car, cons.cdr
}

func xCar(obj lispObject) lispObject {
	return xCons(obj).car
}

func xCdr(obj lispObject) lispObject {
	return xCons(obj).cdr
}

func xCadr(obj lispObject) lispObject {
	return xCar(xCdr(obj))
}

func xCddr(obj lispObject) lispObject {
	return xCdr(xCdr(obj))
}

func xCdddr(obj lispObject) lispObject {
	return xCdr(xCdr(xCdr(obj)))
}

func xSetCar(cell, newCar lispObject) lispObject {
	xCons(cell).car = newCar
	return newCar
}

func xSetCdr(cell, newCdr lispObject) lispObject {
	xCons(cell).cdr = newCdr
	return newCdr
}

func newCons(car lispObject, cdr lispObject) *lispCons {
	return &lispCons{
		car: car,
		cdr: cdr,
	}
}

// Subroutine helpers //

func subroutinep(obj lispObject) bool {
	return obj.getType() == lispTypeSubroutine
}

func xSubroutine(obj lispObject) *lispSubroutine {
	return xCast[*lispSubroutine](obj, "subroutine")
}

// Vector helpers //

func vectorp(obj lispObject) bool {
	return obj.getType() == lispTypeVector
}

func xVector(obj lispObject) *lispVector {
	return xCast[*lispVector](obj, "vector")
}

func xVectorValue(obj lispObject) []lispObject {
	return xVector(obj).val
}

func newVector(val []lispObject) *lispVector {
	return &lispVector{val: val}
}

// Buffer helpers //

func bufferp(obj lispObject) bool {
	return obj.getType() == lispTypeBuffer
}

func xBuffer(obj lispObject) *lispBuffer {
	return xCast[*lispBuffer](obj, "buffer")
}

func newBuffer(name lispObject) *lispBuffer {
	return &lispBuffer{name: name, live: true}
}

// Char Table helpers //

func chartablep(obj lispObject) bool {
	return obj.getType() == lispTypeCharTable
}

func xCharTable(obj lispObject) *lispCharTable {
	return xCast[*lispCharTable](obj, "character table")
}

// String helpers //

func stringp(obj lispObject) bool {
	return obj.getType() == lispTypeString
}

func xString(obj lispObject) *lispString {
	return xCast[*lispString](obj, "string")
}

func xStringMultibytep(obj lispObject) bool {
	return xString(obj).multibyte
}

func xStringValue(obj lispObject) string {
	return xString(obj).str()
}

func xStringEmpty(obj lispObject) bool {
	return xStringValue(obj) == ""
}

func xStringSize(obj lispObject) int {
	return xString(obj).size()
}

func xStringAref(obj lispObject, index int) (lispInt, error) {
	err := errors.New("index out of range")
	if index < 0 {
		return 0, err
	}

	val := xStringValue(obj)
	for i, r := range val {
		if i == index {
			return runeToLispInt(r), nil
		}
	}
	return 0, err
}

// Integer helpers //

func integerp(obj lispObject) bool {
	return obj.getType() == lispTypeInteger
}

func xInteger(obj lispObject) *lispInteger {
	return xCast[*lispInteger](obj, "integer")
}

func xIntegerValue(obj lispObject) lispInt {
	return xInteger(obj).val
}

func xIntegerRune(obj lispObject) rune {
	return lispIntToRune(xIntegerValue(obj))
}

func newInteger(val lispInt) *lispInteger {
	return &lispInteger{
		val: val,
	}
}

// Float helpers //

func floatp(obj lispObject) bool {
	return obj.getType() == lispTypeFloat
}

func xFloat(obj lispObject) *lispFloat {
	return xCast[*lispFloat](obj, "float")
}

func xFloatValue(obj lispObject) lispFp {
	return xFloat(obj).val
}

func newFloat(val lispFp) *lispFloat {
	return &lispFloat{
		val: val,
	}
}

// Channel helpers //

func channelp(obj lispObject) bool {
	return obj.getType() == lispTypeChannel
}

func xChannel(obj lispObject) *lispChannel {
	return xCast[*lispChannel](obj, "channel")
}

// Marker helpers //

func markerp(obj lispObject) bool {
	return obj.getType() == lispTypeMarker
}

// Hash table helpers //

func hashtablep(obj lispObject) bool {
	return obj.getType() == lispTypeHashTable
}

func xHashTable(obj lispObject) *lispHashTable {
	return xCast[*lispHashTable](obj, "hash table")
}

// Cross-type helpers //

func numberp(obj lispObject) bool {
	return integerp(obj) || floatp(obj)
}

func characterp(obj lispObject) bool {
	if !integerp(obj) {
		return false
	}
	return charValidp(xIntegerRune(obj))
}

func arrayp(obj lispObject) bool {
	// TODO: Add more types
	return vectorp(obj) || stringp(obj)
}

func naturalp(obj lispObject) bool {
	return integerp(obj) && xIntegerValue(obj) >= 0
}

func numberOrMarkerp(obj lispObject) bool {
	return numberp(obj) || markerp(obj)
}

func integerOrMarkerp(obj lispObject) bool {
	return integerp(obj) || markerp(obj)
}

// Misc. Lisp Object utilities //

func objAddr(obj lispObject) lispInt {
	u := reflect.ValueOf(obj).Pointer()
	return lispInt(u)
}
