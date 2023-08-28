package proto

type InputEvent interface{}

type InputEventKey struct {
	// TODO: This is too tightly coupled with tcell
	Rune rune
	Key  Key
	Mod  ModMask
}

func (ev *InputEventKey) HasMod(m ModMask) bool {
	return (ev.Mod & m) != 0
}

type DrawOp interface{}

type DrawOpSetText struct {
	Text string
}

type DrawOpSpecial struct {
	Terminate bool
}
