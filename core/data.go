package core

type arithmeticCmp int
type arithmeticOp int

const (
	arithmeticCmpEqual arithmeticCmp = iota + 1
	arithmeticCmpNotEqual
	arithmeticCmpLess
	arithmeticCmpGreater
	arithmeticCmpLessOrEqual
	arithmeticCmpGreaterOrEqual
)

const (
	arithmeticOpAdd arithmeticOp = iota + 1
	arithmeticOpSub
	arithmeticOpMul
	arithmeticOpDiv
	arithmeticOpAnd
	arithmeticOpOr
	arithmeticOpXor
)

func (ec *execContext) null(object lispObject) (lispObject, error) {
	return ec.bool(object == ec.nil_)
}

func (ec *execContext) sequencep(object lispObject) (lispObject, error) {
	return ec.bool(ec.listpBool(object) || arrayp(object))
}

func (ec *execContext) consp(object lispObject) (lispObject, error) {
	return ec.bool(consp(object))
}

func (ec *execContext) listpBool(object lispObject) bool {
	return (object == ec.nil_ || consp(object))
}

func (ec *execContext) listp(object lispObject) (lispObject, error) {
	return ec.bool(ec.listpBool(object))
}

func (ec *execContext) arrayp(object lispObject) (lispObject, error) {
	return ec.bool(arrayp(object))
}

func (ec *execContext) atom(object lispObject) (lispObject, error) {
	return ec.bool(!consp(object))
}

func (ec *execContext) symbolp(object lispObject) (lispObject, error) {
	return ec.bool(symbolp(object))
}

func (ec *execContext) stringp(object lispObject) (lispObject, error) {
	return ec.bool(stringp(object))
}

func (ec *execContext) multibyteStringp(object lispObject) (lispObject, error) {
	return ec.bool(stringp(object) && xStringMultibytep(object))
}

func (ec *execContext) numberOrMarkerp(object lispObject) (lispObject, error) {
	return ec.bool(numberOrMarkerp(object))
}

func (ec *execContext) integerOrMarkerp(object lispObject) (lispObject, error) {
	return ec.bool(integerOrMarkerp(object))
}

func (ec *execContext) charOrStringp(object lispObject) (lispObject, error) {
	return ec.bool(characterp(object) || stringp(object))
}

func (ec *execContext) integerp(object lispObject) (lispObject, error) {
	return ec.bool(integerp(object))
}

func (ec *execContext) numberp(object lispObject) (lispObject, error) {
	return ec.bool(numberp(object))
}

func (ec *execContext) characterp(object, ignore lispObject) (lispObject, error) {
	return ec.bool(characterp(object))
}

func (ec *execContext) bufferp(object lispObject) (lispObject, error) {
	return ec.bool(bufferp(object))
}

func (ec *execContext) vectorp(object lispObject) (lispObject, error) {
	return ec.bool(vectorp(object))
}

func (ec *execContext) charTablep(object lispObject) (lispObject, error) {
	return ec.bool(chartablep(object))
}

func (ec *execContext) channelp(object lispObject) (lispObject, error) {
	return ec.bool(channelp(object))
}

func (ec *execContext) hashTablep(object lispObject) (lispObject, error) {
	return ec.bool(hashtablep(object))
}

func (ec *execContext) boundp(symbol lispObject) (lispObject, error) {
	if !symbolp(symbol) {
		return ec.wrongTypeArgument(ec.s.symbolp, symbol)
	}
	val, err := ec.findSymbolValue(symbol)
	if err != nil {
		return nil, err
	}

	return ec.bool(val != ec.s.unbound)
}

func (ec *execContext) fboundp(symbol lispObject) (lispObject, error) {
	if !symbolp(symbol) {
		return ec.wrongTypeArgument(ec.s.symbolp, symbol)
	}
	return ec.bool(xSymbol(symbol).function != ec.nil_)
}

func (ec *execContext) makunbound(symbol lispObject) (lispObject, error) {
	_, err := ec.set(symbol, ec.s.unbound)
	if err != nil {
		return nil, err
	}
	return symbol, nil
}

func (ec *execContext) fmakunbound(symbol lispObject) (lispObject, error) {
	if !symbolp(symbol) {
		return ec.wrongTypeArgument(ec.s.symbolp, symbol)
	}
	xSymbol(symbol).function = ec.s.unbound
	return symbol, nil
}

