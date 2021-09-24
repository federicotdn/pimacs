package elisp

import (
	"fmt"
	"time"
)

func (ec *execContext) null(object lispObject) (lispObject, error) {
	return ec.bool(object == ec.nil_)
}

func (ec *execContext) sequencep(object lispObject) (lispObject, error) {
	return ec.bool(consp(object) || object == ec.nil_ || arrayp(object))
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

func (ec *execContext) goChannelp(object lispObject) (lispObject, error) {
	return ec.bool(vectorLikep(object, vectorLikeTypeGoChannel))
}

func (ec *execContext) makeGoChannel(size lispObject) (lispObject, error) {
	if !integerp(size) {
		return ec.wrongTypeArgument(ec.g.integerp, size)
	}

	channel := make(chan lispObject, int(xIntegerValue(size)))
	return ec.makeVectorLike(vectorLikeTypeGoChannel, channel), nil
}

func (ec *execContext) goChannelSend(channel, object lispObject) (lispObject, error) {
	if !vectorLikep(channel, vectorLikeTypeGoChannel) {
		return ec.wrongTypeArgument(ec.g.goChannelp, channel)
	}

	goChan := xVectorLike(channel).value.(chan lispObject)
	goChan <- object

	return object, nil
}

func (ec *execContext) goChannelReceive(channel lispObject) (lispObject, error) {
	if !vectorLikep(channel, vectorLikeTypeGoChannel) {
		return ec.wrongTypeArgument(ec.g.goChannelp, channel)
	}

	goChan := xVectorLike(channel).value.(chan lispObject)
	object, ok := <-goChan

	if !ok {
		return ec.g.goChannelClosed, nil
	}

	return object, nil
}

func (ec *execContext) goChannelClose(channel lispObject) (lispObject, error) {
	if !vectorLikep(channel, vectorLikeTypeGoChannel) {
		return ec.wrongTypeArgument(ec.g.goChannelp, channel)
	}

	goChan := xVectorLike(channel).value.(chan lispObject)
	close(goChan)
	return ec.nil_, nil
}

func (ec *execContext) sleepFor(seconds, milliseconds lispObject) (lispObject, error) {
	// TAGS: revise
	time.Sleep(time.Duration(xIntegerValue(seconds)) * time.Second)
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
		return ec.wrongTypeArgument(ec.g.listp, obj)
	}
	return xCons(obj).car, nil
}

func (ec *execContext) cdr(obj lispObject) (lispObject, error) {
	if obj == ec.nil_ {
		return ec.nil_, nil
	}

	if !consp(obj) {
		return ec.wrongTypeArgument(ec.g.listp, obj)
	}
	return xCons(obj).cdr, nil
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
			return ec.nil_, nil
		}
	}

	return ec.t, nil
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

func (ec *execContext) listLength(obj lispObject) (int, error) {
	tail := obj
	num := 0

	for ; consp(tail); tail = xCdr(tail) {
		num += 1
	}

	if tail != ec.nil_ {
		return 0, xErrOnly(ec.wrongTypeArgument(ec.g.listp, obj))
	}

	return num, nil
}

func (ec *execContext) length(obj lispObject) (lispObject, error) {
	// TAGS: incomplete
	num := 0

	switch obj.getType() {
	case lispTypeString:
		num = len(xString(obj).value)
	case lispTypeCons:
		var err error
		num, err = ec.listLength(obj)
		if err != nil {
			return nil, err
		}
	default:
		if obj != ec.nil_ {
			return ec.wrongTypeArgument(ec.g.sequencep, obj)
		}
	}

	return ec.makeInteger(lispInt(num)), nil
}

func (ec *execContext) read(stream lispObject) (lispObject, error) {
	if stream == ec.nil_ {
		stream = xSymbolValue(ec.g.standardInput)
	}

	if stream == ec.t {
		stream = ec.g.readChar
	}

	if stream == ec.g.readChar {
		return ec.funcall(ec.g.readFromMinibuffer, ec.makeString("Lisp expression: "))
	}

	result, _, err := ec.readInternalStart(stream, ec.nil_, ec.nil_)
	return result, err
}

func (ec *execContext) eval(form, lexical lispObject) (lispObject, error) {
	size := ec.stackSize()
	defer ec.stackPopTo(size)

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
		return ec.wrongTypeArgument(ec.g.symbolp, symbol)
	}

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

func (ec *execContext) assq(key, alist lispObject) (lispObject, error) {
	tail := alist
	for ; consp(tail); tail = xCdr(tail) {
		element := xCar(tail)

		if consp(element) && xCar(element) == key {
			return element, nil
		}
	}

	if tail != ec.nil_ {
		return ec.wrongTypeArgument(ec.g.listp, alist)
	}

	return ec.nil_, nil
}

func (ec *execContext) memq(elt, list lispObject) (lispObject, error) {
	tail := list
	for ; consp(tail); tail = xCdr(tail) {
		if xCar(tail) == elt {
			return tail, nil
		}
	}

	if tail != ec.nil_ {
		return ec.wrongTypeArgument(ec.g.listp, list)
	}

	return ec.g.nil_, nil
}

func (ec *execContext) eq(obj1, obj2 lispObject) (lispObject, error) {
	return ec.bool(obj1 == obj2)
}

