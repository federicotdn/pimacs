package core

const (
	eightBitCodeOffset rune = 0x3fff00
	max5ByteChar            = 0x3fff7f
	maxChar                 = 0x3fffff
)

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
