package main

import (
	"fmt"
	"sync"
)

func worker(id int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this goroutine as done
	results <- fmt.Sprintf("Worker %d finished", id)
}

func main() {
	var wg sync.WaitGroup
	results := make(chan string, 10)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, results, &wg)
	}

	go func() {
		wg.Wait()      // Wait for all workers
		close(results) // Close channel when all workers are done
	}()

	for result := range results {
		fmt.Println(result)
	}
}
