// Go’s structs are collections of fields.
// They’re useful for grouping data together to form records.

package main

import (
	"fmt"
)

// this Doctor struct has number, actorName and companions field names
// Caps == exported from the package
type Doctor struct {
	number     int
	actorName  string
	episodes   []string
	companions []string
}

func main() {
	// create a variable aDoctor
	aDoctor := Doctor{

		// create 3 values associated with it
		// listing the names
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}

	// print all
	fmt.Println(aDoctor)

	// just the actor name
	fmt.Println(aDoctor.actorName)

	// print just the first companion like in slices
	fmt.Println(aDoctor.companions[0])

	// using the positional syntax
	fmt.Println()

	//

}
