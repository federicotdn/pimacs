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
	lispTypeVectorLike
	lispTypeCons
	lispTypeFloat
	argsMany      = -1
	argsUnevalled = -2
)

const (
	vectorLikeTypeNormal vectorLikeType = iota + 1
	vectorLikeTypeSubroutine
	vectorLikeTypeBuffer
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

type subroutine struct {
	callabe  lispFn
	minArgs  int
	maxArgs  int
	noReturn bool
	name     string
}

type buffer struct {
	contents string
	live     bool
	name     string
}

type lispSymbol struct {
	name     string
	value    lispObject
	function lispObject
	plist    lispObject
	special  bool
	redirect *forwardValue
}

type forwardValue struct {
	sym     *lispSymbol
	valLisp lispObject
	valBool bool
	valInt  lispInt
	fwdType symbolFwdType
}

type lispCons struct {
	car lispObject
	cdr lispObject
}

type lispInteger struct {
	value lispInt
}

type lispFloat struct {
	value lispFp
}

type lispString struct {
	value string
}

type vectorLikeValue interface{}

type lispVectorLike struct {
	vecType vectorLikeType
	value   vectorLikeValue
}

type lispObject interface {
	getType() lispType
}

func (s *subroutine) setAttrs(noReturn bool) {
	s.noReturn = noReturn
}

func (ls *lispSymbol) getType() lispType {
	return lispTypeSymbol
}

func (ls *lispSymbol) setAttributes(value, function, plist lispObject) {
	ls.value = value
	ls.function = function
	ls.plist = plist
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

func (lv *lispVectorLike) getType() lispType {
	return lispTypeVectorLike
}
