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
	go func(ch <-chan int, wg *sync.WaitGroup) {

		for msg := range ch {
			fmt.Println(msg)
		}

		wg.Done()
	}(ch, wg)

	// send only
	go func(ch chan<- int, wg *sync.WaitGroup) {

		for i := 0; i < 10; i++ {
			ch <- i
		}

		// if you will not close the channel for loop at receiving side panic
		// for together work with close() function to detect when to stop the iteration.
		close(ch)

		wg.Done()
	}(ch, wg)

	wg.Wait()

}
