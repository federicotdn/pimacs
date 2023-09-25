package core

import (
	"errors"
	"testing"
)

func TestHelpersBasic(t *testing.T) {
	t.Parallel()

	sym := &lispSymbol{name: "foo"}

	if !symbolp(sym) {
		t.Fail()
	}

	var obj lispObject = sym
	sym2 := xSymbol(obj)
	if sym.name != sym2.name {
		t.Fail()
	}
}

func TestCastFailure(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	sym := &lispSymbol{name: "foo"}
	var obj lispObject = sym
	xInteger(obj)
}

func TestEnsure(t *testing.T) {
	t.Parallel()
	var obj lispObject = &lispInteger{val: 42}
	xEnsure(obj, nil)
}

func TestEnsureFailure(t *testing.T) {
	t.Parallel()
	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	xEnsure(nil, errors.New("fail"))
}

func TestObjAddr(t *testing.T) {
	t.Parallel()
	s1 := newString("hello")
	s2 := newString("hello")
	s3 := s1

	if objAddr(s1) != objAddr(s3) {
		t.Fail()
	}

	if objAddr(s1) == objAddr(s2) {
		t.Fail()
	}
}
