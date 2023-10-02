package core

import (
	"fmt"
	"github.com/federicotdn/pimacs/proto"
)

type Interpreter struct {
	ec *execContext
}

type InterpreterConfig struct {
	loadPathPrepend []string
}

func terminate(format string, v ...interface{}) {
	panic(fmt.Sprintf(format, v...))
}

func NewInterpreterDefault() (*Interpreter, error) {
	return NewInterpreter(InterpreterConfig{})
}

func NewInterpreter(config InterpreterConfig) (*Interpreter, error) {
	ec, err := newExecContext(config.loadPathPrepend)
	if err != nil {
		return nil, err
	}

	return &Interpreter{ec: ec}, nil
}

func newTestingInterpreter() *Interpreter {
	// When running tests, Go sets the CWD to the package's
	// directory
	ec, err := newExecContext([]string{"../lisp", "../lisp/emacs", "../test/lisp"})
	if err != nil {
		panic(err)
	}
	ec.testing = true
	return &Interpreter{ec: ec}
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

func (inp *Interpreter) LoadFile(filename string) error {
	_, err := inp.ec.load(newString(filename), inp.ec.nil_, inp.ec.nil_, inp.ec.nil_, inp.ec.nil_)
	return err
}

func (inp *Interpreter) ReadPrin1(source string) (string, error) {
	str := newString(source)
	result, err := inp.ec.readFromString(str, inp.ec.nil_, inp.ec.nil_)
	if err != nil {
		return "", err
	}

	printed, err := inp.ec.prin1ToString(xCar(result), inp.ec.nil_, inp.ec.nil_)
	if err != nil {
		return "", err
	}

	return xString(printed).val, nil
}

func (inp *Interpreter) ReadEvalPrin1(source string) (string, error) {
	str := newString(source)
	obj, err := inp.ec.readFromString(str, inp.ec.nil_, inp.ec.nil_)
	if err != nil {
		return "", err
	}

	result, err := inp.ec.evalSub(xCar(obj))
	if err != nil {
		return "", err
	}

	printed, err := inp.ec.prin1ToString(result, inp.ec.nil_, inp.ec.nil_)
	if err != nil {
		return "", err
	}

	return xString(printed).val, nil
}
