package main

import (
	"fmt"
)

func main() {

	// pointer to myStruct
	var ms *myStruct
	// a pointer that is not initialized
	// will be initialzed for you and will hold the value nil
	fmt.Println(ms)
	// using new function
	ms = new(myStruct)
	// ms will be holding the address of 42
	fmt.Println(ms)
}

type myStruct struct {
	foo int
}
