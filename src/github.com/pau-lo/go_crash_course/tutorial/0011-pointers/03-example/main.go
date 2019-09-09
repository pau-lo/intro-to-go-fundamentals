package main

import (
	"fmt"
)

func main() {
	//
	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1] // c := &a[1] pointer arithmetic is not allowed
	// print the value of the array and the value of the 2 pointers
	fmt.Printf("%v %p %p\n", a, b, c)
}
