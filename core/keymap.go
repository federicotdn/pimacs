package core

func (ec *execContext) makeKeymap(str lispObject) (lispObject, error) {
	tail := ec.nil_
	if str != ec.nil_ {
		tail = ec.makeList(str)
	}

	table, err := ec.makeCharTable(ec.s.keymap, ec.s.nil_)
	if err != nil {
		return nil, err
	}
	return newCons(ec.s.keymap, newCons(table, tail)), nil
}

func (ec *execContext) makeSparseKeymap(str lispObject) (lispObject, error) {
	if str != ec.nil_ {
		return ec.makeList(ec.s.keymap, str), nil
	}
	return ec.makeList(ec.s.keymap), nil
}

func (ec *execContext) keymapp(object lispObject) (lispObject, error) {
	keymap, err := ec.getKeymap(object, false, false)
	if err != nil {
		return nil, err
	}
	return ec.bool(keymap != ec.nil_)
}

func (ec *execContext) getKeymap(object lispObject, errIfNotKeymap, autoload bool) (lispObject, error) {
	// TODO: Autoload is currently ignored
	if object == ec.nil_ {
		if errIfNotKeymap {
			return ec.wrongTypeArgument(ec.s.keymapp, object)
		}
		return ec.nil_, nil
	} else if consp(object) && xCar(object) == ec.s.keymap {
		return object, nil
	}

	tem := ec.indirectFunctionInternal(object)
	if consp(tem) {
		if xCar(tem) == ec.s.keymap {
			return tem, nil
		}

		if (autoload || !errIfNotKeymap) && xCar(tem) == ec.s.autoload {
			return ec.pimacsUnimplemented(ec.s.nil_, "no autoload for keymaps")
		}
	}

	if errIfNotKeymap {
		return ec.wrongTypeArgument(ec.s.keymapp, object)
	}
	return ec.nil_, nil
}

func (ec *execContext) defineKey(keymap, key, def, remove lispObject) (lispObject, error) {
	ec.warning("stub invoked: define-key")
	return ec.nil_, nil
}

func (ec *execContext) useGlobalMap(keymap lispObject) (lispObject, error) {
	ec.warning("stub invoked: use-global-map")
	return ec.nil_, nil
}

func (ec *execContext) symbolsOfKeymap() {
	ec.defSym(&ec.s.keymap, "keymap")

	ec.defSubr1(nil, "make-keymap", (*execContext).makeKeymap, 0)
	ec.defSubr1(nil, "make-sparse-keymap", (*execContext).makeSparseKeymap, 0)
	ec.defSubr4(nil, "define-key", (*execContext).defineKey, 3)
	ec.defSubr1(&ec.s.keymapp, "keymapp", (*execContext).keymapp, 1)
	ec.defSubr1(nil, "use-global-map", (*execContext).useGlobalMap, 1)
}
