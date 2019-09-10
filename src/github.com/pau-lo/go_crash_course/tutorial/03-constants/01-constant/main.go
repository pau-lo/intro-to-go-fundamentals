package main

import (
	"fmt"
)

func main() {
	const a = 42
	var b int16 = 27
	// myConst = 47  // cannot assigned new value to a constant
	fmt.Printf("%v, %T\n", a+b, a+b)
}
