package main

import (
	"fmt"
	"sync"
)

/*
Problem with Many Goroutines → Cleaner to use sync.WaitGroup
Imagine:
- 100 workers
- Single jobs channel
- You want to wait till ALL finish.
*/

func main() {
	var wg sync.WaitGroup
	jobs := make(chan int, 100) // Buffered channel to hold jobs

	// Start 100 workers
	for w := 0; w < 100; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				// Process the job
				fmt.Printf("Worker %d processing job %d\n", workerID, job)
			}
		}(w)
	}

	// Send 100 jobs to the jobs channel
	for j := 0; j < 100; j++ {
		jobs <- j
	}

	close(jobs) // Close the channel to signal workers no more jobs

	wg.Wait() // Wait for all workers to finish
}
