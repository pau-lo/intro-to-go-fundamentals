package main

import "fmt"

// create a new type of 'deck'
// which is a slice of strings

type deck []string

func (d deck) print() { // the actual copy of the deck we're working with is available
	// is avail in the func as a var called 'd'

	for i, card := range d {
		fmt.Println(i, card)
	}

}