func (ec *execContext) car(obj lispObject) (lispObject, error) {
	if obj == ec.nil_ {
		return ec.nil_, nil
	}

	if !consp(obj) {
		return ec.wrongTypeArgument(ec.s.listp, obj)
	}
	return xCar(obj), nil
}

func (ec *execContext) cdr(obj lispObject) (lispObject, error) {
	if obj == ec.nil_ {
		return ec.nil_, nil
	}

	if !consp(obj) {
		return ec.wrongTypeArgument(ec.s.listp, obj)
	}
	return xCdr(obj), nil
}

func (ec *execContext) carSafe(obj lispObject) (lispObject, error) {
	if !consp(obj) {
		return ec.nil_, nil
	}
	return xCar(obj), nil
}

func (ec *execContext) cdrSafe(obj lispObject) (lispObject, error) {
	if !consp(obj) {
		return ec.nil_, nil
	}
	return xCdr(obj), nil
}

func (ec *execContext) setCar(obj, newCar lispObject) (lispObject, error) {
	if !consp(obj) {
		return ec.wrongTypeArgument(ec.s.consp, obj)
	}
	xSetCar(obj, newCar)
	return newCar, nil
}

func (ec *execContext) setCdr(obj, newCdr lispObject) (lispObject, error) {
	if !consp(obj) {
		return ec.wrongTypeArgument(ec.s.consp, obj)
	}
	xSetCdr(obj, newCdr)
	return newCdr, nil
}

func (ec *execContext) indirectFunctionInternal(obj lispObject) lispObject {
	for symbolp(obj) && obj != ec.nil_ {
		obj = xSymbol(obj).function
	}
	return obj
}

func (ec *execContext) indirectFunction(obj, _ lispObject) (lispObject, error) {
	return ec.indirectFunctionInternal(obj), nil
}

func (ec *execContext) setDefaultInternal(symbol, val lispObject) error {
	if !symbolp(symbol) {
		return xErrOnly(ec.wrongTypeArgument(ec.s.symbolp, symbol))
	}

	sym := xSymbol(symbol)

	switch sym.redirect {
	case symbolRedirectPlain:
		return ec.setInternal(symbol, val)
	default:
		ec.terminate("unknown symbol redirect type: '%+v'", sym.redirect)
		return nil
	}
}

func (ec *execContext) setDefault(symbol, val lispObject) (lispObject, error) {
	return val, ec.setDefaultInternal(symbol, val)
}

func (ec *execContext) setInternal(symbol, newVal lispObject) error {
	if !symbolp(symbol) {
		return xErrOnly(ec.wrongTypeArgument(ec.s.symbolp, symbol))
	}

	sym := xSymbol(symbol)

	switch sym.redirect {
	case symbolRedirectPlain:
		sym.val = newVal
		return nil
	case symbolRedirectFwd:
		return sym.fwd.setValue(ec, newVal)
	default:
		ec.terminate("unknown symbol redirect type: '%+v'", sym.redirect)
		return nil
	}
}

func (ec *execContext) set(symbol, newVal lispObject) (lispObject, error) {
	err := ec.setInternal(symbol, newVal)
	if err != nil {
		return nil, err
	}
	return newVal, nil
}

func (ec *execContext) fset(symbol, definition lispObject) (lispObject, error) {
	if symbol == ec.nil_ && definition != ec.nil_ {
		return ec.signalN(ec.s.settingConstant, symbol)
	}

	if !symbolp(symbol) {
		return ec.wrongTypeArgument(ec.s.symbolp, symbol)
	}

	xSymbol(symbol).function = definition
	return definition, nil
}

func (ec *execContext) findSymbolValue(symbol lispObject) (lispObject, error) {
	if !symbolp(symbol) {
		return ec.wrongTypeArgument(ec.s.symbolp, symbol)
	}

	sym := xSymbol(symbol)
	var val lispObject

	switch sym.redirect {
	case symbolRedirectPlain:
		val = sym.val
	case symbolRedirectFwd:
		val = sym.fwd.value(ec)
	default:
		ec.terminate("unknown symbol redirect type: '%+v'", sym.redirect)
	}

	return val, nil
}

func (ec *execContext) symbolValue(symbol lispObject) (lispObject, error) {
	val, err := ec.findSymbolValue(symbol)
	if err != nil {
		return nil, err
	}
	if val == ec.s.unbound {
		return ec.signalN(ec.s.voidVariable, symbol)
	}

	return val, nil
}

