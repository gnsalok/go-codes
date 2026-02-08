package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	ch := make(chan int)

	//send
	go foo(ch)

	//receive
	go bar(ch)

	wg.Wait()

}

//send
func foo(ch chan<- int) {
	defer wg.Done()
	ch <- 33

}

// receive
func bar(ch <-chan int) {
	defer wg.Done()
	fmt.Println(<-ch)

}
