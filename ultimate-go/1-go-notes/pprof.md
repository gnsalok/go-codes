## Go Memory Profiling with pprof

Go's `pprof` package provides powerful tools for profiling CPU and memory usage in Go applications. It allows developers to identify performance bottlenecks and memory leaks by generating detailed reports.


### Learn by Example: A Memory Leak Scenario

This scenario is a classic example of a **memory leak via unbounded growth**. In Go, maps don't shrink automatically even if you delete keys, but more importantly, if you never delete keys, the map will grow until the OS kills the process (OOM).

Here is a simplified Go program that mimics this exact leak and how to use `pprof` to catch it.

## 1. The "Leaky" Go Program

This program starts an HTTP server with a `/process` endpoint that simulates adding request metadata to a global map without ever cleaning it up.

```go
package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof" // Standard pprof endpoint
	"sync"
	"time"
)

var (
	// The Culprit: A global map with no eviction logic
	cache = make(map[string]string)
	mu    sync.Mutex
)

func leakyHandler(w http.ResponseWriter, r *http.Request) {
	requestID := fmt.Sprintf("%d", time.Now().UnixNano())
	
	mu.Lock()
	// Mimicking metadata storage - adding ~1KB per request
	cache[requestID] = "Some heavy metadata string..." + string(make([]byte, 1024))
	mu.Unlock()

	fmt.Fprintf(w, "Processed request: %s", requestID)
}

func main() {
	// Register our leaky handler
	http.HandleFunc("/process", leakyHandler)

	fmt.Println("Server starting on :8080...")
	fmt.Println("Access pprof at http://localhost:8080/debug/pprof/")
	
	// Start the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

```

---

## 2. How pprof helps you solve this

`pprof` provides a "living" look into the heap. In your scenario, here is how you would have used it:

### Step A: Capturing the Heap

While the service is running and memory is climbing, you would run the following command to grab a snapshot:
`go tool pprof http://localhost:8080/debug/pprof/heap`

### Step B: Identifying the Source

Once inside the pprof interactive shell, you use the `top` and `list` commands:

1. **`top`**: Shows the functions consuming the most memory. In this case, `leakyHandler` would appear at the top because it's responsible for the map allocations.
2. **`list leakyHandler`**: This is the "aha!" moment. It shows you exactly which line of code is holding onto memory. You would see the `cache[requestID] = ...` line highlighted as the primary allocator.

### Step C: Visualizing the Growth

If the `top` command isn't clear enough, you can use the `-http` flag to see a **Graphviz** visualization:
`go tool pprof -http=:8081 http://localhost:8080/debug/pprof/heap`

This generates a directed graph where the size of the boxes represents the amount of memory held. You would see a giant box pointing toward your `map` operations, confirming that the map is the "sink" where memory goes in but never comes out.

---

## 3. The Fix: TTL and Eviction

As mentioned in your scenario, the fix involves a mechanism to remove stale data. Below is the conceptual fix using a background "janitor" goroutine:

```go
type Item struct {
	Value      string
	Expiration int64
}

func janitor() {
	for {
		time.Sleep(1 * time.Minute)
		now := time.Now().UnixNano()
		mu.Lock()
		for k, v := range cache {
			if now > v.Expiration {
				delete(cache, k) // Explicitly remove stale keys
			}
		}
		mu.Unlock()
	}
}

```

### Why it works?:

* **Boundaries:** By setting a TTL, you ensure that memory usage is a function of **requests per minute**, rather than **total requests since startup**.
* **Observability:** After implementing this, running `pprof` again would show a "sawtooth" memory pattern (rising during traffic, dropping after janitor runs) rather than a never-ending upward slope.