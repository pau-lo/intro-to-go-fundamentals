package main

import (
	"fmt"
)

// main function
func main() {
	// declaring the greeter struct and then
	g := greeter{
		greeting: "Hello",
		name:     "Go!",
	}

	// invoking the method
	g.greet() // calling the function
	fmt.Print("The new name is: ", g.name)
}

// struct called greeter
type greeter struct {
	greeting string
	name     string
}

// this is the method called greet()
// A Method is just a function that is executing in a know context:
// A known context is any type: string, int etc... in go
// looks like a function except for (g greeter)
// which makes a method

// g greeter is the value receiver
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
	// adding a value to the name field
	g.name = "Howard"
}
