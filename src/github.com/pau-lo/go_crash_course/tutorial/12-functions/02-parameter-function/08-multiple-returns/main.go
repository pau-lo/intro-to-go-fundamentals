package main

import (
	"fmt"
)

func main() {
	// d := divide(5.0, 3.0)
	d, err := divide(5.0, 0.0) // using this will get an inference
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}

func divide(a, b float64) (float64, error) {

	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
		// panic("Cannot provide zero as second value")
	}

	return a / b, nil
}
