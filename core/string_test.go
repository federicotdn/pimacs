package core

import "testing"

func TestNewString(t *testing.T) {
	t.Parallel()

	s := newString("foo", false)
	if s.multibytep() {
		t.Fail()
	}

	s = newString("árbol", true)
	if !s.multibytep() {
		t.Fail()
	}
}

func TestUnibyteString(t *testing.T) {
	t.Parallel()

	s := newUniOrMultibyteString("hello")
	if s.size() != 5 {
		t.Fail()
	}
	if s.sizeBytes() != 5 {
		t.Fail()
	}
	if s.str() != "hello" {
		t.Fail()
	}
	if s.multibytep() {
		t.Fail()
	}
}

func TestMultibyteString(t *testing.T) {
	t.Parallel()

	s := newUniOrMultibyteString("ñandú")
	if s.size() != 5 {
		t.Fail()
	}
	if s.sizeBytes() != 7 {
		t.Fail()
	}
	if s.str() != "ñandú" {
		t.Fail()
	}
	if !s.multibytep() {
		t.Fail()
	}
}
