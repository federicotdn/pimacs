package elisp

import (
	"fmt"
)

func (ec *execContext) sequencep(object lispObject) (lispObject, error) {
	if consp(object) || object == ec.nil_ || arrayp(object) {
		return ec.t, nil
	}

	return ec.nil_, nil
}

func (ec *execContext) cons(car lispObject, cdr lispObject) (lispObject, error) {
	return ec.makeCons(car, cdr), nil
}

func (ec *execContext) car(obj lispObject) (lispObject, error) {
	if obj == ec.nil_ {
		return ec.nil_, nil
	}

	if !consp(obj) {
		return nil, fmt.Errorf("object is not a cons")
	}
	return xCons(obj).car, nil
}

func (ec *execContext) cdr(obj lispObject) (lispObject, error) {
	if obj == ec.nil_ {
		return ec.nil_, nil
	}

	if !consp(obj) {
		return nil, fmt.Errorf("object is not a cons")
	}
	return xCons(obj).cdr, nil
}

func (ec *execContext) plusSign(objs ...lispObject) (lispObject, error) {
	var total lispInt = 0
	for i, obj := range objs {
		if !integerp(obj) {
			return nil, fmt.Errorf("argument in position %v is not an integer", i)
		}
		total += xInteger(obj).value
	}

	return ec.makeInteger(total), nil
}

func (ec *execContext) quote(args lispObject) (lispObject, error) {
	cdr := xCdr(args)
	if cdr != ec.nil_ {
		return nil, fmt.Errorf("expected exactly 1 argument")
	}

	return xCar(args), nil
}

func (ec *execContext) if_(args lispObject) (lispObject, error) {
	car_ := xCar(args)
	cdr_ := xCdr(args)

	cond, err := ec.evalSub(car_)
	if err != nil {
		return nil, err
	}

	if cond != ec.nil_ {
		return ec.evalSub(xCar(cdr_))
	}

	else_, err := ec.cdr(cdr_)
	if err != nil {
		return nil, err
	}

	return ec.progn(else_)
}

func (ec *execContext) progn(body lispObject) (lispObject, error) {
	val := ec.nil_

	for ; consp(body); body = xCdr(body) {
		var err error

		form := xCar(body)
		val, err = ec.evalSub(form)
		if err != nil {
			return nil, err
		}
	}

	return val, nil
}

func (ec *execContext) length(obj lispObject) (lispObject, error) {
	num := 0

	switch obj.getType() {
	case lispTypeString:
		num = len(xString(obj).value)
	case lispTypeCons:
		for ; consp(obj); obj = xCdr(obj) {
			num += 1
		}

		if obj != ec.nil_ {
			return nil, fmt.Errorf("length: wrong type argument")
		}
	default:
		if obj != ec.nil_ {
			return nil, fmt.Errorf("length: wrong type argument")
		}
	}

	return ec.makeInteger(lispInt(num)), nil
}

