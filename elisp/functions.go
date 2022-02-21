package elisp

func (ec *execContext) listLength(obj lispObject) (int, error) {
	tail := obj
	num := 0

	for ; consp(tail); tail = xCdr(tail) {
		num += 1
	}

	if tail != ec.nil_ {
		return 0, xErrOnly(ec.wrongTypeArgument(ec.g.listp, obj))
	}

	return num, nil
}

func (ec *execContext) length(obj lispObject) (lispObject, error) {
	// TAGS: incomplete
	num := 0

	switch obj.getType() {
	case lispTypeString:
		num = len(xString(obj).value)
	case lispTypeCons:
		var err error
		num, err = ec.listLength(obj)
		if err != nil {
			return nil, err
		}
	default:
		if obj != ec.nil_ {
			return ec.wrongTypeArgument(ec.g.sequencep, obj)
		}
	}

	return ec.makeInteger(lispInt(num)), nil
}

func (ec *execContext) assq(key, alist lispObject) (lispObject, error) {
	tail := alist
	for ; consp(tail); tail = xCdr(tail) {
		element := xCar(tail)

		if consp(element) && xCar(element) == key {
			return element, nil
		}
	}

	if tail != ec.nil_ {
		return ec.wrongTypeArgument(ec.g.listp, alist)
	}

	return ec.nil_, nil
}

func (ec *execContext) memq(elt, list lispObject) (lispObject, error) {
	tail := list
	for ; consp(tail); tail = xCdr(tail) {
		if xCar(tail) == elt {
			return tail, nil
		}
	}

	if tail != ec.nil_ {
		return ec.wrongTypeArgument(ec.g.listp, list)
	}

	return ec.nil_, nil
}

func (ec *execContext) equal(o1, o2 lispObject) (lispObject, error) {
	// TAGS: incomplete
	if o1 == o2 {
		return ec.true_()
	}

	t1, t2 := o1.getType(), o2.getType()
	if t1 != t2 {
		return ec.false_()
	}

	switch t1 {
	case lispTypeCons:
		for ; consp(o1); o1 = xCdr(o1) {
			if !consp(o2) {
				return ec.false_()
			}

			equal, err := ec.equal(xCar(o1), xCar(o2))
			if err != nil {
				return nil, err
			}

			if equal == ec.nil_ {
				return ec.false_()
			}

			o2 = xCdr(o2)
			if xCdr(o1) == o2 {
				return ec.true_()
			}
		}

		return ec.false_()
	case lispTypeFloat:
		return ec.bool(xFloatValue(o1) == xFloatValue(o2))
	case lispTypeInteger:
		return ec.bool(xIntegerValue(o1) == xIntegerValue(o2))
	case lispTypeString:
		return ec.bool(xStringValue(o1) == xStringValue(o2))
	case lispTypeSymbol:
		// Symbols must match exactly (eq).
		return ec.false_()
	case lispTypeVectorLike:
		vec1, vec2 := xVectorLike(o1), xVectorLike(o2)
		if vec1.vecType != vec2.vecType {
			return ec.false_()
		}

		return ec.pimacsUnimplemented(ec.g.equal, "unknown vector-like type: '%v'", vec1.vecType)
	}

	return ec.false_()
}

func (ec *execContext) plistPut(plist, prop, val lispObject) (lispObject, error) {
	prev := ec.nil_
	tail := plist

	for consp(tail) {
		next := xCdr(tail)

		if !consp(next) {
			break
		} else if prop == xCar(tail) {
			xSetCar(next, val)
			return plist, nil
		}

		prev = tail
		tail = xCdr(next)
	}

	if tail != ec.nil_ {
		return ec.wrongTypeArgument(ec.g.listp, plist)
	}

	if prev == ec.nil_ {
		return ec.makeCons(prop, ec.makeCons(val, plist)), nil
	} else {
		newCell := ec.makeCons(prop, ec.makeCons(val, xCdr(xCdr(prev))))
		xSetCdr(xCdr(prev), newCell)
		return plist, nil
	}
}

func (ec *execContext) plistGet(plist, prop lispObject) (lispObject, error) {
	tail := plist

	for consp(tail) {
		next := xCdr(tail)

		if !consp(next) {
			break
		} else if prop == xCar(tail) {
			return xCar(next), nil
		}

		tail = xCdr(next)
	}

	return ec.nil_, nil
}

func (ec *execContext) get(symbol, propName lispObject) (lispObject, error) {
	plist, err := ec.symbolPlist(symbol)
	if err != nil {
		return nil, err
	}
	return ec.plistGet(plist, propName)
}

func (ec *execContext) put(symbol, propName, value lispObject) (lispObject, error) {
	plist, err := ec.symbolPlist(symbol)
	if err != nil {
		return nil, err
	}

	plist, err = ec.plistPut(plist, propName, value)
	if err != nil {
		return nil, err
	}

	xSymbol(symbol).plist = plist
	return value, nil
}

func (ec *execContext) symbolsOfFunctions() {
	ec.defSubr1("length", ec.length, 1)
	ec.defSubr2("equal", ec.equal, 2)
	ec.defSubr2("assq", ec.assq, 2)
	ec.defSubr2("memq", ec.memq, 2)
	ec.defSubr2("get", ec.get, 2)
	ec.defSubr3("put", ec.put, 3)
	ec.defSubr2("plist-get", ec.plistGet, 2)
	ec.defSubr3("plist-put", ec.plistPut, 3)
}
