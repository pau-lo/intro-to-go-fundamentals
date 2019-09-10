// guess the number

package main

import (
	"fmt"
)

func main() {

	number := 8
	guess := 3

	if guess < 1 {

		fmt.Println("The guess number must be greate than 1.")

	} else if guess > 10 {
		fmt.Println("The guess number must be less than 10.")
	} else {

		if guess < number {
			fmt.Println("Too high, try again!")
		}

		if guess > number {
			fmt.Println("Too high!")
		}

		if guess == number {
			fmt.Println("You guess it!")
		}
	}

	// main output
	fmt.Println(number <= guess, number >= guess, number != guess)

}
