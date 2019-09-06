// switch statements

package main

import (
	"fmt"
)

func main() {
	switch 1 {
	case 1:
		fmt.Println("one")
		// break >>>  is imply
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("not one or two")
	}
}
