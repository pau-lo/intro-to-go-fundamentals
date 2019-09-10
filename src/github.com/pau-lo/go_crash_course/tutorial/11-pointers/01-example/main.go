// an example not using pointers

package main

import (
	"fmt"
)

func main() {
	a := 42
	// b copy the data in a
	b := a
	// 42 will print 2x
	fmt.Println(a, b)
}
