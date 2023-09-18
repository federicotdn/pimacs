package core

func (ec *execContext) null(object lispObject) (lispObject, error) {
	return ec.bool(object == ec.nil_)
}

func (ec *execContext) sequencep(object lispObject) (lispObject, error) {
	return ec.bool(consp(object) || object == ec.nil_ || arrayp(object))
}

func (ec *execContext) consp(object lispObject) (lispObject, error) {
	return ec.bool(consp(object))
}

func (ec *execContext) listp(object lispObject) (lispObject, error) {
	return ec.bool(object == ec.nil_ || consp(object))
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

func (ec *execContext) numberOrMarkerp(object lispObject) (lispObject, error) {
	return ec.bool(numberp(object))
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

	return newString(xSymbolName(symbol)), nil
}

func (ec *execContext) plusSign(objs ...lispObject) (lispObject, error) {
	var total lispInt = 0
	for _, obj := range objs {
		if !numberp(obj) {
			return ec.wrongTypeArgument(ec.s.numberOrMarkerp, obj)
		}
		total += xIntegerValue(obj)
	}

	return newInteger(total), nil
}

func (ec *execContext) lessThanSign(objs ...lispObject) (lispObject, error) {
	for i := 1; i < len(objs); i++ {
		if !numberp(objs[i-1]) {
			return ec.wrongTypeArgument(ec.s.numberOrMarkerp, objs[i-1])
		} else if !numberp(objs[i]) {
			return ec.wrongTypeArgument(ec.s.numberOrMarkerp, objs[i])
		}

		if xIntegerValue(objs[i-1]) >= xIntegerValue(objs[i]) {
			return ec.false_()
		}
	}

	return ec.true_()
}

func (ec *execContext) bareSymbol(sym lispObject) (lispObject, error) {
	// TODO: Is this correct?
	return sym, nil
}

func (ec *execContext) symbolsOfData() {
	ec.defSym(&ec.s.wholeNump, "wholenump")

	ec.defSubr1(nil, "null", (*execContext).null, 1)
	ec.defSubr1(&ec.s.sequencep, "sequencep", (*execContext).sequencep, 1)
	ec.defSubr1(&ec.s.consp, "consp", (*execContext).consp, 1)
	ec.defSubr1(&ec.s.listp, "listp", (*execContext).listp, 1)
	ec.defSubr1(nil, "atom", (*execContext).atom, 1)
	ec.defSubr1(&ec.s.symbolp, "symbolp", (*execContext).symbolp, 1)
	ec.defSubr1(&ec.s.stringp, "stringp", (*execContext).stringp, 1)
	ec.defSubr1(&ec.s.numberOrMarkerp, "number-or-marker-p", (*execContext).numberOrMarkerp, 1)
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
	ec.defSubrM(nil, "<", (*execContext).lessThanSign, 1)
	ec.defSubr1(nil, "bare-symbol", (*execContext).bareSymbol, 1)
}
