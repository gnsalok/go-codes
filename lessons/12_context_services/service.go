package contextservices

import (
	"context"
	"errors"
)

var ErrNoWork = errors.New("no work available")

// Fetch waits for either a service result or context cancellation.
func Fetch(ctx context.Context, result <-chan string) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case value, ok := <-result:
		if !ok {
			return "", ErrNoWork
		}
		return value, nil
	}
}
