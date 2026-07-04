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
