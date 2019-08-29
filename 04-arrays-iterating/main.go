// Array: fixed length list of things

// Slice: An array that can grow or shrink
// every element of a slice must be of same data type

package main

import "fmt"

func main() {

	// creating a slice
	words := []string{"Hello, how are you?", newWords()} // 1
	words = append(words, "See ya!")                     // 3

	// index, current where we are iterating over := take the slice of words and loop over it
	for i, word := range words {
		fmt.Println(i, word) // run it one time for each word in the slice
	}

}

func newWords() string {
	return "Nice to meet you!" // 2
}
