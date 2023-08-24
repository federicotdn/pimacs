package core

import "fmt"

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

func xSymbol(obj lispObject) *lispSymbol {
	val, ok := obj.(*lispSymbol)
	if !ok {
		terminate("object is not a symbol: '%+v'", obj)
	}

	return val
}

func symbolp(obj lispObject) bool {
	return obj.getType() == lispTypeSymbol
}

func xCons(obj lispObject) *lispCons {
	val, ok := obj.(*lispCons)
	if !ok {
		terminate("object is not a cons: '%+v'", obj)
	}

	return val
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

func consp(obj lispObject) bool {
	return obj.getType() == lispTypeCons
}

func xVectorLike(obj lispObject) *lispVectorLike {
	val, ok := obj.(*lispVectorLike)
	if !ok {
		terminate("object is not vector-like: '%+v'", obj)
	}

	return val
}

func vectorLikep(obj lispObject, vecType vectorLikeType) bool {
	vec, ok := obj.(*lispVectorLike)
	if !ok {
		return false
	}

	return vec.vecType == vecType
}

func xSubroutine(obj lispObject) *subroutine {
	vec := xVectorLike(obj)
	if vec.vecType != vectorLikeTypeSubroutine {
		terminate("object is not a subroutine: '%+v'", obj)
	}

	return vec.value.(*subroutine)
}

func subroutinep(obj lispObject) bool {
	return vectorLikep(obj, vectorLikeTypeSubroutine)
}

func xBuffer(obj lispObject) *buffer {
	vec := xVectorLike(obj)
	if vec.vecType != vectorLikeTypeBuffer {
		terminate("object is not a buffer: '%+v'", obj)
	}

	return vec.value.(*buffer)
}

func bufferp(obj lispObject) bool {
	return vectorLikep(obj, vectorLikeTypeBuffer)
}

func xString(obj lispObject) *lispString {
	val, ok := obj.(*lispString)
	if !ok {
		terminate("object is not a string: '%+v'", obj)
	}

	return val
}

func xStringValue(obj lispObject) string {
	return xString(obj).value
}

func xStringChars(obj lispObject) int {
	return len(xStringValue(obj))
}

func stringp(obj lispObject) bool {
	return obj.getType() == lispTypeString
}

func xInteger(obj lispObject) *lispInteger {
	val, ok := obj.(*lispInteger)
	if !ok {
		terminate("object is not an integer: '%+v'", obj)
	}

	return val
}

func xIntegerValue(obj lispObject) lispInt {
	return xInteger(obj).value
}

func xIntegerRune(obj lispObject) rune {
	return rune(xIntegerValue(obj))
}

func integerp(obj lispObject) bool {
	return obj.getType() == lispTypeInteger
}

func xFloat(obj lispObject) *lispFloat {
	val, ok := obj.(*lispFloat)
	if !ok {
		terminate("object is not a float: '%+v'", obj)
	}

	return val
}

func xFloatValue(obj lispObject) lispFp {
	return xFloat(obj).value
}

func floatp(obj lispObject) bool {
	return obj.getType() == lispTypeFloat
}

func numberp(obj lispObject) bool {
	return integerp(obj) || floatp(obj)
}

func characterp(obj lispObject) bool {
	return integerp(obj) && xIntegerValue(obj) <= maxChar
}

func arrayp(obj lispObject) bool {
	return vectorLikep(obj, vectorLikeTypeNormal) || stringp(obj)
}

func debugRepr(obj lispObject) string {
	switch obj.getType() {
	case lispTypeCons:
		return fmt.Sprintf("CONS(%v, %v)", debugRepr(xCar(obj)), debugRepr(xCdr(obj)))
	case lispTypeFloat:
		return fmt.Sprintf("FLOAT(%v)", xFloatValue(obj))
	case lispTypeString:
		return fmt.Sprintf("STRING(%v)", xStringValue(obj))
	case lispTypeInteger:
		return fmt.Sprintf("INTEGER(%v)", xIntegerValue(obj))
	case lispTypeSymbol:
		return fmt.Sprintf("SYMBOL(%v)", xSymbol(obj).name)
	case lispTypeVectorLike:
		val := xVectorLike(obj)

		switch val.vecType {
		case vectorLikeTypeNormal:
			return "VECTOR()"
		case vectorLikeTypeBuffer:
			buf := xBuffer(obj)
			return fmt.Sprintf("BUFFER(name=%v, live=%v)", buf.name, buf.live)
		case vectorLikeTypeSubroutine:
			subr := xSubroutine(obj)
			return fmt.Sprintf("SUBROUTINE(min=%v, max=%v)", subr.minArgs, subr.maxArgs)
		default:
			panic("unknown vector-like type")
		}
	default:
		panic(fmt.Sprintf("unknown object type: %v", obj.getType()))
	}
}
