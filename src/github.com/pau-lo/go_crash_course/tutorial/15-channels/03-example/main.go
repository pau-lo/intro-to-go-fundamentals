package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {

	ch := make(chan int)
	wg.Add(2)

	go func() {
		i := <-ch
		fmt.Println(i)
		ch <- 27
		wg.Done()

	}()
	// pushing a message into the channel
	go func() {
		ch <- 42
		fmt.Println(<-ch)
		wg.Done()
	}()

	wg.Wait()

}