func (ec *execContext) eval(form, lexical lispObject) (lispObject, error) {
	size := ec.stackSize()
	defer ec.unbind(size)

	if !consp(lexical) && lexical != ec.nil_ {
		lexical = ec.makeList(ec.t)
	}

	ec.stackPushLet(ec.g.internalInterpreterEnv, lexical)
	val, err := ec.evalSub(form)
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (ec *execContext) set(symbol, newVal lispObject) (lispObject, error) {
	if !symbolp(symbol) {
		return nil, fmt.Errorf("not a symbol")
	}

	xSymbol(symbol).value = newVal
	return newVal, nil
}

func (ec *execContext) fset(symbol, definition lispObject) (lispObject, error) {
	if symbol == ec.nil_ && definition != ec.nil_ {
		return nil, fmt.Errorf("setting constant nil")
	}

	if !symbolp(symbol) {
		return nil, fmt.Errorf("not a symbol")
	}

	xSymbol(symbol).function = definition
	return definition, nil
}

func (ec *execContext) assq(key, alist lispObject) (lispObject, error) {
	for ; consp(alist); alist = xCdr(alist) {
		element := xCar(alist)

		if consp(element) && xCar(element) == key {
			return element, nil
		}
	}

	if alist != ec.nil_ {
		return nil, fmt.Errorf("assq: wrong type argument %v", alist.getType())
	}

	return ec.nil_, nil
}

func (ec *execContext) memq(elt, list lispObject) (lispObject, error) {
	for ; consp(list); list = xCdr(list) {
		if xCar(list) == elt {
			return list, nil
		}
	}

	if list != ec.nil_ {
		return nil, fmt.Errorf("memq: wrong type argument")
	}

	return ec.g.nil_, nil
}

func (ec *execContext) eq(obj1, obj2 lispObject) (lispObject, error) {
	if obj1 == obj2 {
		return ec.t, nil
	}
	return ec.nil_, nil
}

func (ec *execContext) throw(tag, value lispObject) (lispObject, error) {
	if !ec.catchInStack(tag) {
		// TODO: Revise
		return ec.signal(ec.g.noCatch, ec.makeList(tag, value))
	}

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

	size := ec.stackSize()
	ec.stackPushCatch(tag)
	defer ec.unbind(size)

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

func (ec *execContext) clauseHandlesConditions(clause, errConditions lispObject) (bool, error) {
	clauseConditions := xCar(clause)

	if !consp(clauseConditions) {
		clauseConditions = ec.makeList(clauseConditions)
	}

	for cond := clauseConditions; consp(cond); cond = xCdr(cond) {
		condName := xCar(cond)
		found, err := ec.memq(condName, errConditions)
		if err != nil {
			return false, err
		}

		if found != ec.nil_ || condName == ec.t {
			return true, nil
		}
	}

	return false, nil
}

func (ec *execContext) conditionCase(args lispObject) (lispObject, error) {
	variable := xCar(args)
	bodyForm := xCar(xCdr(args))
	handlers := xCdr(xCdr(args))

	clauses := []lispObject{}

	if !symbolp(variable) {
		return nil, fmt.Errorf("wrong type argument: var")
	}

	for tail := handlers; consp(tail); tail = xCdr(tail) {
		tem := xCar(tail)
		clauses = append(clauses, tem)

		if tem == ec.nil_ {
			continue
		}

		if consp(tem) && (symbolp(xCar(tem)) || consp(xCar(tem))) {
			continue
		}

		return nil, fmt.Errorf("invalid type handler")
	}

	result, err := ec.evalSub(bodyForm)
	if err == nil {
		// No error, return result as normal
		return result, nil
	}

	jmp, ok := err.(*stackJumpSignal)
	if !ok {
		// There's an error, but it's a throw - don't handle that here
		return nil, err
	}

	errConditions, err := ec.get(jmp.errorSymbol, ec.g.errorConditions)
	if err != nil {
		return nil, err
	}

	var body lispObject

	for _, clause := range clauses {
		doesHandle, err := ec.clauseHandlesConditions(clause, errConditions)
		if err != nil {
			return nil, err
		}

		if doesHandle {
			body = xCdr(clause)
			break
		}
	}

	if body == nil {
		// No handler was chosen
		return nil, jmp
	}

	if variable == ec.nil_ {
		return ec.progn(body)
	}

	value := jmp.data
	handlerVar := variable
	env := xSymbolValue(ec.g.internalInterpreterEnv)

	if env != ec.nil_ {
		value = ec.makeCons(ec.makeCons(variable, value), env)
		handlerVar = ec.g.internalInterpreterEnv
	}

	size := ec.stackSize()
	defer ec.unbind(size)
	ec.stackPushLet(handlerVar, value)

	result, err = ec.progn(body)

	if err != nil {
		return nil, err
	}

	return result, nil

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
		s = xSymbol(obj).name
	case lispTypeInteger:
		s = fmt.Sprint(xInteger(obj).value)
	case lispTypeString:
		s = "\"" + xString(obj).value + "\""
	case lispTypeCons:
		// TODO: Clean up when string functions are available (don't type-assert)
		lc := xCons(obj)
		current := lc

		carStr, err := ec.prin1(lc.car, ec.nil_)
		if err != nil {
			return nil, err
		}
		s = "(" + xString(carStr).value

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
		s = fmt.Sprint(xFloat(obj).value)
	case lispTypeVectorLike:
		s = "<vector-like>"
	default:
		panic("unknown type")
	}

	return ec.makeString(s), nil
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

func (ec *execContext) symbolPlist(symbol lispObject) (lispObject, error) {
	if !symbolp(symbol) {
		return nil, fmt.Errorf("wrong type error")
	}

	return xSymbol(symbol).plist, nil
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

func (ec *execContext) list(args ...lispObject) (lispObject, error) {
	return ec.makeList(args...), nil
}

func (ec *execContext) initialDefsFunctions() {
	ec.defSubr1("sequencep", ec.sequencep, 1)
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
	ec.defSubr2("memq", ec.memq, 2)
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
