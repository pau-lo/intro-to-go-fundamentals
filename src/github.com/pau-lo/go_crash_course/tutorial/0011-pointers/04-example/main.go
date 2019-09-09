package main

import (
	"fmt"
)

func main() {

	// pointer to myStruct
	var ms *myStruct
	ms = &myStruct{foo: 42}
	// ms will be holding the address of 42
	fmt.Println(ms)
}

type myStruct struct {
	foo int
}
