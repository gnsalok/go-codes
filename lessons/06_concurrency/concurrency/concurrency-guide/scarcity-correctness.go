package main

import (
	"fmt"
	"sync"
	"time"
)

/*
See the concurrency note for details

Where we discuss about 3 Pilar of Concurrency: Correctness, Coordination, and Scarcity.

In this example, we are demonstrating the "Scarcity" pillar by using a buffered channel as a semaphore to limit the number of concurrent goroutines accessing a shared resource (simulated by the `accessResource` function).
The "Correctness" aspect is maintained by ensuring that no more than 5 goroutines can access the resource at the same time, thus preventing potential issues that could arise from too many concurrent accesses (like overwhelming a database connection pool).

*/

func main() {
	// 1. SCARCITY: Limit concurrency to 5 using a buffered channel
	semaphore := make(chan struct{}, 5)
	var wg sync.WaitGroup

	for i := 1; i <= 20; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Acquire: Push into channel. If full, this blocks.
			semaphore <- struct{}{}

			// 2. CORRECTNESS: The "Work" is protected by the semaphore
			/*
				Note: So the correctness property is: no more than 5 goroutines can execute accessResource(id) concurrently.
				That’s scarcity, not general correctness (ie. Mutext)

			*/
			accessResource(id)

			// Release: Pull from channel to let another goroutine in
			<-semaphore
		}(i)
	}

	wg.Wait()
	fmt.Println("All tasks finished.")
}

func accessResource(id int) {

	fmt.Printf("Worker %d is using a scarce DB connection...\n", id)
	time.Sleep(500 * time.Millisecond) // Simulate I/O
}