func (ec *execContext) equal(o1, o2 lispObject) (lispObject, error) {
	// TAGS: incomplete
	if o1 == o2 {
		return ec.t, nil
	}

	t1, t2 := o1.getType(), o2.getType()
	if t1 != t2 {
		return ec.nil_, nil
	}

	switch t1 {
	case lispTypeCons:
		for ; consp(o1); o1 = xCdr(o1) {
			if !consp(o2) {
				return ec.bool(false)
			}

			equal, err := ec.equal(xCar(o1), xCar(o2))
			if err != nil {
				return nil, err
			}

			if equal == ec.nil_ {
				return ec.bool(false)
			}

			o2 = xCdr(o2)
			if xCdr(o1) == o2 {
				return ec.bool(true)
			}
		}

		return ec.bool(false)
	case lispTypeFloat:
		return ec.bool(xFloatValue(o1) == xFloatValue(o2))
	case lispTypeInteger:
		return ec.bool(xIntegerValue(o1) == xIntegerValue(o2))
	case lispTypeString:
		return ec.bool(xStringValue(o1) == xStringValue(o2))
	case lispTypeSymbol:
		// Symbols must match exactly (eq).
		return ec.bool(false)
	case lispTypeVectorLike:
		vec1, vec2 := xVectorLike(o1), xVectorLike(o2)
		if vec1.vecType != vec2.vecType {
			return ec.bool(false)
		}

		if vec1.vecType == vectorLikeTypeGoChannel {
			return ec.bool(vec1.value == vec2.value)
		}

		return ec.pimacsUnimplemented(ec.g.equal, "unknown vector-like type: '%v'", vec1.vecType)
	}

	return ec.bool(false)
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
	tag, err := ec.evalSub(xCar(args))
	if err != nil {
		return nil, err
	}

	size := ec.stackSize()
	ec.stackPushCatch(tag)
	defer ec.stackPopTo(size)

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

	size := ec.stackSize()
	defer ec.stackPopTo(size)
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

func (ec *execContext) readFromString(str, start, end lispObject) (lispObject, error) {
	if !stringp(str) {
		return ec.wrongTypeArgument(ec.g.stringp, str)
	}

	result, ctx, err := ec.readInternalStart(str, start, end)
	if err != nil {
		return nil, err
	}

	pos := ec.makeInteger(lispInt(ctx.position()))
	return ec.cons(result, pos)
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
		return ec.wrongTypeArgument(ec.g.listp, plist)
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

func (ec *execContext) defalias(symbol, definition, docstring lispObject) (lispObject, error) {
	// TAGS: incomplete
	return ec.fset(symbol, definition)
}

func (ec *execContext) list(args ...lispObject) (lispObject, error) {
	return ec.makeList(args...), nil
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

func (ec *execContext) initialDefsSubroutines() {
	ec.defSubr1("null", ec.null, 1)
	ec.defSubr1("sequencep", ec.sequencep, 1)
	ec.defSubr1("listp", ec.listp, 1)
	ec.defSubr1("symbolp", ec.symbolp, 1)
	ec.defSubr1("stringp", ec.stringp, 1)
	ec.defSubr1("number-or-marker-p", ec.numberOrMarkerp, 1)
	ec.defSubr1("char-or-string-p", ec.charOrStringp, 1)
	ec.defSubr1("integerp", ec.numberOrMarkerp, 1)
	ec.defSubr1("pimacs-go-channel-p", ec.goChannelp, 1)
	ec.defSubr1("pimacs-make-go-channel", ec.makeGoChannel, 1)
	ec.defSubr1("pimacs-go-channel-receive", ec.goChannelReceive, 1)
	ec.defSubr1("pimacs-go-channel-close", ec.goChannelClose, 1)
	ec.defSubr1("car", ec.car, 1)
	ec.defSubr1("cdr", ec.cdr, 1)
	ec.defSubr1("length", ec.length, 1)
	ec.defSubr1("read", ec.read, 0)
	ec.defSubr1("symbol-plist", ec.symbolPlist, 1)
	ec.defSubr1("symbol-name", ec.symbolName, 1)
	ec.defSubr2("sleep-for", ec.sleepFor, 1)
	ec.defSubr2("pimacs-go-channel-send", ec.goChannelSend, 2)
	ec.defSubr2("cons", ec.cons, 2)
	ec.defSubr2("eval", ec.eval, 1)
	ec.defSubr2("set", ec.set, 2)
	ec.defSubr2("fset", ec.fset, 2)
	ec.defSubr2("eq", ec.eq, 2)
	ec.defSubr2("equal", ec.equal, 2)
	ec.defSubr2("assq", ec.assq, 2)
	ec.defSubr2("memq", ec.memq, 2)
	ec.defSubr2("get", ec.get, 2)
	ec.defSubr2("throw", ec.throw, 2).noReturn = true
	ec.defSubr2("signal", ec.signal, 2).noReturn = true
	ec.defSubr2("plist-get", ec.plistGet, 2)
	ec.defSubr3("plist-put", ec.plistPut, 3)
	ec.defSubr3("put", ec.put, 3)
	ec.defSubr3("defalias", ec.defalias, 2)
	ec.defSubr3("read-from-string", ec.readFromString, 1)
	ec.defSubrM("+", ec.plusSign, 0)
	ec.defSubrM("<", ec.lessThanSign, 1)
	ec.defSubrM("list", ec.list, 0)
	ec.defSubrM("funcall", ec.funcall, 1)
	ec.defSubrM("apply", ec.apply, 1)
	ec.defSubrU("catch", ec.catch, 1)
	ec.defSubrU("unwind-protect", ec.unwindProtect, 1)
	ec.defSubrU("quote", ec.quote, 1)
	ec.defSubrU("if", ec.if_, 2)
	ec.defSubrU("while", ec.while, 1)
	ec.defSubrU("progn", ec.progn, 0)
	ec.defSubrU("function", ec.function, 1)
	ec.defSubrU("condition-case", ec.conditionCase, 2)
}
