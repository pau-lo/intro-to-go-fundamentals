package main

import (
	"fmt"
)

func main() {
	a, b := 1, 0
	// int divided by zero
	// this will generate a panic
	ans := a / b
	fmt.Println(ans)
}
