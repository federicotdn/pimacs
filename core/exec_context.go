package core

import (
	"fmt"
	"os"
)

type stackEntryTag int

const (
	entryLet stackEntryTag = iota + 1
	entryCatch
	entryFnLispObject
)

type stackEntry interface {
	tag() stackEntryTag
}

type stackJumpTag struct {
	tag   lispObject
	value lispObject
}

type stackJumpSignal struct {
	errorSymbol lispObject
	data        lispObject
	goStack     string
}

type execContext struct {
	stack      []stackEntry
	nil_       lispObject
	t          lispObject
	g          lispGlobals
	obarray    map[string]*lispSymbol
	currentBuf *buffer
	buffers    lispObject

	errorOnVarRedefine bool
}

type stackEntryLet struct {
	symbol *lispSymbol
	oldVal lispObject
}

type stackEntryCatch struct {
	catchTag lispObject
}

type stackEntryFnLispObject struct {
	function func(lispObject)
	arg      lispObject
}

func (jmp *stackJumpTag) Error() string {
	format := "stack jump: tag: '%+v'"
	if symbolp(jmp.tag) {
		return fmt.Sprintf(format, xSymbol(jmp.tag).name)
	}

	return fmt.Sprintf(format, jmp.tag)
}

func (jmp *stackJumpSignal) Error() string {
	message := "stack jump: signal: "

	if !symbolp(jmp.errorSymbol) {
		message += fmt.Sprintf("'%+v' '<unknown>'", jmp.errorSymbol)
		return message
	}

	name := xSymbol(jmp.errorSymbol).name
	message += "'" + name + "' "

	data := xCdr(jmp.data)

	switch name {
	case "error":
		message = xStringValue(xCar(data))
	case "wrong-type-argument":
		pred := xCar(data)
		val := xCar(xCdr(data))
		message += fmt.Sprintf("'%+v' '%+v'", xSymbol(pred).name, val)
	case "wrong-number-of-arguments":
		fn := xCar(data)

		for ; consp(fn); fn = xCar(fn) {
		}

		if symbolp(fn) {
			message += fmt.Sprintf("'%+v' ", xSymbol(fn).name)
		} else if subroutinep(fn) {
			message += fmt.Sprintf("'%+v' ", xSubroutine(fn).name)
		} else {
			message += "'<function>' "
		}

		count := xCar(xCdr(data))
		message += fmt.Sprintf("'%+v'", xIntegerValue(count))
	}

	message += "\nGo stack for jump:\n" + jmp.goStack

	return message
}

func (se *stackEntryLet) tag() stackEntryTag {
	return entryLet
}

func (se *stackEntryCatch) tag() stackEntryTag {
	return entryCatch
}

func (se *stackEntryFnLispObject) tag() stackEntryTag {
	return entryFnLispObject
}

func (ec *execContext) makeSymbolBase(name string) *lispSymbol {
	return &lispSymbol{
		name:    name,
		special: false,
	}
}

func (ec *execContext) makeSymbol(name string) *lispSymbol {
	base := ec.makeSymbolBase(name)
	base.value = ec.g.unbound
	base.function = ec.nil_
	base.plist = ec.nil_
	return base
}

func (ec *execContext) makeInteger(value lispInt) *lispInteger {
	return &lispInteger{
		value: value,
	}
}

func (ec *execContext) makeCons(car, cdr lispObject) *lispCons {
	return &lispCons{
		car: car,
		cdr: cdr,
	}
}

func (ec *execContext) makeFloat(value lispFp) *lispFloat {
	return &lispFloat{
		value: value,
	}
}

func (ec *execContext) makeString(value string) *lispString {
	return &lispString{
		value: value,
	}
}

func (ec *execContext) emptyString() *lispString {
	return ec.makeString("")
}

func (ec *execContext) makeList(objs ...lispObject) lispObject {
	if len(objs) == 0 {
		return ec.nil_
	}

	tmp := ec.makeCons(objs[0], ec.nil_)
	val := tmp

	for _, obj := range objs[1:] {
		tmp.cdr = ec.makeCons(obj, ec.nil_)
		tmp = xCons(tmp.cdr)
	}

	return val
}

func (ec *execContext) makeVectorLike(vecType vectorLikeType, value vectorLikeValue) *lispVectorLike {
	return &lispVectorLike{
		vecType: vecType,
		value:   value,
	}
}

func (ec *execContext) makeBuffer(buf *buffer) *lispVectorLike {
	return ec.makeVectorLike(vectorLikeTypeBuffer, buf)
}

type listIter struct {
	tail      lispObject
	listHead  lispObject
	err       error
	ec        *execContext
	predicate lispObject
	hasCycle  bool
}

func (ec *execContext) iterate(tail lispObject) *listIter {
	li := &listIter{
		tail:      tail,
		listHead:  tail,
		ec:        ec,
		predicate: ec.g.listp,
	}

	// We may have been given already an object that is not a list
	li.checkTailType()

	return li
}

