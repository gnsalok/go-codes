package main

import (
	"fmt"
	"time"
)

// Scenario: A Chef (Producer) cooking orders and placing them on a
// pickup counter (Buffer). The Waiter (Consumer) delivers them.
// The Chef works faster than the Waiter.

func main() {

	// 1. Create a buffered channel with capacity 3
	// The Chef can cook 3 plates ahead of time without waiting.
	counter := make(chan string, 3)

	// 2. Start the Chef (Producer)
	go func() {
		defer close(counter) // close when done processing

		for i := 1; i <= 5; i++ {
			dish := fmt.Sprintf("Dish #%d", i)
			fmt.Printf("Chef cooked %s\n", dish)

			// SENDING
			// If buffer has space (size < 3), this happens instantly.
			// If buffer is full (size == 3), this BLOCKS until Waiter takes one.
			counter <- dish
		}
		fmt.Println("Chef: All order cooked, I'm going home now.")
	}()

	// 3. (main goroutine) The Waiter (Consumer) - Slower than the Chef
	// Range loop handles the receiving logic automatically
	for dish := range counter {
		fmt.Printf("\tWaiter: Delivering %s...\n", dish)
		time.Sleep(1 * time.Second) // simulate slow delivery
	}

	fmt.Println("Restaurant closed.")
}
