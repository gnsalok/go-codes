# Bounded job processor interview exercise

Implement a function that processes jobs with at most `n` concurrent workers.
It must:

- stop dispatching new jobs promptly when the context is cancelled;
- pass cancellation to jobs that have already started;
- return the first processing error observed;
- wait for every worker it starts;
- never send on a closed channel; and
- make no promise about completion order.

Run the example and its tests:

```bash
go run ./lessons/06_concurrency/concurrency/interview/bounded_job_processor
go test -race ./lessons/06_concurrency/concurrency/interview/bounded_job_processor
```

The executable keeps the call site deliberately small:

```go
func main() {
	jobs := []int{1, 2, 3, 4, 5, 6}
	err := ProcessJobs(context.Background(), jobs, 3, processJob)
	fmt.Println("result:", err)
}
```

Read it aloud as: “Process these integer jobs with at most three workers, using
`processJob` for each job.” The processing logic is a separate named function so
the concurrency setup and the business operation can be understood separately.

## The mental model: bounds, ownership, exits, errors

Do not try to picture every interleaving. Write four small decisions first:

1. **Bounds:** start exactly `min(n, len(jobs))` workers. An unbuffered jobs
   channel also avoids building a hidden work queue.
2. **Ownership:** decide who may send, receive, close, and cancel before writing a
   goroutine.
3. **Exits:** give both the dispatcher and workers a context-cancellation path.
4. **Errors:** use `sync.Once` so only the first error is recorded, then cancel
   the shared worker context.

| Resource | Owner | Allowed operations |
| --- | --- | --- |
| `jobCh` | dispatcher | send, then close exactly once |
| `jobCh` | workers | receive only |
| `firstErr` | workers | `sync.Once` permits exactly one write |
| derived context | `ProcessJobs` | any first-error winner may call its safe, idempotent cancel function |

This ownership table is the proof that there is no send on a closed channel:
the only goroutine that sends jobs is also the only goroutine that closes the
jobs channel, and it closes only after its send loop ends.

## A writing order that works at a whiteboard

Build the answer in layers and explain one invariant per layer:

1. Write the function signature and reject `n <= 0`.
2. Derive a cancellable context, create `jobCh`, and reserve a `firstErr` variable.
3. Write one worker loop: select cancellation or receive; process; report an
   error through `sync.Once`; return.
4. Put that worker loop behind a `WaitGroup` and start at most `n` copies.
5. Write the dispatch loop with the same cancellation select.
6. Close `jobCh`, wait for workers, then prefer a processing error over the
   caller's context error.

At each step, say the local invariant aloud. For example: “Workers do not close
channels,” “every worker defers `Done`,” and “after `Wait`, no worker can still
touch local state.” This is much easier than reasoning about the whole program
at once.

`ProcessJobs` cannot forcibly stop an arbitrary Go function. Prompt cancellation
requires `process` to select on `ctx.Done()` around blocking work. Waiting for all
workers prevents this function from leaking goroutines, but a processor that
ignores cancellation can make the call wait forever. State that contract in an
interview; it is an important limitation, not an implementation failure.

## If already-started jobs must drain after cancellation

Separate “stop dispatching” from “cancel work”:

- use the caller's context only in the dispatcher, so cancellation stops new
  jobs;
- call already-started processors with a non-cancelled work context, such as
  `context.WithoutCancel(ctx)`, if preserving context values is useful;
- still close `jobCh` and wait for all workers; and
- return the first processing error, otherwise return the caller's context error
  after the drain completes.

The same separation can apply to processing errors: record the first error and
stop dispatching, but do not cancel processors already running. In production,
give draining its own grace-period context; otherwise one stuck job can prevent
shutdown forever. This changes the contract from **fail fast** to **stop starting,
then drain**, so cancellation latency is now bounded by the slowest started job
or the grace period.

