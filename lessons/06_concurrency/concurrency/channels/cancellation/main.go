package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	results := make(chan string, 1)
	go func() {
		select {
		case <-time.After(50 * time.Millisecond):
			results <- "work complete"
		case <-ctx.Done():
			return
		}
	}()

	select {
	case result := <-results:
		fmt.Println(result)
	case <-ctx.Done():
		fmt.Println("work cancelled:", ctx.Err())
	}
}
