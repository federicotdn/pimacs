package main

import (
	"bufio"
	"fmt"
	"github.com/federicotdn/pimacs/elisp"
	"os"
	"strings"
)

func main() {
	interpreter := elisp.NewInterpreter()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		source, _ := reader.ReadString('\n')
		source = strings.TrimRight(source, "\r\n")

		if source == "" {
			break
		}

		printed, err := interpreter.ReadEvalPrint(source)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(printed)
		}
	}
}
