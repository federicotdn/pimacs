package core

import (
	"hash/fnv"
	"math"
	"slices"
)

const sxHashMaxDepth = 3

type lispHashTableLookupResult struct {
	entries []lispHashTableEntry
	i       int
	code    lispInt
	err     error
}

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
		num = xStringSize(obj)
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

func (ec *execContext) plistPut(plist, prop, val, predicate lispObject) (lispObject, error) {
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

func (ec *execContext) plistGet(plist, prop, predicate lispObject) (lispObject, error) {
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
	return ec.plistGet(plist, propName, ec.nil_)
}

func (ec *execContext) put(symbol, propName, value lispObject) (lispObject, error) {
	plist, err := ec.symbolPlist(symbol)
	if err != nil {
		return nil, err
	}

	plist, err = ec.plistPut(plist, propName, value, ec.nil_)
	if err != nil {
		return nil, err
	}

	xSymbol(symbol).plist = plist
	return value, nil
}

func (ec *execContext) concat(args ...lispObject) (lispObject, error) {
	result := ""
	multibyte := false

	for _, arg := range args {
		switch {
		case stringp(arg):
			s := xString(arg)
			result += s.str()
			multibyte = multibyte || s.multibytep()
		case arg == ec.nil_:
		case consp(arg):
			fallthrough
		case vectorp(arg):
			fallthrough
		default:
			return ec.wrongTypeArgument(ec.s.sequencep, arg)
		}
	}

	return newString(result, multibyte), nil
}

func (ec *execContext) vconcat(args ...lispObject) (lispObject, error) {
	result := []lispObject{}

	for _, arg := range args {
		switch {
		case vectorp(arg):
			vec := xVector(arg)
			result = append(result, vec.val...)
		case arg == ec.nil_:
		case consp(arg):
			iter := ec.iterate(arg)
			for ; iter.hasNext(); arg = iter.nextCons() {
				result = append(result, xCar(arg))
			}

			if iter.hasError() {
				return iter.error()
			}
		case stringp(arg):
			fallthrough
		default:
			return ec.wrongTypeArgument(ec.s.sequencep, arg)
		}
	}

	return newVector(result), nil
}

func (ec *execContext) append_(args ...lispObject) (lispObject, error) {
	result := ec.nil_
	last := ec.nil_

	for _, arg := range args {
		switch {
		case consp(arg):
			head := newCons(xCar(arg), ec.nil_)
			prev := head
			arg = xCdr(arg)

			iter := ec.iterate(arg)
			for ; iter.hasNext(); arg = iter.nextCons() {
				next := newCons(xCar(arg), ec.nil_)
				xSetCdr(prev, next)
				prev = next
			}

			if iter.hasError() {
				return iter.error()
			}

			if result == ec.nil_ {
				result = head
			} else {
				xSetCdr(last, head)
			}
			last = prev
		case arg == ec.nil_:
		case vectorp(arg):
			vec := xVector(arg)

			for _, elem := range vec.val {
				node := newCons(elem, ec.nil_)
				if result == ec.nil_ {
					result = node
				} else {
					xSetCdr(last, node)
				}
				last = node
			}
		case stringp(arg):
			fallthrough
		default:
			return ec.wrongTypeArgument(ec.s.sequencep, arg)
		}
	}

	return result, nil
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

func (ec *execContext) copySequence(arg lispObject) (lispObject, error) {
	switch {
	case vectorp(arg):
		vec := xVector(arg)
		return newVector(slices.Clone(vec.val)), nil
	case arg == ec.nil_:
		return arg, nil
	case consp(arg):
		val := newCons(xCar(arg), ec.nil_)
		prev := val
		tail := xCdr(arg)

		iter := ec.iterate(tail)
		for ; iter.hasNext(); tail = iter.nextCons() {
			c := newCons(xCar(tail), ec.nil_)
			xSetCdr(prev, c)
			prev = c
		}

		if iter.hasError() {
			return iter.error()
		}

		return val, nil
	case stringp(arg):
		fallthrough
	default:
		return ec.wrongTypeArgument(ec.s.sequencep, arg)
	}
}

func (ec *execContext) provide(feature, subfeatures lispObject) (lispObject, error) {
	if !symbolp(feature) {
		return ec.wrongTypeArgument(ec.s.symbolp, feature)
	} else if !ec.listpBool(subfeatures) {
		return ec.wrongTypeArgument(ec.s.listp, subfeatures)
	}

	tem, err := ec.memq(feature, ec.v.features.val)
	if err != nil {
		return nil, err
	}
	if tem == ec.nil_ {
		ec.v.features.val = newCons(feature, ec.v.features.val)
	}
	if subfeatures != ec.nil_ {
		_, err := ec.put(feature, ec.s.subfeatures, subfeatures)
		if err != nil {
			return nil, err
		}
	}

	return feature, nil
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

func (ec *execContext) delq(element, list lispObject) (lispObject, error) {
	tail := list
	prev := ec.nil_

	iter := ec.iterate(tail)
	for ; iter.hasNext(); tail = iter.nextCons() {
		tem := xCar(tail)
		if element == tem {
			if prev == ec.nil_ {
				list = xCdr(tail)
			} else if _, err := ec.setCdr(prev, xCdr(tail)); err != nil {
				return nil, err
			}
		} else {
			prev = tail
		}
	}

	if iter.hasError() {
		return iter.error()
	}

	return list, nil
}

func (ec *execContext) sxHashObj(obj lispObject, depth int) lispInt {
	if depth > sxHashMaxDepth {
		return 0
	}

	switch obj.getType() {
	case lispTypeInteger:
		return xIntegerValue(obj)
	case lispTypeSymbol:
		return objAddr(obj)
	case lispTypeString:
		// TODO: Revise this - good enough?
		h := fnv.New64a()
		h.Write([]byte(xStringValue(obj)))
		return lispInt(h.Sum64())
	case lispTypeFloat:
		f := float64(xFloatValue(obj))
		return lispInt(math.Float64bits(f))
	default:
		ec.warning("sxhash implementation missing for object type '%+v'", obj.getType())
		return 0
	}
}

func (ec *execContext) sxHashEq(obj lispObject) (lispObject, error) {
	return newInteger(objAddr(obj)), nil
}

func (ec *execContext) sxHashEql(obj lispObject) (lispObject, error) {
	if numberp(obj) {
		return ec.sxHashEqual(obj)
	}
	return ec.sxHashEq(obj)
}

func (ec *execContext) sxHashEqual(obj lispObject) (lispObject, error) {
	return newInteger(ec.sxHashObj(obj, 0)), nil
}

func (ec *execContext) sxHashEqualIncludingProperties(obj lispObject) (lispObject, error) {
	return ec.nil_, nil
}

func (ec *execContext) makeHashTable(args ...lispObject) (lispObject, error) {
	test := ec.hashTestEql
	kwArgs, err := ec.kwPlistToMap(ec.makeList(args...))
	if err != nil {
		return nil, err
	}
	key := xSymbolName(ec.s.cTest)
	testSym := getDefault(kwArgs, key, ec.s.eql)
	delete(kwArgs, key)

	switch testSym {
	case ec.s.eql:
		// Use default value
	case ec.s.eq:
		test = ec.hashTestEq
	case ec.s.equal:
		test = ec.hashTestEqual
	default:
		prop, err := ec.get(testSym, ec.s.hashTableTest)
		if err != nil {
			return nil, err
		}

		if !consp(prop) || !consp(xCdr(prop)) {
			return ec.signalError("Invalid hash table test: '%+v'", prop)
		}

		test = &lispHashTableTest{
			name:         testSym,
			compFunction: xCar(prop),
			hashFunction: xCar(xCdr(prop)),
		}
	}

	key = xSymbolName(ec.s.cWeakness)
	delete(kwArgs, key)

	if len(kwArgs) > 0 {
		return ec.signalError("Invalid arguments list: '%+v'", args)
	}

	return &lispHashTable{
		val:  make(map[lispInt][]lispHashTableEntry),
		test: test,
	}, nil
}

func (ec *execContext) putHash(key, value, table lispObject) (lispObject, error) {
	if !hashtablep(table) {
		return ec.wrongTypeArgument(ec.s.hashTablep, table)
	}

	ht := xHashTable(table)
	result := ec.hashLookup(key, ht)
	if result.err != nil {
		return nil, result.err
	}

	entry := lispHashTableEntry{key: key, val: value, code: result.code}

	if result.entries == nil {
		ht.val[result.code] = []lispHashTableEntry{entry}
	} else if result.i < 0 {
		ht.val[result.code] = append(result.entries, entry)
	} else {
		result.entries[result.i] = entry
	}

	return value, nil
}

func (ec *execContext) hashLookup(key lispObject, ht *lispHashTable) lispHashTableLookupResult {
	codeObj, err := ec.funcall(ht.test.hashFunction, key)
	result := lispHashTableLookupResult{i: -1}

	if err != nil {
		result.err = err
		return result
	} else if !integerp(codeObj) {
		result.err = xErrOnly(ec.signalError("Invalid hash code type"))
		return result
	}

	code := xIntegerValue(codeObj)
	entries := ht.val[code]

	result.code = code
	result.entries = entries

	for i, entry := range entries {
		if entry.key == key {
			result.i = i
			break
		} else if ht.test.compFunction != nil &&
			ht.test.compFunction != ec.nil_ &&
			entry.code == code {
			cmp, err := ec.funcall(ht.test.compFunction, key, entry.key)
			if err != nil {
				result.err = err
				break
			}

			if cmp != ec.nil_ {
				result.i = i
				break
			}
		}
	}

	return result
}

func (ec *execContext) getHash(key, table, default_ lispObject) (lispObject, error) {
	if !hashtablep(table) {
		return ec.wrongTypeArgument(ec.s.hashTablep, table)
	}

	result := ec.hashLookup(key, xHashTable(table))
	if result.err != nil {
		return nil, result.err
	}

	if result.entries == nil || result.i < 0 {
		return default_, nil
	}

	return result.entries[result.i].val, nil
}

func (ec *execContext) remHash(key, table lispObject) (lispObject, error) {
	if !hashtablep(table) {
		return ec.wrongTypeArgument(ec.s.hashTablep, table)
	}

	ht := xHashTable(table)
	result := ec.hashLookup(key, xHashTable(table))
	if result.err != nil {
		return nil, result.err
	}

	if result.entries != nil && result.i >= 0 {
		ht.val[result.code] = slices.Delete(result.entries, result.i, result.i+1)
	}

	return ec.nil_, nil
}

func (ec *execContext) clearHash(table lispObject) (lispObject, error) {
	if !hashtablep(table) {
		return ec.wrongTypeArgument(ec.s.hashTablep, table)
	}

	clear(xHashTable(table).val)
	return table, nil
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
	ec.defSym(&ec.s.subfeatures, "subfeatures")
	ec.defVarLisp(&ec.v.features, "features", ec.makeList(ec.s.emacs))

	ec.defSubr1(nil, "length", (*execContext).length, 1)
	ec.defSubr2(&ec.s.equal, "equal", (*execContext).equal, 2)
	ec.defSubr2(&ec.s.eql, "eql", (*execContext).eql, 2)
	ec.defSubr2(nil, "assq", (*execContext).assq, 2)
	ec.defSubr3(nil, "assoc", (*execContext).assoc, 2)
	ec.defSubr2(nil, "memq", (*execContext).memq, 2)
	ec.defSubr2(nil, "get", (*execContext).get, 2)
	ec.defSubr3(nil, "put", (*execContext).put, 3)
	ec.defSubr1(&ec.s.plistp, "plistp", (*execContext).plistp, 1)
	ec.defSubr3(nil, "plist-get", (*execContext).plistGet, 2)
	ec.defSubr4(nil, "plist-put", (*execContext).plistPut, 3)
	ec.defSubrM(nil, "nconc", (*execContext).nconc, 0)
	ec.defSubrM(nil, "append", (*execContext).append_, 0)
	ec.defSubrM(nil, "concat", (*execContext).concat, 0)
	ec.defSubrM(nil, "vconcat", (*execContext).vconcat, 0)
	ec.defSubr1(nil, "copy-sequence", (*execContext).copySequence, 1)
	ec.defSubr2(nil, "provide", (*execContext).provide, 1)
	ec.defSubr1(nil, "nreverse", (*execContext).nreverse, 1)
	ec.defSubr1(&ec.s.reverse, "reverse", (*execContext).reverse, 1)
	ec.defSubr2(nil, "nthcdr", (*execContext).nthCdr, 2)
	ec.defSubr2(nil, "nth", (*execContext).nth, 2)
	ec.defSubr2(&ec.s.mapCar, "mapcar", (*execContext).mapCar, 2)
	ec.defSubrM(nil, "make-hash-table", (*execContext).makeHashTable, 0)
	ec.defSubr3(nil, "puthash", (*execContext).putHash, 3)
	ec.defSubr3(nil, "gethash", (*execContext).getHash, 2)
	ec.defSubr2(nil, "remhash", (*execContext).remHash, 2)
	ec.defSubr1(nil, "clrhash", (*execContext).clearHash, 1)
	ec.defSubr1(&ec.s.sxHashEq, "sxhash-eq", (*execContext).sxHashEq, 1)
	ec.defSubr1(&ec.s.sxHashEql, "sxhash-eql", (*execContext).sxHashEql, 1)
	ec.defSubr1(&ec.s.sxHashEqual, "sxhash-equal", (*execContext).sxHashEqual, 1)
	ec.defSubr1(
		&ec.s.sxHashEqualIncludingProperties,
		"sxhash-equal-including-properties",
		(*execContext).sxHashEqualIncludingProperties,
		1,
	)
	ec.defSubr2(nil, "delq", (*execContext).delq, 2)

	ec.hashTestEq = &lispHashTableTest{
		name:         ec.s.eq,
		hashFunction: ec.s.sxHashEq,
		compFunction: ec.s.eq,
	}
	ec.hashTestEql = &lispHashTableTest{
		name:         ec.s.eql,
		hashFunction: ec.s.sxHashEql,
		compFunction: ec.s.eql,
	}
	ec.hashTestEqual = &lispHashTableTest{
		name:         ec.s.equal,
		hashFunction: ec.s.sxHashEqual,
		compFunction: ec.s.equal,
	}
}
