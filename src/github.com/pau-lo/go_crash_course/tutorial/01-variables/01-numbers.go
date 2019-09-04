// Variable declaration
// redeclaration and shadowing
// visibility
// naming conventions
// type conversions

package main

import "fmt"

// at the package level
var l int = 99

func main() {
	// 3 ways to declare

	// 1st way
	var i int
	i = 42
	i = 27

	// 2nd way
	var j float32 = 33

	// 3rd way
	k := 45

	fmt.Println(i)
	fmt.Println(j)

	fmt.Printf("%v, %T", j, j)

	fmt.Println(k)

	fmt.Println(l)
}
