package elisp

func (ec *execContext) initialDefsVariables() {
	ec.g.internalInterpreterEnv = ec.nil_
}
