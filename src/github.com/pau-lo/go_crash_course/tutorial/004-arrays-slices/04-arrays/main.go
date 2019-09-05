// array of arrays

package main

import (
	"fmt"
)

func main() {

	a := [...]int{1, 2, 3}
	//b := a // assing to copy of array a
	b := &a // using a pointer a and b are point to the same data
	// changing same data for both since they are the same array
	b[1] = 5 // but assign 5 to the 2nd element
	fmt.Println(a)
	fmt.Println(b)
}
