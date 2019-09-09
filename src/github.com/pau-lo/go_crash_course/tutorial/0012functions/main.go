package main

import "fmt"

func main() {

	car := newCar()
	caryear := carYear()

	fmt.Println(car)
	fmt.Println(caryear)

}

func newCar() string { // new function / tell it is a string datatype
	return "I own a FERRARI"
}

func carYear() string {
	return "The FERRARI was made in 2012"
}
