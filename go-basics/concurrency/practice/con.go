package main

import (
	"fmt"
	"sync"
	"time"
)

// Main goroutine
func main() {

	/*
		- Concurrency is dealing with a lot of work at once. (Threads)
		- Parallelism is doing a lot work at once.

		- Go usage concurrency, it abstract them through a lightweight construct layered on top, called Goroutine.
		  Its not schedule by os it’s Go RunTime which scheduled them.

		- WaitGroups : A WaitGroup waits for a collection of goroutines to finish
			Functions : ws.Add(), defer wg.Done(), wg.Wait()

	*/

	// Write a concurrent program using Go.
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	fmt.Println("Waiting...")

	go func() {
		defer waitGrp.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("First Func")
	}()

	go func() {
		defer waitGrp.Done()
		fmt.Println("Second Func")
	}()

	waitGrp.Wait()
}
