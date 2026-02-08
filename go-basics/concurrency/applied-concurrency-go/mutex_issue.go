package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}
func main() {
	var w sync.WaitGroup

	// as there is not Mutex, results will be inconsitent while every run.
	// to avoid it use Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w)
	}

	w.Wait()
	fmt.Println("final value of x", x)
}
