package main

import (
	"fmt"
	"time"
)

func main() {
	var msg = "Hello"
	go func(msg string) {
		fmt.Println(msg)

	}(msg)
	msg = "Goodbye"
	// using the sleep function is not best practice though
	time.Sleep(100 * time.Millisecond)
}
