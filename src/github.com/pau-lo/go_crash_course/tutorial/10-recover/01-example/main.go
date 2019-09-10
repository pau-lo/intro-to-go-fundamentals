package main

import (
	"fmt"
)

func main() {
	// panic happens after defer statements are executed
	fmt.Println("start")
	defer fmt.Println("this was deferred")
	panic("something bad happened")
	// this line will be unreachable
	fmt.Println("end")
}
