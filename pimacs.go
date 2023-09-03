package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/federicotdn/pimacs/core"
	"github.com/federicotdn/pimacs/ui/tui"
	"os"
	"strings"
)

func repl() {
	interpreter := core.NewInterpreterDefault()

	for {
		reader := bufio.NewReader(os.Stdin)
		eval := true

		fmt.Print("> ")
		source, _ := reader.ReadString('\n')
		source = strings.TrimRight(source, "\r\n")

		if source == "" {
			break
		} else if source[0] == '|' {
			eval = false
			source = source[1:]
			fmt.Println("[input will not be evaluated]")
		}

		var printed string
		var err error

		if eval {
			printed, err = interpreter.ReadEvalPrin1(source)
		} else {
			printed, err = interpreter.ReadPrin1(source)
		}

		if err != nil {
			fmt.Println("repl error:", err)
		} else {
			fmt.Println(printed)
		}
	}
}

func main() {
	var useTui bool
	flag.BoolVar(&useTui, "tui", false, "start Pimacs in TUI mode")
	flag.Parse()

	if useTui {
		tui.Run()
	} else {
		repl()
	}
}
