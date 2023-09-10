package core

import (
	"slices"
	"unicode/utf8"
)

func (ec *execContext) listLength(obj lispObject) (int, error) {
	num := 0

	iter := ec.iterate(obj)
	for ; iter.hasNext(); iter.nextCons() {
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
		// TODO: Probably not correct given EU8 encoding
		num = utf8.RuneCountInString(xString(obj).val)
	case lispTypeCons:
		var err error
		num, err = ec.listLength(obj)
		if err != nil {
			return nil, err
		}
	case lispTypeVector:
		num = len(xVector(obj).val)
	default:
		if obj != ec.nil_ {
			return ec.wrongTypeArgument(ec.s.sequencep, obj)
		}
	}

	return newInteger(lispInt(num)), nil
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

func (ec *execContext) eql(o1, o2 lispObject) (lispObject, error) {
	if o1 == o2 {
		return ec.true_()
	}

	t1, t2 := o1.getType(), o2.getType()
	if t1 != t2 {
		return ec.false_()
	}

	switch t1 {
	case lispTypeInteger:
		if xIntegerValue(o1) == xIntegerValue(o2) {
			return ec.true_()
		}
	case lispTypeFloat:
		// TODO: Probably not correct, needs to match Emacs eql
		if xFloatValue(o1) == xFloatValue(o2) {
			return ec.true_()
		}
	}

	return ec.false_()
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

			equal, err = ec.equal(xCdr(o1), xCdr(o2))
			if err != nil {
				return nil, err
			}
			if equal != ec.nil_ {
				return ec.true_()
			}

			o2 = xCdr(o2)
		}

		if iter.hasError() {
			return iter.error()
		}

		return ec.false_()
	case lispTypeFloat:
		fallthrough
	case lispTypeInteger:
		return ec.eql(o1, o2)
	case lispTypeString:
		return ec.bool(xStringValue(o1) == xStringValue(o2))
	case lispTypeSymbol:
		// Symbols must match exactly (eq).
		return ec.false_()
	case lispTypeVector:
		v1 := xVector(o1)
		v2 := xVector(o2)

		if len(v1.val) != len(v2.val) {
			return ec.false_()
		}

		for i := 0; i < len(v1.val); i++ {
			equal, err := ec.equal(v1.val[i], v2.val[i])
			if err != nil {
				return nil, err
			}
			if equal == ec.nil_ {
				return ec.false_()
			}
		}

		return ec.true_()
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
			return ec.wrongTypeArgument(ec.s.plistp, plist)
		}

		if prop == xCar(tail) {
			_, err := ec.setCar(xCdr(tail), val)
			if err != nil {
				return nil, err
			}
			return plist, nil
		}

		prev = tail
		// Advance in pairs when traversing plist
		iter.nextCons()
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

		// Advance in pairs when traversing plist
		iter.nextCons()
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

func (ec *execContext) provide(feature, subfeatures lispObject) (lispObject, error) {
	return ec.nil_, nil
}

func (ec *execContext) nreverse(seq lispObject) (lispObject, error) {
	return ec.reverse(seq)
}

func (ec *execContext) reverse(seq lispObject) (lispObject, error) {
	switch seq.getType() {
	case lispTypeSymbol:
		if seq != ec.nil_ {
			break
		}
		fallthrough
	case lispTypeCons:
		result, err := ec.listToSlice(seq)
		if err != nil {
			return nil, err
		}
		slices.Reverse(result)
		return ec.makeList(result...), nil
	case lispTypeVector:
		copy := slices.Clone(xVector(seq).val)
		slices.Reverse(copy)
		return newVector(copy), nil
	case lispTypeString:
		return ec.pimacsUnimplemented(ec.s.reverse, "no reverse for string")
	}

	return ec.wrongTypeArgument(ec.s.sequencep, seq)
}

func (ec *execContext) require(feature, filename, noerror lispObject) (lispObject, error) {
	// TODO: Implement
	return ec.nil_, nil
}

func (ec *execContext) nthCdr(n, list lispObject) (lispObject, error) {
	if !integerp(n) {
		return ec.wrongTypeArgument(ec.s.integerp, n)
	}
	num := xIntegerValue(n)
	tail := list

	for ; 0 < num; num, tail = num-1, xCdr(tail) {
		if !consp(tail) {
			if tail != ec.nil_ {
				return ec.wrongTypeArgument(ec.s.listp, list)
			}
			return ec.nil_, nil
		}
	}

	return tail, nil
}

func (ec *execContext) nth(n, list lispObject) (lispObject, error) {
	cdr, err := ec.nthCdr(n, list)
	if err != nil {
		return nil, err
	}
	return ec.car(cdr)
}

func (ec *execContext) mapCarInternal(seq lispObject, length lispObject, function lispObject) ([]lispObject, error) {
	if !integerp(length) {
		return nil, xErrOnly(ec.wrongTypeArgument(ec.s.integerp, length))
	}

	leni := xIntegerValue(length)

	switch {
	case seq == ec.nil_:
		return []lispObject{}, nil
	case consp(seq):
		result := []lispObject{}
		tail := seq
		for i := lispInt(0); i < leni; i++ {
			if !consp(tail) {
				return result, nil
			}

			tmp, err := ec.funcall(function, xCar(tail))
			if err != nil {
				return nil, err
			}

			result = append(result, tmp)
			tail = xCdr(tail)
		}

		return result, nil
	default:
		return nil, xErrOnly(ec.pimacsUnimplemented(ec.s.mapCar, "mapcar unimplemented for this object: '%+v'", seq))
	}
}

