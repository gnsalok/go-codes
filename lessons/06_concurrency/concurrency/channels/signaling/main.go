package main

import "fmt"

func main() {
	waitForTask()
	waitForResult()
}

// waitForTask uses an unbuffered send as both a data handoff and a
// synchronization point: the send completes only after the receiver is ready.
func waitForTask() {
	tasks := make(chan string)
	done := make(chan struct{})

	go func() {
		fmt.Println("worker received:", <-tasks)
		close(done)
	}()

	tasks <- "review pull request"
	<-done
}

// waitForResult reverses ownership: the caller waits for a worker result.
func waitForResult() {
	results := make(chan string)

	go func() {
		results <- "review complete"
	}()

	fmt.Println("caller received:", <-results)
}
