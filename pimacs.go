package main

import (
	"log"
)

func main() {
	obj, err := readFile("test.el")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Value:")
	log.Println(obj.printObj())
}
