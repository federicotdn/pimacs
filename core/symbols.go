package core

import (
	"reflect"
)

type symbols struct {
	// Essential runtime objects
	nil_    lispObject
	t       lispObject
	unbound lispObject

	// Subroutine symbols
	sequencep                      lispObject
	listp                          lispObject
	plistp                         lispObject
	consp                          lispObject
	symbolp                        lispObject
	stringp                        lispObject
	channelp                       lispObject
	numberOrMarkerp                lispObject
	integerOrMarkerp               lispObject
	charOrStringp                  lispObject
	hashTablep                     lispObject
	integerp                       lispObject
	bufferp                        lispObject
	keymapp                        lispObject
	characterp                     lispObject
	quote                          lispObject
	backquote                      lispObject
	comma                          lispObject
	commaAt                        lispObject
	function                       lispObject
	read                           lispObject
	equal                          lispObject
	eval                           lispObject
	setq                           lispObject
	prin1                          lispObject
	readFromMinibuffer             lispObject
	recursiveEdit                  lispObject
	charTablep                     lispObject
	reverse                        lispObject
	mapCar                         lispObject
	eql                            lispObject
	eq                             lispObject
	sxHashEq                       lispObject
	sxHashEql                      lispObject
	sxHashEqual                    lispObject
	sxHashEqualIncludingProperties lispObject

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
	arithError             lispObject

	// Misc. symbols
	errorConditions       lispObject
	errorMessage          lispObject
	lambda                lispObject
	closure               lispObject
	macro                 lispObject
	andRest               lispObject
	andOptional           lispObject
	readChar              lispObject
	charTableExtraSlots   lispObject
	wholeNump             lispObject
	cTest                 lispObject
	cSize                 lispObject
	cPureCopy             lispObject
	cRehashSize           lispObject
	cRehashThreshold      lispObject
	cWeakness             lispObject
	key                   lispObject
	value                 lispObject
	hashTableTest         lispObject
	keyOrValue            lispObject
	keyAndValue           lispObject
	variableDocumentation lispObject
	riskyLocalVariable    lispObject
	emacs                 lispObject
	subfeatures           lispObject
	keymap                lispObject
}

type vars struct {
	// Variables with static Go values
	nonInteractive forwardBool
	standardOutput forwardLispObj
	standardInput  forwardLispObj
	loadPath       forwardLispObj
	pimacsRepo     forwardLispObj
	features       forwardLispObj
	systemType     forwardLispObj
}

func (ec *execContext) initSymbols() {
	s := ec.s

	// Set up nil and unbound first so that we can use ec.defSym()
	// and other symbol-related functions
	unbound := ec.makeSymbol("unbound", false)
	nil_ := ec.makeSymbol("nil", false)

	// Set their attributes
	unbound.setAttributes(unbound, nil_, nil_, false)
	nil_.setAttributes(nil_, nil_, nil_, true)

	// Intern nil
	ec.obarray.loadOrStoreSymbol(nil_)

	// nil and unbound are now complete, next up set them in
	// execContext.symbols
	s.unbound = unbound
	s.nil_ = nil_

	// Create t
	t := ec.defSym(&s.t, "t")
	t.val = t
	t.special = true
	s.t = t

	// Set up convenience accessors in execContext
	ec.t = s.t
	ec.nil_ = s.nil_
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
