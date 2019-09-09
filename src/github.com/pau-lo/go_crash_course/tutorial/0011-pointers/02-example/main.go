// an example not using pointers

package main

import (
	"fmt"
)

func main() {
	var a int = 42
	// point to an integer *int
	// b is going  to use the address of operator
	// b is not holding the value 42
	// b is holding the memory location that is holding
	// a data
	var b *int = &a

	// 42 and its memory address
	fmt.Println(a, b)
	// checking if the memory values are the same
	fmt.Println(&a, b)
}
