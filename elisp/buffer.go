package elisp

func (ec *execContext) insert(args ...lispObject) (lispObject, error) {
	// TAGS: revise
	for _, arg := range args {
		if characterp(arg) {
			ec.currentBuffer.contents += string(xIntegerRune(arg))
		} else if stringp(arg) {
			ec.currentBuffer.contents += xStringValue(arg)
		} else {
			return ec.wrongTypeArgument(ec.g.charOrStringp, arg)
		}
	}

	return ec.nil_, nil
}

func (ec *execContext) bufferString() (lispObject, error) {
	return ec.makeString(ec.currentBuffer.contents), nil
}

func (ec *execContext) symbolsOfBuffer() {
	ec.defSubr0("buffer-string", ec.bufferString)
	ec.defSubrM("insert", ec.insert, 0)
}
