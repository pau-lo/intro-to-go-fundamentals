package main

import (
	"fmt"
)

func main() {

	// i scope to the main function not
	// the for loop
	i := 0

	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break

		}
	}

}
