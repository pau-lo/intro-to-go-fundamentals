// channels basic example

package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// receiving data from the channel and
	// assign it to i
	ch := make(chan int)
	wg.Add(2)

	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()

	}()

	go func() {
		ch <- 42
		wg.Done()
	}()

	wg.Wait()

}
