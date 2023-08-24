package core

import (
	"fmt"
)

func (ec *execContext) signalN(errorSymbol lispObject, args ...lispObject) (lispObject, error) {
	list := ec.makeList(args...)
	return ec.signal(errorSymbol, list)
}

func (ec *execContext) signalError(format string, v ...interface{}) (lispObject, error) {
	message := ec.makeString(fmt.Sprintf(format, v...))
	return ec.signalN(ec.s.error_, message)
}

func (ec *execContext) wrongTypeArgument(predicate, value lispObject) (lispObject, error) {
	return ec.signalN(ec.s.wrongTypeArgument, predicate, value)
}

func (ec *execContext) wrongNumberOfArguments(fn lispObject, count lispInt) (lispObject, error) {
	return ec.signalN(ec.s.wrongNumberofArguments, fn, ec.makeInteger(count))
}

func (ec *execContext) pimacsUnimplemented(fn lispObject, format string, v ...interface{}) (lispObject, error) {
	message := ec.makeString(fmt.Sprintf(format, v...))
	return ec.signalN(ec.s.pimacsUnimplemented, message)
}

func (ec *execContext) invalidReadSyntax(format string, v ...interface{}) (lispObject, error) {
	message := ec.makeString(fmt.Sprintf(format, v...))
	return ec.signalN(ec.s.invalidReadSyntax, message)
}

func (ec *execContext) putError(sym, tail lispObject, msg string) {
	xEnsure(ec.put(sym, ec.s.errorConditions, ec.makeCons(sym, tail)))
	xEnsure(ec.put(sym, ec.s.errorMessage, ec.makeString(msg)))
}

func (ec *execContext) symbolsOfErrors() {
	ec.defSym(&ec.s.error_, "error")
	ec.defSym(&ec.s.quit, "quit")
	ec.defSym(&ec.s.userError, "user-error")
	ec.defSym(&ec.s.wrongLengthArgument, "wrong-length-argument")
	ec.defSym(&ec.s.wrongTypeArgument, "wrong-type-argument")
	ec.defSym(&ec.s.argsOutOfRange, "args-out-of-range")
	ec.defSym(&ec.s.voidFunction, "void-function")
	ec.defSym(&ec.s.invalidFunction, "invalid-function")
	ec.defSym(&ec.s.voidVariable, "void-variable")
	ec.defSym(&ec.s.wrongNumberofArguments, "wrong-number-of-arguments")
	ec.defSym(&ec.s.endOfFile, "end-of-file")
	ec.defSym(&ec.s.noCatch, "no-catch")
	ec.defSym(&ec.s.settingConstant, "setting-constant")
	ec.defSym(&ec.s.invalidReadSyntax, "invalid-read-syntax")
	ec.defSym(&ec.s.pimacsUnimplemented, "pimacs-unimplemented")
	ec.defSym(&ec.s.circularList, "circular-list")
	ec.defSym(&ec.s.fileMissing, "file-missing")
	ec.defSym(&ec.s.errorConditions, "error-conditions")
	ec.defSym(&ec.s.errorMessage, "error-message")

	errorTail := ec.makeList(ec.s.error_)

	xEnsure(ec.put(ec.s.error_, ec.s.errorConditions, ec.makeList(ec.s.error_)))
	xEnsure(ec.put(ec.s.error_, ec.s.errorMessage, ec.makeString("error")))

	ec.putError(ec.s.quit, ec.nil_, "Quit")
	ec.putError(ec.s.userError, errorTail, "")
	ec.putError(ec.s.wrongLengthArgument, errorTail, "Wrong length argument")
	ec.putError(ec.s.wrongTypeArgument, errorTail, "Wrong type argument")
	ec.putError(ec.s.argsOutOfRange, errorTail, "Args out of range")
	ec.putError(ec.s.voidFunction, errorTail, "Symbol's function definition is void")
	ec.putError(ec.s.invalidFunction, errorTail, "Invalid function")
	ec.putError(ec.s.voidVariable, errorTail, "Symbol's value as variable is void")
	ec.putError(ec.s.wrongNumberofArguments, errorTail, "Wrong number of arguments")
	ec.putError(ec.s.endOfFile, errorTail, "End of file during parsing")
	ec.putError(ec.s.noCatch, errorTail, "No catch for tag")
	ec.putError(ec.s.settingConstant, errorTail, "Attempt to set a constant symbol")
	ec.putError(ec.s.invalidReadSyntax, errorTail, "Invalid read syntax")
	ec.putError(ec.s.pimacsUnimplemented, errorTail, "Unimplemented feature")
	ec.putError(ec.s.circularList, errorTail, "List contains a loop")
	ec.putError(ec.s.fileMissing, errorTail, "No such file or directory")
}
