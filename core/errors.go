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
	return ec.signalN(ec.g.error_, message)
}

func (ec *execContext) wrongTypeArgument(predicate, value lispObject) (lispObject, error) {
	return ec.signalN(ec.g.wrongTypeArgument, predicate, value)
}

func (ec *execContext) wrongNumberOfArguments(fn lispObject, count lispInt) (lispObject, error) {
	return ec.signalN(ec.g.wrongNumberofArguments, fn, ec.makeInteger(count))
}

func (ec *execContext) pimacsUnimplemented(fn lispObject, format string, v ...interface{}) (lispObject, error) {
	message := ec.makeString(fmt.Sprintf(format, v...))
	return ec.signalN(ec.g.pimacsUnimplemented, message)
}

func (ec *execContext) invalidReadSyntax(format string, v ...interface{}) (lispObject, error) {
	message := ec.makeString(fmt.Sprintf(format, v...))
	return ec.signalN(ec.g.invalidReadSyntax, message)
}

func (ec *execContext) putError(sym, tail lispObject, msg string) {
	xEnsure(ec.put(sym, ec.g.errorConditions, ec.makeCons(sym, tail)))
	xEnsure(ec.put(sym, ec.g.errorMessage, ec.makeString(msg)))
}

func (ec *execContext) symbolsOfErrors() {
	ec.defVar(&ec.g.error_, "error", ec.g.unbound)
	ec.defVar(&ec.g.quit, "quit", ec.g.unbound)
	ec.defVar(&ec.g.userError, "user-error", ec.g.unbound)
	ec.defVar(&ec.g.wrongLengthArgument, "wrong-length-argument", ec.g.unbound)
	ec.defVar(&ec.g.wrongTypeArgument, "wrong-type-argument", ec.g.unbound)
	ec.defVar(&ec.g.argsOutOfRange, "args-out-of-range", ec.g.unbound)
	ec.defVar(&ec.g.voidFunction, "void-function", ec.g.unbound)
	ec.defVar(&ec.g.invalidFunction, "invalid-function", ec.g.unbound)
	ec.defVar(&ec.g.voidVariable, "void-variable", ec.g.unbound)
	ec.defVar(&ec.g.wrongNumberofArguments, "wrong-number-of-arguments", ec.g.unbound)
	ec.defVar(&ec.g.endOfFile, "end-of-file", ec.g.unbound)
	ec.defVar(&ec.g.noCatch, "no-catch", ec.g.unbound)
	ec.defVar(&ec.g.settingConstant, "setting-constant", ec.g.unbound)
	ec.defVar(&ec.g.invalidReadSyntax, "invalid-read-syntax", ec.g.unbound)
	ec.defVar(&ec.g.pimacsUnimplemented, "pimacs-unimplemented", ec.g.unbound)
	ec.defVar(&ec.g.circularList, "circular-list", ec.g.unbound)
	ec.defVar(&ec.g.fileMissing, "file-missing", ec.g.unbound)
	ec.defVarUninterned(&ec.g.errorConditions, "error-conditions", ec.g.unbound)
	ec.defVarUninterned(&ec.g.errorMessage, "error-message", ec.g.unbound)

	errorTail := ec.makeList(ec.g.error_)

	xEnsure(ec.put(ec.g.error_, ec.g.errorConditions, ec.makeList(ec.g.error_)))
	xEnsure(ec.put(ec.g.error_, ec.g.errorMessage, ec.makeString("error")))

	ec.putError(ec.g.quit, ec.nil_, "Quit")
	ec.putError(ec.g.userError, errorTail, "")
	ec.putError(ec.g.wrongLengthArgument, errorTail, "Wrong length argument")
	ec.putError(ec.g.wrongTypeArgument, errorTail, "Wrong type argument")
	ec.putError(ec.g.argsOutOfRange, errorTail, "Args out of range")
	ec.putError(ec.g.voidFunction, errorTail, "Symbol's function definition is void")
	ec.putError(ec.g.invalidFunction, errorTail, "Invalid function")
	ec.putError(ec.g.voidVariable, errorTail, "Symbol's value as variable is void")
	ec.putError(ec.g.wrongNumberofArguments, errorTail, "Wrong number of arguments")
	ec.putError(ec.g.endOfFile, errorTail, "End of file during parsing")
	ec.putError(ec.g.noCatch, errorTail, "No catch for tag")
	ec.putError(ec.g.settingConstant, errorTail, "Attempt to set a constant symbol")
	ec.putError(ec.g.invalidReadSyntax, errorTail, "Invalid read syntax")
	ec.putError(ec.g.pimacsUnimplemented, errorTail, "Unimplemented feature")
	ec.putError(ec.g.circularList, errorTail, "List contains a loop")
	ec.putError(ec.g.fileMissing, errorTail, "No such file or directory")
}
