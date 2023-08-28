package core

import (
	"github.com/federicotdn/pimacs/proto"
)

func (ec *execContext) recursiveEdit() (lispObject, error) {
	source := ""
loop:
	for {
		rawEv := <-ec.events
		switch ev := rawEv.(type) {
		case *proto.InputEventKey:
			if ev.Key == proto.KeyCtrlC || ev.Key == proto.KeyEscape {
				ec.ops <- &proto.DrawOpSpecial{Terminate: true}
				break loop
			}

			if ev.Key == proto.KeyEnter {
				str := ec.makeString(source)
				source = ""
				obj, err := ec.readFromString(str, ec.nil_, ec.nil_)
				if err != nil {
					ec.ops <- &proto.DrawOpSetText{Text: err.Error()}
					break
				}

				result, err := ec.evalSub(xCar(obj))
				if err != nil {
					ec.ops <- &proto.DrawOpSetText{Text: err.Error()}
					break
				}

				printed, err := ec.prin1ToString(result, ec.nil_)
				if err != nil {
					ec.ops <- &proto.DrawOpSetText{Text: err.Error()}
					break
				}

				ec.ops <- &proto.DrawOpSetText{Text: xStringValue(printed)}
			} else {
				source += string(ev.Rune)
				ec.ops <- &proto.DrawOpSetText{Text: source}
			}
		}

	}
	return ec.nil_, nil
}

func (ec *execContext) symbolsOfKeyboard() {
	ec.defSubr0(&ec.s.recursiveEdit, "recursive-edit", ec.recursiveEdit)
}
