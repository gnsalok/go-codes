package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	// 1. Create a root context that we can cancel.
	// This is the 'Master Signal'.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 2. Use errgroup for 'All-or-Nothing' orchestration.
	// If one function returns an error, ctx is cancelled for everyone.
	g, ctx := errgroup.WithContext(ctx)

	// Orchestrating Component A: Database
	g.Go(func() error {
		return startComponent(ctx, "Database")
	})

	// Orchestrating Component B: File Watcher
	g.Go(func() error {
		return startComponent(ctx, "File Watcher")
	})

	// Orchestrating Component C: API Server
	g.Go(func() error {
		// Let's pretend the API server fails after 2 seconds
		time.Sleep(2 * time.Second)
		fmt.Println("!!! API Server encountered a critical error !!!")
		return fmt.Errorf("api failure")
	})

	// 3. Wait for everything to finish or for a failure to occur
	if err := g.Wait(); err != nil {
		fmt.Printf("Orchestrator: Shutting down entire system due to: %v\n", err)
	}
}

func startComponent(ctx context.Context, name string) error {
	for {
		select {
		case <-ctx.Done(): // 4. This is the 'Listen' part of orchestration
			fmt.Printf("Shutting down %s cleanly...\n", name)
			return ctx.Err()
		default:
			fmt.Printf("%s is running...\n", name)
			time.Sleep(1 * time.Second)
		}
	}
}
