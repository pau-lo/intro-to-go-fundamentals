// passing parameters as value types
// passing the name variables by value : data is not change
// passing pointers: values will change
// working with maps and sclices will not work here

package main

import (
	"fmt"
)

func main() {
	// passing 2 strings
	greeting := "Hello"
	name := "Stacey"
	sayGreeting(&greeting, &name)
	fmt.Println(name)
}

// compile infer that both are strings
// * passing in pointers
// & acquiring the memory address
func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name) // dereference the pointers here
	*name = "Ted"
	fmt.Println(*name)
}
