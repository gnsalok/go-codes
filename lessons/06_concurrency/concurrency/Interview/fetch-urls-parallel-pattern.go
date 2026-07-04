package main

import (
	"context"
	"errors"
	"io"
	"net/http"
	"sync"
)

type Resp struct {
	Data string
	Err  error
}

// Pre-indexed results (no extra locking)
func call(ctx context.Context, url string) Resp {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Resp{Err: err}
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return Resp{Err: err}
	}
	return Resp{Data: string(b)}
}

func FetchAll(ctx context.Context, urls []string) ([]Resp, error) {
	if len(urls) != 3 { /* or generalize */
	}
	out := make([]Resp, len(urls))

	var wg sync.WaitGroup
	wg.Add(len(urls))

	for i, u := range urls {
		i, u := i, u // capture
		go func() {
			defer wg.Done()
			out[i] = call(ctx, u) // distinct index => no mutex needed
		}()
	}

	wg.Wait()

	// aggregate error with context
	var errs []error
	for _, r := range out {
		if r.Err != nil {
			errs = append(errs, r.Err)
		}
	}
	if len(errs) > 0 {
		return out, errors.Join(errs...)
	}
	return out, nil
}

// Fan-in channel (simple, order-agnostic)
func FetchAllFanIn(ctx context.Context, urls []string) ([]Resp, error) {
	ch := make(chan Resp, len(urls)) // buffered to avoid goroutine leaks
	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, u := range urls {
		u := u
		go func() {
			defer wg.Done()
			ch <- call(ctx, u)
		}()
	}

	wg.Wait()
	close(ch)

	var out []Resp
	var errs []error
	for r := range ch {
		out = append(out, r)
		if r.Err != nil {
			errs = append(errs, r.Err)
		}
	}
	if len(errs) > 0 {
		return out, errors.Join(errs...)
	}
	return out, nil
}

func main() {
	ctx := context.Background()
	urls := []string{"https://example.com/1", "https://example.com/2", "https://example.com/3"}

	// FetchAll
	res, err := FetchAll(ctx, urls)
	if err != nil {
		// handle error
	}
	// use res

	// FetchAllFanIn
	res, err = FetchAllFanIn(ctx, urls)
	if err != nil {
		// handle error
	}
	// use res
}
