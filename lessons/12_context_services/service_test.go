package contextservices

import (
	"context"
	"errors"
	"testing"
)

func TestFetchReturnsContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := Fetch(ctx, make(chan string))
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("Fetch() error = %v; want context.Canceled", err)
	}
}

func TestFetchReturnsResult(t *testing.T) {
	result := make(chan string, 1)
	result <- "ready"

	got, err := Fetch(context.Background(), result)
	if err != nil || got != "ready" {
		t.Fatalf("Fetch() = %q, %v; want %q, nil", got, err, "ready")
	}
}
