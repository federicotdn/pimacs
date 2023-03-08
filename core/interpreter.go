package core

import (
	"fmt"
)

type Interpreter interface {
	ReadPrin1(source string) (string, error)
	ReadEvalPrin1(source string) (string, error)
}

type interpreter struct {
	ec *execContext
}

func terminate(format string, v ...interface{}) {
	panic(fmt.Sprintf(format, v...))
}

func NewInterpreter() Interpreter {
	return &interpreter{
		ec: newExecContext(),
	}
}

func (inp *interpreter) ReadPrin1(source string) (string, error) {
	str := inp.ec.makeString(source)
	result, err := inp.ec.readFromString(str, inp.ec.nil_, inp.ec.nil_)
	if err != nil {
		return "", err
	}

	printed, err := inp.ec.prin1ToString(xCar(result), inp.ec.nil_)
	if err != nil {
		return "", err
	}

	return xString(printed).value, nil
}

func (inp *interpreter) ReadEvalPrin1(source string) (string, error) {
	str := inp.ec.makeString(source)
	obj, err := inp.ec.readFromString(str, inp.ec.nil_, inp.ec.nil_)
	if err != nil {
		return "", err
	}

	result, err := inp.ec.evalSub(xCar(obj))
	if err != nil {
		return "", err
	}

	printed, err := inp.ec.prin1ToString(result, inp.ec.nil_)
	if err != nil {
		return "", err
	}

	return xString(printed).value, nil
}
