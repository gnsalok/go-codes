package mutexprofile

import (
	"sync"
	"testing"
	"time"
)

func TestMutexProfile(t *testing.T) {
	var (
		mu   sync.RWMutex
		data = make([]string, 1000)
		wg   sync.WaitGroup
	)

	const pairs = 100
	wg.Add(2 * pairs)

	for range pairs {
		go func() {
			defer wg.Done()
			for range 10 {
				mu.Lock()
				data = append(data, "A")
				time.Sleep(time.Millisecond)
				mu.Unlock()
			}
		}()

		go func() {
			defer wg.Done()
			for range 10 {
				mu.RLock()
				_ = len(data)
				time.Sleep(time.Millisecond)
				mu.RUnlock()
			}
		}()
	}

	wg.Wait()
	if got, want := len(data), 1000+pairs*10; got != want {
		t.Fatalf("len(data) = %d, want %d", got, want)
	}
}
