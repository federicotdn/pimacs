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
		if escapeFlag {
			s = "\"" + xString(obj).value + "\""
		} else {
			s = xString(obj).value
		}
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
	// TAGS: revise
	buf := &buffer{}
	vec := ec.makeVectorLike(vectorLikeTypeBuffer, buf)

	_, err := ec.prin1(obj, vec)
	if err != nil {
		return nil, err
	}

	return ec.makeString(buf.contents), nil
}

func (ec *execContext) printGeneric(obj, printCharFn lispObject, escapeFlag, newlines bool) (lispObject, error) {
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

	if newlines {
		err := ec.printString("\n", printCharFn)
		if err != nil {
			return nil, err
		}
	}

	_, err := ec.printInternal(obj, printCharFn, escapeFlag)
	if err != nil {
		return nil, err
	}

	if newlines {
		err := ec.printString("\n", printCharFn)
		if err != nil {
			return nil, err
		}
	}

	return obj, nil
}

func (ec *execContext) prin1(obj, printCharFn lispObject) (lispObject, error) {
	return ec.printGeneric(obj, printCharFn, true, false)
}

func (ec *execContext) print_(obj, printCharFn lispObject) (lispObject, error) {
	return ec.printGeneric(obj, printCharFn, true, true)
}

func (ec *execContext) princ(obj, printCharFn lispObject) (lispObject, error) {
	return ec.printGeneric(obj, printCharFn, false, false)
}

func (ec *execContext) symbolsOfPrint() {
	ec.defSubr2("prin1", ec.prin1, 1)
	ec.defSubr2("print", ec.print_, 1)
	ec.defSubr2("princ", ec.princ, 1)
	ec.defSubr2("prin1-to-string", ec.prin1ToString, 1)

	ec.defVar(ec.g.standardOutput, ec.t)
}
