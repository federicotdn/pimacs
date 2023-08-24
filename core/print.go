package core

import (
	"fmt"
)

func (ec *execContext) printString(str string, printCharFn lispObject) error {
	if printCharFn == ec.nil_ {
		_, err := ec.insert(ec.makeString(str))
		return err
	}

	if ec.g.nonInteractive.val && printCharFn == ec.t {
		fmt.Printf("%v", str)
		return nil
	}

	if printCharFn == ec.t {
		return xErrOnly(ec.pimacsUnimplemented(ec.g.prin1, "unknown print char function"))
	}

	_, err := ec.funcall(printCharFn, ec.makeString(str))
	return err
}

func (ec *execContext) printInternal(obj, printCharFn lispObject, escapeFlag bool) error {
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
		err := ec.printString("(", printCharFn)
		if err != nil {
			return err
		}

		for ; consp(obj); obj = xCdr(obj) {
			err := ec.printInternal(xCar(obj), printCharFn, escapeFlag)
			if err != nil {
				return err
			}

			if xCdr(obj) != ec.nil_ {
				err = ec.printString(" ", printCharFn)
				if err != nil {
					return err
				}
			}
		}

		if obj != ec.nil_ {
			err = ec.printString(". ", printCharFn)
			if err != nil {
				return err
			}

			err = ec.printInternal(obj, printCharFn, escapeFlag)
			if err != nil {
				return err
			}
		}

		return ec.printString(")", printCharFn)
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
		return err
	}

	return nil
}

func (ec *execContext) prin1ToString(obj, noEscape lispObject) (lispObject, error) {
	// Should this buffer be created via get-buffer-create?
	// Needs to be hidden from buffer list though
	buf := ec.createBuffer(" prin1")
	bufObj := ec.makeBuffer(buf)

	_, err := ec.prin1(obj, bufObj)
	if err != nil {
		return nil, err
	}

	return ec.makeString(buf.contents), nil
}

func (ec *execContext) printGeneric(obj, printCharFn lispObject, escapeFlag, newlines bool) (lispObject, error) {
	if printCharFn == ec.nil_ {
		printCharFn = ec.g.standardOutput.val
	}

	if printCharFn == ec.nil_ {
		printCharFn = ec.t
	}

	if bufferp(printCharFn) {
		defer ec.unwind()()
		ec.stackPushCurrentBuffer()
		_, err := ec.setBuffer(printCharFn)
		if err != nil {
			return nil, err
		}

		printCharFn = ec.nil_
	}

	if newlines {
		err := ec.printString("\n", printCharFn)
		if err != nil {
			return nil, err
		}
	}

	err := ec.printInternal(obj, printCharFn, escapeFlag)
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
	ec.defVarLisp(&ec.g.standardOutput, "standard-output", ec.t)

	ec.defSubr2(&ec.g.prin1, "prin1", ec.prin1, 1)
	ec.defSubr2(nil, "print", ec.print_, 1)
	ec.defSubr2(nil, "princ", ec.princ, 1)
	ec.defSubr2(nil, "prin1-to-string", ec.prin1ToString, 1)
}
