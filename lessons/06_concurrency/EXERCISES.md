# Lesson 06 Exercises: Production Goroutines and Channels

Use these exercises after reading the examples. The goal is not only to make code run; the goal is to make the lifecycle, cancellation path, and error behavior obvious.

## Exercise 1: Concurrent tasks with ordered results

**Goal:** Run several independent tasks concurrently and return results in the same order as the input.

**Constraints:**

- Do not append to a shared slice from multiple goroutines.
- Either write each result to a unique index or send indexed results through a channel.
- Return an aggregated error if any task fails.

**Production skill:** Avoiding data races while preserving caller expectations.

## Exercise 2: Add a timeout

**Goal:** Add `context.WithTimeout` to Exercise 1.

**Constraints:**

- Each task must receive the context.
- A task must stop early when `ctx.Done()` is closed.
- The caller must not wait forever after timeout.

**Production skill:** Preventing goroutine leaks after a caller gives up.

## Exercise 3: Limit concurrency

**Goal:** Process 1,000 jobs with at most 5 workers.

**Constraints:**

- Do not start 1,000 goroutines.
- Use a worker pool or semaphore.
- Close the jobs channel from the producer.
- Close the results channel only after all workers finish.

**Production skill:** Protecting memory, CPU, database connections, and external APIs.

## Exercise 4: Fail fast

**Goal:** Stop all workers when the first job fails.

**Constraints:**

- Use `errgroup.WithContext` or equivalent cancellation logic.
- Stop feeding new jobs after cancellation.
- Workers must exit when context is cancelled.

**Production skill:** Returning first error while avoiding wasted work.

## Exercise 5: Stream partial results

**Goal:** Return results as soon as they complete.

**Constraints:**

- Results should be received through a channel.
- The caller should be able to stop early through context cancellation.
- Senders must not block forever if the caller stops reading.

**Production skill:** Building responsive streaming APIs.

## Exercise 6: Find and fix the leak

**Broken pattern:**

```go
func slow(ctx context.Context) <-chan string {
	ch := make(chan string)
	go func() {
		time.Sleep(time.Second)
		ch <- "done"
	}()
	return ch
}
```

**Goal:** Fix the goroutine leak when the caller times out before the send.

**Hint:** Use a `select` with `ctx.Done()` around the send, or use a buffered channel only when the result count is known and bounded.

**Production skill:** Understanding that cancellation must be observed by the goroutine that can block.
