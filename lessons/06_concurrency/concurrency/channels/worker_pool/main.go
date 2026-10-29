package main

import (
	"fmt"
	"sync"
)

type result struct {
	jobID int
	value int
}

func main() {
	jobs := make(chan int)
	results := make(chan result)

	var workers sync.WaitGroup
	for workerID := 1; workerID <= 2; workerID++ {
		workers.Add(1)
		go func() {
			defer workers.Done()
			for jobID := range jobs {
				results <- result{jobID: jobID, value: jobID * jobID}
			}
		}()
	}

	go func() {
		for jobID := 1; jobID <= 5; jobID++ {
			jobs <- jobID
		}
		close(jobs)
	}()

	go func() {
		workers.Wait()
		close(results)
	}()

	for completed := range results {
		fmt.Printf("job %d result=%d\n", completed.jobID, completed.value)
	}
}
