package core

import (
	"time"
)

func (ec *execContext) makeGoroutine(function, name lispObject) (lispObject, error) {
	newEc := ec.copyExecContext()

	go func() {
		// TODO: Set up internalInterpreterEnv etc.
		_, err := newEc.funcall(function)
		if err != nil {
			newEc.terminate("error in goroutine: '%+v'", err)
		}
	}()

	return ec.nil_, nil
}

func (ec *execContext) sleep(ms lispObject) (lispObject, error) {
	if !naturalp(ms) {
		return ec.wrongTypeArgument(ec.s.integerp, ms)
	}

	time.Sleep(time.Duration(xIntegerValue(ms)) * time.Millisecond)

	return ec.nil_, nil
}

func (ec *execContext) symbolsOfGoroutine() {
	ec.defSubr2(nil, "pimacs-go", (*execContext).makeGoroutine, 1)
	ec.defSubr1(nil, "pimacs-sleep", (*execContext).sleep, 1)
}
