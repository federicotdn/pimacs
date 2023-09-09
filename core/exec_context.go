package core

import (
	"fmt"
	"github.com/federicotdn/pimacs/proto"
	"sync"
)

type stackEntryTag int

const (
	entryLet stackEntryTag = iota + 1
	entryLetForwarded
	entryCatch
	entryFnLispObject
	entryBacktrace
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
	lispStack   string
	ec          *execContext
}

type goroutineLocals struct {
	stack                  []stackEntry
	currentBuf             *lispBuffer
	internalInterpreterEnv forwardLispObj
	lexicalBinding         forwardLispObj
}

type execContext struct {
	gl *goroutineLocals

	nil_ lispObject
	t    lispObject
	s    *symbols
	v    *vars

	obarray     map[string]*lispSymbol
	obarrayLock *sync.Mutex

	buffers     map[string]*lispBuffer
	buffersLock *sync.RWMutex

	stubs *emacsStubs

	hashTestEq     *lispHashTableTest
	hashTestEql    *lispHashTableTest
	hashTableEqual *lispHashTableTest

	events chan proto.InputEvent
	ops    chan proto.DrawOp
	done   chan bool

	testing bool

	init struct {
		errorOnVarRedefine bool
		subrs              int
		stubs              int
	}
}

type emacsStubs struct {
}

type stackEntryLet struct {
	symbol *lispSymbol
	oldVal lispObject
}

type stackEntryLetForwarded struct {
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

type stackEntryBacktrace struct {
	debugOnExit bool
	function    lispObject
	args        []lispObject
	evaluated   bool
}

func (jmp *stackJumpTag) Error() string {
	format := "stack jump: tag: '%+v'"
	if symbolp(jmp.tag) {
		return fmt.Sprintf(format, xSymbolName(jmp.tag))
	}

	return fmt.Sprintf(format, jmp.tag)
}

func (jmp *stackJumpSignal) Is(target error) bool {
	other, ok := target.(*stackJumpSignal)
	if !ok {
		return false
	}

	if jmp.errorSymbol != other.errorSymbol {
		return false
	}

	ec := jmp.ec

	switch jmp.errorSymbol {
	case ec.s.wrongTypeArgument:
		ourData := xCdr(jmp.data)
		otherData := xCdr(other.data)
		if !consp(ourData) || !consp(otherData) {
			return false
		}

		ourPred := xCar(ourData)
		otherPred := xCar(otherData)
		return ourPred == otherPred
	default:
		// By default, two stack jumps with the same error
		// symbol are considered equal, unless a specific
		// handler for that error symbol is implemented
		return true
	}
}

func (jmp *stackJumpSignal) Error() string {
	message := "stack jump: signal:"
	ending := "\nbacktrace:\n" + jmp.lispStack

	if !symbolp(jmp.errorSymbol) {
		message += fmt.Sprintf(" '%+v' '<unknown>'", jmp.errorSymbol)
		return message
	}

	name := xSymbolName(jmp.errorSymbol)
	message += fmt.Sprintf(" '%+v'", name)

	originalData := xCdr(jmp.data)
	data := originalData
	if !consp(data) {
		if stringp(data) {
			message += fmt.Sprintf(" '%+v'", xStringValue(data))
		} else {
			message += fmt.Sprintf(" '%+v'", data)
		}
		return message + ending
	}

	ec := jmp.ec

	switch jmp.errorSymbol {
	case ec.s.error_, ec.s.pimacsUnimplemented, ec.s.fileMissing:
		message += fmt.Sprintf(" '%+v'", xStringValue(xCar(data)))
	case ec.s.voidVariable:
		message += fmt.Sprintf(" '%+v'", xSymbolName(xCar(data)))
	case ec.s.voidFunction:
		message += fmt.Sprintf(" '%+v'", xCar(data))
	case ec.s.wrongTypeArgument:
		pred := xCar(data)
		val := xCar(xCdr(data))
		message += fmt.Sprintf(" '%+v' '%+v'", xSymbolName(pred), val)
	case ec.s.wrongNumberofArguments:
		fn := xCar(data)

		for ; consp(fn); fn = xCar(fn) {
		}

		if symbolp(fn) {
			message += fmt.Sprintf(" '%+v'", xSymbolName(fn))
		} else if subroutinep(fn) {
			message += fmt.Sprintf(" '%+v'", xSubroutine(fn).name)
		} else {
			message += " '<function>'"
		}

		count := xCar(xCdr(data))
		message += fmt.Sprintf(" '%+v'", xIntegerValue(count))
	default:
		message += fmt.Sprintf(" '%+v'", originalData)

	}

	return message + ending
}

func (se *stackEntryLet) tag() stackEntryTag {
	return entryLet
}

func (se *stackEntryLetForwarded) tag() stackEntryTag {
	return entryLetForwarded
}

func (se *stackEntryCatch) tag() stackEntryTag {
	return entryCatch
}

func (se *stackEntryFnLispObject) tag() stackEntryTag {
	return entryFnLispObject
}

func (se *stackEntryBacktrace) tag() stackEntryTag {
	return entryBacktrace
}

func (ec *execContext) makeSymbol(name string, init bool) *lispSymbol {
	sym := &lispSymbol{name: name}
	if init {
		sym.val = ec.s.unbound
		sym.function = ec.nil_
		sym.plist = ec.nil_
	}

	return sym
}

func (ec *execContext) makeList(objs ...lispObject) lispObject {
	if len(objs) == 0 {
		return ec.nil_
	}

	tmp := newCons(objs[0], ec.nil_)
	val := tmp

	for _, obj := range objs[1:] {
		tmp.cdr = newCons(obj, ec.nil_)
		tmp = xCons(tmp.cdr)
	}

	return val
}

func (ec *execContext) makePlist(objs map[string]lispObject) (lispObject, error) {
	if len(objs) == 0 {
		return ec.nil_, nil
	}

	val := ec.nil_

	for key, obj := range objs {
		val = newCons(obj, val)
		sym, err := ec.intern(newString(":"+key), ec.nil_)
		if err != nil {
			return nil, err
		}
		val = newCons(sym, val)
	}

	return val, nil
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
		predicate: ec.s.listp,
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
		li.err = xErrOnly(li.ec.wrongTypeArgument(li.predicate, li.listHead))
	}
}

