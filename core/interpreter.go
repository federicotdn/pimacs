package core

import (
	"fmt"
	"github.com/federicotdn/pimacs/proto"
)

type Interpreter struct {
	ec *execContext
}

func terminate(format string, v ...interface{}) {
	panic(fmt.Sprintf(format, v...))
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		ec: newExecContext(),
	}
}

func (inp *Interpreter) InputEventsChan() chan<- proto.InputEvent {
	return inp.ec.events
}

func (inp *Interpreter) DrawOpsChan() <-chan proto.DrawOp {
	return inp.ec.ops
}

func (inp *Interpreter) Done() <-chan bool {
	return inp.ec.done
}

func (inp *Interpreter) RecursiveEdit() {
	xEnsure(inp.ec.recursiveEdit())
	inp.ec.done <- true
}

func (inp *Interpreter) ReadPrin1(source string) (string, error) {
	str := inp.ec.makeString(source)
	result, err := inp.ec.readFromString(str, inp.ec.nil_, inp.ec.nil_)
	if err != nil {
		return "", err
	}

	printed, err := inp.ec.prin1ToString(xCar(result), inp.ec.nil_)
	if err != nil {
		return "", err
	}

	return xString(printed).val, nil
}

func (inp *Interpreter) ReadEvalPrin1(source string) (string, error) {
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

	return xString(printed).val, nil
}
