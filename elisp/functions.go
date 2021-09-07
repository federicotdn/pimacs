package elisp

import (
	"fmt"
)

func (inp *interpreter) xCar(obj lispObject) lispObject {
	return obj.(*lispCons).car
}

func (inp *interpreter) xCdr(obj lispObject) lispObject {
	return obj.(*lispCons).cdr
}

func (inp *interpreter) cons(car lispObject, cdr lispObject) (lispObject, error) {
	return inp.makeCons(car, cdr), nil
}

func (inp *interpreter) setCdr(obj lispObject, newCdr lispObject) (lispObject, error) {
	obj.(*lispCons).cdr = newCdr

	return newCdr, nil
}

func (inp *interpreter) car(obj lispObject) (lispObject, error) {
	if obj == inp.nil_ {
		return inp.nil_, nil
	}

	cons, ok := obj.(*lispCons)
	if !ok {
		return nil, fmt.Errorf("object is not a cons")
	}
	return cons.car, nil
}

func (inp *interpreter) cdr(obj lispObject) (lispObject, error) {
	if obj == inp.nil_ {
		return inp.nil_, nil
	}

	cons, ok := obj.(*lispCons)
	if !ok {
		return nil, fmt.Errorf("object is not a cons")
	}
	return cons.cdr, nil
}

func (inp *interpreter) plusSign(objs ...lispObject) (lispObject, error) {
	var total lispInt = 0
	for i, obj := range objs {
		number, ok := obj.(*lispInteger)
		if !ok {
			return nil, fmt.Errorf("argument in position %v is not an integer", i)
		}
		total += number.value
	}

	return inp.makeInteger(total), nil
}

func (inp *interpreter) quote(args lispObject) (lispObject, error) {
	cdr, _ := inp.cdr(args)
	if cdr != inp.nil_ {
		return nil, fmt.Errorf("expected exactly 1 argument")
	}

	return inp.car(args)
}

func (inp *interpreter) if_(args lispObject) (lispObject, error) {
	car := inp.xCar(args)
	cdr := inp.xCdr(args)

	cond, err := inp.evalSub(car)
	if err != nil {
		return nil, err
	}

	if cond != inp.nil_ {
		then, err := inp.car(cdr)
		if err != nil {
			return nil, err
		}

		return inp.evalSub(then)
	}

	else_, err := inp.cdr(cdr)
	if err != nil {
		return nil, err
	}

	return inp.progn(else_)
}

func (inp *interpreter) progn(body lispObject) (lispObject, error) {
	var err error
	val := inp.nil_

	for body.getType() == cons {
		form := inp.xCar(body)
		body = inp.xCdr(body)
		val, err = inp.evalSub(form)
		if err != nil {
			return nil, err
		}
	}

	return val, nil
}

func (inp *interpreter) listLength(obj lispObject) (lispObject, error) {
	total := 0

	for obj.getType() == cons {
		total += 1
		obj = inp.xCdr(obj)
	}

	if obj != inp.nil_ {
		return nil, fmt.Errorf("wrong type argument")
	}

	return inp.makeInteger(lispInt(total)), nil
}

func (inp *interpreter) length(obj lispObject) (lispObject, error) {
	num := 0

	switch obj.getType() {
	case string_:
		num = len(obj.(*lispString).value)
	case cons:
		return inp.listLength(obj)
	default:
		if obj != inp.nil_ {
			return nil, fmt.Errorf("wrong type argument")
		}
	}

	return inp.makeInteger(lispInt(num)), nil
}

func (inp *interpreter) initialDefsFunctions() {
	inp.defun1("car", inp.car, 1)
	inp.defun1("length", inp.length, 1)
	inp.defunM("+", inp.plusSign, 0)
	inp.defun2("cons", inp.cons, 2)
	inp.defunU("quote", inp.quote, 1)
	inp.defunU("if", inp.if_, 2)
	inp.defunU("progn", inp.progn, 0)

	// TODO:
	// backquote (`)
}
