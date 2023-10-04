package core

type enum int
type lispType enum
type symbolRedirectType enum

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
	lispTypeCharTable
	lispTypeHashTable
	lispTypeChannel
	lispTypeRecord
	lispTypeMarker
	argsMany      = -1
	argsUnevalled = -2
)

const (
	symbolRedirectPlain symbolRedirectType = iota
	symbolRedirectAlias
	symbolRedirectLocal
	symbolRedirectFwd
)

type lispFn0 func(*execContext) (lispObject, error)
type lispFn1 func(*execContext, lispObject) (lispObject, error)
type lispFn2 func(*execContext, lispObject, lispObject) (lispObject, error)
type lispFn3 func(*execContext, lispObject, lispObject, lispObject) (lispObject, error)
type lispFn4 func(*execContext, lispObject, lispObject, lispObject, lispObject) (lispObject, error)
type lispFn5 func(*execContext, lispObject, lispObject, lispObject, lispObject, lispObject) (lispObject, error)
type lispFn6 func(*execContext, lispObject, lispObject, lispObject, lispObject, lispObject, lispObject) (lispObject, error)
type lispFn7 func(*execContext, lispObject, lispObject, lispObject, lispObject, lispObject, lispObject, lispObject) (lispObject, error)
type lispFn8 func(*execContext, lispObject, lispObject, lispObject, lispObject, lispObject, lispObject, lispObject, lispObject) (lispObject, error)
type lispFnM func(*execContext, ...lispObject) (lispObject, error)
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
	redirect symbolRedirectType
	fwd      forward
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

type lispRecord struct {
	val   []lispObject
	type_ lispObject
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

type lispCharTableEntry struct {
	start lispInt
	end   lispInt
	val   lispObject
}

type lispCharTable struct {
	val        []lispCharTableEntry
	subtype    *lispSymbol
	parent     *lispCharTable
	defaultVal lispObject
	extraSlots lispInt
}

type lispHashTableTest struct {
	name         lispObject
	hashFunction lispObject
	compFunction lispObject
}

type lispHashTableEntry struct {
	key  lispObject
	val  lispObject
	code lispInt
}

type lispHashTable struct {
	val  map[lispInt][]lispHashTableEntry
	test *lispHashTableTest
}

type lispChannel struct {
	val chan lispObject
}

func (s *lispSubroutine) setAttrs(noReturn bool) {
	s.noReturn = noReturn
}

func (ls *lispSymbol) getType() lispType {
	return lispTypeSymbol
}

func (ls *lispSymbol) setAttributes(value, function, plist lispObject, special bool) {
	ls.val = value
	ls.function = function
	ls.plist = plist
	ls.special = special
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

func (lb *lispBuffer) getType() lispType {
	return lispTypeBuffer
}

func (ct *lispCharTable) getType() lispType {
	return lispTypeCharTable
}

func (ht *lispHashTable) getType() lispType {
	return lispTypeHashTable
}

func (ch *lispChannel) getType() lispType {
	return lispTypeChannel
}

func (e *lispCharTableEntry) contains(c lispInt) bool {
	return e.start <= c && c <= e.end
}
