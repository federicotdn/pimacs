package core

const (
	eightBitCodeOffset rune = 0x3fff00
	max5ByteChar       rune = 0x3fff7f
	maxChar            rune = 0x3fffff

	charAlt rune = 1 << (22 + iota)
	charSuper
	charHyper
	charShift
	charCtrl
	charMeta
	charModMask rune = charAlt | charSuper | charHyper | charShift | charCtrl | charMeta
)

func runeToLispInt(c rune) lispInt {
	return lispInt(c)
}

func lispIntToRune(i lispInt) rune {
	return rune(i)
}

func charValidp(c rune) bool {
	return 0 <= c && c <= maxChar
}

func byte8toChar(c rune) rune {
	return c + eightBitCodeOffset
}

func charByte8(c rune) bool {
	return c > max5ByteChar
}

func charToByte8(c rune) rune {
	if charByte8(c) {
		return c - eightBitCodeOffset
	}
	return c & 0xff
}
