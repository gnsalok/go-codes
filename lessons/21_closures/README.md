# Lesson 21: Closures and Captured State

## Learning goals

- Explain what a closure captures and how long captured state remains reachable.
- Distinguish independent closures from closures sharing one variable.
- Use closures for configuration, callbacks, and HTTP middleware.
- Identify capture-related loop, lifetime, allocation, and race hazards.

## Mental model

A closure is a function value that refers to variables from an enclosing scope. It
captures variables, not frozen snapshots of their values. The compiler decides
whether captured storage can remain on a stack or must escape; use escape analysis
when that allocation matters.

Each call to `NewCounter` creates a different `count` variable, so returned counters
are independent. A counter closure is not safe for concurrent use merely because
its state is private: multiple goroutines would still race on the captured variable.

```bash
go run ./lessons/21_closures
go test ./lessons/21_closures
```

## Middleware example

HTTP middleware commonly captures the next handler and configuration:

```go
func withHeader(name, value string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(name, value)
		next.ServeHTTP(w, r)
	})
}
```

The outer function runs while the handler chain is assembled. The returned closure
runs for every request and retains access to `name`, `value`, and `next`.

## Production tradeoffs and common mistakes

- A closure can keep a large object graph reachable longer than expected.
- Capturing a mutable loop or shared variable can produce incorrect results or a
  data race; reason about the variable's lifetime and synchronization.
- Hidden state is convenient but can make tests and ownership unclear. A named type
  is often better once behavior grows beyond one small operation.
- Do not compare function values for identity; functions are only comparable to nil.

## Interview prompts

1. What does a closure capture: a variable or its current value?
2. When does captured state escape to the heap?
3. How would you make `NewCounter` safe for concurrent callers?
4. Why are closures a natural fit for middleware construction?

## Practice

Extend `NewCounter` with a starting value and configurable delta. Then implement a
concurrency-safe version and prove it with a race-detector test.
