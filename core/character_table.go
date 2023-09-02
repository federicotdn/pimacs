package core

import (
	"slices"
)

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

func (ec *execContext) charTableLookup(table *lispCharTable, c lispInt) (int, bool) {
	cmp := func(e lispCharTableEntry, v lispInt) int {
		return int(e.start - v)
	}
	index, found := slices.BinarySearchFunc(table.val, c, cmp)
	if found {
		// Found an entry with .start == c
		return index, true
	}

	if index > 0 && table.val[index-1].contains(c) {
		// Went one element back and it contains c
		return index - 1, true
	}

	return index, false
}

func (ec *execContext) charTableSet(table *lispCharTable, from, to lispInt, value lispObject) {
	if to < from {
		to = from
	}
	rep := []lispCharTableEntry{{start: from, end: to, val: value}}
	index, found := ec.charTableLookup(table, from)
	elem := lispCharTableEntry{start: -1, end: -1}

	if !found && index > 0 {
		elem = table.val[index-1]
	} else if len(table.val) > 0 {
		elem = table.val[index]
	}

	if elem.contains(from) && elem.start < from {
		rep = append([]lispCharTableEntry{{
			start: elem.start,
			end:   from - 1,
			val:   elem.val,
		}}, rep...)
	}

	index2, found := ec.charTableLookup(table, to)

	if !found && index2 > 0 {
		elem = table.val[index2-1]
	} else if len(table.val) > 0 {
		elem = table.val[index2]
	}

	if elem.contains(to) && to < elem.end {
		rep = append(rep, lispCharTableEntry{
			start: to + 1,
			end:   elem.end,
			val:   elem.val,
		})
	}

	if index2 < len(table.val) {
		index2++
	}
	table.val = slices.Replace(table.val, index, index2, rep...)
}

func (ec *execContext) charTableRange(table, range_ lispObject) (lispObject, error) {
	if !chartablep(table) {
		return ec.wrongTypeArgument(ec.s.charTablep, table)
	}

	ct := xCharTable(table)

	if range_ == ec.nil_ {
		return ct.defaultVal, nil
	} else if characterp(range_) {
		c := xIntegerValue(range_)
		index, found := ec.charTableLookup(ct, c)
		if found {
			return ct.val[index].val, nil
		}
		return ct.defaultVal, nil
	} else if consp(range_) {
		if !characterp(xCar(range_)) {
			return ec.wrongTypeArgument(ec.s.characterp, xCar(range_))
		} else if !characterp(xCdr(range_)) {
			return ec.wrongTypeArgument(ec.s.characterp, xCdr(range_))
		}

		// Even though we have received a cons of (FROM . TO), we only use
		// the FROM part - Emacs does the same-ish.
		// See: https://lists.gnu.org/archive/html/emacs-devel/2023-09/msg00014.html
		from := xIntegerValue(xCar(range_))
		index, found := ec.charTableLookup(ct, from)
		if found {
			return ct.val[index].val, nil
		}
		return ct.defaultVal, nil
	} else {
		return ec.signalError("Invalid RANGE argument to `char-table-range'")
	}
}

func (ec *execContext) setCharTableRange(table, range_, value lispObject) (lispObject, error) {
	if !chartablep(table) {
		return ec.wrongTypeArgument(ec.s.charTablep, table)
	}

	ct := xCharTable(table)

	if range_ == ec.t {
		ec.charTableSet(ct, 0, maxChar, value)
	} else if range_ == ec.nil_ {
		ct.defaultVal = value
	} else if characterp(range_) {
		c := xInteger(range_)
		ec.charTableSet(ct, c.val, c.val, value)
	} else if consp(range_) {
		if !characterp(xCar(range_)) {
			return ec.wrongTypeArgument(ec.s.characterp, xCar(range_))
		} else if !characterp(xCdr(range_)) {
			return ec.wrongTypeArgument(ec.s.characterp, xCdr(range_))
		}

		ec.charTableSet(ct, xIntegerValue(xCar(range_)), xIntegerValue(xCdr(range_)), value)
	} else {
		return ec.signalError("Invalid RANGE argument to `set-char-table-range'")
	}

	return value, nil
}

func (ec *execContext) symbolsOfCharacterTable() {
	ec.defSym(&ec.s.charTableExtraSlots, "char-table-extra-slots")

	ec.defSubr2(nil, "make-char-table", ec.makeCharTable, 1)
	ec.defSubr2(nil, "char-table-range", ec.charTableRange, 2)
	ec.defSubr3(nil, "set-char-table-range", ec.setCharTableRange, 3)
}
