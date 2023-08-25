package proto

type inputEventType int
type drawOpType int

type InputEvent interface{}

type InputEventKey struct {
	// TODO: This is too tightly coupled with tcell
	Rune rune
	Key  int16
	Mod  int16
}

type DrawOp interface{}

type DrawOpSetText struct {
	Text string
}

type DrawOpSpecial struct {
	Terminate bool
}
