// simple loop

package main

import (
	"fmt"
)

func main() {

	// setting i = 0 >> initializer , boolean result to determine
	// if is done looping or not, increment by 2
	for l := 0; l < 6; l = l + 2 {
		fmt.Println(l)
	}

	fmt.Println("------------------------------")

	for i, j := 0, 0; i < 6; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	fmt.Println("------------------------------")

	for k := 0; k < 6; k++ {
		fmt.Println(k)

		if k%2 == 0 {
			k /= 2
		} else {
			k = 2*k + 1
		}
	}

}
