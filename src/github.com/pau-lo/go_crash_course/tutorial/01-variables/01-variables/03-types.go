package main

import "fmt"

func main() {
	var age int32 = 21
	const isDope = true

	name, email := "Paul", "paul@gmail.com"
	size := 1.3 // automatically gives a float64

	fmt.Println(name, email, age, isDope)
	fmt.Printf("%T\n", size)

}
