package main

import (
	"fmt"
	"sync"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	// Simulate some work
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	// Launch several goroutines using WaitGroup.Go
	for i := 1; i <= 5; i++ {
		// wg.Go handles the Add(1) and defer wg.Done() calls internally
		wg.Go(func() {
			worker(i)
		})
	}

	// Block until all the goroutines started by wg are done
	wg.Wait()
	fmt.Println("All workers finished")
}
