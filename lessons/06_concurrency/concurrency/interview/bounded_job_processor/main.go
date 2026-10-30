package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// ErrInvalidWorkerCount reports a worker limit that cannot make progress.
var ErrInvalidWorkerCount = errors.New("worker count must be positive")

// ProcessJobs runs process for every job with at most workerCount calls active
// at once. The first processing error cancels the remaining work.
//
// A process function must observe ctx and return for cancellation to be prompt.
// ProcessJobs waits for all started calls before returning, so none of its
// goroutines can outlive the call.
func ProcessJobs(
	ctx context.Context,
	jobs []int,
	workerCount int,
	process func(context.Context, int) error,
) error {
	if workerCount <= 0 {
		return ErrInvalidWorkerCount
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	if len(jobs) == 0 {
		return nil
	}
	if workerCount > len(jobs) {
		workerCount = len(jobs)
	}

	workerCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	jobCh := make(chan int)
	var firstErr error
	var firstErrorOnce sync.Once

	var workers sync.WaitGroup
	workers.Add(workerCount)
	for range workerCount {
		go func() {
			defer workers.Done()
			for {
				select {
				case <-workerCtx.Done():
					return
				case job, ok := <-jobCh:
					if !ok {
						return
					}
					if workerCtx.Err() != nil {
						return
					}
					if err := process(workerCtx, job); err != nil {
						firstErrorOnce.Do(func() {
							firstErr = err
							cancel()
						})
						return
					}
				}
			}
		}()
	}

	sendJobs(workerCtx, jobCh, jobs)
	workers.Wait()

	if firstErr != nil {
		return firstErr
	}
	return ctx.Err()
}

// sendJobs is the only function that sends to or closes jobCh.
func sendJobs(ctx context.Context, jobCh chan<- int, jobs []int) {
	defer close(jobCh)

	for _, job := range jobs {
		select {
		case <-ctx.Done():
			return
		case jobCh <- job:
		}
	}
}

func processJob(ctx context.Context, job int) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(50 * time.Millisecond):
	}

	if job == 5 {
		return fmt.Errorf("process job %d: rejected", job)
	}

	fmt.Println("processed", job)
	return nil
}

func main() {
	jobs := []int{1, 2, 3, 4, 5, 6}
	err := ProcessJobs(context.Background(), jobs, 3, processJob)

	fmt.Println("result:", err)
}
