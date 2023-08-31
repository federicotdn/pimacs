package core

func (ec *execContext) makeCharTable(purpose, init lispObject) (lispObject, error) {
	if !symbolp(purpose) {
		return ec.wrongTypeArgument(ec.s.symbolp, purpose)
	}

	n, err := ec.get(purpose, ec.s.charTableExtraSlots)
	if err != nil {
		return nil, err
	}

	extra := lispInt(0)
	if n != ec.nil_ {
		if naturalp(n) && xIntegerValue(n) <= 10 {
			extra = xIntegerValue(n)
		} else {
			return ec.wrongTypeArgument(ec.s.wholeNump, n)
		}
	}

	return &lispCharTable{
		subtype:    xSymbol(purpose),
		defaultVal: init,
		extraSlots: extra,
	}, nil
}

func (ec *execContext) charTableRange(char_table, range_ lispObject) (lispObject, error) {
	if !chartablep(char_table) {
		return ec.wrongTypeArgument(ec.s.charTablep, char_table)
	}
	return nil, nil
}

func (ec *execContext) setCharTableRange(char_table, range_, value lispObject) (lispObject, error) {
	if !chartablep(char_table) {
		return ec.wrongTypeArgument(ec.s.charTablep, char_table)
	}
	return value, nil
}

func (ec *execContext) symbolsOfCharacterTable() {
	ec.defSym(&ec.s.charTableExtraSlots, "char-table-extra-slots")

	ec.defSubr2(nil, "make-char-table", ec.makeCharTable, 1)
	ec.defSubr2(nil, "char-table-range", ec.charTableRange, 2)
	ec.defSubr3(nil, "set-char-table-range", ec.setCharTableRange, 3)
}
