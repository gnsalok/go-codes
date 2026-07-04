package runtimememorymodel

import (
	"sync"
	"testing"
)

func TestCounterSynchronizesConcurrentWrites(t *testing.T) {
	var counter Counter
	var wg sync.WaitGroup
	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Add(1)
		}()
	}
	wg.Wait()

	if got := counter.Value(); got != 100 {
		t.Fatalf("Value() = %d; want 100", got)
	}
}
