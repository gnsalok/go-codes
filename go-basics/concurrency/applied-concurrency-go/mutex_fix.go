package main

import (
	"fmt"
	"sync"
)

var x = 0

// A Mutex is used to provide a locking mechanism to ensure that only one
// Goroutine is running the critical section of code at any point in time
// to prevent race conditions from happening.

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	defer m.Unlock()
	x = x + 1
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
