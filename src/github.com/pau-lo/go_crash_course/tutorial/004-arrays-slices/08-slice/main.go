// stack operations

// array of arrays

package main

import (
	"fmt"
)

func main() {

	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	b := append(a[:2], a[3:]...)
	fmt.Println(b)
	fmt.Println(a) // 5 is duplicated

}
