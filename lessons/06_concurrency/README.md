# Lesson 06: Concurrency

Go makes concurrency easy to start, but production concurrency is about **ownership, cancellation, backpressure, and correctness**. This lesson teaches goroutines and channels as tools for building bounded, observable, and stoppable systems.

## What you should be able to do after this lesson

After completing the lesson, you should be able to:

- Start goroutines only when you know who owns them and how they stop.
- Use channels to pass data, stream results, coordinate shutdown, and apply backpressure.
- Use `sync.WaitGroup` to wait for completion without confusing it with communication or error handling.
- Use `context.Context` to propagate cancellation and deadlines across goroutines and API boundaries.
- Use worker pools and semaphores to avoid unbounded goroutine creation.
- Choose between channels, mutexes, atomics, `WaitGroup`, and `errgroup` based on the problem.
- Review concurrent code for leaks, races, deadlocks, swallowed errors, and unclear channel ownership.

## Production decision checklist

Ask these questions before adding a goroutine:

1. **Who owns this goroutine?** The caller, a service, a worker pool, or the process?
2. **How does it stop?** Channel close, context cancellation, finite input, or process shutdown?
3. **How are errors reported?** Result channel, `errgroup`, callback, log, or returned aggregate error?
4. **What shared state does it touch?** If state is shared, protect it with ownership transfer, a mutex, or atomic operations.
5. **What bounds the work?** Fixed worker count, semaphore, queue size, timeout, or rate limit?

| Need | Prefer |
| --- | --- |
| Wait until goroutines finish | `sync.WaitGroup` |
| Return first error and cancel sibling work | `errgroup.WithContext` |
| Pass data or stream results | Channel |
| Protect complex shared state | `sync.Mutex` |
| Count simple numeric events | `sync/atomic` |
| Cancel request-scoped work | `context.Context` |
| Limit concurrency | Worker pool or buffered-channel semaphore |

## Production rules

- Every goroutine must have a clear owner and a clear stop path.
- Do not launch unbounded goroutines for unbounded input.
- Prefer `context.Context` for cancellation across function and service boundaries.
- Use channels for communication and coordination, not as a replacement for every shared variable.
- Close a channel from the sending side, and only when no more values will be sent.
- Do not close a channel just to stop one sender; use context or another explicit signal.
- Do not use a bigger buffer to hide a deadlock. Fix the ownership and lifecycle problem.
- Always decide what happens on success, error, timeout, cancellation, and partial completion.

## Recommended reading order

1. `concurrency/production/bounded_fetch.go` — a production-style example that combines context, bounded workers, channels, errors, and ordered results.
2. `concurrency/concurrency-guide/0-concurrency-guide.md` — core mental model for buffered and unbuffered channels.
3. `concurrency/concurrency-guide/unbuffered-synchronization.go` — rendezvous and hand-off semantics.
4. `concurrency/concurrency-guide/buffered-decoupling.go` — buffering and backpressure.
5. `concurrency/concurrency-guide/worker-pool-pattern.go` — fan-out/fan-in with explicit worker lifecycle.
6. `concurrency/concurrency-guide/worker-pool-with-context.go` — cancellation and first-error propagation.
7. `concurrency/context/` — timeouts, request cancellation, and context values.
8. `concurrency/mutex/` — protecting shared memory with mutexes and atomics.
9. `concurrency/Interview/` — practice patterns and discussion examples.
10. `../notes/concurrency.md` and `../notes/channels.md` — production notes and reusable templates.

## Run the lesson

From the repository root:

```bash
go run ./lessons/06_concurrency/concurrency/production/bounded_fetch.go
```

You can also run individual standalone examples:

```bash
go run ./lessons/06_concurrency/concurrency/concurrency-guide/worker-pool-pattern.go
```

For directories with multiple standalone `main` functions, run the specific file you want to study.

## Run tests and checks

This lesson currently has standalone `main` programs rather than package-level tests. Check specific examples with `go run`, and check package compilation where applicable:

```bash
go test ./lessons/06_concurrency/...
```

## Production confidence checklist

Before shipping concurrent Go code, verify:

- [ ] Every goroutine has an owner.
- [ ] Every goroutine exits on success, error, cancellation, or closed input.
- [ ] Blocking sends and receives cannot leak after the caller gives up.
- [ ] Channels have one clear closing owner.
- [ ] Worker counts and queue sizes are bounded by real constraints.
- [ ] Shared mutable state is protected or avoided.
- [ ] Errors are returned, joined, or intentionally handled.
- [ ] Context deadlines and cancellation are respected by I/O calls.
- [ ] The code is understandable without relying on `time.Sleep` for synchronization.
- [ ] The race detector is used when changing concurrent code.
