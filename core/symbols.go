package core

import (
	"reflect"
)

type symbols struct {
	// Essential runtime objects
	nil_        lispObject
	t           lispObject
	unbound     lispObject
	emptySymbol lispObject

	// Subroutine symbols
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
	function           lispObject
	read               lispObject
	equal              lispObject
	eval               lispObject
	setq               lispObject
	prin1              lispObject
	readFromMinibuffer lispObject

	// Errors
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
	fileMissing            lispObject

	// Misc. symbols
	errorConditions lispObject
	errorMessage    lispObject
	lambda          lispObject
	closure         lispObject
	macro           lispObject
	andRest         lispObject
	andOptional     lispObject
	readChar        lispObject

	// Variables with static Go values
	internalInterpreterEnv forwardLispObj
	nonInteractive         forwardBool
	standardOutput         forwardLispObj
	standardInput          forwardLispObj
	lexicalBinding         forwardLispObj
	loadPath               forwardLispObj
	pimacsRepo             forwardLispObj
}

func (ec *execContext) initSymbols() {
	g := &ec.g

	// Set up nil and unbound first so that we can use
	// ec.makeSymbol()
	unbound := ec.makeSymbolBase("unbound")
	nil_ := ec.makeSymbolBase("nil")

	// Set the attributes and intern nil
	unbound.setAttributes(unbound, nil_, nil_)
	nil_.setAttributes(nil_, nil_, nil_)
	ec.internNewSymbol(nil_, true)

	// nil and unbound are now complete

	// Set them in globals
	g.unbound = unbound
	g.nil_ = nil_
	ec.nil_ = g.nil_ // Convenience accessor

	// Create t
	t := ec.defSym(&g.t, "t")
	t.value = t
	g.t = t
	ec.t = ec.g.t // Convenience accessor

	// Set up other essential symbols
	ec.defSym(&ec.g.emptySymbol, "")
}

func (ec *execContext) checkSymbolValues() {
	v := reflect.ValueOf(ec.g)

	for i := 0; i < v.NumField(); i++ {
		// field := v.Field(i)
		// if field.IsNil() {
		// 	ec.terminate("initialization error: lispGlobals field not set: '%+v'", v.Type().Field(i).Name)
		// }
	}
}
