// skeleton of a basic hello go file

// collection of files: executable file
// generates a file that we can run
package main // package declaration

// standard library package format
// linked to main ... A fmt is the raw formatter used by Printf etc.
// Package fmt implements formatted I/O
import "fmt" // import packages/libraries that we need

// func tells go that we are about to create a function
func main() { // main sets the name of the function () list of arguments to pass to the function

	fmt.Println("Hi, there!")
}
