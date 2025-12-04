package main

import (
	"fmt"
	"sync"
	"time"
)

// worker function
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for j := range jobs {
		fmt.Println("worker %d started job %d\n", j)
		time.Sleep(500 * time.Microsecond) // simulate work
		results <- j * 2                   // send result
	}

}

func main() {

	const numJobs = 10
	const numWorkers = 3

	// 1. Create Channels
	// Buffered channels are great here to prevent blocking the generator
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 2. Setup waitGroup
	var wg sync.WaitGroup

	// 3. Spawn Workers
	// We start 3 goroutines that all listen to the SAME jobs channel
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// 4. Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Signal to workers that no more jobs are coming

	// 5. Wait for completion in a separate goroutine
	// Why separate? Because we need to close 'results' channel. If you will not close the results channel:
	/*
			* A range loop only stops when the channel is explicitly closed. If you don't close it,
		// the loop finishes reading the 10 results and then sits there, waiting for an 11th result that will never come.
		// 	* Runtime Panic: The Go runtime sees that the main goroutine is waiting on results,
		// but there are no other active goroutines left to send data to it.
		// It detects this impossible state and crashes with: fatal error: all goroutines are asleep - deadlock!
	*/
	// to let the main function's range loop terminate.
	go func() {
		wg.Wait()      // Wait for all workers to finish
		close(results) // Safe to close results now
	}()

	// 6. Collect Results
	// Note : Read above why results need to be closed before range over channel
	for res := range results {
		fmt.Println("Result:", res)
	}

}
