package main

import "fmt"

var I int = 45 // cap letter will make it gobally visible

func main() {

	var i int = 45 // this var is scoped to to the block
	// never visible outside the main block
	fmt.Println(i)
}
