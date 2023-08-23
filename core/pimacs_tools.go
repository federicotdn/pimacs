package core

func (ec *execContext) pimacsSymbolDebug(symbol lispObject) (lispObject, error) {
	if !symbolp(symbol) {
		return ec.wrongTypeArgument(ec.g.symbolp, symbol)
	}

	sym := xSymbol(symbol)
	val := xEnsure(ec.findSymbolValue(symbol))
	redirect := ec.nil_
	if sym.redirect != nil {
		redirect = sym.redirect.value(ec)
	}

	return ec.makePlist(map[string]lispObject{
		"value":    val,
		"function": sym.function,
		"name":     ec.makeString(sym.name),
		"special":  xEnsure(ec.bool(sym.special)),
		"plist":    sym.plist,
		"redirect": redirect,
	}), nil
}

func (ec *execContext) symbolsOfPimacsTools() {
	ec.defVarLisp(
		&ec.g.pimacsRepo,
		"pimacs-repo",
		ec.makeString("https://github.com/federicotdn/pimacs"),
	)

	ec.defSubr1(nil, "pimacs-symbol-debug", ec.pimacsSymbolDebug, 1)
}
