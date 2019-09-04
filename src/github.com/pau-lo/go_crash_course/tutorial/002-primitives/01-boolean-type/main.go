package main

import "fmt"

func main() {

	// every value in go has a 0 value
	var f bool
	fmt.Printf("%v, %T\n ", f, f)

	// using logical operators
	p := 1 == 1
	m := 1 == 2

	fmt.Printf("%v, %T\n ", p, p)
	fmt.Printf("%v, %T\n ", m, m)

	var n bool = false

	fmt.Printf("%v, %T\n ", n, n)

}
