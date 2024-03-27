package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Channels Demo...")

	wg.Add(2)

	ch := make(chan int)
	input := 100

	go send(ch, input)

	go receive(ch)

	wg.Wait()

}

func send(ch chan<- int, input int) {
	defer wg.Done()
	ch <- input
}

func receive(ch <-chan int) {
	defer wg.Done()
	fmt.Println(<-ch)
}
