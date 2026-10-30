package main

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
)

func TestProcessJobsRejectsInvalidWorkerCount(t *testing.T) {
	err := ProcessJobs(context.Background(), []int{1}, 0, func(context.Context, int) error {
		return nil
	})

	if !errors.Is(err, ErrInvalidWorkerCount) {
		t.Fatalf("ProcessJobs() error = %v, want %v", err, ErrInvalidWorkerCount)
	}
}

func TestProcessJobsBoundsConcurrencyAndProcessesEveryJob(t *testing.T) {
	const workerCount = 3
	jobs := []int{1, 2, 3, 4, 5, 6, 7, 8}
	release := make(chan struct{})
	started := make(chan struct{}, len(jobs))

	var active atomic.Int32
	var maximum atomic.Int32
	var completed atomic.Int32

	done := make(chan error, 1)
	go func() {
		done <- ProcessJobs(context.Background(), jobs, workerCount, func(context.Context, int) error {
			current := active.Add(1)
			defer active.Add(-1)
			for {
				observed := maximum.Load()
				if current <= observed || maximum.CompareAndSwap(observed, current) {
					break
				}
			}
			started <- struct{}{}
			<-release
			completed.Add(1)
			return nil
		})
	}()

	for range workerCount {
		<-started
	}
	if got := active.Load(); got != workerCount {
		t.Fatalf("active jobs = %d, want %d", got, workerCount)
	}
	close(release)

	if err := <-done; err != nil {
		t.Fatalf("ProcessJobs() error = %v, want nil", err)
	}
	if got := maximum.Load(); got != workerCount {
		t.Errorf("maximum concurrency = %d, want %d", got, workerCount)
	}
	if got := completed.Load(); got != int32(len(jobs)) {
		t.Errorf("completed jobs = %d, want %d", got, len(jobs))
	}
}

func TestProcessJobsReturnsFirstProcessingErrorAndStopsDispatch(t *testing.T) {
	wantErr := errors.New("job failed")
	var started atomic.Int32

	err := ProcessJobs(context.Background(), []int{1, 2, 3}, 1, func(context.Context, int) error {
		started.Add(1)
		return wantErr
	})

	if !errors.Is(err, wantErr) {
		t.Fatalf("ProcessJobs() error = %v, want %v", err, wantErr)
	}
	if got := started.Load(); got != 1 {
		t.Fatalf("started jobs = %d, want 1", got)
	}
}

func TestProcessJobsProcessingErrorCancelsAlreadyStartedJobs(t *testing.T) {
	wantErr := errors.New("job failed")
	started := make(chan struct{}, 2)
	fail := make(chan struct{})

	done := make(chan error, 1)
	go func() {
		done <- ProcessJobs(context.Background(), []int{1, 2}, 2, func(ctx context.Context, job int) error {
			started <- struct{}{}
			if job == 1 {
				<-fail
				return wantErr
			}
			<-ctx.Done()
			return ctx.Err()
		})
	}()

	<-started
	<-started
	close(fail)

	if err := <-done; !errors.Is(err, wantErr) {
		t.Fatalf("ProcessJobs() error = %v, want %v", err, wantErr)
	}
}

func TestProcessJobsCancellationStopsDispatchAndWorkers(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	const workerCount = 2
	started := make(chan struct{}, workerCount)
	var count atomic.Int32

	done := make(chan error, 1)
	go func() {
		done <- ProcessJobs(ctx, []int{1, 2, 3, 4}, workerCount, func(ctx context.Context, _ int) error {
			count.Add(1)
			started <- struct{}{}
			<-ctx.Done()
			return ctx.Err()
		})
	}()

	for range workerCount {
		<-started
	}
	cancel()

	if err := <-done; !errors.Is(err, context.Canceled) {
		t.Fatalf("ProcessJobs() error = %v, want %v", err, context.Canceled)
	}
	if got := count.Load(); got != workerCount {
		t.Fatalf("started jobs = %d, want %d", got, workerCount)
	}
}
