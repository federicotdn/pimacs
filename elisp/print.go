package elisp

import (
	"fmt"
)

func (ec *execContext) printString(str string, printCharFn lispObject) error {
	// TAGS: incomplete,revise
	if printCharFn == ec.nil_ {
		_, err := ec.insert(ec.makeString(str))
		return err
	}

	if xSymbolValue(ec.g.nonInteractive) == ec.t && printCharFn == ec.t {
		fmt.Printf("%v", str)
		return nil
	}

	if printCharFn == ec.t {
		return xErrOnly(ec.pimacsUnimplemented(ec.g.prin1, "unknown print char function"))
	}

	_, err := ec.funcall(printCharFn, ec.makeString(str))
	return err
}

func (ec *execContext) printInternal(obj, printCharFn lispObject, escapeFlag bool) (lispObject, error) {
	// TAGS: incomplete
	lispType := obj.getType()
	var s string

	switch lispType {
	case lispTypeSymbol:
		s = xSymbol(obj).name
	case lispTypeInteger:
		s = fmt.Sprint(xIntegerValue(obj))
	case lispTypeString:
		s = "\"" + xString(obj).value + "\""
	case lispTypeCons:
		// TODO: This should not use prin1ToString, but printString directly.
		s = "("

		for ; consp(obj); obj = xCdr(obj) {
			car, err := ec.prin1ToString(xCar(obj), printCharFn)
			if err != nil {
				return nil, err
			}

			s += xStringValue(car)

			if xCdr(obj) != ec.nil_ {
				s += " "
			}
		}

		if obj != ec.nil_ {
			end, err := ec.prin1ToString(obj, printCharFn)
			if err != nil {
				return nil, err
			}

			s += ". " + xStringValue(end)
		}

		s += ")"
	case lispTypeFloat:
		s = fmt.Sprint(xFloat(obj).value)
	case lispTypeVectorLike:
		vec := xVectorLike(obj)
		s = fmt.Sprintf("#<vector-like type '%v' value '%+v'>", vec.vecType, vec.value)
	default:
		s = fmt.Sprintf("#<INVALID DATATYPE '%+v'>", obj)
	}

	err := ec.printString(s, printCharFn)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (ec *execContext) prin1ToString(obj, noEscape lispObject) (lispObject, error) {
	size := ec.stackSize()
	ec.stackPushCurrentBuffer(&buffer{})
	defer ec.stackPopTo(size)

	_, err := ec.printInternal(obj, ec.nil_, noEscape == ec.nil_)
	if err != nil {
		return nil, err
	}

	return ec.bufferString()
}

func (ec *execContext) prin1(obj, printCharFn lispObject) (lispObject, error) {
	// TAGS: incomplete
	if printCharFn == ec.nil_ {
		printCharFn = xSymbolValue(ec.g.standardOutput)
	}

	if printCharFn == ec.nil_ {
		printCharFn = ec.t
	}

	if bufferp(printCharFn) {
		size := ec.stackSize()
		ec.stackPushCurrentBuffer(xBuffer(printCharFn))
		defer ec.stackPopTo(size)

		printCharFn = ec.nil_
	}

	_, err := ec.printInternal(obj, printCharFn, true)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (ec *execContext) symbolsOfPrint() {
	ec.defSubr2("prin1", ec.prin1, 1)
	ec.defSubr2("prin1-to-string", ec.prin1ToString, 1)

	ec.defVar(ec.g.standardOutput, ec.t)
}
