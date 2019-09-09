package main

import (
	"fmt"
	"log"
)

func main() {
	// panic happens after defer statements are executed
	fmt.Println("start")

	// custom func anonymous function
	//
	defer func() { // defer takes a function call
		if err := recover(); err != nil {
			log.Println("Error: ", err)
		}
	}() // making the function execute invoking defer

	panic("Something bad happend")
	// this will not be print out
	fmt.Println("end")

}
