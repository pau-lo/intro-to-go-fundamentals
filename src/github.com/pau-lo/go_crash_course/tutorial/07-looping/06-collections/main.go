// in collection we can use the range function
// for arrays and slices

package main

import (
	"fmt"
)

func main() {

	// for a

	s := []int{1, 2, 3, 4}

	for k, v := range s {

		fmt.Println(k, v)

	}

	// for an array

	a := [3]int{1, 2, 3}

	for k, v := range a {

		fmt.Println(k, v)
	}

	// for map
	someNames := map[string]int{
		"John":     23,
		"Johnny":   24,
		"Jon":      25,
		"Jonathan": 26,
		"Joe":      27,
		"Joan":     28,
		"Joey":     29,
		"Jorge":    20,
		"Joris":    21,
	}

	// getting all the values
	for k, v := range someNames {
		fmt.Println(k, v)
	}

	// getting only the names
	for k := range someNames {
		fmt.Println(k)
	}

	// getting the values only
	for _, k := range someNames {
		fmt.Println(k)
	}

	// using a string

	h := "Hello Go!"

	for k, v := range h {
		fmt.Println(k, string(v))
	}

}
