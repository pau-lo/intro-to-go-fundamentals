package main

import (
	"fmt"
)

// the main function will be calling another function sayMessage()
func main() { // body functions
	sayMessage("Hello, you") // "Hello, you" is the argument that will
	// pass on to sayMessage
}

// takes in msg == name of the parameter of type string
// msg ==
func sayMessage(msg string) {
	fmt.Println(msg) // the msg is the message acquired from sayMessage
}
