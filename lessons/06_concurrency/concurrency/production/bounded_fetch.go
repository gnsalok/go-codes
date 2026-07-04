package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const maxBodyBytes = 1 << 20 // 1 MiB keeps the example safe for unknown responses.

// FetchResult contains the result for one input URL. Index lets callers restore
// the original input order even when workers finish out of order.
type FetchResult struct {
	Index      int
	URL        string
	StatusCode int
	Bytes      int
	Err        error
}

type fetchJob struct {
	index int
	url   string
}

// FetchURLs fetches URLs with a bounded number of workers.
//
// Production choices demonstrated here:
//   - The caller owns cancellation through ctx.
//   - maxConcurrent bounds memory, sockets, and external-service pressure.
//   - The job producer is the only goroutine that closes jobs.
//   - The result channel is closed only after every worker has stopped sending.
//   - Results are written to unique indexes, so no mutex is required.
//   - Errors are collected and joined; this example returns partial results.
func FetchURLs(ctx context.Context, urls []string, maxConcurrent int) ([]FetchResult, error) {
	if maxConcurrent <= 0 {
		return nil, fmt.Errorf("maxConcurrent must be positive: %d", maxConcurrent)
	}
	if len(urls) == 0 {
		return []FetchResult{}, nil
	}
	if maxConcurrent > len(urls) {
		maxConcurrent = len(urls)
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	jobs := make(chan fetchJob)
	results := make(chan FetchResult)

	// Producer owns closing jobs. It also listens to ctx so it does not block
	// forever if workers stop early.
	go func() {
		defer close(jobs)
		for i, url := range urls {
			select {
			case jobs <- fetchJob{index: i, url: url}:
			case <-ctx.Done():
				return
			}
		}
	}()

	done := make(chan struct{})
	for workerID := 1; workerID <= maxConcurrent; workerID++ {
		go func(id int) {
			defer func() { done <- struct{}{} }()
			for {
				select {
				case <-ctx.Done():
					return
				case job, ok := <-jobs:
					if !ok {
						return
					}

					result := fetchOne(ctx, job.index, job.url)
					select {
					case results <- result:
					case <-ctx.Done():
						return
					}
				}
			}
		}(workerID)
	}

	// Close results after all workers report that they are done sending.
	go func() {
		for i := 0; i < maxConcurrent; i++ {
			<-done
		}
		close(results)
	}()

	out := make([]FetchResult, len(urls))
	var errs []error
	for result := range results {
		out[result.Index] = result
		if result.Err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", result.URL, result.Err))
		}
	}

	if err := ctx.Err(); err != nil && !errors.Is(err, context.Canceled) {
		errs = append(errs, err)
	}
	return out, errors.Join(errs...)
}

func fetchOne(ctx context.Context, index int, url string) FetchResult {
	result := FetchResult{Index: index, URL: url}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		result.Err = err
		return result
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		result.Err = err
		return result
	}
	defer resp.Body.Close()

	result.StatusCode = resp.StatusCode
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		result.Err = fmt.Errorf("unexpected status: %s", resp.Status)
		return result
	}

	body, err := io.ReadAll(io.LimitReader(resp.Body, maxBodyBytes))
	if err != nil {
		result.Err = err
		return result
	}
	result.Bytes = len(body)
	return result
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	urls := []string{
		"https://example.com",
		"https://www.iana.org/domains/reserved",
		"https://www.rfc-editor.org/rfc/rfc9110.txt",
	}

	results, err := FetchURLs(ctx, urls, 2)
	if err != nil {
		fmt.Println("completed with error:", err)
	}

	for _, result := range results {
		if result.Err != nil {
			fmt.Printf("[%d] %s failed: %v\n", result.Index, result.URL, result.Err)
			continue
		}
		fmt.Printf("[%d] %s status=%d bytes=%d\n", result.Index, result.URL, result.StatusCode, result.Bytes)
	}
}
