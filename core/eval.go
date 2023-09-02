package core

import (
	"fmt"
	"runtime"
)

const (
	eofRune rune = -1
	nbsp         = '\u00A0'
)

func (ec *execContext) applySubroutine(fn, originalArgs lispObject) (lispObject, error) {
	sub := xSubroutine(fn)

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
		result, err = sub.callabe.(lispFn1)(originalArgs)
	}

	if result == nil && err == nil {
		ec.terminate("subroutine returned no value and no error: '%+v'", sub.name)
	}

	if err == nil && sub.noReturn {
		ec.terminate("subroutine with noreturn returned value: '%+v'", sub.name)
	}

	return result, err
}

func (ec *execContext) funcallSubroutine(fn lispObject, args ...lispObject) (lispObject, error) {
	sub := xSubroutine(fn)

	if sub.maxArgs == argsUnevalled {
		return ec.signalN(ec.s.invalidFunction, fn)
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
		result, err = sub.callabe.(lispFn0)()
	case 1:
		result, err = sub.callabe.(lispFn1)(args[0])
	case 2:
		result, err = sub.callabe.(lispFn2)(args[0], args[1])
	case 3:
		result, err = sub.callabe.(lispFn3)(args[0], args[1], args[2])
	case 4:
		result, err = sub.callabe.(lispFn4)(args[0], args[1], args[2], args[3])
	case 5:
		result, err = sub.callabe.(lispFn5)(args[0], args[1], args[2], args[3], args[4])
	case 6:
		result, err = sub.callabe.(lispFn6)(args[0], args[1], args[2], args[3], args[4], args[5])
	case 7:
		result, err = sub.callabe.(lispFn7)(args[0], args[1], args[2], args[3], args[4], args[5], args[6])
	case 8:
		result, err = sub.callabe.(lispFn8)(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7])
	case argsMany:
		result, err = sub.callabe.(lispFnM)(args...)
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
		if xCar(fn) == ec.s.closure {
			cdr := xCdr(fn) // Drop 'closure'
			if !consp(cdr) {
				return ec.signalN(ec.s.invalidFunction, fn)
			}

			fn = cdr
			lexEnv = xCar(fn)
		}

		symsLeft = xCdr(fn)
		if !consp(symsLeft) {
			return ec.signalN(ec.s.invalidFunction, fn)
		}

		symsLeft = xCar(symsLeft)
	} else {
		ec.terminate("lambda is not a cons cell")
	}

	rest, optional := false, false
	i := 0

	iter := ec.iterate(symsLeft)
	for ; iter.hasNext(); symsLeft = iter.nextCons() {
		next := xCar(symsLeft)

		if !symbolp(next) {
			return ec.signalN(ec.s.invalidFunction, fn)
		}

		if next == ec.s.andRest {
			if rest {
				return ec.signalN(ec.s.invalidFunction, fn)
			}
			rest = true
		} else if next == ec.s.andOptional {
			if optional {
				return ec.signalN(ec.s.invalidFunction, fn)
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

	if iter.circular() {
		return iter.error()
	}

	if symsLeft != ec.nil_ {
		return ec.signalN(ec.s.invalidFunction, fn)
	} else if i < len(args) {
		return ec.wrongNumberOfArguments(fn, lispInt(len(args)))
	}

	if lexEnv != ec.v.internalInterpreterEnv.val {
		ec.stackPushLet(ec.v.internalInterpreterEnv.sym, lexEnv)
	}

	if !consp(fn) {
		return ec.pimacsUnimplemented(ec.s.eval, "unknown lambda body type")
	}

	return ec.progn(xCdr(xCdr(fn)))
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
		return ec.signalN(ec.s.voidFunction, originalFn)
	}

	if !consp(fn) {
		return ec.signalN(ec.s.invalidFunction, originalFn)
	}

	fnCar := xCar(fn)

	if !symbolp(fnCar) {
		return ec.signalN(ec.s.invalidFunction, originalFn)
	}

	if fnCar == ec.s.lambda || fnCar == ec.s.closure {
		return ec.funcallLambda(fn, args[1:]...)
	}

	return ec.signalN(ec.s.invalidFunction, originalFn)
}

func (ec *execContext) throw(tag, value lispObject) (lispObject, error) {
	if !ec.catchInStack(tag) {
		return ec.signalN(ec.s.noCatch, tag, value)
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

	iter := ec.iterate(clauseConditions)
	for ; iter.hasNext(); clauseConditions = iter.nextCons() {
		condName := xCar(clauseConditions)
		found, err := ec.memq(condName, errConditions)
		if err != nil {
			return false, err
		}

		if found != ec.nil_ || condName == ec.t {
			return true, nil
		}
	}

	if iter.hasError() {
		return false, xErrOnly(iter.error())
	}

	return false, nil
}

func (ec *execContext) conditionCase(args lispObject) (lispObject, error) {
	variable := xCar(args)
	bodyForm := xCar(xCdr(args))
	handlers := xCdr(xCdr(args))

	clauses := []lispObject{}

	if !symbolp(variable) {
		return ec.wrongTypeArgument(ec.s.symbolp, variable)
	}

	iter := ec.iterate(handlers)
	for ; iter.hasNext(); handlers = iter.nextCons() {
		tem := xCar(handlers)
		clauses = append(clauses, tem)

		if tem == ec.nil_ {
			continue
		}

		if consp(tem) && (symbolp(xCar(tem)) || consp(xCar(tem))) {
			continue
		}

		return nil, fmt.Errorf("invalid type handler")
	}

	if iter.hasError() {
		return iter.error()
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

	errConditions, err := ec.get(jmp.errorSymbol, ec.s.errorConditions)
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
	env := ec.v.internalInterpreterEnv.val

	if env != ec.nil_ {
		value = ec.makeCons(ec.makeCons(variable, value), env)
		handlerVar = ec.v.internalInterpreterEnv.sym
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
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)

	if errorSymbol == ec.nil_ && data == ec.nil_ {
		errorSymbol = ec.s.error_
	}

	return nil, &stackJumpSignal{
		errorSymbol: errorSymbol,
		data:        ec.makeCons(errorSymbol, data),
		goStack:     string(buf[:n]),
	}
}

func (ec *execContext) quote(args lispObject) (lispObject, error) {
	cdr := xCdr(args)
	if cdr != ec.nil_ {
		length, err := ec.length(args)
		if err != nil {
			return nil, err
		}
		return ec.wrongNumberOfArguments(ec.s.quote, xIntegerValue(length))
	}

	return xCar(args), nil
}

func (ec *execContext) backquote(args lispObject) (lispObject, error) {
	return ec.nil_, nil
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

	iter := ec.iterate(body)
	for ; iter.hasNext(); body = iter.nextCons() {
		var err error

		form := xCar(body)
		val, err = ec.evalSub(form)
		if err != nil {
			return nil, err
		}
	}

	if iter.hasError() {
		return iter.error()
	}

	return val, nil
}

func (ec *execContext) setq(originalArgs lispObject) (lispObject, error) {
	args, err := ec.listToSlice(originalArgs)
	if err != nil {
		return nil, err
	}

	val := originalArgs

	for i, elem := range args {
		if i%2 == 0 && i == len(args)-1 {
			return ec.wrongNumberOfArguments(ec.s.setq, lispInt(i+1))
		}

		if i%2 == 1 {
			continue
		}

		sym := elem
		val, err = ec.evalSub(args[i+1])
		if err != nil {
			return nil, err
		}

		lexBinding := ec.nil_
		if symbolp(sym) {
			lexBinding, err = ec.assq(sym, ec.v.internalInterpreterEnv.val)
		}

		if lexBinding != ec.nil_ {
			xSetCdr(lexBinding, val) // Lexically bound
		} else {
			_, err = ec.set(sym, val)
			if err != nil {
				return nil, err
			}
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
		return ec.wrongNumberOfArguments(ec.s.function, xIntegerValue(length))
	}

	env := ec.v.internalInterpreterEnv.val
	if env != ec.nil_ &&
		consp(quoted) &&
		xCar(quoted) == ec.s.lambda {
		cdr := xCdr(quoted)

		return ec.makeCons(ec.s.closure, ec.makeCons(env, cdr)), nil
	}

	return quoted, nil
}

func (ec *execContext) let(args lispObject) (lispObject, error) {
	defer ec.unwind()()

	varList, err := ec.listToSlice(xCar(args))
	if err != nil {
		return nil, err
	}

	temps := make([]lispObject, len(varList))

	for argnum, elt := range varList {
		if symbolp(elt) {
			temps[argnum] = ec.nil_
		} else {
			cdr, err := ec.cdr(elt)
			if err != nil {
				return nil, err
			}

			cddr, err := ec.cdr(cdr)
			if err != nil {
				return nil, err
			}

			if cddr != ec.nil_ {
				return ec.signalError("`let' bindings can only have one value-form '%+v'", elt)
			}

			cadr, err := ec.car(cdr)
			if err != nil {
				return nil, err
			}

			val, err := ec.evalSub(cadr)
			if err != nil {
				return nil, err
			}

			temps[argnum] = val
		}
	}

	lexEnv := ec.v.internalInterpreterEnv.val

	for argnum, elt := range varList {
		var variable lispObject
		if symbolp(elt) {
			variable = elt
		} else {
			var err error
			variable, err = ec.car(elt)
			if err != nil {
				return nil, err
			}
		}

		tem := temps[argnum]
		included, err := ec.memq(variable, ec.v.internalInterpreterEnv.val)
		if err != nil {
			return nil, err
		}

		if lexEnv != ec.nil_ && symbolp(variable) &&
			!xSymbol(variable).special &&
			included == ec.nil_ {
			// Bind variable lexically
			lexEnv = ec.makeCons(ec.makeCons(variable, tem), lexEnv)
		} else {
			// Bind variable dynamically
			ec.stackPushLet(variable, tem)
		}
	}

	if lexEnv != ec.v.internalInterpreterEnv.sym {
		ec.stackPushLet(ec.v.internalInterpreterEnv.sym, lexEnv)
	}

	return ec.progn(xCdr(args))
}

func (ec *execContext) eval(form, lexical lispObject) (lispObject, error) {
	defer ec.unwind()()

	if !consp(lexical) && lexical != ec.nil_ {
		lexical = ec.makeList(ec.t)
	}

	ec.stackPushLet(ec.v.internalInterpreterEnv.sym, lexical)
	val, err := ec.evalSub(form)
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (ec *execContext) evalSub(form lispObject) (lispObject, error) {
	var err error

	if symbolp(form) {
		lex, err := ec.assq(form, ec.v.internalInterpreterEnv.val)
		if err != nil {
			return nil, err
		}

		if lex != ec.nil_ {
			return xCdr(lex), nil
		}

		return ec.symbolValue(form)
	} else if !consp(form) {
		return form, nil
	}

	originalFn := xCar(form)
	originalArgs := xCdr(form)
	if !consp(originalArgs) && originalArgs != ec.nil_ {
		return ec.wrongTypeArgument(ec.s.listp, originalArgs)
	}
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
		return ec.signalN(ec.s.voidFunction, originalFn)
	} else if !consp(fn) {
		return ec.signalN(ec.s.invalidFunction, originalFn)
	}

	fnCar := xCar(fn)
	if !symbolp(fnCar) {
		return ec.signalN(ec.s.invalidFunction, originalFn)
	}

	if fnCar == ec.s.macro {
		val := ec.t
		if ec.v.internalInterpreterEnv.val == ec.nil_ {
			val = ec.nil_
		}

		defer ec.unwind()()
		ec.stackPushLet(ec.v.lexicalBinding.sym, val)

		exp, err := ec.apply([]lispObject{xCdr(fn), originalArgs}...)
		if err != nil {
			return nil, err
		}

		return ec.evalSub(exp)
	} else if fnCar == ec.s.lambda || fnCar == ec.s.closure {
		return ec.applyLambda(fn, originalArgs)
	} else {
		return ec.signalN(ec.s.invalidFunction, originalFn)
	}
}

func (ec *execContext) symbolsOfEval() {
	// TODO: Unintern this
	ec.defVarLisp(&ec.v.internalInterpreterEnv, "internal-interpreter-environment", ec.nil_)
	ec.defSym(&ec.s.lambda, "lambda")
	ec.defSym(&ec.s.closure, "closure")
	ec.defSym(&ec.s.macro, "macro")
	ec.defSym(&ec.s.andRest, "&rest")
	ec.defSym(&ec.s.andOptional, "&optional")

	ec.defSubr2(&ec.s.eval, "eval", ec.eval, 1)
	ec.defSubrM(nil, "funcall", ec.funcall, 1)
	ec.defSubrM(nil, "apply", ec.apply, 1)
	ec.defSubrU(nil, "progn", ec.progn, 0)
	ec.defSubrU(&ec.s.setq, "setq", ec.setq, 0)
	ec.defSubrU(nil, "if", ec.if_, 2)
	ec.defSubrU(nil, "while", ec.while, 1)
	ec.defSubrU(&ec.s.quote, "quote", ec.quote, 1)
	ec.defSubrU(&ec.s.backquote, "`", ec.backquote, 1)
	ec.defSubrU(&ec.s.function, "function", ec.function, 1)
	ec.defSubrU(nil, "let", ec.let, 1)
	ec.defSubrU(nil, "catch", ec.catch, 1)
	ec.defSubrU(nil, "unwind-protect", ec.unwindProtect, 1)
	ec.defSubrU(nil, "condition-case", ec.conditionCase, 2)
	ec.defSubr2(nil, "throw", ec.throw, 2).setAttrs(true)
	ec.defSubr2(nil, "signal", ec.signal, 2).setAttrs(true)
}
