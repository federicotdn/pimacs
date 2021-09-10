package elisp

import (
	"fmt"
)

func (ec *execContext) cons(car lispObject, cdr lispObject) (lispObject, error) {
	return ec.makeCons(car, cdr), nil
}

func (ec *execContext) car(obj lispObject) (lispObject, error) {
	if obj == ec.nil_ {
		return ec.nil_, nil
	}

	cons, ok := obj.(*lispCons)
	if !ok {
		return nil, fmt.Errorf("object is not a cons")
	}
	return cons.car, nil
}

func (ec *execContext) cdr(obj lispObject) (lispObject, error) {
	if obj == ec.nil_ {
		return ec.nil_, nil
	}

	cons, ok := obj.(*lispCons)
	if !ok {
		return nil, fmt.Errorf("object is not a cons")
	}
	return cons.cdr, nil
}

func (ec *execContext) plusSign(objs ...lispObject) (lispObject, error) {
	var total lispInt = 0
	for i, obj := range objs {
		number, ok := obj.(*lispInteger)
		if !ok {
			return nil, fmt.Errorf("argument in position %v is not an integer", i)
		}
		total += number.value
	}

	return ec.makeInteger(total), nil
}

func (ec *execContext) quote(args lispObject) (lispObject, error) {
	cdr, _ := ec.cdr(args)
	if cdr != ec.nil_ {
		return nil, fmt.Errorf("expected exactly 1 argument")
	}

	return ec.car(args)
}

func (ec *execContext) if_(args lispObject) (lispObject, error) {
	car_ := xCar(args)
	cdr_ := xCdr(args)

	cond, err := ec.evalSub(car_)
	if err != nil {
		return nil, err
	}

	if cond != ec.nil_ {
		then, err := ec.car(cdr_)
		if err != nil {
			return nil, err
		}

		return ec.evalSub(then)
	}

	else_, err := ec.cdr(cdr_)
	if err != nil {
		return nil, err
	}

	return ec.progn(else_)
}

func (ec *execContext) progn(body lispObject) (lispObject, error) {
	var err error
	val := ec.nil_

	for body.getType() == lispTypeCons {
		form := xCar(body)
		body = xCdr(body)
		val, err = ec.evalSub(form)
		if err != nil {
			return nil, err
		}
	}

	return val, nil
}

func (ec *execContext) listLength(obj lispObject) (lispObject, error) {
	total := 0

	for obj.getType() == lispTypeCons {
		total += 1
		obj = xCdr(obj)
	}

	if obj != ec.nil_ {
		return nil, fmt.Errorf("length: wrong type argument")
	}

	return ec.makeInteger(lispInt(total)), nil
}

func (ec *execContext) length(obj lispObject) (lispObject, error) {
	num := 0

	switch obj.getType() {
	case lispTypeStr:
		num = len(obj.(*lispString).value)
	case lispTypeCons:
		return ec.listLength(obj)
	default:
		if obj != ec.nil_ {
			return nil, fmt.Errorf("length: wrong type argument")
		}
	}

	return ec.makeInteger(lispInt(num)), nil
}

