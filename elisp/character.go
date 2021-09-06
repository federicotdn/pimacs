package elisp

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
