package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	// using the in build panic function
	panic("Something bad happened")
	fmt.Println("end")
}
