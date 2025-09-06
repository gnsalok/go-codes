package main

import (
	"fmt"
	"runtime"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		// Perform some CPU-bound task
		result := j * 2

		// Simulate some processing time
		for i := 0; i < 1000000; i++ {
			result += i
		}

		// Send result back to the results channel
		results <- result
	}
}

func main() {
	// Set the maximum number of operating system threads to utilize
	runtime.GOMAXPROCS(4) // Set to the number of CPU cores you want to utilize

	numJobs := 10
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Create some worker goroutines
	var wg sync.WaitGroup
	for i := 1; i <= 4; i++ { // Adjust the number of goroutines based on the number of cores
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			worker(workerID, jobs, results)
		}(i)
	}

	// Send jobs to the jobs channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results from the results channel
	go func() {
		wg.Wait()
		close(results)
	}()

	// Print results
	for result := range results {
		fmt.Println("Result:", result)
	}
}
