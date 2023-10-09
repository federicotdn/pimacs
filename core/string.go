package core

import (
	"errors"
	"unicode/utf8"
)

var indexErr = errors.New("index out of range")

// |------------+------------------------------------+-------------------------------------|
// |            | string-backed (val)                | []byte-backed (valMut != nil)       |
// |------------+------------------------------------+-------------------------------------|
// | size_ < 0  | Unibyte immutable string           | Unibyte mutable string              |
// |            | Ideal for storing ASCII runes      | Good for use as a raw bytes buffer. |
// |            | that will not be modified.         | size_ is -1.                        |
// |            | size_ is -1.                       |                                     |
// |------------+------------------------------------+-------------------------------------|
// | size_ >= 0 | Multibyte immutable string         | Multibyte mutable string            |
// |            | Ideal for storing Unicode runes    | Will copy bytes when calling str()! |
// |            | that will not be modified.         |                                     |
// |            | size_ contains the number of runes | size_ is 0.                         |
// |            | the string has.                    |                                     |
// |------------+------------------------------------+-------------------------------------|
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
	size_ := 0

	tmp := val
	for len(tmp) > 0 {
		r, width := utf8.DecodeRuneInString(tmp)
		if r == utf8.RuneError {
			size_ = len(val)
			break
		}

		size_++
		tmp = tmp[width:]
	}

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

func (ls *lispString) aref(i int) (lispInt, error) {
	if i < 0 {
		return 0, indexErr
	}
	if ls.multibytep() {
		if ls.valMut == nil {
			// We are multibyte - size check is only cheap
			// if we are immutable
			if i >= ls.size() {
				return 0, indexErr
			}

			for j, c := range ls.val {
				if j == i {
					return runeToLispInt(c), nil
				}
			}

			// We should never get here
			return 0, indexErr
		}

		return 0, errors.New("string aref unimplemented")
	}

	// We are unibyte, so size check is cheap
	if i >= ls.size() {
		return 0, indexErr
	}

	if ls.valMut == nil {
		return lispInt(ls.val[i]), nil
	}
	return lispInt(ls.valMut[i]), nil
}

// str returns the lispString as a Go string.  if size_ is negative
// (unibyte string), the string will simply contain the byte contents
// of the lispString in no specific encoding. If size_ is
// non-negative, str will return a UTF-8 encoded string.
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

func (ls *lispString) emptyp() bool {
	return ls.sizeBytes() == 0
}

func (ls *lispString) getType() lispType {
	return lispTypeString
}
