package core

type lispType int
type vectorLikeType int
type symbolFwdType int

type lispInt int64
type lispFp float64

const (
	lispTypeSymbol lispType = iota + 1
	lispTypeInteger
	lispTypeString
	lispTypeCons
	lispTypeFloat
	lispTypeVector
	lispTypeSubroutine
	lispTypeBuffer
	argsMany      = -1
	argsUnevalled = -2
)

const (
	symbolFwdTypeLispObj symbolFwdType = iota + 1
	symbolFwdTypeInt
	symbolFwdTypeBool
)

type lispFn0 func() (lispObject, error)
type lispFn1 func(lispObject) (lispObject, error)
type lispFn2 func(lispObject, lispObject) (lispObject, error)
type lispFn3 func(lispObject, lispObject, lispObject) (lispObject, error)
type lispFn4 func(lispObject, lispObject, lispObject, lispObject) (lispObject, error)
type lispFn5 func(lispObject, lispObject, lispObject, lispObject, lispObject) (lispObject, error)
type lispFn6 func(lispObject, lispObject, lispObject, lispObject, lispObject, lispObject) (lispObject, error)
type lispFn7 func(lispObject, lispObject, lispObject, lispObject, lispObject, lispObject, lispObject) (lispObject, error)
type lispFn8 func(lispObject, lispObject, lispObject, lispObject, lispObject, lispObject, lispObject, lispObject) (lispObject, error)
type lispFnM func(...lispObject) (lispObject, error)
type lispFn interface{}

type lispObject interface {
	getType() lispType
}

type lispSymbol struct {
	name     string
	val      lispObject
	function lispObject
	plist    lispObject
	special  bool
	redirect forward
}

type forward interface {
	value(*execContext) lispObject
	setValue(*execContext, lispObject) error
}

type forwardBase struct {
	sym *lispSymbol
}

type forwardBool struct {
	forwardBase
	val bool
}

type forwardLispObj struct {
	forwardBase
	val lispObject
}

type lispCons struct {
	car lispObject
	cdr lispObject
}

type lispInteger struct {
	val lispInt
}

type lispFloat struct {
	val lispFp
}

type lispString struct {
	val string
}

type lispVector struct {
	val []lispObject
}

type lispSubroutine struct {
	callabe  lispFn
	minArgs  int
	maxArgs  int
	noReturn bool
	name     string
}

type lispBuffer struct {
	contents string
	live     bool
	name     string
}

func (s *lispSubroutine) setAttrs(noReturn bool) {
	s.noReturn = noReturn
}

func (ls *lispSymbol) getType() lispType {
	return lispTypeSymbol
}

func (ls *lispSymbol) setAttributes(value, function, plist lispObject) {
	ls.val = value
	ls.function = function
	ls.plist = plist
}

func (fb *forwardBool) value(ec *execContext) lispObject {
	if fb.val {
		return ec.t
	}
	return ec.nil_
}

func (fb *forwardBool) setValue(ec *execContext, val lispObject) error {
	if val == ec.nil_ {
		fb.val = false
	} else {
		fb.val = true
	}
	return nil
}

func (fl *forwardLispObj) value(_ *execContext) lispObject {
	return fl.val
}

func (fl *forwardLispObj) setValue(ec *execContext, val lispObject) error {
	fl.val = val
	return nil
}

func (lc *lispCons) getType() lispType {
	return lispTypeCons
}

func (li *lispInteger) getType() lispType {
	return lispTypeInteger
}

func (lf *lispFloat) getType() lispType {
	return lispTypeFloat
}

func (ls *lispString) getType() lispType {
	return lispTypeString
}

func (lv *lispVector) getType() lispType {
	return lispTypeVector
}

func (ls *lispSubroutine) getType() lispType {
	return lispTypeSubroutine
}

func (ls *lispBuffer) getType() lispType {
	return lispTypeBuffer
}
