package elisp

import (
	"fmt"
)

const (
	eofRune rune = -1
	nbsp         = '\u00A0'
)

func (ec *execContext) applySubroutine(fn, originalArgs lispObject) (lispObject, error) {
	sub := xVectorLike(fn).value.(*subroutine)

	args, err := ec.listToSlice(originalArgs)
	if err != nil {
		return nil, err
	}

	if len(args) < sub.minArgs || (sub.maxArgs >= 0 && len(args) > sub.maxArgs) {
		return ec.wrongNumberOfArguments(fn, lispInt(len(args)))
	}

	var result lispObject
	if sub.maxArgs != argsUnevalled {
		for i := 0; i < len(args); i++ {
			evalled, err := ec.evalSub(args[i])
			if err != nil {
				return nil, err
			}
			args[i] = evalled
		}

		result, err = ec.funcallSubroutine(fn, args...)
	} else {
		// Handle case: maxArgs == argsUnevalled
		result, err = sub.callabe1(originalArgs)
	}

	if result == nil && err == nil {
		ec.terminate("subroutine returned no value and no error")
	}

	if err == nil && sub.noReturn {
		ec.terminate("subroutine with noreturn returned value")
	}

	return result, err
}

func (ec *execContext) funcallSubroutine(fn lispObject, args ...lispObject) (lispObject, error) {
	sub := xVectorLike(fn).value.(*subroutine)

	if sub.maxArgs == argsUnevalled {
		return ec.signalN(ec.g.invalidFunction, fn)
	}

	if len(args) < sub.minArgs || (sub.maxArgs >= 0 && len(args) > sub.maxArgs) {
		return ec.wrongNumberOfArguments(fn, lispInt(len(args)))
	}

	missing := sub.maxArgs - len(args)
	for i := 0; i < missing && sub.maxArgs >= 0; i++ {
		args = append(args, ec.nil_)
	}

	var err error
	var result lispObject
	count := ec.stackSize()

	switch sub.maxArgs {
	case 0:
		result, err = sub.callabe0()
	case 1:
		result, err = sub.callabe1(args[0])
	case 2:
		result, err = sub.callabe2(args[0], args[1])
	case 3:
		result, err = sub.callabe3(args[0], args[1], args[2])
	case 4:
		result, err = sub.callabe4(args[0], args[1], args[2], args[3])
	case 5:
		result, err = sub.callabe5(args[0], args[1], args[2], args[3], args[4])
	case 6:
		result, err = sub.callabe6(args[0], args[1], args[2], args[3], args[4], args[5])
	case 7:
		result, err = sub.callabe7(args[0], args[1], args[2], args[3], args[4], args[5], args[6])
	case 8:
		result, err = sub.callabe8(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7])
	case argsMany:
		result, err = sub.callabem(args...)
	default:
		ec.terminate("unknown subroutine maxargs value")
	}

	if count != ec.stackSize() {
		ec.terminate("subroutine did not pop one or more stack items")
	}

	if err != nil {
		return nil, err
	}

	if sub.noReturn {
		ec.terminate("subroutine with noreturn returned value")
	}

	return result, nil
}

func (ec *execContext) applyLambda(fn, originalArgs lispObject) (lispObject, error) {
	args, err := ec.listToSlice(originalArgs)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(args); i++ {
		evalled, err := ec.evalSub(args[i])
		if err != nil {
			return nil, err
		}
		args[i] = evalled
	}

	return ec.funcallLambda(fn, args...)
}

