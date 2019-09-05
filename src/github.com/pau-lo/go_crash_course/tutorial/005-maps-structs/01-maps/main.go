// maps

package main

import (
	"fmt"
)

func main() {

	statePopulations := map[string]int{
		"California":   39536653,
		"Texas":        28304596,
		"Florida":      20984400,
		"New York":     19849399,
		"Pennsylvania": 12805537,
		"Illinois":     12802023,
		"Ohio":         11658609,
		"Georgia":      10429379,
	}

	fmt.Println(statePopulations)

}
