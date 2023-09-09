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
	channelp           lispObject
	numberOrMarkerp    lispObject
	charOrStringp      lispObject
	integerp           lispObject
	bufferp            lispObject
	characterp         lispObject
	quote              lispObject
	backquote          lispObject
	comma              lispObject
	commaAt            lispObject
	function           lispObject
	read               lispObject
	equal              lispObject
	eval               lispObject
	setq               lispObject
	prin1              lispObject
	readFromMinibuffer lispObject
	recursiveEdit      lispObject
	charTablep         lispObject
	provide            lispObject
	reverse            lispObject
	mapCar             lispObject
	eql                lispObject

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
	errorConditions     lispObject
	errorMessage        lispObject
	lambda              lispObject
	closure             lispObject
	macro               lispObject
	andRest             lispObject
	andOptional         lispObject
	readChar            lispObject
	charTableExtraSlots lispObject
	wholeNump           lispObject
	cTest               lispObject
	cSize               lispObject
	cPureCopy           lispObject
	cRehashSize         lispObject
	cRehashThreshold    lispObject
	cWeakness           lispObject
	key                 lispObject
	value               lispObject
	hashTableTest       lispObject
	keyOrValue          lispObject
	keyAndValue         lispObject
}

type vars struct {
	// Variables with static Go values
	nonInteractive forwardBool
	standardOutput forwardLispObj
	standardInput  forwardLispObj
	loadPath       forwardLispObj
	pimacsRepo     forwardLispObj
}

func (ec *execContext) initSymbols() {
	g := ec.s

	// Set up nil and unbound first so that we can use
	// ec.makeSymbol()
	unbound := ec.makeSymbol("unbound", false)
	nil_ := ec.makeSymbol("nil", false)

	// Set the attributes and intern nil
	unbound.setAttributes(unbound, nil_, nil_)
	nil_.setAttributes(nil_, nil_, nil_)
	// NOTE: This should be the only access to the obarray not done
	// via ec.internInternal().
	ec.obarray.values["nil"] = nil_

	// nil and unbound are now complete

	// Set them in globals
	g.unbound = unbound
	g.nil_ = nil_
	ec.nil_ = g.nil_ // Convenience accessor

	// Create t
	t := ec.defSym(&g.t, "t")
	t.val = t
	g.t = t
	ec.t = ec.s.t // Convenience accessor

	// Set up other essential symbols
	ec.defSym(&ec.s.emptySymbol, "")
}

func (ec *execContext) checkSymbolValues() {
	v := reflect.ValueOf(*ec.s)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.IsNil() {
			ec.terminate("initialization error: symbol not initialized: 'symbols.%+v'", v.Type().Field(i).Name)
		}
	}
}

func (ec *execContext) checkVarValues() {
	v := reflect.ValueOf(*ec.v)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		sym := field.FieldByName("sym")
		if sym == (reflect.Value{}) {
			ec.terminate("initialization error: invalid forward type: 'vars.%+v'", v.Type().Field(i).Name)
		} else if sym.IsNil() {
			ec.terminate("initialization error: forward value not initialized: 'vars.%+v'", v.Type().Field(i).Name)
		}
	}
}
