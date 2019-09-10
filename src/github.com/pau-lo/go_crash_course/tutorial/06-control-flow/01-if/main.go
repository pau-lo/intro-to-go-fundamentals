// if-statements

package main

import (
	"fmt"
)

func main() {
	someState := map[string]int{
		"California": 1234567890,
	}
	if pop, ok := someState["California"]; ok {
		fmt.Println(pop)
	}
}
