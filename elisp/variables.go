package elisp

func (ec *execContext) initialDefsVariables() {
	xSymbol(ec.g.internalInterpreterEnv).value = ec.nil_
}
