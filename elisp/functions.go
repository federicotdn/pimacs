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

	for body.getType() == cons {
		form := ec.xCar(body)
		body = ec.xCdr(body)
		val, err = ec.evalSub(form)
		if err != nil {
			return nil, err
		}
	}

	return val, nil
}

func (ec *execContext) listLength(obj lispObject) (lispObject, error) {
	total := 0

	for obj.getType() == cons {
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
	case string_:
		num = len(obj.(*lispString).value)
	case cons:
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

	if lexical.getType() != cons && lexical != ec.nil_ {
		lexical = ec.makeList(ec.t)
	}

	ec.specBind(ec.env, lexical)
	val, err := ec.evalSub(form)

	return ec.unbindTo(size, val, err)
}

func (ec *execContext) set(symbol, newVal lispObject) (lispObject, error) {
	return ec.setInternal(symbol, newVal)
}

func (ec *execContext) assq(key, alist lispObject) (lispObject, error) {
	for alist.getType() == cons {
		element := ec.xCar(alist)

		if element.getType() == cons && ec.xCar(element) == key {
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

func (ec *execContext) initialDefsFunctions() {
	ec.defun1("car", ec.car, 1)
	ec.defun1("cdr", ec.cdr, 1)
	ec.defun1("length", ec.length, 1)
	ec.defun2("cons", ec.cons_, 2)
	ec.defun2("eval", ec.eval, 1)
	ec.defun2("set", ec.set, 2)
	ec.defun2("eq", ec.eq, 2)
	ec.defun2("assq", ec.assq, 2)
	ec.defunM("+", ec.plusSign, 0)
	ec.defunU("quote", ec.quote, 1)
	ec.defunU("if", ec.if_, 2)
	ec.defunU("progn", ec.progn, 0)

	// TODO:
	// backquote (`)
}