func (li *listIter) withPredicate(predicate lispObject) *listIter {
	li.predicate = predicate
	return li
}

func (li *listIter) hasNext() bool {
	return consp(li.tail) && !li.hasError()
}

func (li *listIter) checkTailType() {
	if !consp(li.tail) && li.tail != li.ec.nil_ {
		// List does not end with nil.
		// Signal using list start, not tail!
		li.err = xErrOnly(li.ec.wrongTypeArgument(li.ec.g.listp, li.head()))
	}
}

func (li *listIter) next() lispObject {
	// TODO: incomplete
	// Need cycle detection still.
	li.tail = xCdr(li.tail)

	li.checkTailType()

	return li.tail
}

func (li *listIter) head() lispObject {
	return li.listHead
}

func (li *listIter) hasError() bool {
	return li.err != nil
}

func (li *listIter) circular() bool {
	return li.hasCycle
}

func (li *listIter) error() (lispObject, error) {
	return nil, li.err
}

func (ec *execContext) defSubrInternal(symbol *lispObject, fn lispFn, name string, minArgs, maxArgs int) *subroutine {
	sub := &subroutine{
		callabe: fn,
		name:    name,
		minArgs: minArgs,
		maxArgs: maxArgs,
	}

	if symbol != nil && *symbol != nil {
		if ec.errorOnVarRedefine {
			ec.terminate("subroutine value already set: '%+v'", *symbol)
		}
		return sub
	}

	if sub.maxArgs >= 0 && sub.minArgs > sub.maxArgs {
		ec.terminate("min args must be smaller or equal to max args (subroutine: '%+v')", name)
	}

	sym := ec.defVar(symbol, sub.name, ec.g.unbound) // Create a variable with value = unbound
	vec := ec.makeVectorLike(vectorLikeTypeSubroutine, sub)
	sym.function = vec

	if symbol != nil {
		*symbol = sym
	}

	return sub
}

func (ec *execContext) defSubr0(symbol *lispObject, name string, fn lispFn0) *subroutine {
	return ec.defSubrInternal(symbol, fn, name, 0, 0)
}

func (ec *execContext) defSubr1(symbol *lispObject, name string, fn lispFn1, minArgs int) *subroutine {
	return ec.defSubrInternal(symbol, fn, name, minArgs, 1)
}

func (ec *execContext) defSubr2(symbol *lispObject, name string, fn lispFn2, minArgs int) *subroutine {
	return ec.defSubrInternal(symbol, fn, name, minArgs, 2)
}

func (ec *execContext) defSubr3(symbol *lispObject, name string, fn lispFn3, minArgs int) *subroutine {
	return ec.defSubrInternal(symbol, fn, name, minArgs, 3)
}

func (ec *execContext) defSubr4(symbol *lispObject, name string, fn lispFn4, minArgs int) *subroutine {
	return ec.defSubrInternal(symbol, fn, name, minArgs, 4)
}

func (ec *execContext) defSubr5(symbol *lispObject, name string, fn lispFn5, minArgs int) *subroutine {
	return ec.defSubrInternal(symbol, fn, name, minArgs, 5)
}

func (ec *execContext) defSubr6(symbol *lispObject, name string, fn lispFn6, minArgs int) *subroutine {
	return ec.defSubrInternal(symbol, fn, name, minArgs, 6)
}

func (ec *execContext) defSubr7(symbol *lispObject, name string, fn lispFn7, minArgs int) *subroutine {
	return ec.defSubrInternal(symbol, fn, name, minArgs, 7)
}

func (ec *execContext) defSubr8(symbol *lispObject, name string, fn lispFn8, minArgs int) *subroutine {
	return ec.defSubrInternal(symbol, fn, name, minArgs, 8)
}

func (ec *execContext) defSubrM(symbol *lispObject, name string, fn lispFnM, minArgs int) *subroutine {
	return ec.defSubrInternal(symbol, fn, name, minArgs, argsMany)
}

func (ec *execContext) defSubrU(symbol *lispObject, name string, fn lispFn1, minArgs int) *subroutine {
	return ec.defSubrInternal(symbol, fn, name, minArgs, argsUnevalled)
}

func (ec *execContext) defVarUninterned(symbol *lispObject, name string, value lispObject) *lispSymbol {
	if symbol != nil && *symbol != nil {
		ec.terminate("variable value already set: '%+v'", *symbol)
	}

	sym := ec.makeSymbol(name)
	sym.value = value

	if symbol != nil {
		*symbol = sym
	}
	return sym
}

func (ec *execContext) defVar(symbol *lispObject, name string, value lispObject) *lispSymbol {
	sym := ec.defVarUninterned(symbol, name, value)
	ec.internNewSymbol(sym, ec.errorOnVarRedefine)
	return sym
}