func (ec *execContext) funcallLambda(fn lispObject, args ...lispObject) (lispObject, error) {
	symsLeft := ec.nil_
	lexEnv := ec.nil_
	defer ec.unwind()()

	if consp(fn) {
		if xCar(fn) == ec.g.closure {
			cdr := xCdr(fn) // Drop 'closure'
			if !consp(cdr) {
				return ec.signalN(ec.g.invalidFunction, fn)
			}

			fn = cdr
			lexEnv = xCar(fn)
		}

		symsLeft = xCdr(fn)
		if !consp(symsLeft) {
			return ec.signalN(ec.g.invalidFunction, fn)
		}

		symsLeft = xCar(symsLeft)
	} else {
		ec.terminate("lambda is not a cons cell")
	}

	rest, optional := false, false
	i := 0

	for ; consp(symsLeft); symsLeft = xCdr(symsLeft) {
		next := xCar(symsLeft)

		if !symbolp(next) {
			return ec.signalN(ec.g.invalidFunction, fn)
		}

		if next == ec.g.andRest {
			if rest {
				return ec.signalN(ec.g.invalidFunction, fn)
			}
			rest = true
		} else if next == ec.g.andOptional {
			if optional {
				return ec.signalN(ec.g.invalidFunction, fn)
			}
			optional = true
		} else {
			var arg lispObject

			if rest {
				arg = ec.makeList(args[i:]...)
				i = len(args)
			} else if i < len(args) {
				arg = args[i]
				i++
			} else if !optional {
				return ec.wrongNumberOfArguments(fn, lispInt(len(args)))
			} else {
				arg = ec.nil_
			}

			if lexEnv != ec.nil_ {
				lexEnv = ec.makeCons(ec.makeCons(next, arg), lexEnv)
			} else {
				ec.stackPushLet(next, arg)
			}
		}
	}

	if symsLeft != ec.nil_ {
		return ec.signalN(ec.g.invalidFunction, fn)
	} else if i < len(args) {
		return ec.wrongNumberOfArguments(fn, lispInt(len(args)))
	}

	if lexEnv != xSymbolValue(ec.g.internalInterpreterEnv) {
		ec.stackPushLet(ec.g.internalInterpreterEnv, lexEnv)
	}

	if consp(fn) {
		return ec.progn(xCdr(xCdr(fn)))
	}

	return ec.pimacsUnimplemented(ec.g.eval, "unknown lambda body type")
}

func (ec *execContext) progIgnore(body lispObject) error {
	_, err := ec.progn(body)
	return err
}

func (ec *execContext) apply(args ...lispObject) (lispObject, error) {
	argsSlice, err := ec.listToSlice(args[len(args)-1])
	if err != nil {
		return nil, err
	}

	finalArgs := append(args[:len(args)-1], argsSlice...)

	return ec.funcall(finalArgs...)
}

func (ec *execContext) funcall(args ...lispObject) (lispObject, error) {
	originalFn := args[0]
	fn := originalFn

	if symbolp(fn) && fn != ec.nil_ {
		fn = xSymbol(fn).function
	}

	if subroutinep(fn) {
		return ec.funcallSubroutine(fn, args[1:]...)
	}

	if fn == ec.nil_ {
		return ec.signalN(ec.g.voidFunction, originalFn)
	}

	if !consp(fn) {
		return ec.signalN(ec.g.invalidFunction, originalFn)
	}

	fnCar := xCar(fn)

	if !symbolp(fnCar) {
		return ec.signalN(ec.g.invalidFunction, originalFn)
	}

	if fnCar == ec.g.lambda || fnCar == ec.g.closure {
		return ec.funcallLambda(fn, args[1:]...)
	}

	return ec.signalN(ec.g.invalidFunction, originalFn)
}

func (ec *execContext) throw(tag, value lispObject) (lispObject, error) {
	if !ec.catchInStack(tag) {
		return ec.signalN(ec.g.noCatch, tag, value)
	}

	return nil, &stackJumpTag{
		tag:   tag,
		value: value,
	}
}

