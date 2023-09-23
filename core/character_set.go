package core

func (ec *execContext) symbolsOfCharacterSet() {
	ec.defSym(&ec.s.emacs, "emacs")
}
