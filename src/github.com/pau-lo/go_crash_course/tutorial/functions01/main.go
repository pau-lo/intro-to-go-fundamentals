package main

import "fmt"

// our first function
func greeting(name string) string {
	return "Hello, " + name
}

// 2nd function

func howAreYou(reply string) string {
	return "I hope you are " + reply
}

// 3rd function
func getDifference(num1 int, num2 int) int {
	return num1 - num2
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

// main function: will run here
func main() {

	fmt.Println(greeting("Paul."))
	fmt.Println(howAreYou("fine!"))
	fmt.Println(getDifference(12, 3))
	fmt.Println(getSum(9, 8))
}
