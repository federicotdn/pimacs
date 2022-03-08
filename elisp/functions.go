package elisp

func (ec *execContext) listLength(obj lispObject) (int, error) {
	num := 0

	iter := ec.iterate(obj)
	for ; iter.hasNext(); obj = iter.next() {
		num += 1
	}

	if iter.hasError() {
		return 0, xErrOnly(iter.error())
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
	iter := ec.iterate(alist)
	for ; iter.hasNext(); alist = iter.next() {
		element := xCar(alist)

		if consp(element) && xCar(element) == key {
			return element, nil
		}
	}

	if iter.hasError() {
		return iter.error()
	}

	return ec.nil_, nil
}

func (ec *execContext) assoc(key, alist, testFn lispObject) (lispObject, error) {
	// TAGS: incomplete
	iter := ec.iterate(alist)
	for ; iter.hasNext(); alist = iter.next() {
		element := xCar(alist)

		if consp(element) {
			equal, err := ec.equal(xCar(element), key)
			if err != nil {
				return nil, err
			}

			if equal == ec.t {
				return element, nil
			}
		}
	}

	if iter.hasError() {
		return iter.error()
	}

	return ec.nil_, nil
}

func (ec *execContext) memq(elt, list lispObject) (lispObject, error) {
	iter := ec.iterate(list)
	for ; iter.hasNext(); list = iter.next() {
		if xCar(list) == elt {
			return list, nil
		}
	}

	if iter.hasError() {
		return iter.error()
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
		iter := ec.iterate(o1)
		for ; iter.hasNext(); o1 = iter.next() {
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

		if iter.hasError() {
			return iter.error()
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
	iter := ec.iterate(tail).withPredicate(ec.g.plistp)

	for ; iter.hasNext(); tail = iter.next() {
		if !consp(xCdr(tail)) {
			break
		}

		if prop == xCar(tail) {
			_, err := ec.setCar(xCdr(tail), val)
			if err != nil {
				return nil, err
			}
			return plist, nil
		}

		prev = tail
		tail = xCdr(tail)
	}

	if iter.hasError() {
		return iter.error()
	}

	if prev == ec.nil_ {
		return ec.makeCons(prop, ec.makeCons(val, plist)), nil
	} else {
		newCell := ec.makeCons(prop, ec.makeCons(val, xCdr(xCdr(prev))))
		_, err := ec.setCdr(xCdr(prev), newCell)
		if err != nil {
			return nil, err
		}

		return plist, nil
	}
}

func (ec *execContext) plistGet(plist, prop lispObject) (lispObject, error) {
	iter := ec.iterate(plist)

	for ; iter.hasNext(); plist = iter.next() {
		if !consp(xCdr(plist)) {
			break
		}

		if prop == xCar(plist) {
			return xCar(xCdr(plist)), nil
		}

		plist = iter.next()
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

func (ec *execContext) nconc(args ...lispObject) (lispObject, error) {
	val := ec.nil_

	for i, tem := range args {
		if tem == ec.nil_ {
			continue
		}

		if val == ec.nil_ {
			val = tem
		}

		if i+1 == len(args) {
			// Break off early in last iteration
			break
		}

		if !consp(tem) {
			return ec.wrongTypeArgument(ec.g.consp, tem)
		}

		var tail lispObject
		for aux := tem; consp(aux); aux = xCdr(aux) {
			tail = aux
		}

		tem = args[i+1]
		_, err := ec.setCdr(tail, tem)
		if err != nil {
			return nil, err
		}

		if tem == ec.nil_ {
			args[i+1] = tail
		}
	}

	return val, nil
}

func (ec *execContext) symbolsOfFunctions() {
	ec.defSubr1("length", ec.length, 1)
	ec.defSubr2("equal", ec.equal, 2)
	ec.defSubr2("assq", ec.assq, 2)
	ec.defSubr3("assoc", ec.assoc, 2)
	ec.defSubr2("memq", ec.memq, 2)
	ec.defSubr2("get", ec.get, 2)
	ec.defSubr3("put", ec.put, 3)
	ec.defSubr2("plist-get", ec.plistGet, 2)
	ec.defSubr3("plist-put", ec.plistPut, 3)
	ec.defSubrM("nconc", ec.nconc, 0)
}
