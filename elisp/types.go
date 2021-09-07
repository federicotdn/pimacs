package elisp

type lispType int
type lispInt int64

const (
	symbol lispType = iota
	integer
	string_
	vectorLike
	cons
	float
	argsMany      = -1
	argsUnevalled = -2
)

type lispFn0 func() (lispObject, error)
type lispFn1 func(lispObject) (lispObject, error)
type lispFn2 func(lispObject, lispObject) (lispObject, error)
type lispFnM func(...lispObject) (lispObject, error)

type builtInFunction struct {
	callabe0 lispFn0
	callabe1 lispFn1
	callabe2 lispFn2
	callabem lispFnM
	minArgs  int
	maxArgs  int
}

type lispSymbol struct {
	name  string
	value lispObject

	// TODO: Should be lispObject instead (?)
	function *builtInFunction
}

type lispCons struct {
	car lispObject
	cdr lispObject
}

type lispInteger struct {
	value lispInt
}

type lispFloat struct {
	value float64
}

type lispString struct {
	value string
}

type lispObject interface {
	getType() lispType
}

func (ls *lispSymbol) getType() lispType {
	return symbol
}

func (lc *lispCons) getType() lispType {
	return cons
}

func (li *lispInteger) getType() lispType {
	return integer
}

func (lf *lispFloat) getType() lispType {
	return float
}

func (ls *lispString) getType() lispType {
	return string_
}