func (ec *execContext) eval(form, lexical lispObject) (lispObject, error) {
	size := ec.bindingsSize()
	defer ec.unbind(size)

	if lexical.getType() != lispTypeCons && lexical != ec.nil_ {
		lexical = ec.makeList(ec.globals.t)
	}

	ec.specBind(ec.globals.internalInterpreterEnv.(*lispSymbol), lexical)
	val, err := ec.evalSub(form)
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (ec *execContext) set(symbol, newVal lispObject) (lispObject, error) {
	sym, ok := symbol.(*lispSymbol)
	if !ok {
		return nil, fmt.Errorf("not a symbol")
	}

	sym.value = newVal
	return newVal, nil
}

func (ec *execContext) fset(symbol, definition lispObject) (lispObject, error) {
	if symbol == ec.nil_ && definition != ec.nil_ {
		return nil, fmt.Errorf("setting constant nil")
	}

	if symbol.getType() != lispTypeSymbol {
		return nil, fmt.Errorf("not a symbol")
	}

	symbol.(*lispSymbol).function = definition
	return definition, nil
}

func (ec *execContext) assq(key, alist lispObject) (lispObject, error) {
	for alist.getType() == lispTypeCons {
		element := xCar(alist)

		if element.getType() == lispTypeCons && xCar(element) == key {
			return element, nil
		}

		alist = xCdr(alist)
	}

	if alist != ec.nil_ {
		return nil, fmt.Errorf("assq: wrong type argument %v", alist.getType())
	}

	return ec.nil_, nil
}

func (ec *execContext) eq(obj1, obj2 lispObject) (lispObject, error) {
	if obj1 == obj2 {
		return ec.globals.t, nil
	}
	return ec.nil_, nil
}

func (ec *execContext) throw(tag, value lispObject) (lispObject, error) {
	return nil, &stackJumpTag{
		tag:   tag,
		value: value,
	}
}

func (ec *execContext) catch(args lispObject) (lispObject, error) {
	tag, err := ec.evalSub(xCar(args))
	if err != nil {
		return nil, err
	}

	obj, err := ec.progn(xCdr(args))
	if err != nil {
		jmp, ok := err.(*stackJumpTag)
		if !ok {
			return nil, err
		}

		if jmp.tag == tag {
			return jmp.value, nil
		}

		return nil, err
	}
	return obj, nil
}

func (ec *execContext) unwindProtect(args lispObject) (lispObject, error) {
	val, err := ec.evalSub(xCar(args))

	unwindErr := ec.progIgnore(xCdr(args))
	if unwindErr != nil {
		// TODO: Check if this is correct
		return nil, unwindErr
	}

	return val, err
}

func (ec *execContext) conditionCase(args lispObject) (lispObject, error) {
	variable := xCar(args)
	bodyForm := xCar(xCdr(args))
	// handlers := xCdr(xCdr(args))

	if variable.getType() != lispTypeSymbol {
		return nil, fmt.Errorf("wrong type argument: var")
	}

	// TODO: check structure of handlers

	result, err := ec.evalSub(bodyForm)
	if err == nil {
		return result, nil
	}

	jmp, ok := err.(*stackJumpSignal)
	if !ok {
		return nil, err
	}

	return nil, jmp
}

func (ec *execContext) signal(errorSymbol, data lispObject) (lispObject, error) {
	return nil, &stackJumpSignal{
		errorSymbol: errorSymbol,
		data:        data,
	}
}

func (ec *execContext) prin1(obj, printCharFun lispObject) (lispObject, error) {
	lispType := obj.getType()
	var s string

	switch lispType {
	case lispTypeSymbol:
		s = obj.(*lispSymbol).name
	case lispTypeInt:
		s = fmt.Sprint(obj.(*lispInteger).value)
	case lispTypeStr:
		s = "\"" + obj.(*lispString).value + "\""
	case lispTypeVecLike:
		return nil, fmt.Errorf("prin1 unimplemented")
	case lispTypeCons:
		// TODO: Clean up when string functions are avaiable (don't type-assert)
		lc := obj.(*lispCons)
		current := lc

		carStr, err := ec.prin1(lc.car, ec.nil_)
		if err != nil {
			return nil, err
		}
		s = "(" + carStr.(*lispString).value

		for {
			next, ok := current.cdr.(*lispCons)
			if ok {
				nextStr, err := ec.prin1(next.car, ec.nil_)
				if err != nil {
					return nil, err
				}

				s += " " + nextStr.(*lispString).value
				current = next
			} else {
				if current.cdr != ec.nil_ {
					cdrStr, err := ec.prin1(current.cdr, ec.nil_)
					if err != nil {
						return nil, err
					}
					s += " . " + cdrStr.(*lispString).value
				}

				break
			}
		}

		s += ")"
	case lispTypeFloat:
		s = fmt.Sprint(obj.(*lispFloat).value)
	default:
		return nil, fmt.Errorf("prin1 unimplemented for '%v'", lispType)
	}

	return ec.makeString(s), nil
}

func (ec *execContext) plistPut(plist, prop, val lispObject) (lispObject, error) {
	prev := ec.nil_
	tail := plist

	for tail.getType() == lispTypeCons {
		next := xCdr(tail)
		if next.getType() != lispTypeCons {
			break
		}

		if prop == xCar(tail) {
			xSetCar(next, val)
			return plist, nil
		}

		prev = tail
		tail = xCdr(next)
	}

	if tail != ec.nil_ {
		return nil, fmt.Errorf("plist-put: wrong type argument")
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

	for tail.getType() == lispTypeCons {
		next := xCdr(tail)

		if next.getType() != lispTypeCons {
			break
		} else if prop == xCar(tail) {
			return xCar(next), nil
		}

		tail = xCdr(next)
	}

	return ec.nil_, nil
}

func (ec *execContext) symbolPlist(symbol lispObject) (lispObject, error) {
	if symbol.getType() != lispTypeSymbol {
		return nil, fmt.Errorf("wrong type error")
	}

	return symbol.(*lispSymbol).plist, nil
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

	symbol.(*lispSymbol).plist = plist
	return value, nil
}

func (ec *execContext) list(args ...lispObject) (lispObject, error) {
	return ec.makeList(args...), nil
}

func (ec *execContext) initialDefsFunctions() {
	ec.defSubr1("car", ec.car, 1)
	ec.defSubr1("cdr", ec.cdr, 1)
	ec.defSubr1("length", ec.length, 1)
	ec.defSubr1("symbol-plist", ec.symbolPlist, 1)
	ec.defSubr2("cons", ec.cons, 2)
	ec.defSubr2("eval", ec.eval, 1)
	ec.defSubr2("set", ec.set, 2)
	ec.defSubr2("fset", ec.fset, 2)
	ec.defSubr2("eq", ec.eq, 2)
	ec.defSubr2("assq", ec.assq, 2)
	ec.defSubr2("get", ec.get, 2)
	ec.defSubr2("prin1", ec.prin1, 1)
	ec.defSubr2("throw", ec.throw, 2).noReturn = true
	ec.defSubr2("signal", ec.signal, 2).noReturn = true
	ec.defSubr2("plist-get", ec.plistGet, 2)
	ec.defSubr3("plist-put", ec.plistPut, 3)
	ec.defSubr3("put", ec.put, 3)
	ec.defSubrM("+", ec.plusSign, 0)
	ec.defSubrM("list", ec.list, 0)
	ec.defSubrU("catch", ec.catch, 1)
	ec.defSubrU("unwind-protect", ec.unwindProtect, 1)
	ec.defSubrU("quote", ec.quote, 1)
	ec.defSubrU("if", ec.if_, 2)
	ec.defSubrU("progn", ec.progn, 0)
	ec.defSubrU("condition-case", ec.conditionCase, 2)
}
