package main

import "fmt"

type result struct {
	jobID int
	value int
}

func main() {
	const jobs = 5
	results := make(chan result, jobs)

	for jobID := 1; jobID <= jobs; jobID++ {
		go func(id int) {
			results <- result{jobID: id, value: id * id}
		}(jobID)
	}

	for range jobs {
		completed := <-results
		fmt.Printf("job %d result=%d\n", completed.jobID, completed.value)
	}
}
