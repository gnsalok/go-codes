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
		// fmt.Println("value from the channel", <-ch) // get 0

		// if control flow
		if msg, ok := <-ch; ok {
			fmt.Println("value from the channel", msg)
		}

		wg.Done()
	}(ch, wg)

	// send only
	go func(ch chan int, wg *sync.WaitGroup) {

		ch <- 12

		wg.Done()
	}(ch, wg)

	wg.Wait()

}
