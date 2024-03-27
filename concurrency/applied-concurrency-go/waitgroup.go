package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	// always pass waitGroup by reference otherwise it will set the counter to 0 in function and when it calls defer wg.Done()
	// counter value becomes -ve and program panics
	go hello(&wg)

	wg.Wait() // it will wait for counter to be 0

	goodbye()

}

func hello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello")
}

func goodbye() {
	fmt.Println("Goodbye")
}
