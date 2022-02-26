package elisp

import (
	"embed"
	"fmt"
)

//go:embed scripts
var scripts embed.FS

type stackEntryTag int

const (
	entryLet stackEntryTag = iota + 1
	entryCatch
	entryCurrentBuffer
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
}

type execContext struct {
	stack         []stackEntry
	nil_          lispObject
	t             lispObject
	g             lispGlobals
	obarray       map[string]*lispSymbol
	currentBuffer *buffer
}

type stackEntryLet struct {
	symbol *lispSymbol
	oldVal lispObject
}

type stackEntryCatch struct {
	catchTag lispObject
}

type stackEntryCurrentBuffer struct {
	prevCurrentBuffer *buffer
}

func (jmp *stackJumpTag) Error() string {
	format := "stack jump: tag: '%+v'"
	if symbolp(jmp.tag) {
		return fmt.Sprintf(format, xSymbol(jmp.tag).name)
	}

	return fmt.Sprintf(format, jmp.tag)
}

func (jmp *stackJumpSignal) Error() string {
	format := "stack jump: signal: '%+v'"
	if symbolp(jmp.errorSymbol) {
		return fmt.Sprintf(format, xSymbol(jmp.errorSymbol).name)
	}

	return fmt.Sprintf(format, jmp.errorSymbol)
}

func (sel *stackEntryLet) tag() stackEntryTag {
	return entryLet
}

func (sec *stackEntryCatch) tag() stackEntryTag {
	return entryCatch
}

func (secb *stackEntryCurrentBuffer) tag() stackEntryTag {
	return entryCurrentBuffer
}

