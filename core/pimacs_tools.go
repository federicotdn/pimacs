package core

func (ec *execContext) pimacsSymbolDebug(symbol lispObject) (lispObject, error) {
	if !symbolp(symbol) {
		return ec.wrongTypeArgument(ec.s.symbolp, symbol)
	}

	sym := xSymbol(symbol)
	val := xEnsure(ec.findSymbolValue(symbol))
	fwd := ec.nil_
	if sym.fwd != nil {
		fwd = sym.fwd.value(ec)
	}

	return ec.makeKwPlist(map[string]lispObject{
		"value":    val,
		"function": sym.function,
		"name":     newString(sym.name),
		"special":  xEnsure(ec.bool(sym.special)),
		"plist":    sym.plist,
		"redirect": newInteger(lispInt(sym.redirect)),
		"fwd":      fwd,
	})
}

func (ec *execContext) symbolsOfPimacsTools() {
	ec.defVarLisp(
		&ec.v.pimacsRepo,
		"pimacs--repo",
		newString("https://github.com/federicotdn/pimacs"),
	)

	ec.defSubr1(nil, "pimacs--symbol-debug", (*execContext).pimacsSymbolDebug, 1)
}
