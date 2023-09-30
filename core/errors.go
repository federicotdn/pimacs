package core

import (
	"fmt"
)

func (ec *execContext) signalN(errorSymbol lispObject, args ...lispObject) (lispObject, error) {
	list := ec.makeList(args...)
	return ec.signal(errorSymbol, list)
}

func (ec *execContext) signalError(format string, v ...interface{}) (lispObject, error) {
	message := newString(fmt.Sprintf(format, v...))
	return ec.signalN(ec.s.error_, message)
}

func (ec *execContext) wrongTypeArgument(predicate, value lispObject) (lispObject, error) {
	return ec.signalN(ec.s.wrongTypeArgument, predicate, value)
}

func (ec *execContext) wrongNumberOfArguments(fn lispObject, count lispInt) (lispObject, error) {
	return ec.signalN(ec.s.wrongNumberofArguments, fn, newInteger(count))
}

func (ec *execContext) pimacsUnimplemented(fn lispObject, format string, v ...interface{}) (lispObject, error) {
	message := newString(fmt.Sprintf(format, v...))
	return ec.signalN(ec.s.pimacsUnimplemented, message)
}

func (ec *execContext) invalidReadSyntax(format string, v ...interface{}) (lispObject, error) {
	message := newString(fmt.Sprintf(format, v...))
	return ec.signalN(ec.s.invalidReadSyntax, message)
}

// func (ec *execContext) putError(sym, tail lispObject, msg string) {
// 	xEnsure(ec.put(sym, ec.s.errorConditions, newCons(sym, tail)))
// 	xEnsure(ec.put(sym, ec.s.errorMessage, newString(msg)))
// }

func (ec *execContext) defError(symbol *lispObject, tail lispObject, name, msg string) {
	ec.defSym(symbol, name)
	xEnsure(ec.put(*symbol, ec.s.errorConditions, newCons(*symbol, tail)))
	xEnsure(ec.put(*symbol, ec.s.errorMessage, newString(msg)))
}

func (ec *execContext) symbolsOfErrors() {
	ec.defSym(&ec.s.errorConditions, "error-conditions")
	ec.defSym(&ec.s.errorMessage, "error-message")

	ec.defError(&ec.s.error_, ec.nil_, "error", "error")
	errorTail := ec.makeList(ec.s.error_)

	ec.defError(&ec.s.quit, ec.nil_, "quit", "Quit")
	ec.defError(&ec.s.userError, errorTail, "user-error", "")
	ec.defError(&ec.s.wrongLengthArgument, errorTail, "wrong-length-argument", "Wrong length argument")
	ec.defError(&ec.s.wrongTypeArgument, errorTail, "wrong-type-argument", "Wrong type argument")
	ec.defError(&ec.s.argsOutOfRange, errorTail, "args-out-of-range", "Args out of range")
	ec.defError(&ec.s.voidFunction, errorTail, "void-function", "Symbol's function definition is void")
	ec.defError(&ec.s.invalidFunction, errorTail, "invalid-function", "Invalid function")
	ec.defError(&ec.s.voidVariable, errorTail, "void-variable", "Symbol's value as variable is void")
	ec.defError(&ec.s.wrongNumberofArguments, errorTail, "wrong-number-of-arguments", "Wrong number of arguments")
	ec.defError(&ec.s.endOfFile, errorTail, "end-of-file", "End of file during parsing")
	ec.defError(&ec.s.noCatch, errorTail, "no-catch", "No catch for tag")
	ec.defError(&ec.s.settingConstant, errorTail, "setting-constant", "Attempt to set a constant symbol")
	ec.defError(&ec.s.invalidReadSyntax, errorTail, "invalid-read-syntax", "Invalid read syntax")
	ec.defError(&ec.s.pimacsUnimplemented, errorTail, "pimacs-unimplemented", "Unimplemented feature")
	ec.defError(&ec.s.circularList, errorTail, "circular-list", "List contains a loop")
	ec.defError(&ec.s.fileMissing, errorTail, "file-missing", "No such file or directory")
	ec.defError(&ec.s.arithError, errorTail, "arith-error", "Arithmetic error")
}
