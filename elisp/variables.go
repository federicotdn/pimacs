package elisp

func (ec *execContext) initialDefsVariables() {
	ec.globals.internalInterpreterEnv = ec.nil_
}
