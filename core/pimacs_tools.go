package core

func (ec *execContext) pimacsSymbolDebug(symbol lispObject) (lispObject, error) {
	if !symbolp(symbol) {
		return ec.wrongTypeArgument(ec.s.symbolp, symbol)
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
		"name":     newString(sym.name),
		"special":  xEnsure(ec.bool(sym.special)),
		"plist":    sym.plist,
		"redirect": redirect,
	})
}

func (ec *execContext) pimacsImplementationStats() (lispObject, error) {
	data := map[string]lispObject{
		"subroutines": newInteger(lispInt(ec.init.subrs)),
		"stubs":       newInteger(lispInt(ec.init.stubs)),
		"completion":  newFloat(lispFp(100.0 * float32(ec.init.subrs) / float32(ec.init.stubs))),
	}
	return ec.makePlist(data)
}

func (ec *execContext) symbolsOfPimacsTools() {
	ec.defVarLisp(
		&ec.v.pimacsRepo,
		"pimacs--repo",
		newString("https://github.com/federicotdn/pimacs"),
	)

	ec.defSubr1(nil, "pimacs--symbol-debug", ec.pimacsSymbolDebug, 1)
	ec.defSubr0(nil, "pimacs--impl-stats", ec.pimacsImplementationStats)
}
