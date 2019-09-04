package main

import (
	"fmt"
)

// declaring variables
//var actorName string = "Elisabeth Sladen"
//var companion string = "Sara Jane Smith"
//var doctorName int = 3
//var season int = 11

// other way: inside a block at the package level

var (
	actorName  string = "Elisabeth Sladen"
	companion  string = "Sara Jane Smith"
	doctorName int    = 3
	season     int    = 11
)

func main() {

	fmt.Println(actorName, companion, doctorName, season)

}
