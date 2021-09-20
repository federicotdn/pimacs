package elisp

func (ec *execContext) initialDefsVariables() {
	// TAGS: revise
	xSymbol(ec.g.internalInterpreterEnv).value = ec.nil_
	xSymbol(ec.g.lexicalBinding).value = ec.nil_
	xSymbol(ec.g.nonInteractive).value = ec.t
}