func (ec *execContext) makeSymbolBase(name string) *lispSymbol {
	return &lispSymbol{
		name: name,
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

func (ec *execContext) defSubr(name string, sub *subroutine) {
	if sub.maxArgs >= 0 && sub.minArgs > sub.maxArgs {
		ec.terminate("min args must be smaller or equal to max args (subroutine: '%+v')", sub)
	}

	symbol := xEnsure(ec.intern(ec.makeString(name), ec.nil_))
	sym := xSymbol(symbol)

	if sym.function != ec.nil_ {
		ec.terminate("function value already set: '%v'", sym.name)
	}

	vec := ec.makeVectorLike(vectorLikeTypeSubroutine, sub)
	sym.function = vec
}

func (ec *execContext) defSubr0(name string, fn lispFn0) *subroutine {
	sub := &subroutine{
		callabe0: fn,
	}
	ec.defSubr(name, sub)
	return sub
}

func (ec *execContext) defSubr1(name string, fn lispFn1, minArgs int) *subroutine {
	sub := &subroutine{
		callabe1: fn,
		minArgs:  minArgs,
		maxArgs:  1,
	}
	ec.defSubr(name, sub)
	return sub
}

func (ec *execContext) defSubr2(name string, fn lispFn2, minArgs int) *subroutine {
	sub := &subroutine{
		callabe2: fn,
		minArgs:  minArgs,
		maxArgs:  2,
	}
	ec.defSubr(name, sub)
	return sub
}

func (ec *execContext) defSubr3(name string, fn lispFn3, minArgs int) *subroutine {
	sub := &subroutine{
		callabe3: fn,
		minArgs:  minArgs,
		maxArgs:  3,
	}
	ec.defSubr(name, sub)
	return sub
}

func (ec *execContext) defSubr4(name string, fn lispFn4, minArgs int) *subroutine {
	sub := &subroutine{
		callabe4: fn,
		minArgs:  minArgs,
		maxArgs:  4,
	}
	ec.defSubr(name, sub)
	return sub
}

func (ec *execContext) defSubr5(name string, fn lispFn5, minArgs int) *subroutine {
	sub := &subroutine{
		callabe5: fn,
		minArgs:  minArgs,
		maxArgs:  5,
	}
	ec.defSubr(name, sub)
	return sub
}

func (ec *execContext) defSubr6(name string, fn lispFn6, minArgs int) *subroutine {
	sub := &subroutine{
		callabe6: fn,
		minArgs:  minArgs,
		maxArgs:  6,
	}
	ec.defSubr(name, sub)
	return sub
}

func (ec *execContext) defSubr7(name string, fn lispFn7, minArgs int) *subroutine {
	sub := &subroutine{
		callabe7: fn,
		minArgs:  minArgs,
		maxArgs:  7,
	}
	ec.defSubr(name, sub)
	return sub
}

func (ec *execContext) defSubr8(name string, fn lispFn8, minArgs int) *subroutine {
	sub := &subroutine{
		callabe8: fn,
		minArgs:  minArgs,
		maxArgs:  8,
	}
	ec.defSubr(name, sub)
	return sub
}

func (ec *execContext) defSubrM(name string, fn lispFnM, minArgs int) *subroutine {
	sub := &subroutine{
		callabem: fn,
		minArgs:  minArgs,
		maxArgs:  argsMany,
	}
	ec.defSubr(name, sub)
	return sub
}

func (ec *execContext) defSubrU(name string, fn lispFn1, minArgs int) *subroutine {
	sub := &subroutine{
		callabe1: fn,
		minArgs:  minArgs,
		maxArgs:  argsUnevalled,
	}
	ec.defSubr(name, sub)
	return sub
}

func (ec *execContext) defVar(symbol, value lispObject) {
	sym := xSymbol(symbol)
	if sym.value != ec.g.unbound {
		ec.terminate("variable value already set: '%v'", sym.name)
	}

	sym.value = value
}

func newExecContext() *execContext {
	ec := execContext{}

	ec.obarray = make(map[string]*lispSymbol)
	ec.stack = []stackEntry{}
	ec.currentBuffer = &buffer{}

	ec.createSymbols()       // symbols.go
	ec.symbolsOfErrors()     // errors.gmo
	ec.symbolsOfRead()       // read.go
	ec.symbolsOfEval()       // eval.go
	ec.symbolsOfPrint()      // print.go
	ec.symbolsOfData()       // data.go
	ec.symbolsOfAllocation() // allocation.go
	ec.symbolsOfFunctions()  // functions.go
	ec.symbolsOfBuffer()     // buffer.go
	ec.symbolsOfMinibuffer() // minibuffer.go

	ec.defVar(ec.g.nonInteractive, ec.t)

	contents, err := scripts.ReadFile("scripts/pimacs/base.el")
	if err != nil {
		ec.terminate("open script error: '%v'", err)
	}

	source := ec.makeString(string(contents))
	result, _, err := ec.readInternalStart(source, ec.nil_, ec.nil_)
	xEnsure(ec.evalSub(xEnsure(result, err)))

	return &ec
}

func (ec *execContext) internNewSymbol(symbol *lispSymbol) {
	prev, ok := ec.obarray[symbol.name]
	if ok && prev != symbol {
		ec.terminate("different symbol with that name already interned, name: '%v'", symbol.name)
	}
	ec.obarray[symbol.name] = symbol
}

func (ec *execContext) listToSlice(list lispObject) ([]lispObject, error) {
	result := []lispObject{}
	tail := list

	for ; consp(tail); tail = xCdr(tail) {
		result = append(result, xCar(tail))
	}

	if tail != ec.nil_ {
		return nil, xErrOnly(ec.wrongTypeArgument(ec.g.listp, list))
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

func (ec *execContext) stackPushCurrentBuffer(buf *buffer) {
	ec.stack = append(ec.stack, &stackEntryCurrentBuffer{
		prevCurrentBuffer: ec.currentBuffer,
	})

	ec.currentBuffer = buf
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
		case entryCurrentBuffer:
			ec.currentBuffer = current.(*stackEntryCurrentBuffer).prevCurrentBuffer
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

func (ec *execContext) terminate(format string, v ...interface{}) {
	// TAGS: incomplete
	terminate(format, v...)
}
