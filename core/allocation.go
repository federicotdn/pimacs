package core

func (ec *execContext) cons(car lispObject, cdr lispObject) (lispObject, error) {
	return newCons(car, cdr), nil
}

func (ec *execContext) list(args ...lispObject) (lispObject, error) {
	return ec.makeList(args...), nil
}

func (ec *execContext) makeVector(length, init lispObject) (lispObject, error) {
	if !integerp(length) {
		return ec.wrongTypeArgument(ec.s.integerp, length)
	}
	val := make([]lispObject, xIntegerValue(length))
	for i := 0; i < len(val); i++ {
		val[i] = init
	}

	return newVector(val), nil
}

func (ec *execContext) pureCopy(obj lispObject) (lispObject, error) {
	// TODO: Without having a build/dump phase, maybe this implementation
	// is good enough?
	return obj, nil
}

func (ec *execContext) symbolsOfAllocation() {
	ec.defSubr2(nil, "cons", (*execContext).cons, 2)
	ec.defSubrM(nil, "list", (*execContext).list, 0)
	ec.defSubr2(nil, "make-vector", (*execContext).makeVector, 2)
	ec.defSubr1(nil, "purecopy", (*execContext).pureCopy, 1)
}
