package main

import "fmt"

func worker1(id int, results chan<- string, done chan<- struct{}) {
	// results is buffered to numWorkers in main. That guarantees each worker can
	// publish its single result before sending the done signal.
	results <- fmt.Sprintf("Worker %d finished", id)
	done <- struct{}{}
}

func main() {
	numWorkers := 5

	results := make(chan string, numWorkers)
	done := make(chan struct{})

	for i := 1; i <= numWorkers; i++ {
		go worker1(i, results, done)
	}

	go func() {
		for i := 0; i < numWorkers; i++ {
			<-done
		}
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}
}
