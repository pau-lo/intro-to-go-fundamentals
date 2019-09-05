package main

import (
	"fmt"
)

// at package level

// const a = iota // iota is a counter we can use when using enumerated constants

const (
	a = iota // since a pattern has been established
	b        // = iota
	c        // = iota
)

func main() {

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)

}
