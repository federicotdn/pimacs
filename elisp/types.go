package elisp

type lispType int
type lispInt int64

const (
	symbol lispType = iota
	integer
	string_
	vector
	cons
	float
	argsMany      = -1
	argsUnevalled = -2
)

type lispSymbol struct {
	name  string
	value lispObject

	// Should be lispObject instead
	function struct {
		// 'lispObject' or something more expressive? i.e. a struct
		// Return error?
		callabe func(args ...lispObject) (lispObject, error)
		minArgs int
		maxArgs int
	}
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
