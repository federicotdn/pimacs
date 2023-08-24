package core

func (ec *execContext) insert(args ...lispObject) (lispObject, error) {
	for _, arg := range args {
		if characterp(arg) {
			ec.currentBuf.contents += string(xIntegerRune(arg))
		} else if stringp(arg) {
			ec.currentBuf.contents += xStringValue(arg)
		} else {
			return ec.wrongTypeArgument(ec.s.charOrStringp, arg)
		}
	}

	return ec.nil_, nil
}

func (ec *execContext) bufferString() (lispObject, error) {
	return ec.makeString(ec.currentBuf.contents), nil
}

func (ec *execContext) currentBuffer() (lispObject, error) {
	return ec.makeBuffer(ec.currentBuf), nil
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

func (ec *execContext) setBufferInternal(buf *buffer) {
	ec.currentBuf = buf
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

func (ec *execContext) getBuffer(bufferOrName lispObject) (lispObject, error) {
	if bufferp(bufferOrName) {
		return bufferOrName, nil
	}

	if !stringp(bufferOrName) {
		return ec.wrongTypeArgument(ec.s.stringp, bufferOrName)
	}

	elem, err := ec.assoc(bufferOrName, ec.buffers, ec.nil_)
	if err != nil {
		return nil, err
	}

	if elem == ec.nil_ {
		return ec.nil_, nil
	}

	return xCdr(elem), nil
}

func (ec *execContext) createBuffer(name string) *buffer {
	return &buffer{
		name: name,
		live: true,
	}
}

func (ec *execContext) getBufferCreate(bufferOrName, inhibitBufferHooks lispObject) (lispObject, error) {
	obj, err := ec.getBuffer(bufferOrName)
	if err != nil {
		return nil, err
	}

	if obj != ec.nil_ {
		return obj, nil
	}

	if xStringChars(bufferOrName) == 0 {
		return ec.signalError("Empty string for buffer name is not allowed")
	}

	buf := ec.makeBuffer(ec.createBuffer(xStringValue(bufferOrName)))
	newList, err := ec.nconc(ec.buffers, ec.makeList(ec.makeCons(bufferOrName, buf)))
	ec.buffers = newList

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

	return ec.makeString(buf.name), nil
}

func (ec *execContext) bufferList(frame lispObject) (lispObject, error) {
	buffers := []lispObject{}

	for tail := ec.buffers; consp(tail); tail = xCdr(tail) {
		elem := xCar(tail)
		buffers = append(buffers, xCdr(elem))
	}

	return ec.makeList(buffers...), nil
}

func (ec *execContext) symbolsOfBuffer() {
	ec.defSubr0(nil, "buffer-string", ec.bufferString)
	ec.defSubrM(nil, "insert", ec.insert, 0)
	ec.defSubr0(nil, "current-buffer", ec.currentBuffer)
	ec.defSubr1(nil, "set-buffer", ec.setBuffer, 1)
	ec.defSubr1(nil, "get-buffer", ec.getBuffer, 1)
	ec.defSubr1(nil, "buffer-name", ec.bufferName, 0)
	ec.defSubr1(nil, "buffer-list", ec.bufferList, 0)
	ec.defSubr2(nil, "get-buffer-create", ec.getBufferCreate, 1)
}

func (ec *execContext) initBuffer() {
	ec.buffers = ec.makeList()
	buf := xEnsure(ec.getBufferCreate(ec.makeString("*scratch*"), ec.nil_))
	ec.currentBuf = xBuffer(buf)
}
