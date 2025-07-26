package main

import (
	"fmt"
	"sync"
)

func main() {

	// write channel to add even and odd numbers and read channel to read them
	evenCh := make(chan int)
	oddCh := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	// Goroutine to send even numbers
	go func() {
		defer wg.Done()
		for i := 0; i <= 10; i++ {
			if i%2 == 0 {
				evenCh <- i
			}
		}
		close(evenCh)
	}()

	// Goroutine to send odd numbers
	go func() {
		defer wg.Done()
		for i := 0; i <= 10; i++ {
			if i%2 != 0 {
				oddCh <- i
			}
		}
		close(oddCh)
	}()

	// Goroutine to read even numbers
	go func() {
		for num := range evenCh {
			fmt.Println("Even:", num)
		}
	}()

	// Goroutine to read odd numbers
	go func() {
		for num := range oddCh {
			fmt.Println("Odd:", num)
		}
	}()

	wg.Wait()
}
