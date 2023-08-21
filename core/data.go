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

func (ec *execContext) bufferp(object lispObject) (lispObject, error) {
	return ec.bool(bufferp(object))
}

func (ec *execContext) boundp(symbol lispObject) (lispObject, error) {
	if !symbolp(symbol) {
		return ec.wrongTypeArgument(ec.g.symbolp, symbol)
	}
	return ec.bool(xSymbolValue(symbol) != ec.g.unbound)
}

func (ec *execContext) fboundp(symbol lispObject) (lispObject, error) {
	if !symbolp(symbol) {
		return ec.wrongTypeArgument(ec.g.symbolp, symbol)
	}
	return ec.bool(xSymbol(symbol).function != ec.g.unbound)
}

func (ec *execContext) makunbound(symbol lispObject) (lispObject, error) {
	if !symbolp(symbol) {
		return ec.wrongTypeArgument(ec.g.symbolp, symbol)
	}
	xSymbol(symbol).value = ec.g.unbound
	return symbol, nil
}

func (ec *execContext) fmakunbound(symbol lispObject) (lispObject, error) {
	if !symbolp(symbol) {
		return ec.wrongTypeArgument(ec.g.symbolp, symbol)
	}
	xSymbol(symbol).function = ec.g.unbound
	return symbol, nil
}

func (ec *execContext) car(obj lispObject) (lispObject, error) {
	if obj == ec.nil_ {
		return ec.nil_, nil
	}

	if !consp(obj) {
		return ec.wrongTypeArgument(ec.g.listp, obj)
	}
	return xCar(obj), nil
}

func (ec *execContext) cdr(obj lispObject) (lispObject, error) {
	if obj == ec.nil_ {
		return ec.nil_, nil
	}

	if !consp(obj) {
		return ec.wrongTypeArgument(ec.g.listp, obj)
	}
	return xCdr(obj), nil
}

func (ec *execContext) setCar(obj, newCar lispObject) (lispObject, error) {
	if !consp(obj) {
		return ec.wrongTypeArgument(ec.g.consp, obj)
	}
	xSetCar(obj, newCar)
	return newCar, nil
}

func (ec *execContext) setCdr(obj, newCdr lispObject) (lispObject, error) {
	if !consp(obj) {
		return ec.wrongTypeArgument(ec.g.consp, obj)
	}
	xSetCdr(obj, newCdr)
	return newCdr, nil
}

func (ec *execContext) set(symbol, newVal lispObject) (lispObject, error) {
	if !symbolp(symbol) {
		return ec.wrongTypeArgument(ec.g.symbolp, symbol)
	}

	// TODO: Handle symbol value forwarding
	xSymbol(symbol).value = newVal
	return newVal, nil
}

func (ec *execContext) fset(symbol, definition lispObject) (lispObject, error) {
	if symbol == ec.nil_ && definition != ec.nil_ {
		return ec.signalN(ec.g.settingConstant, symbol)
	}

	if !symbolp(symbol) {
		return ec.wrongTypeArgument(ec.g.symbolp, symbol)
	}

	xSymbol(symbol).function = definition
	return definition, nil
}

func (ec *execContext) symbolValue(symbol lispObject) (lispObject, error) {
	if !symbolp(symbol) {
		return ec.wrongTypeArgument(ec.g.symbolp, symbol)
	}

	sym := xSymbol(symbol)
	var val lispObject
	if sym.redirect != nil {
		switch sym.redirect.fwdType {
		case symbolFwdTypeLispObj:
			val = sym.redirect.valLisp
		case symbolFwdTypeInt:
			val = ec.makeInteger(sym.redirect.valInt)
		case symbolFwdTypeBool:
			val = xEnsure(ec.bool(sym.redirect.valBool))
		default:
			ec.terminate("unknown symbol forward value type: '%+v'", sym.redirect.fwdType)
		}
	} else {
		val = sym.value
	}

	if val == ec.g.unbound {
		return ec.signalN(ec.g.voidVariable, symbol)
	}

	return val, nil
}

func (ec *execContext) symbolFunction(symbol lispObject) (lispObject, error) {
	if !symbolp(symbol) {
		return ec.wrongTypeArgument(ec.g.symbolp, symbol)
	}

	return xSymbol(symbol).function, nil
}