func (ec *execContext) symbolFunction(symbol lispObject) (lispObject, error) {
	if !symbolp(symbol) {
		return ec.wrongTypeArgument(ec.s.symbolp, symbol)
	}

	return xSymbol(symbol).function, nil
}

func (ec *execContext) eq(obj1, obj2 lispObject) (lispObject, error) {
	return ec.bool(obj1 == obj2)
}

func (ec *execContext) defalias(symbol, definition, docstring lispObject) (lispObject, error) {
	return ec.fset(symbol, definition)
}

func (ec *execContext) symbolPlist(symbol lispObject) (lispObject, error) {
	if !symbolp(symbol) {
		return ec.wrongTypeArgument(ec.s.symbolp, symbol)
	}

	return xSymbol(symbol).plist, nil
}

func (ec *execContext) symbolName(symbol lispObject) (lispObject, error) {
	if !symbolp(symbol) {
		return ec.wrongTypeArgument(ec.s.symbolp, symbol)
	}

	return newUniOrMultibyteString(xSymbolName(symbol)), nil
}

func (ec *execContext) plusSign(objs ...lispObject) (lispObject, error) {
	if len(objs) == 0 {
		return newInteger(0), nil
	}
	if !integerOrMarkerp(objs[0]) {
		return ec.wrongTypeArgument(ec.s.integerOrMarkerp, objs[0])
	}
	return ec.arithmeticOperate(arithmeticOpAdd, objs[0], objs[1:]...)
}

func (ec *execContext) logiOr(objs ...lispObject) (lispObject, error) {
	if len(objs) == 0 {
		return newInteger(0), nil
	}
	if !integerOrMarkerp(objs[0]) {
		return ec.wrongTypeArgument(ec.s.integerOrMarkerp, objs[0])
	}
	return ec.arithmeticOperate(arithmeticOpOr, objs[0], objs[1:]...)
}

func (ec *execContext) onePlus(number lispObject) (lispObject, error) {
	if !numberOrMarkerp(number) {
		return ec.wrongTypeArgument(ec.s.numberOrMarkerp, number)
	}

	if integerp(number) {
		return newInteger(xIntegerValue(number) + 1), nil
	}
	return newFloat(xFloatValue(number) + 1.0), nil
}

func (ec *execContext) arithmeticFloatOperate(op arithmeticOp, val lispObject, objs ...lispObject) (lispObject, error) {
	return ec.pimacsUnimplemented(ec.nil_, "arithmetic operator for float is unimplemented")
}

func (ec *execContext) arithmeticOperate(op arithmeticOp, val lispObject, objs ...lispObject) (lispObject, error) {
	if !numberOrMarkerp(val) {
		return ec.wrongTypeArgument(ec.s.numberOrMarkerp, val)
	} else if markerp(val) {
		return ec.pimacsUnimplemented(ec.nil_, "arithmetic operator for markers is unimplemented")
	} else if floatp(val) {
		return ec.arithmeticFloatOperate(op, val, objs...)
	}

	accum := xIntegerValue(val)
	for _, obj := range objs {
		if !numberOrMarkerp(obj) {
			return ec.wrongTypeArgument(ec.s.numberOrMarkerp, obj)
		} else if !integerp(obj) {
			return ec.pimacsUnimplemented(ec.nil_, "arithmetic operator for markers/float is unimplemented")
		}

		next := xIntegerValue(obj)

		switch op {
		case arithmeticOpAdd:
			accum += next
		case arithmeticOpSub:
			accum -= next
		case arithmeticOpMul:
			accum *= next
		case arithmeticOpDiv:
			if next == 0 {
				return ec.signalN(ec.s.arithError)
			}
			accum /= next
		case arithmeticOpAnd:
			accum &= next
		case arithmeticOpOr:
			accum |= next
		case arithmeticOpXor:
			accum ^= next
		default:
			return ec.signalError("Unknown arithmetic operator")
		}
	}

	return newInteger(accum), nil
}

