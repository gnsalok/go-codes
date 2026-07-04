package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

type job struct {
	ID int
}

// processJob simulates work that can fail. In production this could be an API
// call, database write, or CPU-heavy transformation.
func processJob(ctx context.Context, j job) error {
	select {
	case <-time.After(100 * time.Millisecond):
		if j.ID == 7 {
			return fmt.Errorf("job %d failed", j.ID)
		}
		fmt.Println("processed job", j.ID)
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func runWorkerPool(ctx context.Context, jobs []job, workerCount int) error {
	g, ctx := errgroup.WithContext(ctx)
	jobCh := make(chan job)

	for workerID := 1; workerID <= workerCount; workerID++ {
		workerID := workerID
		g.Go(func() error {
			for {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case j, ok := <-jobCh:
					if !ok {
						return nil
					}
					if err := processJob(ctx, j); err != nil {
						return fmt.Errorf("worker %d: %w", workerID, err)
					}
				}
			}
		})
	}

	// The producer owns closing jobCh. It stops feeding new work when any worker
	// returns an error because errgroup cancels ctx.
	g.Go(func() error {
		defer close(jobCh)
		for _, j := range jobs {
			select {
			case jobCh <- j:
			case <-ctx.Done():
				return ctx.Err()
			}
		}
		return nil
	})

	return g.Wait()
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	jobs := make([]job, 10)
	for i := range jobs {
		jobs[i] = job{ID: i + 1}
	}

	if err := runWorkerPool(ctx, jobs, 3); err != nil {
		fmt.Println("worker pool stopped:", err)
		return
	}
	fmt.Println("all jobs processed")
}