func (li *listIter) nextCons() lispObject {
	// TODO: incomplete
	// Need cycle detection still.
	li.tail = xCdr(li.tail)

	li.checkTailType()

	return li.tail
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

func (ec *execContext) defSubrInternal(symbol *lispObject, fn lispFn, name string, minArgs, maxArgs int) *lispSubroutine {
	sub := &lispSubroutine{
		callabe: fn,
		name:    name,
		minArgs: minArgs,
		maxArgs: maxArgs,
	}

	if ec.init.errorOnVarRedefine {
		ec.init.subrs++
	} else {
		ec.init.stubs++
	}

	if symbol != nil && *symbol != nil {
		if ec.init.errorOnVarRedefine {
			ec.terminate("subroutine value already set: '%+v'", *symbol)
		}
		return sub
	}

	if sub.maxArgs >= 0 && sub.minArgs > sub.maxArgs {
		ec.terminate("min args must be smaller or equal to max args (subroutine: '%+v')", name)
	}

	sym := ec.defSym(symbol, sub.name) // Create a variable with value = unbound
	sym.function = sub

	if symbol != nil {
		*symbol = sym
	}

	return sub
}

func (ec *execContext) defSubr0(symbol *lispObject, name string, fn lispFn0) *lispSubroutine {
	return ec.defSubrInternal(symbol, fn, name, 0, 0)
}

func (ec *execContext) defSubr1(symbol *lispObject, name string, fn lispFn1, minArgs int) *lispSubroutine {
	return ec.defSubrInternal(symbol, fn, name, minArgs, 1)
}

func (ec *execContext) defSubr2(symbol *lispObject, name string, fn lispFn2, minArgs int) *lispSubroutine {
	return ec.defSubrInternal(symbol, fn, name, minArgs, 2)
}

func (ec *execContext) defSubr3(symbol *lispObject, name string, fn lispFn3, minArgs int) *lispSubroutine {
	return ec.defSubrInternal(symbol, fn, name, minArgs, 3)
}

func (ec *execContext) defSubr4(symbol *lispObject, name string, fn lispFn4, minArgs int) *lispSubroutine {
	return ec.defSubrInternal(symbol, fn, name, minArgs, 4)
}

func (ec *execContext) defSubr5(symbol *lispObject, name string, fn lispFn5, minArgs int) *lispSubroutine {
	return ec.defSubrInternal(symbol, fn, name, minArgs, 5)
}

func (ec *execContext) defSubr6(symbol *lispObject, name string, fn lispFn6, minArgs int) *lispSubroutine {
	return ec.defSubrInternal(symbol, fn, name, minArgs, 6)
}

func (ec *execContext) defSubr7(symbol *lispObject, name string, fn lispFn7, minArgs int) *lispSubroutine {
	return ec.defSubrInternal(symbol, fn, name, minArgs, 7)
}

func (ec *execContext) defSubr8(symbol *lispObject, name string, fn lispFn8, minArgs int) *lispSubroutine {
	return ec.defSubrInternal(symbol, fn, name, minArgs, 8)
}

func (ec *execContext) defSubrM(symbol *lispObject, name string, fn lispFnM, minArgs int) *lispSubroutine {
	return ec.defSubrInternal(symbol, fn, name, minArgs, argsMany)
}
func (ec *execContext) defSubrU(symbol *lispObject, name string, fn lispFn1, minArgs int) *lispSubroutine {
	return ec.defSubrInternal(symbol, fn, name, minArgs, argsUnevalled)
}

func (ec *execContext) defSym(symbol *lispObject, name string) *lispSymbol {
	if symbol != nil && *symbol != nil {
		ec.terminate("symbol already initialized: '%+v'", *symbol)
	}

	sym := ec.makeSymbol(name, true)
	if symbol != nil {
		*symbol = sym
	}

	_, existed := ec.internInstanceInternal(sym)
	if existed && ec.init.errorOnVarRedefine {
		ec.terminate("symbol already initialized: '%+v'", *symbol)
	}
	return sym
}

func (ec *execContext) defVarLisp(fwd *forwardLispObj, name string, value lispObject) {
	if fwd.sym != nil {
		ec.terminate("variable already initialized: '%+v'", fwd)
	}
	sym := ec.defSym(nil, name)
	sym.special = true
	fwd.sym = sym
	fwd.val = value
	sym.redirect = fwd
}

func (ec *execContext) defVarBool(fwd *forwardBool, name string, value bool) {
	if fwd.sym != nil {
		ec.terminate("variable already initialized: '%+v'", fwd)
	}
	sym := ec.defSym(nil, name)
	sym.special = true
	fwd.sym = sym
	fwd.val = value
	sym.redirect = fwd
}

func (ec *execContext) internInternal(name string) (*lispSymbol, bool) {
	sym := ec.makeSymbol(name, true)
	return ec.internInstanceInternal(sym)
}

func (ec *execContext) internInstanceInternal(sym *lispSymbol) (*lispSymbol, bool) {
	ec.obarrayLock.Lock()
	defer ec.obarrayLock.Unlock()

	prev, existed := ec.obarray[sym.name]
	if existed {
		return prev, true
	}

	ec.obarray[sym.name] = sym
	return sym, false
}

func (ec *execContext) uninternInternal(name string) bool {
	ec.obarrayLock.Lock()
	defer ec.obarrayLock.Unlock()

	_, existed := ec.obarray[name]
	if !existed {
		return false
	}

	delete(ec.obarray, name)
	return true
}

func (ec *execContext) initGoroutineLocals() {
	ec.defVarLisp(&ec.gl.lexicalBinding, "lexical-binding", ec.nil_)
	ec.defVarLisp(&ec.gl.internalInterpreterEnv, "internal-interpreter-environment", ec.nil_)
	ec.uninternInternal("internal-interpreter-environment")

	ec.initBufferGoroutineLocals() // buffer.go
}

func newExecContext(loadPathPrepend []string) (*execContext, error) {
	ec := execContext{
		gl:          &goroutineLocals{},
		s:           &symbols{},
		v:           &vars{},
		obarray:     make(map[string]*lispSymbol),
		obarrayLock: &sync.Mutex{},
		buffers:     make(map[string]*lispBuffer),
		buffersLock: &sync.RWMutex{},
		// TODO: Move '10' to config value
		events: make(chan proto.InputEvent, 10),
		ops:    make(chan proto.DrawOp, 10),
		done:   make(chan bool),
	}

	ec.init.errorOnVarRedefine = true

	// Symbol and vars basic initialization
	ec.initSymbols()             // symbols.go
	ec.symbolsOfExecContext()    // exec_context.go
	ec.symbolsOfErrors()         // errors.go
	ec.symbolsOfRead()           // read.go
	ec.symbolsOfEval()           // eval.go
	ec.symbolsOfPrint()          // print.go
	ec.symbolsOfData()           // data.go
	ec.symbolsOfAllocation()     // allocation.go
	ec.symbolsOfFunctions()      // functions.go
	ec.symbolsOfBuffer()         // buffer.go
	ec.symbolsOfMinibuffer()     // minibuffer.go
	ec.symbolsOfCallProc()       // callproc.go
	ec.symbolsOfKeyboard()       // keyboard.go
	ec.symbolsOfCharacterTable() // character_table.go
	ec.symbolsOfGoroutine()      // goroutine.go
	ec.symbolsOfPimacsTools()    // pimacs_tools.go

	// Goroutine-specific initialization
	ec.initGoroutineLocals()

	// Load Emacs stubs at the end, without erroring out
	// on repeated defuns or defvars
	ec.init.errorOnVarRedefine = false
	ec.symbolsOfEmacs_autogen()

	ec.checkSymbolValues()
	ec.checkVarValues()

	// We are now ready to evaluate Elisp code
	return &ec, ec.loadElisp(loadPathPrepend)
}

func (ec *execContext) symbolsOfExecContext() {
	ec.defVarBool(&ec.v.nonInteractive, "noninteractive", true)
}

func (ec *execContext) copyExecContext() *execContext {
	copy := *ec
	copy.gl = &goroutineLocals{}
	copy.initGoroutineLocals()
	return &copy
}

func (ec *execContext) loadElisp(loadPathPrepend []string) error {
	loadPath := []lispObject{}
	for _, elem := range loadPathPrepend {
		loadPath = append(loadPath, newString(elem))
	}
	loadPath = append(loadPath, newString("lisp"), newString("lisp/emacs"))

	ec.v.loadPath.val = ec.makeList(loadPath...)

	_, err := ec.load(newString("loadup-pimacs.el"), ec.nil_, ec.nil_, ec.nil_, ec.nil_)
	if err != nil {
		return err
	}

	return nil
}

func (ec *execContext) listToSlice(list lispObject) ([]lispObject, error) {
	result := []lispObject{}

	iter := ec.iterate(list)
	for ; iter.hasNext(); list = iter.nextCons() {
		result = append(result, xCar(list))
	}

	if iter.hasError() {
		return nil, xErrOnly(iter.error())
	}

	return result, nil
}

func (ec *execContext) stackPushLet(symbol lispObject, value lispObject) error {
	sym := xSymbol(symbol)

	if sym.redirect == nil {
		ec.gl.stack = append(ec.gl.stack, &stackEntryLet{
			symbol: sym,
			oldVal: sym.val,
		})
		sym.val = value
	} else {
		entry := &stackEntryLetForwarded{
			symbol: sym,
			oldVal: sym.redirect.value(ec),
		}

		err := sym.redirect.setValue(ec, value)
		if err != nil {
			return err
		}
		ec.gl.stack = append(ec.gl.stack, entry)
	}

	return nil
}

func (ec *execContext) stackPushCatch(tag lispObject) {
	ec.gl.stack = append(ec.gl.stack, &stackEntryCatch{
		catchTag: tag,
	})
}

func (ec *execContext) stackPushCurrentBuffer() {
	arg := ec.currentBufferInternal()
	ec.stackPushFnLispObject(ec.setBufferIfLive, arg)
}

func (ec *execContext) stackPushFnLispObject(function func(lispObject), arg lispObject) {
	ec.gl.stack = append(ec.gl.stack, &stackEntryFnLispObject{
		function: function,
		arg:      arg,
	})
}

func (ec *execContext) stackPushBacktrace(function lispObject, args []lispObject, evaluated bool) {
	ec.gl.stack = append(ec.gl.stack, &stackEntryBacktrace{
		function:  function,
		args:      args,
		evaluated: evaluated,
	})
}

func (ec *execContext) catchInStack(tag lispObject) bool {
	for _, binding := range ec.gl.stack {
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
	return len(ec.gl.stack)
}

func (ec *execContext) stackPopTo(target int) {
	size := len(ec.gl.stack)
	if target < 0 || size < target {
		ec.terminate("unable to pop back to '%v', size is '%v'", target, size)
	}

	for len(ec.gl.stack) > target {
		current := ec.gl.stack[len(ec.gl.stack)-1]

		switch current.tag() {
		case entryLet:
			let := current.(*stackEntryLet)
			let.symbol.val = let.oldVal
		case entryLetForwarded:
			let := current.(*stackEntryLetForwarded)
			// Should not fail as we're just setting the old value back
			err := let.symbol.redirect.setValue(ec, let.oldVal)
			if err != nil {
				ec.terminate("could not restore forwarded symbol value: '%+v'", let)
			}
		case entryFnLispObject:
			entry := current.(*stackEntryFnLispObject)
			entry.function(entry.arg)
		case entryCatch:
		case entryBacktrace:
		default:
			ec.terminate("unknown stack item tag: '%v'", current.tag())
		}

		ec.gl.stack = ec.gl.stack[:len(ec.gl.stack)-1]
	}
}

func (ec *execContext) printLispStack() string {
	lispStack := ""
	for i := len(ec.gl.stack) - 1; i >= 0; i-- {

		switch elem := ec.gl.stack[i].(type) {
		case *stackEntryBacktrace:
			functionName := "<unknown function>"
			if symbolp(elem.function) {
				functionName = xSymbolName(elem.function)
			}

			lispStack += fmt.Sprintf("  - bt: %v(", functionName)

			for j, arg := range elem.args {
				printed := debugRepr(arg)
				if len(printed) > 10 {
					printed = printed[:10] + "[...]"
				}
				lispStack += printed

				if j < len(elem.args)-1 {
					lispStack += " "
				}
			}

			lispStack += ")"
		case *stackEntryLet:
			lispStack += fmt.Sprintf("  - let: %v = %v", debugRepr(elem.symbol), debugRepr(elem.oldVal))
		case *stackEntryLetForwarded:
			lispStack += fmt.Sprintf("  - letfwd: %v = %v", debugRepr(elem.symbol), debugRepr(elem.oldVal))
		default:
			lispStack += "  - other"
		}

		if i > 0 {
			lispStack += "\n"
		}

	}
	return lispStack
}

func (ec *execContext) boolVal(b bool) lispObject {
	if b {
		return ec.t
	}
	return ec.nil_
}

func (ec *execContext) bool(b bool) (lispObject, error) {
	return ec.boolVal(b), nil
}

func (ec *execContext) true_() (lispObject, error) {
	return ec.bool(true)
}

func (ec *execContext) false_() (lispObject, error) {
	return ec.bool(false)
}

func (ec *execContext) stub(name string) (lispObject, error) {
	ec.terminate("stub invoked: '%v'", name)
	return ec.nil_, nil
}

func (ec *execContext) terminate(format string, v ...interface{}) {
	if !ec.testing {
		stack := ec.printLispStack()
		fmt.Println("backtrace:")
		fmt.Println(stack)
	}
	terminate(format, v...)
}
