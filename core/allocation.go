package core

func (ec *execContext) cons(car lispObject, cdr lispObject) (lispObject, error) {
	return newCons(car, cdr), nil
}

func (ec *execContext) list(args ...lispObject) (lispObject, error) {
	return ec.makeList(args...), nil
}

func (ec *execContext) symbolsOfAllocation() {
	ec.defSubr2(nil, "cons", (*execContext).cons, 2)
	ec.defSubrM(nil, "list", (*execContext).list, 0)
}
