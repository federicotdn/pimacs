package core

import (
	"unicode/utf8"
)

type lispString struct {
	valMut    []byte
	val       string
	multibyte bool
}

func newString(val string, multibyte bool) *lispString {
	return &lispString{val: val, multibyte: multibyte}
}

func newUniOrMultibyteString(val string) *lispString {
	for _, r := range val {
		if !asciiCharp(r) {
			return newString(val, true)
		}
	}
	return newString(val, false)
}

func (ls *lispString) str() string {
	if ls.valMut == nil {
		return ls.val
	}
	// Do a copy of the mutable value
	return string(ls.valMut)
}

func (ls *lispString) size() int {
	if ls.multibyte {
		if ls.valMut == nil {
			return utf8.RuneCountInString(ls.val)
		}
		return utf8.RuneCount(ls.valMut)
	}
	return ls.sizeBytes()
}

func (ls *lispString) sizeBytes() int {
	if ls.valMut == nil {
		return len(ls.val)
	}
	return len(ls.valMut)
}

func (ls *lispString) getType() lispType {
	return lispTypeString
}
