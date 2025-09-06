package main

import (
	"fmt"
)

func worker1(id int, results chan string, done chan bool) {
	results <- fmt.Sprintf("Worker %d finished", id)
	done <- true // Signal completion
}

func main() {
	results := make(chan string) // Unbuffered channel
	done := make(chan bool)      // Channel to signal completion

	numWorkers := 5

	for i := 1; i <= numWorkers; i++ {
		go worker1(i, results, done)
	}

	go func() {
		for i := 0; i < numWorkers; i++ {
			<-done // Wait for each worker to signal completion
		}
		close(results) // Close results channel after all workers are done
	}()

	for result := range results {
		fmt.Println(result)
	}
}
