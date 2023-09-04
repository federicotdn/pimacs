package tui

import (
	"github.com/federicotdn/pimacs/core"
	"github.com/federicotdn/pimacs/proto"
	"github.com/gdamore/tcell/v2"
)

func Run() {
	s, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}

	inp, err := core.NewInterpreterDefault()
	if err != nil {
		panic(err)
	}

	tcellEvents := make(chan tcell.Event)
	quit := make(chan struct{})

	go s.ChannelEvents(tcellEvents, quit)
	go inp.RecursiveEdit()

	drawOpsChan := inp.DrawOpsChan()
	inputEventsChan := inp.InputEventsChan()

loop:
	for {
		select {
		case event := <-tcellEvents:
			switch ev := event.(type) {
			case *tcell.EventKey:
				mod, key, ch := ev.Modifiers(), ev.Key(), ev.Rune()
				inputEventsChan <- &proto.InputEventKey{
					Rune: ch,
					Key:  proto.Key(key),
					Mod:  proto.ModMask(mod),
				}
			}
		case rawOp := <-drawOpsChan:
			switch op := rawOp.(type) {
			case *proto.DrawOpSetText:
				s.Clear()

				col := 0
				row := 0
				width, height := s.Size()
				for _, r := range []rune(op.Text) { //nolint: staticcheck,gosimple
					if r == '\n' {
						col = 0
						row++
					} else {
						s.SetContent(col, row, r, nil, tcell.StyleDefault)
						col++
					}

					if col >= width {
						col = 0
						row++
					}

					if row >= height {
						col = 0
						row = 0
					}
				}

				s.Show()
			case *proto.DrawOpSpecial:
				break loop
			}
		}
	}

	<-inp.Done()
	close(quit)
	s.Fini()
}
