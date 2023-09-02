package core

import (
	"unicode/utf8"
)

func (ec *execContext) listLength(obj lispObject) (int, error) {
	num := 0

	iter := ec.iterate(obj)
	for ; iter.hasNext(); obj = iter.nextCons() {
		num += 1
	}

	if iter.hasError() {
		return 0, xErrOnly(iter.error())
	}

	return num, nil
}

func (ec *execContext) length(obj lispObject) (lispObject, error) {
	num := 0

	switch obj.getType() {
	case lispTypeString:
		num = utf8.RuneCountInString(xString(obj).val)
	case lispTypeCons:
		var err error
		num, err = ec.listLength(obj)
		if err != nil {
			return nil, err
		}
	default:
		if obj != ec.nil_ {
			return ec.wrongTypeArgument(ec.s.sequencep, obj)
		}
	}

	return ec.makeInteger(lispInt(num)), nil
}

func (ec *execContext) assq(key, alist lispObject) (lispObject, error) {
	iter := ec.iterate(alist)
	for ; iter.hasNext(); alist = iter.nextCons() {
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
	iter := ec.iterate(alist)
	for ; iter.hasNext(); alist = iter.nextCons() {
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
	for ; iter.hasNext(); list = iter.nextCons() {
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
		for ; iter.hasNext(); o1 = iter.nextCons() {
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
	default:
		return ec.pimacsUnimplemented(ec.s.equal, "implementation missing for object type '%v'", t1)
	}
}

func (ec *execContext) plistp(object lispObject) (lispObject, error) {
	return ec.nil_, nil
}

func (ec *execContext) plistPut(plist, prop, val lispObject) (lispObject, error) {
	prev := ec.nil_
	tail := plist
	iter := ec.iterate(tail).withPredicate(ec.s.plistp)

	for ; iter.hasNext(); tail = iter.nextCons() {
		next := xCdr(tail)
		if !consp(next) {
			if next == ec.nil_ {
				break
			} else {
				return ec.wrongTypeArgument(ec.s.plistp, plist)
			}
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
		return newCons(prop, newCons(val, plist)), nil
	} else {
		newCell := newCons(prop, newCons(val, xCdr(xCdr(prev))))
		_, err := ec.setCdr(xCdr(prev), newCell)
		if err != nil {
			return nil, err
		}

		return plist, nil
	}
}

func (ec *execContext) plistGet(plist, prop lispObject) (lispObject, error) {
	iter := ec.iterate(plist)

	for ; iter.hasNext(); plist = iter.nextCons() {
		if !consp(xCdr(plist)) {
			break
		}

		if prop == xCar(plist) {
			return xCar(xCdr(plist)), nil
		}

		plist = iter.nextCons()
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
			return ec.wrongTypeArgument(ec.s.consp, tem)
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
	ec.defSubr1(nil, "length", ec.length, 1)
	ec.defSubr2(&ec.s.equal, "equal", ec.equal, 2)
	ec.defSubr2(nil, "assq", ec.assq, 2)
	ec.defSubr3(nil, "assoc", ec.assoc, 2)
	ec.defSubr2(nil, "memq", ec.memq, 2)
	ec.defSubr2(nil, "get", ec.get, 2)
	ec.defSubr3(nil, "put", ec.put, 3)
	ec.defSubr1(&ec.s.plistp, "plistp", ec.plistp, 1)
	ec.defSubr2(nil, "plist-get", ec.plistGet, 2)
	ec.defSubr3(nil, "plist-put", ec.plistPut, 3)
	ec.defSubrM(nil, "nconc", ec.nconc, 0)
}
