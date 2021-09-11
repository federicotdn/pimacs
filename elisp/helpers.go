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

func xString(obj lispObject) *lispString {
	return obj.(*lispString)
}

func xInteger(obj lispObject) *lispInteger {
	return obj.(*lispInteger)
}

func integerp(obj lispObject) bool {
	return obj.getType() == lispTypeInt
}

func xFloat(obj lispObject) *lispFloat {
	return obj.(*lispFloat)
}
