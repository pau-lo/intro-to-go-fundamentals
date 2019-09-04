// interger types

package main

import "fmt"

func main() {

	// some arithmetic operations

	a := 12
	b := 3

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	// we can't add two different integers
	// var x int = 10
	// var y int8 = 3
	// (mismatched types int and int8)
	// fmt.Println(x + y)
	var x int = 10
	var y int8 = 3

	fmt.Println(x + int(y))

}
