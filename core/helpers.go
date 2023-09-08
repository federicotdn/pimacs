package core

import (
	"fmt"
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
	return xSymbol(obj).name
}

// Cons helpers //

func consp(obj lispObject) bool {
	return obj.getType() == lispTypeCons
}

func xCons(obj lispObject) *lispCons {
	return xCast[*lispCons](obj, "cons")
}

func xCar(obj lispObject) lispObject {
	return xCons(obj).car
}

func xCdr(obj lispObject) lispObject {
	return xCons(obj).cdr
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

func xStringValue(obj lispObject) string {
	return xString(obj).val
}

func xStringChars(obj lispObject) int {
	return len(xStringValue(obj))
}

func newString(val string) *lispString {
	return &lispString{
		val: val,
	}
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

// Debug utilities //

func debugRepr(obj lispObject) string {
	return debugReprInternal(obj, 20)
}

func debugReprInternal(obj lispObject, depth int) string {
	if depth <= 0 {
		return "(...)"
	}
	depth--

	switch obj.getType() {
	case lispTypeCons:
		return fmt.Sprintf("(%v . %v)", debugReprInternal(xCar(obj), depth), debugReprInternal(xCdr(obj), depth))
	case lispTypeFloat:
		return fmt.Sprintf("%vf", xFloatValue(obj))
	case lispTypeString:
		return fmt.Sprintf("\"%v\"", xStringValue(obj))
	case lispTypeInteger:
		return fmt.Sprintf("int(%v)", xIntegerValue(obj))
	case lispTypeSymbol:
		sym := xSymbol(obj)
		if sym.val == sym || sym.function == sym {
			return sym.name
		}

		if sym.redirect == nil {
			return fmt.Sprintf("sym(%v, v=%v, f=%v)", sym.name, debugReprInternal(sym.val, depth), debugReprInternal(sym.function, depth))
		} else {
			return fmt.Sprintf("sym(%v, v=FWD, f=%v)", sym.name, debugReprInternal(sym.function, depth))
		}
	case lispTypeVector:
		s := "["
		for _, elem := range xVector(obj).val {
			s += debugReprInternal(elem, depth) + ", "
		}
		return s + "]"
	case lispTypeBuffer:
		buf := xBuffer(obj)
		return fmt.Sprintf("buf(name=%v, live=%v)", buf.name, buf.live)
	case lispTypeSubroutine:
		subr := xSubroutine(obj)
		return fmt.Sprintf("subr(min=%v, max=%v)", subr.minArgs, subr.maxArgs)
	default:
		panic(fmt.Sprintf("unknown object type: %v", obj.getType()))
	}
}