func newExecContext() *execContext {
	ec := execContext{
		obarray:            make(map[string]*lispSymbol),
		stack:              []stackEntry{},
		errorOnVarRedefine: true,
	}

	ec.initSymbols()         // symbols.go
	ec.symbolsOfErrors()     // errors.gmo
	ec.symbolsOfRead()       // read.go
	ec.symbolsOfEval()       // eval.go
	ec.symbolsOfPrint()      // print.go
	ec.symbolsOfData()       // data.go
	ec.symbolsOfAllocation() // allocation.go
	ec.symbolsOfFunctions()  // functions.go
	ec.symbolsOfBuffer()     // buffer.go
	ec.symbolsOfMinibuffer() // minibuffer.go
	ec.symbolsOfCallProc()   // callproc.go

	ec.defVar(&ec.g.nonInteractive, "noninteractive", ec.t)

	// Load Emacs stubs at the end, without erroring out
	// on repeated defuns or defvars
	ec.errorOnVarRedefine = false
	ec.symbolsOfEmacs_autogen()
	ec.errorOnVarRedefine = true

	ec.checkSymbols()

	ec.initBuffer() // buffer.go

	loadPath, ok := os.LookupEnv("PIMACS_LISP")
	if !ok {
		// Use relative path from CWD
		loadPath = "lisp"
	}
	xSymbol(ec.g.loadPath).value = ec.makeList(ec.makeString(loadPath))

	err := ec.loadElisp()
	if err != nil {
		panic(err)
	}

	return &ec
}

func (ec *execContext) loadElisp() error {
	files := []string{"pimacs.el"}

	for _, path := range files {
		_, err := ec.load(ec.makeString(path), ec.nil_, ec.nil_, ec.nil_, ec.nil_)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ec *execContext) internNewSymbol(symbol *lispSymbol, errorOnExisting bool) {
	prev, ok := ec.obarray[symbol.name]
	if ok && prev != symbol {
		if errorOnExisting {
			ec.terminate("different symbol with that name already interned, name: '%v'", symbol.name)
		}
		return
	}
	ec.obarray[symbol.name] = symbol
}

func (ec *execContext) listToSlice(list lispObject) ([]lispObject, error) {
	result := []lispObject{}

	iter := ec.iterate(list)
	for ; iter.hasNext(); list = iter.next() {
		result = append(result, xCar(list))
	}

	if iter.hasError() {
		return nil, xErrOnly(iter.error())
	}

	return result, nil
}

func (ec *execContext) stackPushLet(symbol lispObject, value lispObject) {
	sym := xSymbol(symbol)
	ec.stack = append(ec.stack, &stackEntryLet{
		symbol: sym,
		oldVal: sym.value,
	})

	sym.value = value
}

func (ec *execContext) stackPushCatch(tag lispObject) {
	ec.stack = append(ec.stack, &stackEntryCatch{
		catchTag: tag,
	})
}

func (ec *execContext) stackPushCurrentBuffer() {
	arg := ec.currentBufferInternal()
	ec.stackPushFnLispObject(ec.setBufferIfLive, arg)
}

func (ec *execContext) stackPushFnLispObject(function func(lispObject), arg lispObject) {
	ec.stack = append(ec.stack, &stackEntryFnLispObject{
		function: function,
		arg:      arg,
	})
}

func (ec *execContext) catchInStack(tag lispObject) bool {
	for _, binding := range ec.stack {
		if binding.tag() != entryCatch {
			continue
		}

		catchTag := binding.(*stackEntryCatch).catchTag
		if catchTag == tag {
			return true
		}
	}

	return false
}

// unwind returns an anonymous function that when called
// will unwind the execution context stack to the point
// where it was when unwind was called.
// The most common way of invoking this function is:
// defer ec.unwind()()
func (ec *execContext) unwind() func() {
	size := ec.stackSize()
	return func() {
		ec.stackPopTo(size)
	}
}

func (ec *execContext) stackSize() int {
	return len(ec.stack)
}

func (ec *execContext) stackPopTo(target int) {
	size := len(ec.stack)
	if target < 0 || size < target {
		ec.terminate("unable to pop back to '%v', size is '%v'", target, size)
	}

	for len(ec.stack) > target {
		current := ec.stack[len(ec.stack)-1]

		switch current.tag() {
		case entryLet:
			let := current.(*stackEntryLet)
			let.symbol.value = let.oldVal
		case entryCatch:
		case entryFnLispObject:
			entry := current.(*stackEntryFnLispObject)
			entry.function(entry.arg)
		default:
			ec.terminate("unknown stack item tag: '%v'", current.tag())
		}

		ec.stack = ec.stack[:len(ec.stack)-1]
	}
}

func (ec *execContext) bool(b bool) (lispObject, error) {
	if b {
		return ec.t, nil
	}
	return ec.nil_, nil
}

func (ec *execContext) true_() (lispObject, error) {
	return ec.t, nil
}

func (ec *execContext) false_() (lispObject, error) {
	return ec.nil_, nil
}

func (ec *execContext) stub(name string) (lispObject, error) {
	println("stub invoked:", name)
	return ec.nil_, nil
}

func (ec *execContext) terminate(format string, v ...interface{}) {
	// TAGS: incomplete
	terminate(format, v...)
}
