package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	attempts := 0
	err := retry(ctx, 10*time.Millisecond, func() error {
		attempts++
		if attempts < 3 {
			return errors.New("not ready")
		}
		return nil
	})
	fmt.Printf("attempts=%d error=%v\n", attempts, err)
}

func retry(ctx context.Context, interval time.Duration, check func() error) error {
	for {
		if err := check(); err == nil {
			return nil
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(interval):
		}
	}
}
