package core

import (
	"os"
)

func (ec *execContext) getenvInternal(variable, env lispObject) (lispObject, error) {
	if !stringp(variable) {
		return ec.wrongTypeArgument(ec.s.stringp, variable)
	}

	val, ok := os.LookupEnv(xStringValue(variable))
	if !ok {
		return ec.nil_, nil
	}

	return ec.makeString(val), nil
}

func (ec *execContext) symbolsOfCallProc() {
	ec.defSubr2(nil, "getenv-internal", ec.getenvInternal, 1)
}
