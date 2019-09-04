// converting form int to strings

package main

import (
	"fmt"
	"strconv" // string conversion package
)

func main() {
	var i int = 42
	fmt.Printf("%v , %T\n", i, i)

	var j string
	j = strconv.Itoa(i)
	fmt.Printf("%v , %T\n", j, j)
}
