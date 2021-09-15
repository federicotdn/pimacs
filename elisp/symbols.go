package elisp

import (
	"fmt"
	"reflect"
)

type symbolInit struct {
	loc      *lispObject
	name     string
	unintern bool
}

type globals struct {
	// 1. Essential runtime objects
	nil_                   lispObject
	t                      lispObject
	internalInterpreterEnv lispObject
	unbound                lispObject
	// 2. Subroutine symbols
	sequencep       lispObject
	listp           lispObject
	symbolp         lispObject
	numberOrMarkerp lispObject
	goChannelp      lispObject
	integerp        lispObject
	quote           lispObject
	progn           lispObject
	function        lispObject
	// 3. Errors
	error_                 lispObject
	quit                   lispObject
	userError              lispObject
	wrongLengthArgument    lispObject
	wrongTypeArgument      lispObject
	argsOutOfRange         lispObject
	voidFunction           lispObject
	voidVariable           lispObject
	wrongNumberofArguments lispObject
	endOfFile              lispObject
	noCatch                lispObject
	settingConstant        lispObject
	invalidReadSyntax      lispObject
	// 4. Misc. symbols
	errorConditions lispObject
	errorMessage    lispObject
	lambda          lispObject
	closure         lispObject
	// 5. Pimacs
	goChannelClosed lispObject
}

func (ec *execContext) initialDefsSymbols() {
	g := &ec.g

	syms := []symbolInit{
		// 1
		{loc: &g.unbound, name: "unbound", unintern: true},
		{loc: &g.nil_, name: "nil"},
		{loc: &g.t, name: "t"},
		{
			loc:      &g.internalInterpreterEnv,
			name:     "internal-interpreter-environment",
			unintern: true,
		},
		// 2
		{loc: &g.sequencep, name: "sequencep"},
		{loc: &g.listp, name: "listp"},
		{loc: &g.symbolp, name: "symbolp"},
		{loc: &g.numberOrMarkerp, name: "number-or-marker-p"},
		{loc: &g.goChannelp, name: "pimacs-go-channel-p"},
		{loc: &g.integerp, name: "integerp"},
		{loc: &g.quote, name: "quote"},
		{loc: &g.progn, name: "progn"},
		{loc: &g.function, name: "function"},
		// 3
		{loc: &g.error_, name: "error"},
		{loc: &g.quit, name: "quit"},
		{loc: &g.userError, name: "user-error"},
		{loc: &g.wrongLengthArgument, name: "wrong-length-argument"},
		{loc: &g.wrongTypeArgument, name: "wrong-type-argument"},
		{loc: &g.argsOutOfRange, name: "args-out-of-range"},
		{loc: &g.voidFunction, name: "void-function"},
		{loc: &g.voidVariable, name: "void-variable"},
		{loc: &g.wrongNumberofArguments, name: "wrong-number-of-arguments"},
		{loc: &g.endOfFile, name: "end-of-file"},
		{loc: &g.noCatch, name: "no-catch"},
		{loc: &g.settingConstant, name: "setting-constant"},
		{loc: &g.invalidReadSyntax, name: "invalid-read-syntax"},
		// 4
		{loc: &g.errorConditions, name: "error-conditions"},
		{loc: &g.errorMessage, name: "error-message"},
		{loc: &g.lambda, name: "lambda"},
		{loc: &g.closure, name: "closure"},
		// 5
		{loc: &g.goChannelClosed, name: "go-channel-closed"},
	}

	ec.initializeSymbols(syms)

	xSymbol(ec.g.nil_).value = ec.g.nil_
	xSymbol(ec.g.t).value = ec.g.t

	ec.nil_ = ec.g.nil_
	ec.t = ec.g.t
}

func (ec *execContext) initializeSymbols(syms []symbolInit) {
	// TAGS: errors
	count := ec.countGlobals()
	if len(syms) != count {
		panic(fmt.Sprintf("%v globals exist but got %v initializers", count, len(syms)))
	}

	names := make(map[string]bool)

	for _, sym := range syms {
		if sym.name == "" {
			panic("no symbol name defined")
		}

		_, ok := names[sym.name]
		if ok {
			panic(fmt.Sprintf("repeated symbol name: %v", sym.name))
		}
		names[sym.name] = true

		*sym.loc = ec.makeSymbolBase(sym.name)
		symbol := xSymbol(*sym.loc)

		if ec.g.unbound == nil {
			panic("unbound not initialized yet")
		}

		symbol.value = ec.g.unbound

		if !sym.unintern {
			ec.internSymbol(symbol)
		}
	}

	for _, sym := range syms {
		symbol := xSymbol(*sym.loc)
		symbol.function = ec.g.nil_
		symbol.plist = ec.g.nil_
	}
}

func (ec *execContext) countGlobals() int {
	return reflect.ValueOf(ec.g).NumField()
}
