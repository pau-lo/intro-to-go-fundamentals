package main

import (
	"fmt"
)

func main() {
	// creating an anonymous struct
	//         struct keyword & structure/ data into the struct
	aDoctor := struct{ name string }{name: "John Pertwee"}

	// creating a copy
	// anotherDoctor := aDoctor
	// pointing to the same underlie data
	anotherDoctor := &aDoctor
	// new doctor name
	anotherDoctor.name = "Tom Baker"
	// everything is independent but if using a pointer it will
	// point to the same underlie data
	fmt.Println(aDoctor)
	fmt.Println(anotherDoctor)
}
