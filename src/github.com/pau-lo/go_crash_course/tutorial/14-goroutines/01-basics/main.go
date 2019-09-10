// Goroutines concurrent programming
// working on multiple things

package main

import (
	"fmt"
	"time"
)

func main() {
	// by putting go here
	// it will spin off a green thread
	// os threads: abstraction of a thread
	go sayHello()
	time.Sleep(100 * time.Millisecond)
}

func sayHello() {
	fmt.Println("Hello")
}
