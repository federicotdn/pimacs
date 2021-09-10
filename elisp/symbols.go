package elisp

type symbolInit struct {
	loc      *lispObject
	name     string
	unintern bool
}

func (ec *execContext) initializeSymbols(syms []symbolInit) {
	for _, sym := range syms {
		if sym.name == "" {
			panic("no symbol name defined")
		}

		*sym.loc = ec.makeSymbolBase(sym.name)
		symbol := xSymbol(*sym.loc)

		if ec.globals.unbound == nil {
			panic("unbound not initialized yet")
		}

		symbol.value = ec.globals.unbound

		if !sym.unintern {
			ec.internSymbol(symbol)
		}
	}

	for _, sym := range syms {
		symbol := xSymbol(*sym.loc)
		symbol.function = ec.globals.nil_
		symbol.plist = ec.globals.nil_
	}
}

func (ec *execContext) initialDefsSymbols() {
	g := &ec.globals

	syms := []symbolInit{
		{loc: &g.unbound, name: "unbound", unintern: true},
		{loc: &g.nil_, name: "nil"},
		{loc: &g.t, name: "t"},
		{
			loc:      &g.internalInterpreterEnv,
			name:     "internal-interpreter-environment",
			unintern: true,
		},
	}

	ec.initializeSymbols(syms)

	xSymbol(ec.globals.nil_).value = ec.globals.nil_
	xSymbol(ec.globals.t).value = ec.globals.t

	ec.nil_ = ec.globals.nil_
}
