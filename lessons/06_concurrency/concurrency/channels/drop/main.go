package main

import "fmt"

func main() {
	const capacity = 2
	work := make(chan int, capacity)

	for jobID := 1; jobID <= 5; jobID++ {
		select {
		case work <- jobID:
			fmt.Println("accepted:", jobID)
		default:
			fmt.Println("dropped:", jobID)
		}
	}
	close(work)

	for jobID := range work {
		fmt.Println("processed:", jobID)
	}
}
