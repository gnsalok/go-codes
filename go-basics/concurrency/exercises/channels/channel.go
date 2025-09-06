package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	// receive only
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println("Value from the channel: ", <-ch)
		wg.Done()
	}(ch, wg)

	// send only
	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 12
		// closing the channel
		close(ch)
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
