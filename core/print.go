package core

import (
	"fmt"
	"strconv"
	"strings"
)

func (ec *execContext) printStringE(str string, printCharFn lispObject, err error) error {
	if err != nil {
		return err
	}
	return ec.printString(str, printCharFn)
}

func (ec *execContext) printString(str string, printCharFn lispObject) error {
	if printCharFn == ec.nil_ {
		_, err := ec.insert(newString(str))
		return err
	}

	if ec.v.nonInteractive.val && printCharFn == ec.t {
		fmt.Printf("%v", str)
		return nil
	}

	if printCharFn == ec.t {
		return xErrOnly(ec.pimacsUnimplemented(ec.s.prin1, "unknown print char function"))
	}

	_, err := ec.funcall(printCharFn, newString(str))
	return err
}

func (ec *execContext) printInternalE(obj, printCharFn lispObject, escapeFlag bool, err error) error {
	if err != nil {
		return err
	}
	return ec.printInternal(obj, printCharFn, escapeFlag)
}

func (ec *execContext) printInternal(obj, printCharFn lispObject, escapeFlag bool) error {
	lispType := obj.getType()
	var s string

	switch lispType {
	case lispTypeSymbol:
		s = xSymbolName(obj)
		if s == "" {
			s = "##"
			break
		}

		confusing := false
		if _, err := strconv.ParseFloat(s, 64); err == nil {
			confusing = true
		}
		if strings.HasPrefix(s, ".") || strings.HasPrefix(s, "?") || strings.HasPrefix(s, "\\") {
			confusing = true
		}

		if confusing && escapeFlag {
			s = "\\" + s
		}
	case lispTypeInteger:
		s = fmt.Sprint(xIntegerValue(obj))
	case lispTypeString:
		if escapeFlag {
			s = "\"" + xString(obj).val + "\""
		} else {
			s = xString(obj).val
		}
	case lispTypeCons:
		err := ec.printString("(", printCharFn)
		for ; consp(obj); obj = xCdr(obj) {
			err = ec.printInternalE(xCar(obj), printCharFn, escapeFlag, err)
			if xCdr(obj) != ec.nil_ {
				err = ec.printStringE(" ", printCharFn, err)
			}
		}

		if obj != ec.nil_ {
			err = ec.printStringE(". ", printCharFn, err)
			err = ec.printInternalE(obj, printCharFn, escapeFlag, err)
		}

		return ec.printStringE(")", printCharFn, err)
	case lispTypeFloat:
		s = fmt.Sprint(xFloat(obj).val)
	case lispTypeVector:
		err := ec.printString("[", printCharFn)
		vec := xVector(obj)
		for i, obj := range vec.val {
			err = ec.printInternalE(obj, printCharFn, escapeFlag, err)

			if i < len(vec.val)-1 {
				err = ec.printStringE(" ", printCharFn, err)
			}
		}
		return ec.printStringE("]", printCharFn, err)
	case lispTypeCharTable:
		ct := xCharTable(obj)
		err := ec.printString("#^[", printCharFn)
		err = ec.printInternalE(ct.subtype, printCharFn, escapeFlag, err)
		err = ec.printStringE(" ", printCharFn, err)
		err = ec.printInternalE(ct.defaultVal, printCharFn, escapeFlag, err)
		err = ec.printStringE(" ", printCharFn, err)
		err = ec.printInternalE(newInteger(ct.extraSlots), printCharFn, escapeFlag, err)
		err = ec.printStringE(" ", printCharFn, err)

		parent := ec.nil_
		if ct.parent != nil {
			parent = ct.parent
		}
		err = ec.printInternalE(parent, printCharFn, escapeFlag, err)

		for _, elem := range ct.val {
			err = ec.printStringE(" ", printCharFn, err)
			vec := []lispObject{
				newInteger(elem.start),
				newInteger(elem.end),
				elem.val,
			}
			err = ec.printInternalE(newVector(vec), printCharFn, escapeFlag, err)
		}

		return ec.printStringE("]", printCharFn, err)
	default:
		s = fmt.Sprintf("#<unknown datatype '%+v'>", obj)
	}

	err := ec.printString(s, printCharFn)
	if err != nil {
		return err
	}

	return nil
}

func (ec *execContext) prin1ToString(obj, noEscape lispObject) (lispObject, error) {
	// TODO: Should this buffer be created via get-buffer-create?
	// Needs to be hidden from buffer list though
	bufObj := ec.makeBuffer(" prin1")

	_, err := ec.prin1(obj, bufObj)
	if err != nil {
		return nil, err
	}

	return newString(bufObj.contents), nil
}

func (ec *execContext) printGeneric(obj, printCharFn lispObject, escapeFlag, newlines bool) (lispObject, error) {
	if printCharFn == ec.nil_ {
		printCharFn = ec.v.standardOutput.val
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
	ec.defVarLisp(&ec.v.standardOutput, "standard-output", ec.t)

	ec.defSubr2(&ec.s.prin1, "prin1", ec.prin1, 1)
	ec.defSubr2(nil, "print", ec.print_, 1)
	ec.defSubr2(nil, "princ", ec.princ, 1)
	ec.defSubr2(nil, "prin1-to-string", ec.prin1ToString, 1)
}
