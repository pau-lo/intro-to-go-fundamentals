package main

import "fmt"

func main() {
	var n complex64 = 1 + 2i
	fmt.Printf("%v , %T\n", n, n)

	// arithmetic
	a := 1 + 2i
	b := 2 + 5.2i
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
}
