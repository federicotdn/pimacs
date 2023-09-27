package core

import (
	"time"
)

func (parentEc *execContext) makeGoroutine(function, name lispObject) (lispObject, error) {
	newEc := parentEc.copyExecContext()

	go func() {
		defer newEc.unwind()()

		// Lexical binding by default
		if err := newEc.stackPushLet(newEc.gl.lexicalBinding.sym, newEc.t); err != nil {
			newEc.terminate("error in goroutine: '%+v'", err)
		}
		if err := newEc.setupInternalInterpreterEnv(); err != nil {
			newEc.terminate("error in goroutine: '%+v'", err)
		}

		// TODO: Should not actually terminate, but having this for now
		// helps in finding bugs with makeGoroutine.
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

func (ec *execContext) makeChannel(capacity lispObject) (lispObject, error) {
	c := 0
	if capacity != ec.nil_ && naturalp(capacity) {
		c = int(xIntegerValue(capacity))
	} else {
		return ec.wrongTypeArgument(ec.s.integerp, capacity)
	}

	return &lispChannel{
		val: make(chan lispObject, c),
	}, nil
}

func (ec *execContext) channelSend(channel, obj lispObject) (lispObject, error) {
	if !channelp(channel) {
		return ec.wrongTypeArgument(ec.s.channelp, channel)
	}

	xChannel(channel).val <- obj
	return obj, nil
}

func (ec *execContext) channelReceive(channel lispObject) (lispObject, error) {
	if !channelp(channel) {
		return ec.wrongTypeArgument(ec.s.channelp, channel)
	}

	obj, ok := <-xChannel(channel).val
	if !ok {
		obj = ec.nil_
	}
	return newCons(obj, ec.boolVal(ok)), nil
}

func (ec *execContext) channelClose(channel lispObject) (lispObject, error) {
	if !channelp(channel) {
		return ec.wrongTypeArgument(ec.s.channelp, channel)
	}

	close(xChannel(channel).val)
	return ec.nil_, nil
}

func (ec *execContext) symbolsOfGoroutine() {
	ec.defSubr2(nil, "pimacs-go", (*execContext).makeGoroutine, 1)
	ec.defSubr1(nil, "pimacs-sleep", (*execContext).sleep, 1)
	ec.defSubr1(nil, "pimacs-chan", (*execContext).makeChannel, 0)
	ec.defSubr2(nil, "pimacs-send", (*execContext).channelSend, 2)
	ec.defSubr1(nil, "pimacs-receive", (*execContext).channelReceive, 1)
	ec.defSubr1(nil, "pimacs-close", (*execContext).channelClose, 1)
}
