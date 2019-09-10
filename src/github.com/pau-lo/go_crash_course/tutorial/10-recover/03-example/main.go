package main

import (
	"fmt"
	"log"
)

func main() {
	// panic happens after defer statements are executed
	fmt.Println("start")
	panicker()

	fmt.Println("end")

}

func panicker() {
	fmt.Println("about to panic")
	defer func() { // defer takes a function call
		if err := recover(); err != nil {
			log.Println("Error: ", err)
			panic(err)
		}
	}()

	panic("Something bad happend")
	// this will not be print out
	fmt.Println("end")

}
