package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	const usage = "load Emacs Lisp FILE using the load function"
	var loadFile string
	flag.StringVar(&loadFile, "load", "", usage)
	flag.StringVar(&loadFile, "l", "", usage+" (shorthand)")
	flag.Parse()

	if loadFile != "" {
		data, err := os.ReadFile(loadFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		source := string(data)
		obj, err := readString(source)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(obj.printObj())
	} else {
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("> ")
			source, _ := reader.ReadString('\n')

			if source == "" {
				break
			}

			obj, err := readString(source)
			if err != nil {
				fmt.Println("error:", err)
			} else {
				fmt.Println(obj.printObj())
			}
		}
	}
}
