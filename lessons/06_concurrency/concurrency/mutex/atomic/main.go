// Sample program to show how to use the atomic package to
// provide safe access to numeric types.

// https://github.com/gnsalok/gotraining/tree/master/topics/go/concurrency/data_race
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64

func main() {
	const goroutines = 2

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for range goroutines {
		go func() {
			defer wg.Done()
			for i := 0; i < 2; i++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