func (ec *execContext) eq(obj1, obj2 lispObject) (lispObject, error) {
	return ec.bool(obj1 == obj2)
}

func (ec *execContext) defalias(symbol, definition, docstring lispObject) (lispObject, error) {
	// TAGS: incomplete
	return ec.fset(symbol, definition)
}

func (ec *execContext) symbolPlist(symbol lispObject) (lispObject, error) {
	if !symbolp(symbol) {
		return ec.wrongTypeArgument(ec.g.symbolp, symbol)
	}

	return xSymbol(symbol).plist, nil
}

func (ec *execContext) symbolName(symbol lispObject) (lispObject, error) {
	if !symbolp(symbol) {
		return ec.wrongTypeArgument(ec.g.symbolp, symbol)
	}

	return ec.makeString(xSymbol(symbol).name), nil
}

func (ec *execContext) plusSign(objs ...lispObject) (lispObject, error) {
	// TAGS: incomplete
	var total lispInt = 0
	for _, obj := range objs {
		if !numberp(obj) {
			return ec.wrongTypeArgument(ec.g.numberOrMarkerp, obj)
		}
		total += xIntegerValue(obj)
	}

	return ec.makeInteger(total), nil
}

func (ec *execContext) lessThanSign(objs ...lispObject) (lispObject, error) {
	// TAGS: incomplete
	for i := 1; i < len(objs); i++ {
		if !numberp(objs[i-1]) {
			return ec.wrongTypeArgument(ec.g.numberOrMarkerp, objs[i-1])
		} else if !numberp(objs[i]) {
			return ec.wrongTypeArgument(ec.g.numberOrMarkerp, objs[i])
		}

		if xIntegerValue(objs[i-1]) >= xIntegerValue(objs[i]) {
			return ec.false_()
		}
	}

	return ec.true_()
}

func (ec *execContext) symbolsOfData() {
	ec.defSubr1(nil, "null", ec.null, 1)
	ec.defSubr1(&ec.g.sequencep, "sequencep", ec.sequencep, 1)
	ec.defSubr1(&ec.g.consp, "consp", ec.consp, 1)
	ec.defSubr1(&ec.g.listp, "listp", ec.listp, 1)
	ec.defSubr1(&ec.g.symbolp, "symbolp", ec.symbolp, 1)
	ec.defSubr1(&ec.g.stringp, "stringp", ec.stringp, 1)
	ec.defSubr1(&ec.g.numberOrMarkerp, "number-or-marker-p", ec.numberOrMarkerp, 1)
	ec.defSubr1(&ec.g.charOrStringp, "char-or-string-p", ec.charOrStringp, 1)
	ec.defSubr1(&ec.g.integerp, "integerp", ec.numberOrMarkerp, 1)
	ec.defSubr1(&ec.g.bufferp, "bufferp", ec.bufferp, 1)
	ec.defSubr1(nil, "boundp", ec.boundp, 1)
	ec.defSubr1(nil, "fboundp", ec.fboundp, 1)
	ec.defSubr1(nil, "makunbound", ec.makunbound, 1)
	ec.defSubr1(nil, "fmakunbound", ec.fmakunbound, 1)
	ec.defSubr1(nil, "car", ec.car, 1)
	ec.defSubr1(nil, "cdr", ec.cdr, 1)
	ec.defSubr2(nil, "setcar", ec.setCar, 2)
	ec.defSubr2(nil, "setcdr", ec.setCdr, 2)
	ec.defSubr1(nil, "symbol-plist", ec.symbolPlist, 1)
	ec.defSubr1(nil, "symbol-name", ec.symbolName, 1)
	ec.defSubr2(nil, "set", ec.set, 2)
	ec.defSubr2(nil, "fset", ec.fset, 2)
	ec.defSubr1(nil, "symbol-value", ec.symbolValue, 1)
	ec.defSubr1(nil, "symbol-function", ec.symbolFunction, 1)
	ec.defSubr2(nil, "eq", ec.eq, 2)
	ec.defSubr3(nil, "defalias", ec.defalias, 2)
	ec.defSubrM(nil, "+", ec.plusSign, 0)
	ec.defSubrM(nil, "<", ec.lessThanSign, 1)
}
