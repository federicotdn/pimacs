package elisp

type lispType int
type vectorLikeType int

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

type subroutine struct {
	callabe0 lispFn0
	callabe1 lispFn1
	callabe2 lispFn2
	callabe3 lispFn3
	callabe4 lispFn4
	callabe5 lispFn5
	callabe6 lispFn6
	callabe7 lispFn7
	callabe8 lispFn8
	callabem lispFnM
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

func (ls *lispSymbol) getType() lispType {
	return lispTypeSymbol
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
