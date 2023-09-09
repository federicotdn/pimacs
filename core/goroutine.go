package core

import (
	"time"
)

func (parentEc *execContext) makeGoroutine(function, name lispObject) (lispObject, error) {
	newEc := parentEc.copyExecContext()

	go func() {
		if err := newEc.stackPushLet(newEc.gl.lexicalBinding.sym, newEc.t); err != nil {
			newEc.terminate("error in goroutine: '%+v'", err)
		}
		if err := newEc.setupInternalInterpreterEnv(); err != nil {
			newEc.terminate("error in goroutine: '%+v'", err)
		}

		_, err := newEc.funcall(function)
		if err != nil {
			newEc.terminate("error in goroutine: '%+v'", err)
		}
	}()

	return parentEc.nil_, nil
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
