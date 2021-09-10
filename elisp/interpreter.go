package elisp

type Interpreter interface {
	ReadPrin1(source string) (string, error)
	ReadEvalPrin1(source string) (string, error)
}

type interpreter struct {
	ec *execContext
}

func NewInterpreter() Interpreter {
	return &interpreter{
		ec: newExecContext(),
	}
}

func (inp *interpreter) ReadPrin1(source string) (string, error) {
	result, err := inp.ec.readFromString(source)
	if err != nil {
		return "", err
	}

	printed, err := inp.ec.prin1(result, inp.ec.nil_)
	if err != nil {
		return "", err
	}

	return xString(printed).value, nil
}

func (inp *interpreter) ReadEvalPrin1(source string) (string, error) {
	obj, err := inp.ec.readFromString(source)
	if err != nil {
		return "", err
	}

	result, err := inp.ec.evalSub(obj)
	if err != nil {
		return "", err
	}

	printed, err := inp.ec.prin1(result, inp.ec.nil_)
	if err != nil {
		return "", err
	}

	return xString(printed).value, nil
}
