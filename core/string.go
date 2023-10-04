package core

import (
	"unicode/utf8"
)

// |------------+---------------------------------+-------------------------------------|
// |            | immutable (val only)            | mutable (valMut only, != nil)       |
// |------------+---------------------------------+-------------------------------------|
// | size_ < 0  | Unibyte immutable string        | Unibyte mutable string              |
// |            | Ideal for storing ASCII runes   | Good for use as a raw bytes buffer. |
// |            | that will not be modified.      | Will copy bytes when calling str()! |
// |------------+---------------------------------+-------------------------------------|
// | size_ >= 0 | Multibyte immutable string      | Multibyte mutable string            |
// |            | Ideal for storing Unicode runes | Will copy bytes when calling str()! |
// |            | that will not be modified.      |                                     |
// |------------+---------------------------------+-------------------------------------|
type lispString struct {
	valMut []byte
	val    string
	size_  int
}

func newStringInternal(val string, multibyte bool, size_ int) *lispString {
	if multibyte && size_ < 0 {
		size_ = utf8.RuneCountInString(val)
	}
	return &lispString{
		val:   val,
		size_: size_,
	}
}

func newString(val string, multibyte bool) *lispString {
	return newStringInternal(val, multibyte, -1)
}

func newUniOrMultibyteString(val string) *lispString {
	multibyte := false
	size_ := utf8.RuneCountInString(val)
	if size_ < len(val) {
		multibyte = true
	} else {
		size_ = -1
	}
	return newStringInternal(val, multibyte, size_)
}

func (ls *lispString) multibytep() bool {
	return ls.size_ >= 0
}

func (ls *lispString) str() string {
	if ls.valMut == nil {
		return ls.val
	}
	// Do a copy of the mutable value
	return string(ls.valMut)
}

func (ls *lispString) size() int {
	if ls.multibytep() {
		if ls.valMut == nil {
			return ls.size_
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
