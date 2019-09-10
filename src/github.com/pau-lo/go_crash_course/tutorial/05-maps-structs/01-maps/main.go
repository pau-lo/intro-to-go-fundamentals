// maps

package main

import (
	"fmt"
)

func main() {

	// statePopulations := make(map[string]int)
	statePopulations := map[string]int{
		"California":   39536653,
		"Texas":        28304596,
		"Florida":      20984400,
		"New York":     19849399,
		"Pennsylvania": 12805537,
		"Illinois":     12802023,
		"Ohio":         11658609,
	}

	// adding values to our map
	// statePopulations["Georgia"] = 10429379
	// fmt.Println(statePopulations["Georgia"])

	// printing all to see that Georgia was added
	// fmt.Println(statePopulations)

	// deleting from map
	//delete(statePopulations, "Georgia")
	//fmt.Println(statePopulations)

	// checking to see if Georgia still there
	// fmt.Println(statePopulations["Georgia"]) // we get a zero? == missing from map

	// checking what's on our map with the comma ok syntax
	//_, ok := statePopulations["Ohio"]
	//fmt.Println(ok) // population, true === it exists

	// checking the length of the map
	//fmt.Println(len(statePopulations)) // 7 states

	// making a copy of the original map
	sp := statePopulations
	delete(sp, "Ohio") // deleting ohio from sp

	fmt.Println(sp)               // ohio should not be here
	fmt.Println(statePopulations) // original

}
