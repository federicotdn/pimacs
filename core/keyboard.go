package core

import (
	"fmt"
	"github.com/federicotdn/pimacs/proto"
)

type commandLoopState struct {
	inputHist []*proto.InputEventKey
	debug     bool
}

func (ec *execContext) handleInputEventKey(state *commandLoopState, ev *proto.InputEventKey) bool {
	state.inputHist = append(state.inputHist, ev)

	if ev.Rune == '`' {
		state.debug = !state.debug
		ec.gl.currentBuf.contents = ""
	}

	if state.debug {
		ec.gl.currentBuf.contents = "KEY: "

		if ev.HasMod(proto.ModAlt) {
			ec.gl.currentBuf.contents += "Alt "
		}
		if ev.HasMod(proto.ModCtrl) {
			ec.gl.currentBuf.contents += "Ctrl "
		}
		if ev.HasMod(proto.ModShift) {
			ec.gl.currentBuf.contents += "Shift "
		}
		if ev.HasMod(proto.ModMeta) {
			ec.gl.currentBuf.contents += "Meta "
		}

		ec.gl.currentBuf.contents += fmt.Sprintf("key: <%v> | rune: <%v> <%c>", ev.Key, ev.Rune, ev.Rune)
		ec.ops <- &proto.DrawOpSetText{Text: ec.gl.currentBuf.contents}

		return false
	}

	if ev.Key == proto.KeyCtrlC || ev.Key == proto.KeyEscape {
		ec.ops <- &proto.DrawOpSpecial{Terminate: true}
		return true
	}

	if ev.Key == proto.KeyEnter {
		// TODO: Don't manipulate buffer contents like this
		str := newString(ec.gl.currentBuf.contents)
		ec.gl.currentBuf.contents = ""

		obj, err := ec.readFromString(str, ec.nil_, ec.nil_)
		if err != nil {
			ec.ops <- &proto.DrawOpSetText{Text: err.Error()}
			return false
		}

		result, err := ec.evalSub(xCar(obj))
		if err != nil {
			ec.ops <- &proto.DrawOpSetText{Text: err.Error()}
			return false
		}

		printed, err := ec.prin1ToString(result, ec.nil_)
		if err != nil {
			ec.ops <- &proto.DrawOpSetText{Text: err.Error()}
			return false
		}

		ec.ops <- &proto.DrawOpSetText{Text: xStringValue(printed)}
	} else {
		ec.gl.currentBuf.contents += string(ev.Rune)
		ec.ops <- &proto.DrawOpSetText{Text: ec.gl.currentBuf.contents}
	}
	return false
}

func (ec *execContext) recursiveEdit() (lispObject, error) {
	state := commandLoopState{}
loop:
	for {
		rawEv := <-ec.events
		switch ev := rawEv.(type) {
		case *proto.InputEventKey:
			terminate := ec.handleInputEventKey(&state, ev)
			if terminate {
				break loop
			}
		}

	}
	return ec.nil_, nil
}

func (ec *execContext) symbolsOfKeyboard() {
	ec.defSubr0(&ec.s.recursiveEdit, "recursive-edit", ec.recursiveEdit)
}
