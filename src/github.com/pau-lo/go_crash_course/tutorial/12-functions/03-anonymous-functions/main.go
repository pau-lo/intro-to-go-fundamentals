package main

import (
	"fmt"
)

func main() { // not good practice todo this
	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println("Hello, go!")
		}(i) // parenthesis to invoke this func
	}
}
