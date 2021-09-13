package elisp

func xCar(obj lispObject) lispObject {
	return obj.(*lispCons).car
}

func xCdr(obj lispObject) lispObject {
	return obj.(*lispCons).cdr
}

func xSetCar(cell, newCar lispObject) lispObject {
	cell.(*lispCons).car = newCar
	return newCar
}

func xSetCdr(cell, newCdr lispObject) lispObject {
	cell.(*lispCons).cdr = newCdr
	return newCdr
}

func xEnsure(obj lispObject, err error) lispObject {
	if err != nil {
		panic(err)
	}

	return obj
}

func xSymbol(obj lispObject) *lispSymbol {
	return obj.(*lispSymbol)
}

func xSymbolValue(obj lispObject) lispObject {
	return xSymbol(obj).value
}

func symbolp(obj lispObject) bool {
	return obj.getType() == lispTypeSymbol
}

func xCons(obj lispObject) *lispCons {
	return obj.(*lispCons)
}

func consp(obj lispObject) bool {
	return obj.getType() == lispTypeCons
}

func xVectorLike(obj lispObject) *lispVectorLike {
	return obj.(*lispVectorLike)
}

func vectorLikep(obj lispObject, vecType vectorLikeType) bool {
	vec, ok := obj.(*lispVectorLike)
	if !ok {
		return false
	}

	return vec.vecType == vecType
}

func xString(obj lispObject) *lispString {
	return obj.(*lispString)
}

func xStringValue(obj lispObject) string {
	return xString(obj).value
}

func stringp(obj lispObject) bool {
	return obj.getType() == lispTypeString
}

func xInteger(obj lispObject) *lispInteger {
	return obj.(*lispInteger)
}

func xIntegerValue(obj lispObject) lispInt {
	return xInteger(obj).value
}

func integerp(obj lispObject) bool {
	return obj.getType() == lispTypeInteger
}

func xFloat(obj lispObject) *lispFloat {
	return obj.(*lispFloat)
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
	// TODO: Incomplete!
	return vectorLikep(obj, vectorLikeTypeNormal) || stringp(obj)
}
