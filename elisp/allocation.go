package elisp

func (ec *execContext) cons(car lispObject, cdr lispObject) (lispObject, error) {
	return ec.makeCons(car, cdr), nil
}

func (ec *execContext) list(args ...lispObject) (lispObject, error) {
	return ec.makeList(args...), nil
}

func (ec *execContext) symbolsOfAllocation() {
	ec.defSubr2("cons", ec.cons, 2)
	ec.defSubrM("list", ec.list, 0)
}
