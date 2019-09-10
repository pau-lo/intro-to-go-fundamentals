// example of a variadic parameter
// when using it you can only have one and has to be the last one

package main

import (
	"fmt"
)

func main() {
	// geneeric function passingthe values
	sum(1, 2, 3, 4, 5)

}

// preceeding typ e with the ... dots
func sum(values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is ", result)
}