func (ec *execContext) arithmeticCompare(cmp arithmeticCmp, objs ...lispObject) (lispObject, error) {
	for i := 1; i < len(objs); i++ {
		if !numberp(objs[i-1]) {
			return ec.wrongTypeArgument(ec.s.numberOrMarkerp, objs[i-1])
		} else if !numberp(objs[i]) {
			return ec.wrongTypeArgument(ec.s.numberOrMarkerp, objs[i])
		}

		if !integerp(objs[i-1]) || !integerp(objs[i]) {
			return ec.pimacsUnimplemented(ec.nil_, "arithmetic comparison for floats is unimplemented")
		}

		i1 := xIntegerValue(objs[i-1])
		i2 := xIntegerValue(objs[i])
		cond := false

		switch cmp {
		case arithmeticCmpEqual:
			cond = (i1 == i2)
		case arithmeticCmpNotEqual:
			cond = (i1 != i2)
		case arithmeticCmpLess:
			cond = (i1 < i2)
		case arithmeticCmpGreater:
			cond = (i1 > i2)
		case arithmeticCmpLessOrEqual:
			cond = (i1 <= i2)
		case arithmeticCmpGreaterOrEqual:
			cond = (i1 >= i2)
		default:
			return ec.signalError("Unknown arithmetic comparison operator")
		}

		if !cond {
			return ec.false_()
		}
	}

	return ec.true_()
}

func (ec *execContext) lessThanSign(objs ...lispObject) (lispObject, error) {
	return ec.arithmeticCompare(arithmeticCmpLess, objs...)
}

func (ec *execContext) greaterThanSign(objs ...lispObject) (lispObject, error) {
	return ec.arithmeticCompare(arithmeticCmpGreater, objs...)
}

func (ec *execContext) equalsSign(objs ...lispObject) (lispObject, error) {
	return ec.arithmeticCompare(arithmeticCmpEqual, objs...)
}

func (ec *execContext) notEqualsSign(obj1, obj2 lispObject) (lispObject, error) {
	return ec.arithmeticCompare(arithmeticCmpNotEqual, obj1, obj2)
}

func (ec *execContext) lessThanEqualsSign(objs ...lispObject) (lispObject, error) {
	return ec.arithmeticCompare(arithmeticCmpLessOrEqual, objs...)
}

func (ec *execContext) greaterThanEqualsSign(objs ...lispObject) (lispObject, error) {
	return ec.arithmeticCompare(arithmeticCmpGreaterOrEqual, objs...)
}

func (ec *execContext) bareSymbol(sym lispObject) (lispObject, error) {
	// TODO: Is this correct?
	return sym, nil
}

func (ec *execContext) aref(array, index lispObject) (lispObject, error) {
	if !integerp(index) {
		return ec.wrongTypeArgument(ec.s.integerp, index)
	}

	idx := int(xIntegerValue(index))
	if idx < 0 {
		return ec.argsOutOfRange(array, index)
	}

	switch array.getType() {
	case lispTypeString:
		val, err := xStringAref(array, idx)
		if err != nil {
			return ec.argsOutOfRange(array, index)
		}
		return newInteger(val), nil
	case lispTypeVector:
		val := xVectorValue(array)
		if idx >= len(val) {
			return ec.argsOutOfRange(array, index)
		}
		return val[idx], nil
	default:
		return ec.pimacsUnimplemented(ec.nil_, "aref unimplemented")
	}
}

func (ec *execContext) aset(array, index, elem lispObject) (lispObject, error) {
	if !integerp(index) {
		return ec.wrongTypeArgument(ec.s.integerp, index)
	}

	idx := int(xIntegerValue(index))
	if idx < 0 {
		return ec.argsOutOfRange(array, index)
	}

	switch array.getType() {
	case lispTypeVector:
		val := xVectorValue(array)
		if idx >= len(val) {
			return ec.argsOutOfRange(array, index)
		}
		val[idx] = elem
	default:
		return ec.pimacsUnimplemented(ec.nil_, "aset unimplemented")
	}

	return elem, nil
}

