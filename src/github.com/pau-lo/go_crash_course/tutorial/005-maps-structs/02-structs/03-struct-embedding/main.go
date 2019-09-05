package main

import (
	"fmt"
)

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	// embedding the Animal struct into the Bird struct
	Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {

	// on way to declare it
	//b := Bird{}
	//b.Name = "Emu"
	//b.Origin = "Australia"
	//b.SpeedKPH = 48
	//b.CanFly = false

	b := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 48,
		CanFly:   false,
	}

	fmt.Println(b)
	fmt.Println(b.Name)

}
