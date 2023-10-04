package core

func (ec *execContext) insert(args ...lispObject) (lispObject, error) {
	for _, arg := range args {
		if characterp(arg) {
			ec.gl.currentBuf.contents += string(xIntegerRune(arg))
		} else if stringp(arg) {
			ec.gl.currentBuf.contents += xStringValue(arg)
		} else {
			return ec.wrongTypeArgument(ec.s.charOrStringp, arg)
		}
	}

	return ec.nil_, nil
}

func (ec *execContext) bufferString() (lispObject, error) {
	return newString(ec.gl.currentBuf.contents, true), nil
}

func (ec *execContext) currentBuffer() (lispObject, error) {
	return ec.gl.currentBuf, nil
}

func (ec *execContext) currentBufferInternal() lispObject {
	return xEnsure(ec.currentBuffer())
}

func (ec *execContext) setBufferIfLive(obj lispObject) {
	buf := xBuffer(obj)
	if buf.live {
		ec.setBufferInternal(buf)
	}
}

func (ec *execContext) setBufferInternal(buf *lispBuffer) {
	ec.gl.currentBuf = buf
}

func (ec *execContext) setBuffer(bufferOrName lispObject) (lispObject, error) {
	obj, err := ec.getBuffer(bufferOrName)
	if err != nil {
		return nil, err
	}

	if obj == ec.nil_ {
		return ec.signalError("No buffer named %v", bufferOrName)
	}

	buf := xBuffer(obj)
	if !buf.live {
		return ec.signalError("Selecting deleted buffer")
	}

	ec.setBufferInternal(buf)
	return obj, nil
}

func (ec *execContext) loadOrStoreBuffer(name string, buf *lispBuffer) (*lispBuffer, bool) {
	ec.buffersLock.Lock()
	defer ec.buffersLock.Unlock()

	prev, existed := ec.buffers[name]
	if existed {
		return prev, true
	}

	if buf != nil {
		ec.buffers[name] = buf
	}
	return buf, false
}

func (ec *execContext) getBuffer(bufferOrName lispObject) (lispObject, error) {
	if bufferp(bufferOrName) {
		return bufferOrName, nil
	} else if !stringp(bufferOrName) {
		return ec.wrongTypeArgument(ec.s.stringp, bufferOrName)
	}

	buf, existed := ec.loadOrStoreBuffer(xStringValue(bufferOrName), nil)
	if !existed {
		return ec.nil_, nil
	}
	return buf, nil
}

func (ec *execContext) getBufferCreate(bufferOrName, inhibitBufferHooks lispObject) (lispObject, error) {
	if bufferp(bufferOrName) {
		return bufferOrName, nil
	} else if !stringp(bufferOrName) {
		return ec.wrongTypeArgument(ec.s.stringp, bufferOrName)
	}

	if xStringEmpty(bufferOrName) {
		return ec.signalError("Empty string for buffer name is not allowed")
	}

	buf := newBuffer(bufferOrName)
	buf, _ = ec.loadOrStoreBuffer(xStringValue(bufferOrName), buf)

	return buf, nil
}

func (ec *execContext) bufferName(obj lispObject) (lispObject, error) {
	if obj == ec.nil_ {
		obj = ec.currentBufferInternal()
	} else if !bufferp(obj) {
		return ec.wrongTypeArgument(ec.s.bufferp, obj)
	}

	buf := xBuffer(obj)
	if !buf.live {
		return ec.nil_, nil
	}

	return buf.name, nil
}

func (ec *execContext) bufferList(frame lispObject) (lispObject, error) {
	ec.buffersLock.RLock()
	defer ec.buffersLock.RUnlock()

	buffers := make([]lispObject, 0, len(ec.buffers))
	for _, buf := range ec.buffers {
		buffers = append(buffers, buf)
	}

	return ec.makeList(buffers...), nil
}

func (ec *execContext) symbolsOfBuffer() {
	ec.defSubr0(nil, "buffer-string", (*execContext).bufferString)
	ec.defSubrM(nil, "insert", (*execContext).insert, 0)
	ec.defSubr0(nil, "current-buffer", (*execContext).currentBuffer)
	ec.defSubr1(nil, "set-buffer", (*execContext).setBuffer, 1)
	ec.defSubr1(nil, "get-buffer", (*execContext).getBuffer, 1)
	ec.defSubr1(nil, "buffer-name", (*execContext).bufferName, 0)
	ec.defSubr1(nil, "buffer-list", (*execContext).bufferList, 0)
	ec.defSubr2(nil, "get-buffer-create", (*execContext).getBufferCreate, 1)
}

func (ec *execContext) initBufferGoroutineLocals() {
	// TODO: Use another buffer in new goroutines?
	// TODO: What to do on error?
	buf := xEnsure(ec.getBufferCreate(newString("*scratch*", false), ec.nil_))
	ec.gl.currentBuf = xBuffer(buf)
}
