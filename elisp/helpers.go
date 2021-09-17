package elisp

func xEnsure(obj lispObject, err error) lispObject {
	if err != nil {
		terminate("cannot return value due to error: '%v'", err)
	}

	return obj
}

func xErrOnly(obj lispObject, err error) error {
	if obj != nil {
		terminate("was handed a non-nil Lisp object: '%v'", obj)
	}

	return err
}

func xSymbol(obj lispObject) *lispSymbol {
	val, ok := obj.(*lispSymbol)
	if !ok {
		terminate("object is not a symbol: '%v'", obj)
	}

	return val
}

func xSymbolValue(obj lispObject) lispObject {
	return xSymbol(obj).value
}

func symbolp(obj lispObject) bool {
	return obj.getType() == lispTypeSymbol
}

func xCons(obj lispObject) *lispCons {
	val, ok := obj.(*lispCons)
	if !ok {
		terminate("object is not a cons: '%v'", obj)
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
		terminate("object is not vector-like: '%v'", obj)
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

func subrp(obj lispObject) bool {
	return vectorLikep(obj, vectorLikeTypeSubroutine)
}

func xString(obj lispObject) *lispString {
	val, ok := obj.(*lispString)
	if !ok {
		terminate("object is not a string: '%v'", obj)
	}

	return val
}

func xStringValue(obj lispObject) string {
	return xString(obj).value
}

func stringp(obj lispObject) bool {
	return obj.getType() == lispTypeString
}

func xInteger(obj lispObject) *lispInteger {
	val, ok := obj.(*lispInteger)
	if !ok {
		terminate("object is not an integer: '%v'", obj)
	}

	return val
}

func xIntegerValue(obj lispObject) lispInt {
	return xInteger(obj).value
}

func integerp(obj lispObject) bool {
	return obj.getType() == lispTypeInteger
}

func xFloat(obj lispObject) *lispFloat {
	val, ok := obj.(*lispFloat)
	if !ok {
		terminate("object is not a float: '%v'", obj)
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

func arrayp(obj lispObject) bool {
	// TAGS: incomplete
	return vectorLikep(obj, vectorLikeTypeNormal) || stringp(obj)
}
