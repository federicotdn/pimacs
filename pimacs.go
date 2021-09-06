package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/federicotdn/pimacs/elisp"
	"os"
	"strings"
)

func main() {
	const usage = "load Emacs Lisp FILE using the load function"
	var loadFile string
	flag.StringVar(&loadFile, "load", "", usage)
	flag.StringVar(&loadFile, "l", "", usage+" (shorthand)")
	flag.Parse()

	interpreter := elisp.NewInterpreter()

	if loadFile != "" {
		data, err := os.ReadFile(loadFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		source := string(data)
		obj, err := interpreter.ReadString(source)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		printed, err := interpreter.Print(obj)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(printed)
	} else {
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
}
