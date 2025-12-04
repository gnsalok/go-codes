package main

import (
	"fmt"
	"time"
)

// Scenario: A relay race. The baton (data) must be physically handed
// from one runner to the next. They must meet at the same point.

func main() {

	// 1. create unbuffered channel
	baton := make(chan string)

	// 2. Start the Runner (Goroutine)
	go func() {
		fmt.Println("Runner: waiting for the baton...")

		// This Receive BLOCKS until main sends the data.
		// This guarantees synchronization.
		// receive happen before the send
		obj := <-baton

		fmt.Printf("Runner: Got the baton! Running with: %v\n", obj)

		time.Sleep(1 * time.Second)

		fmt.Println("Runner: Finished race.")
	}()

	fmt.Println("Main: Getting ready to pass baton...")
	time.Sleep(2 * time.Second)

	// 3. Send data.
	// This BLOCKS until the Runner is ready to receive.
	// We know for a fact the Runner is listening when this line completes.
	baton <- "Gold Baton"

	fmt.Println("Main: Baton passed successfully.")

	// Keep main alive just long enough to see the runner finish
	time.Sleep(2 * time.Second)

}
