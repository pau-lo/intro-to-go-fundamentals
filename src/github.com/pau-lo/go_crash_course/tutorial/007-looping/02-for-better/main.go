package main

import (
	"fmt"
)

func main() {

	// i scope to the main function not
	// the for loop
	i := 0

	for i < 5 {
		fmt.Println(i)
		i++
	}

}
