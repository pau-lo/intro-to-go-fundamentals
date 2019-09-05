package main

import (
	"fmt"
)

func main() {

	// name = size / type
	// grades := [3]int{100, 98, 85} // declaring the array 2x
	grades := [...]int{100, 98, 85} // ... means create an array large enough for the data i'm about to put in

	fmt.Printf("Grades: %v", grades)

}