func (ec *execContext) symbolsOfData() {
	ec.defSym(&ec.s.wholeNump, "wholenump")

	ec.defSubr1(nil, "null", (*execContext).null, 1)
	ec.defSubr1(&ec.s.sequencep, "sequencep", (*execContext).sequencep, 1)
	ec.defSubr1(&ec.s.consp, "consp", (*execContext).consp, 1)
	ec.defSubr1(&ec.s.listp, "listp", (*execContext).listp, 1)
	ec.defSubr1(&ec.s.arrayp, "arrayp", (*execContext).arrayp, 1)
	ec.defSubr1(nil, "atom", (*execContext).atom, 1)
	ec.defSubr1(&ec.s.symbolp, "symbolp", (*execContext).symbolp, 1)
	ec.defSubr1(&ec.s.stringp, "stringp", (*execContext).stringp, 1)
	ec.defSubr1(nil, "multibyte-string-p", (*execContext).multibyteStringp, 1)
	ec.defSubr1(&ec.s.numberOrMarkerp, "number-or-marker-p", (*execContext).numberOrMarkerp, 1)
	ec.defSubr1(&ec.s.integerOrMarkerp, "integer-or-marker-p", (*execContext).integerOrMarkerp, 1)
	ec.defSubr1(&ec.s.charOrStringp, "char-or-string-p", (*execContext).charOrStringp, 1)
	ec.defSubr1(&ec.s.integerp, "integerp", (*execContext).numberOrMarkerp, 1)
	ec.defSubr1(nil, "numberp", (*execContext).numberp, 1)
	ec.defSubr1(&ec.s.bufferp, "bufferp", (*execContext).bufferp, 1)
	ec.defSubr2(&ec.s.characterp, "characterp", (*execContext).characterp, 1)
	ec.defSubr1(&ec.s.charTablep, "char-table-p", (*execContext).charTablep, 1)
	ec.defSubr1(&ec.s.channelp, "channelp", (*execContext).channelp, 1)
	ec.defSubr1(nil, "vectorp", (*execContext).vectorp, 1)
	ec.defSubr1(&ec.s.hashTablep, "hashtablep", (*execContext).hashTablep, 1)
	ec.defSubr1(nil, "boundp", (*execContext).boundp, 1)
	ec.defSubr1(nil, "fboundp", (*execContext).fboundp, 1)
	ec.defSubr1(nil, "makunbound", (*execContext).makunbound, 1)
	ec.defSubr1(nil, "fmakunbound", (*execContext).fmakunbound, 1)
	ec.defSubr2(nil, "indirect-function", (*execContext).indirectFunction, 1)
	ec.defSubr1(nil, "car", (*execContext).car, 1)
	ec.defSubr1(nil, "cdr", (*execContext).cdr, 1)
	ec.defSubr1(nil, "car-safe", (*execContext).carSafe, 1)
	ec.defSubr1(nil, "cdr-safe", (*execContext).cdrSafe, 1)
	ec.defSubr2(nil, "setcar", (*execContext).setCar, 2)
	ec.defSubr2(nil, "setcdr", (*execContext).setCdr, 2)
	ec.defSubr1(nil, "symbol-plist", (*execContext).symbolPlist, 1)
	ec.defSubr1(nil, "symbol-name", (*execContext).symbolName, 1)
	ec.defSubr2(nil, "set", (*execContext).set, 2)
	ec.defSubr2(nil, "set-default", (*execContext).setDefault, 2)
	ec.defSubr2(nil, "fset", (*execContext).fset, 2)
	ec.defSubr1(nil, "symbol-value", (*execContext).symbolValue, 1)
	ec.defSubr1(nil, "symbol-function", (*execContext).symbolFunction, 1)
	ec.defSubr2(&ec.s.eq, "eq", (*execContext).eq, 2)
	ec.defSubr3(nil, "defalias", (*execContext).defalias, 2)
	ec.defSubrM(nil, "+", (*execContext).plusSign, 0)
	ec.defSubrM(nil, "logior", (*execContext).logiOr, 0)
	ec.defSubrM(nil, "<", (*execContext).lessThanSign, 1)
	ec.defSubrM(nil, ">", (*execContext).greaterThanSign, 1)
	ec.defSubrM(nil, "=", (*execContext).equalsSign, 1)
	ec.defSubr2(nil, "/=", (*execContext).notEqualsSign, 2)
	ec.defSubrM(nil, "<=", (*execContext).lessThanEqualsSign, 1)
	ec.defSubrM(nil, ">=", (*execContext).greaterThanEqualsSign, 1)
	ec.defSubr1(nil, "bare-symbol", (*execContext).bareSymbol, 1)
	ec.defSubr2(nil, "aref", (*execContext).aref, 2)
	ec.defSubr3(nil, "aset", (*execContext).aset, 3)
	ec.defSubr1(nil, "1+", (*execContext).onePlus, 1)
}
