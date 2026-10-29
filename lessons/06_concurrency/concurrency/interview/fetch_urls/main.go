package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
)

const responseLimit = 1 << 20

type Resp struct {
	Index int
	URL   string
	Data  string
	Err   error
}

func call(ctx context.Context, index int, url string) Resp {
	resp := Resp{Index: index, URL: url}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		resp.Err = err
		return resp
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		resp.Err = err
		return resp
	}
	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		resp.Err = fmt.Errorf("unexpected status: %s", res.Status)
		return resp
	}

	b, err := io.ReadAll(io.LimitReader(res.Body, responseLimit))
	if err != nil {
		resp.Err = err
		return resp
	}
	resp.Data = string(b)
	return resp
}

// FetchAll writes each response to a distinct index. No mutex is needed because
// each goroutine owns one slot in the output slice.
func FetchAll(ctx context.Context, urls []string) ([]Resp, error) {
	out := make([]Resp, len(urls))

	var wg sync.WaitGroup
	wg.Add(len(urls))

	for i, u := range urls {
		i, u := i, u
		go func() {
			defer wg.Done()
			out[i] = call(ctx, i, u)
		}()
	}

	wg.Wait()
	return out, joinResponseErrors(out)
}

// FetchAllFanIn streams results into a channel as workers finish. Collection
// starts immediately, so workers do not depend on a buffer large enough to hold
// every response.
func FetchAllFanIn(ctx context.Context, urls []string) ([]Resp, error) {
	ch := make(chan Resp)
	var wg sync.WaitGroup

	for i, u := range urls {
		i, u := i, u
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp := call(ctx, i, u)
			select {
			case ch <- resp:
			case <-ctx.Done():
			}
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	out := make([]Resp, len(urls))
	for r := range ch {
		out[r.Index] = r
	}
	return out, joinResponseErrors(out)
}

func joinResponseErrors(out []Resp) error {
	var errs []error
	for _, r := range out {
		if r.Err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", r.URL, r.Err))
		}
	}
	return errors.Join(errs...)
}

func main() {
	ctx := context.Background()
	urls := []string{"https://example.com", "https://www.iana.org/domains/reserved", "https://www.rfc-editor.org/rfc/rfc9110.txt"}

	ordered, err := FetchAll(ctx, urls)
	fmt.Println("ordered responses:", len(ordered), "error:", err)

	fanIn, err := FetchAllFanIn(ctx, urls)
	fmt.Println("fan-in responses:", len(fanIn), "error:", err)
}
