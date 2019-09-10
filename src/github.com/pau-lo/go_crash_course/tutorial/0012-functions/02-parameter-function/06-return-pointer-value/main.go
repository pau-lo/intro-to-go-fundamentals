// example of a variadic parameter
// when using it you can only have one and has to be the last one

package main

import (
	"fmt"
)

func main() {
	// geneeric function passingthe values
	s := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum is ", *s)

}

// type with the ... dots
// s == integer type
func sum(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	// instead of returning the result
	// will return the address of the result
	// it will be the same way
	return &result
}
