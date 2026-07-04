package main

import (
	"context"

	"golang.org/x/sync/errgroup"
)

// Task defines the work to be done
type Task func(ctx context.Context) error

// RunWorkerPool demonstrates a robust concurrency pattern
func RunWorkerPool(ctx context.Context, tasks []Task, workerCount int) error {
	// errgroup handles error propagation and synchronization
	g, ctx := errgroup.WithContext(ctx)
	taskChan := make(chan Task)

	// 1. Start Workers
	for i := 0; i < workerCount; i++ {
		g.Go(func() error {
			for task := range taskChan {
				select {
				case <-ctx.Done():
					return ctx.Err()
				default:
					if err := task(ctx); err != nil {
						return err // This cancels the context for other workers
					}
				}
			}
			return nil
		})
	}

	// 2. Feed the workers
	g.Go(func() error {
		defer close(taskChan)
		for _, task := range tasks {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case taskChan <- task:
			}
		}
		return nil
	})

	return g.Wait()
}