func (ec *execContext) catch(args lispObject) (lispObject, error) {
	defer ec.unwind()()

	tag, err := ec.evalSub(xCar(args))
	if err != nil {
		return nil, err
	}

	ec.stackPushCatch(tag)

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
		return ec.wrongTypeArgument(ec.g.symbolp, variable)
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

	defer ec.unwind()()
	ec.stackPushLet(handlerVar, value)

	result, err = ec.progn(body)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (ec *execContext) signal(errorSymbol, data lispObject) (lispObject, error) {
	// TAGS: incomplete
	return nil, &stackJumpSignal{
		errorSymbol: errorSymbol,
		data:        ec.makeCons(errorSymbol, data),
	}
}

func (ec *execContext) quote(args lispObject) (lispObject, error) {
	cdr := xCdr(args)
	if cdr != ec.nil_ {
		length, err := ec.length(args)
		if err != nil {
			return nil, err
		}
		return ec.wrongNumberOfArguments(ec.g.quote, xIntegerValue(length))
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

func (ec *execContext) while(args lispObject) (lispObject, error) {
	test := xCar(args)
	body := xCdr(args)

	for {
		result, err := ec.evalSub(test)
		if err != nil {
			return nil, err
		}
		if result == ec.nil_ {
			break
		}

		err = ec.progIgnore(body)
		if err != nil {
			return nil, err
		}
	}

	return ec.nil_, nil
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

func (ec *execContext) function(args lispObject) (lispObject, error) {
	quoted := xCar(args)

	if xCdr(args) != ec.nil_ {
		length, err := ec.length(args)
		if err != nil {
			return nil, err
		}
		return ec.wrongNumberOfArguments(ec.g.function, xIntegerValue(length))
	}

	env := xSymbolValue(ec.g.internalInterpreterEnv)
	if env != ec.nil_ &&
		consp(quoted) &&
		xCar(quoted) == ec.g.lambda {
		cdr := xCdr(quoted)

		return ec.makeCons(ec.g.closure, ec.makeCons(env, cdr)), nil
	}

	return quoted, nil
}

func (ec *execContext) eval(form, lexical lispObject) (lispObject, error) {
	defer ec.unwind()()

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

func (ec *execContext) evalSub(form lispObject) (lispObject, error) {
	// TAGS: incomplete
	var err error

	if symbolp(form) {
		var lex lispObject

		envVal := xSymbolValue(ec.g.internalInterpreterEnv)
		if envVal != ec.nil_ {
			lex, err = ec.assq(form, envVal)
			if err != nil {
				return nil, err
			}
		} else {
			lex = ec.nil_
		}

		if lex != ec.nil_ {
			return xCdr(lex), nil
		}

		sym := xSymbol(form)
		if sym.value == ec.g.unbound {
			return ec.signalN(ec.g.voidVariable, form)
		}

		return sym.value, nil
	} else if !consp(form) {
		return form, nil
	}

	originalFn := xCar(form)
	originalArgs := xCdr(form)
	fn := originalFn

	if !symbolp(fn) {
		fn, err = ec.function(ec.makeList(fn))
		if err != nil {
			return nil, err
		}
	} else {
		fn = xSymbol(fn).function
	}

	if subroutinep(fn) {
		return ec.applySubroutine(fn, originalArgs)
	} else if fn == ec.nil_ {
		return ec.signalN(ec.g.voidFunction, originalFn)
	} else if !consp(fn) {
		return ec.signalN(ec.g.invalidFunction, originalFn)
	}

	fnCar := xCar(fn)
	if !symbolp(fnCar) {
		return ec.signalN(ec.g.invalidFunction, originalFn)
	}

	if fnCar == ec.g.macro {
		val := ec.t
		if xSymbolValue(ec.g.internalInterpreterEnv) == ec.nil_ {
			val = ec.nil_
		}

		defer ec.unwind()()
		ec.stackPushLet(ec.g.lexicalBinding, val)

		exp, err := ec.funcall(xCdr(fn), originalArgs)
		if err != nil {
			return nil, err
		}

		return ec.evalSub(exp)
	} else if fnCar == ec.g.lambda || fnCar == ec.g.closure {
		return ec.applyLambda(fn, originalArgs)
	} else {
		return ec.signalN(ec.g.invalidFunction, originalFn)
	}
}

func (ec *execContext) symbolsOfEval() {
	ec.defSubr2("eval", ec.eval, 1)
	ec.defSubrM("funcall", ec.funcall, 1)
	ec.defSubrM("apply", ec.apply, 1)
	ec.defSubrU("progn", ec.progn, 0)
	ec.defSubrU("if", ec.if_, 2)
	ec.defSubrU("while", ec.while, 1)
	ec.defSubrU("quote", ec.quote, 1)
	ec.defSubrU("function", ec.function, 1)
	ec.defSubrU("catch", ec.catch, 1)
	ec.defSubrU("unwind-protect", ec.unwindProtect, 1)
	ec.defSubrU("condition-case", ec.conditionCase, 2)
	ec.defSubr2("throw", ec.throw, 2).noReturn = true
	ec.defSubr2("signal", ec.signal, 2).noReturn = true

	ec.defVar(ec.g.internalInterpreterEnv, ec.nil_)
}
