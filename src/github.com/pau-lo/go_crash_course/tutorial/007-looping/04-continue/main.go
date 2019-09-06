package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		// printing the odd numbers only
		fmt.Println(i)
	}

}
