package elisp

import (
	"fmt"
)

func (ec *execContext) cons_(car lispObject, cdr lispObject) (lispObject, error) {
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
	car_ := ec.xCar(args)
	cdr_ := ec.xCdr(args)

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
		form := ec.xCar(body)
		body = ec.xCdr(body)
		val, err = ec.evalSub(form)
		if err != nil {
			return nil, err
		}
	}

	return val, nil
}

func (ec *execContext) progIgnore(body lispObject) error {
	_, err := ec.progn(body)
	return err
}

func (ec *execContext) listLength(obj lispObject) (lispObject, error) {
	total := 0

	for obj.getType() == lispTypeCons {
		total += 1
		obj = ec.xCdr(obj)
	}

	if obj != ec.nil_ {
		return nil, fmt.Errorf("wrong type argument")
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
			return nil, fmt.Errorf("wrong type argument")
		}
	}

	return ec.makeInteger(lispInt(num)), nil
}

func (ec *execContext) eval(form, lexical lispObject) (lispObject, error) {
	size := ec.bindingsSize()
	defer ec.unbind(size)

	if lexical.getType() != lispTypeCons && lexical != ec.nil_ {
		lexical = ec.makeList(ec.t)
	}

	ec.specBind(ec.env, lexical)
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
		element := ec.xCar(alist)

		if element.getType() == lispTypeCons && ec.xCar(element) == key {
			return element, nil
		}

		alist = ec.xCdr(alist)
	}

	if alist != ec.nil_ {
		return nil, fmt.Errorf("wrong type argument")
	}

	return ec.nil_, nil
}

func (ec *execContext) eq(obj1, obj2 lispObject) (lispObject, error) {
	if obj1 == obj2 {
		return ec.t, nil
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
	tag, err := ec.evalSub(ec.xCar(args))
	if err != nil {
		return nil, err
	}

	obj, err := ec.progn(ec.xCdr(args))
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
	val, err := ec.evalSub(ec.xCar(args))

	unwindErr := ec.progIgnore(ec.xCdr(args))
	if unwindErr != nil {
		// TODO: Check if this is correct
		return nil, unwindErr
	}

	return val, err
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

func (ec *execContext) initialDefsFunctions() {
	ec.defSubr1("car", ec.car, 1)
	ec.defSubr1("cdr", ec.cdr, 1)
	ec.defSubr1("length", ec.length, 1)
	ec.defSubr2("cons", ec.cons_, 2)
	ec.defSubr2("eval", ec.eval, 1)
	ec.defSubr2("set", ec.set, 2)
	ec.defSubr2("fset", ec.fset, 2)
	ec.defSubr2("eq", ec.eq, 2)
	ec.defSubr2("assq", ec.assq, 2)
	ec.defSubr2("prin1", ec.prin1, 1)
	ec.defSubr2("throw", ec.throw, 2)
	ec.defSubrM("+", ec.plusSign, 0)
	ec.defSubrU("catch", ec.catch, 1)
	ec.defSubrU("unwind-protect", ec.unwindProtect, 1)
	ec.defSubrU("quote", ec.quote, 1)
	ec.defSubrU("if", ec.if_, 2)
	ec.defSubrU("progn", ec.progn, 0)

	// TODO:
	// backquote (`)
}
