package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"github.com/federicotdn/pimacs/elisp"
)

func main() {
	const usage = "load Emacs Lisp FILE using the load function"
	var loadFile string
	flag.StringVar(&loadFile, "load", "", usage)
	flag.StringVar(&loadFile, "l", "", usage+" (shorthand)")
	flag.Parse()

	parser := elisp.Parser{}

	if loadFile != "" {
		data, err := os.ReadFile(loadFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		source := string(data)
		obj, err := parser.ReadString(source)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(obj.PrintObj())
	} else {
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("> ")
			source, _ := reader.ReadString('\n')
			source = strings.TrimRight(source, "\r\n")

			if source == "" {
				break
			}

			obj, err := parser.ReadString(source)
			if err != nil {
				fmt.Println("error:", err)
			} else {
				fmt.Println(obj.PrintObj())
			}
		}
	}
}