func (ec *execContext) mapCar(function, sequence lispObject) (lispObject, error) {
	length, err := ec.length(sequence)
	if err != nil {
		return nil, err
	}

	if chartablep(sequence) {
		return ec.wrongTypeArgument(ec.s.listp, sequence)
	}

	result, err := ec.mapCarInternal(sequence, length, function)
	if err != nil {
		return nil, err
	}
	return ec.makeList(result...), nil
}

func (ec *execContext) sxHashEq(obj lispObject) (lispObject, error) {
	return ec.nil_, nil
}

func (ec *execContext) sxHashEql(obj lispObject) (lispObject, error) {
	return ec.nil_, nil
}

func (ec *execContext) sxHashEqual(obj lispObject) (lispObject, error) {
	return ec.nil_, nil
}

func (ec *execContext) sxHashEqualIncludingProperties(obj lispObject) (lispObject, error) {
	return ec.nil_, nil
}

func (ec *execContext) getKeyArg(key lispObject, args ...lispObject) (lispObject, bool) {
	for i, arg := range args {
		if arg == key && i+1 < len(args) {
			return args[i+1], true
		}
	}

	return nil, false
}

func (ec *execContext) makeHashTable(args ...lispObject) (lispObject, error) {
	test, found := ec.getKeyArg(ec.s.cTest)
	if !found {
		test = ec.s.eql
	}
	return test, nil
}

func (ec *execContext) symbolsOfFunctions() {
	ec.defSym(&ec.s.cTest, ":test")
	ec.defSym(&ec.s.cSize, ":size")
	ec.defSym(&ec.s.cPureCopy, ":purecopy")
	ec.defSym(&ec.s.cRehashSize, ":rehash-size")
	ec.defSym(&ec.s.cRehashThreshold, ":rehash-threshold")
	ec.defSym(&ec.s.cWeakness, ":weakness")
	ec.defSym(&ec.s.key, "key")
	ec.defSym(&ec.s.value, "value")
	ec.defSym(&ec.s.hashTableTest, "hash-table-test")
	ec.defSym(&ec.s.keyOrValue, "key-or-value")
	ec.defSym(&ec.s.keyAndValue, "key-and-value")

	ec.defSubr1(nil, "length", (*execContext).length, 1)
	ec.defSubr2(&ec.s.equal, "equal", (*execContext).equal, 2)
	ec.defSubr2(&ec.s.eql, "eql", (*execContext).eql, 2)
	ec.defSubr2(nil, "assq", (*execContext).assq, 2)
	ec.defSubr3(nil, "assoc", (*execContext).assoc, 2)
	ec.defSubr2(nil, "memq", (*execContext).memq, 2)
	ec.defSubr2(nil, "get", (*execContext).get, 2)
	ec.defSubr3(nil, "put", (*execContext).put, 3)
	ec.defSubr1(&ec.s.plistp, "plistp", (*execContext).plistp, 1)
	ec.defSubr2(nil, "plist-get", (*execContext).plistGet, 2)
	ec.defSubr3(nil, "plist-put", (*execContext).plistPut, 3)
	ec.defSubrM(nil, "nconc", (*execContext).nconc, 0)
	ec.defSubr2(&ec.s.provide, "provide", (*execContext).provide, 1)
	ec.defSubr1(nil, "nreverse", (*execContext).nreverse, 1)
	ec.defSubr1(&ec.s.reverse, "reverse", (*execContext).reverse, 1)
	ec.defSubr3(nil, "require", (*execContext).require, 1)
	ec.defSubr2(nil, "nthcdr", (*execContext).nthCdr, 2)
	ec.defSubr2(nil, "nth", (*execContext).nth, 2)
	ec.defSubr2(&ec.s.mapCar, "mapcar", (*execContext).mapCar, 2)
	ec.defSubrM(nil, "make-hash-table", (*execContext).makeHashTable, 0)
	ec.defSubr1(nil, "sxhash-eq", (*execContext).sxHashEq, 1)
	ec.defSubr1(nil, "sxhash-eql", (*execContext).sxHashEql, 1)
	ec.defSubr1(nil, "sxhash-equal", (*execContext).sxHashEqual, 1)
	ec.defSubr1(nil, "sxhash-equal-including-properties", (*execContext).sxHashEqualIncludingProperties, 1)

	ec.hashTestEq = &lispHashTableTest{
		name:         ec.s.eq,
		hashFunction: (*execContext).sxHashEq,
		compFunction: (*execContext).eq,
	}
	ec.hashTestEql = &lispHashTableTest{
		name:         ec.s.eql,
		hashFunction: (*execContext).sxHashEql,
		compFunction: (*execContext).eql,
	}
	ec.hashTestEqual = &lispHashTableTest{
		name:         ec.s.equal,
		hashFunction: (*execContext).sxHashEqual,
		compFunction: (*execContext).equal,
	}
}
