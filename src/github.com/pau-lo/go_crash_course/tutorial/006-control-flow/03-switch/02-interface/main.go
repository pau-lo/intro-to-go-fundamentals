package main

import "fmt"

func main() {
	var i interface{} = 3.14 // [3]int{}
	switch i.(type) {        // tell go to pull the actual
	// underline type for that interface
	case int:
		fmt.Println("i is an int")
	case string:
		fmt.Println("is a string")
	case float64:
		fmt.Println("i is a float64")
	default:
		fmt.Println("i is another type")
	}
}
