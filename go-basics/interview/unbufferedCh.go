/*

By default channels are unbuffered, meaning that they will only accept sends (chan <-),
if there is a corresponding receive (<- chan) ready to receive the sent value.
Buffered channels accept a limited number of values without a corresponding receiver for those values.

*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	// // Deadlock : bcz channel is un-buffered
	// msg := make(chan string)
	// msg <- "hi"
	// fmt.Println(<-msg)

	// // works fine with buffered channel
	// msg1 := make(chan string, 1)
	// msg <- "hi"
	// fmt.Println(<-msg1)

	// TO make it work

	wg.Add(2)

	ch := make(chan string)
	go send(ch, "hi")
	go receive(ch)

	wg.Wait()

}

func send(ch chan<- string, input string) {
	defer wg.Done()
	ch <- input
}

func receive(ch <-chan string) {
	defer wg.Done()
	fmt.Println(<-ch)
}
