package elisp

import (
	"reflect"
)

type symbolInit struct {
	loc      *lispObject
	name     string
	unintern bool
}

type lispGlobals struct {
	// 1. Essential runtime objects
	nil_                   lispObject
	t                      lispObject
	internalInterpreterEnv lispObject
	unbound                lispObject
	lexicalBinding         lispObject
	// 2. Subroutine symbols
	sequencep          lispObject
	listp              lispObject
	plistp             lispObject
	consp              lispObject
	symbolp            lispObject
	stringp            lispObject
	numberOrMarkerp    lispObject
	charOrStringp      lispObject
	integerp           lispObject
	bufferp            lispObject
	quote              lispObject
	backquote          lispObject
	progn              lispObject
	function           lispObject
	read               lispObject
	equal              lispObject
	eval               lispObject
	if_                lispObject
	prin1              lispObject
	readFromMinibuffer lispObject
	// 3. Errors
	error_                 lispObject
	quit                   lispObject
	userError              lispObject
	wrongLengthArgument    lispObject
	wrongTypeArgument      lispObject
	argsOutOfRange         lispObject
	voidFunction           lispObject
	invalidFunction        lispObject
	voidVariable           lispObject
	wrongNumberofArguments lispObject
	endOfFile              lispObject
	noCatch                lispObject
	settingConstant        lispObject
	invalidReadSyntax      lispObject
	pimacsUnimplemented    lispObject
	circularList           lispObject
	// 4. Misc. symbols
	errorConditions lispObject
	errorMessage    lispObject
	lambda          lispObject
	closure         lispObject
	macro           lispObject
	andRest         lispObject
	andOptional     lispObject
	readChar        lispObject
	// 5. Variables
	nonInteractive lispObject
	standardInput  lispObject
	standardOutput lispObject
}

func (ec *execContext) createSymbols() {
	g := &ec.g

	syms := []symbolInit{
		// 1
		{loc: &g.unbound, name: "unbound", unintern: true},
		{loc: &g.nil_, name: "nil"},
		{loc: &g.t, name: "t"},
		{loc: &g.internalInterpreterEnv, name: "internal-interpreter-environment", unintern: true},
		{loc: &g.lexicalBinding, name: "lexical-binding"},
		// 2
		{loc: &g.sequencep, name: "sequencep"},
		{loc: &g.listp, name: "listp"},
		{loc: &g.plistp, name: "plistp"},
		{loc: &g.consp, name: "consp"},
		{loc: &g.symbolp, name: "symbolp"},
		{loc: &g.stringp, name: "stringp"},
		{loc: &g.numberOrMarkerp, name: "number-or-marker-p"},
		{loc: &g.charOrStringp, name: "char-or-string-p"},
		{loc: &g.integerp, name: "integerp"},
		{loc: &g.bufferp, name: "bufferp"},
		{loc: &g.quote, name: "quote"},
		{loc: &g.backquote, name: "`"},
		{loc: &g.progn, name: "progn"},
		{loc: &g.function, name: "function"},
		{loc: &g.read, name: "read"},
		{loc: &g.equal, name: "equal"},
		{loc: &g.eval, name: "eval"},
		{loc: &g.if_, name: "if"},
		{loc: &g.prin1, name: "prin1"},
		{loc: &g.readFromMinibuffer, name: "read-from-minibuffer"},
		// 3
		{loc: &g.error_, name: "error"},
		{loc: &g.quit, name: "quit"},
		{loc: &g.userError, name: "user-error"},
		{loc: &g.wrongLengthArgument, name: "wrong-length-argument"},
		{loc: &g.wrongTypeArgument, name: "wrong-type-argument"},
		{loc: &g.argsOutOfRange, name: "args-out-of-range"},
		{loc: &g.voidFunction, name: "void-function"},
		{loc: &g.invalidFunction, name: "invalid-function"},
		{loc: &g.voidVariable, name: "void-variable"},
		{loc: &g.wrongNumberofArguments, name: "wrong-number-of-arguments"},
		{loc: &g.endOfFile, name: "end-of-file"},
		{loc: &g.noCatch, name: "no-catch"},
		{loc: &g.settingConstant, name: "setting-constant"},
		{loc: &g.invalidReadSyntax, name: "invalid-read-syntax"},
		{loc: &g.pimacsUnimplemented, name: "pimacs-unimplemented"},
		{loc: &g.circularList, name: "circular-list"},
		// 4
		{loc: &g.errorConditions, name: "error-conditions"},
		{loc: &g.errorMessage, name: "error-message"},
		{loc: &g.lambda, name: "lambda"},
		{loc: &g.closure, name: "closure"},
		{loc: &g.macro, name: "macro"},
		{loc: &g.andRest, name: "&rest"},
		{loc: &g.andOptional, name: "&optional"},
		{loc: &g.readChar, name: "read-char"},
		// 5
		{loc: &g.nonInteractive, name: "noninteractive"},
		{loc: &g.standardInput, name: "standard-input"},
		{loc: &g.standardOutput, name: "standard-output"},
	}

	ec.initializeSymbols(syms)

	xSymbol(ec.g.nil_).value = ec.g.nil_
	xSymbol(ec.g.t).value = ec.g.t

	ec.nil_ = ec.g.nil_
	ec.t = ec.g.t
}

func (ec *execContext) initializeSymbols(syms []symbolInit) {
	count := ec.countGlobals()
	if len(syms) != count {
		ec.terminate("'%v' globals exist but got '%v' initializers", count, len(syms))
	}

	names := make(map[string]bool)

	for _, sym := range syms {
		if sym.name == "" {
			ec.terminate("no symbol name defined")
		}

		_, ok := names[sym.name]
		if ok {
			ec.terminate("repeated symbol name: '%v'", sym.name)
		}
		names[sym.name] = true

		*sym.loc = ec.makeSymbolBase(sym.name)
		symbol := xSymbol(*sym.loc)

		if ec.g.unbound == nil {
			ec.terminate("'unbound' not initialized yet")
		}

		symbol.value = ec.g.unbound

		if !sym.unintern {
			ec.internNewSymbol(symbol)
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
