# Lesson 06: Production Concurrency

## Learning goals

- Give every goroutine an owner and a finite stop path.
- Distinguish communication, synchronization, cancellation, and resource bounds.
- Choose among channels, mutexes, atomics, `WaitGroup`, and `errgroup`.
- Detect leaks, races, deadlocks, swallowed errors, and unbounded concurrency.

## Mental model

Concurrency is a lifecycle design problem. Before starting a goroutine, answer:

1. Who owns it?
2. What input keeps it alive?
3. How does it stop on success, error, timeout, and caller cancellation?
4. How does it report results and errors?
5. What bounds goroutines, queued work, memory, sockets, and downstream pressure?

An unbuffered channel transfers a value and synchronizes sender with receiver. A
buffered channel decouples them only up to its capacity; that capacity is a policy,
not a deadlock fix. The sender that knows no more values will be produced owns
closing the channel.

## Decision guide

| Requirement | Prefer |
| --- | --- |
| Wait for independent goroutines | `sync.WaitGroup` |
| Cancel sibling work after an error | `errgroup.WithContext` |
| Transfer ownership or stream values | Channel |
| Protect compound shared state | `sync.Mutex` or `sync.RWMutex` after measurement |
| Update one numeric value | Typed atomic operation |
| Propagate request cancellation | `context.Context` |
| Bound parallel work | Fixed worker pool or semaphore |

## Reading and running order

1. Read `concurrency/concurrency-guide/0-concurrency-guide.md`.
2. Run the focused channel examples:

```bash
go run ./lessons/06_concurrency/concurrency/channels/signaling
go run ./lessons/06_concurrency/concurrency/channels/fan_out
go run ./lessons/06_concurrency/concurrency/channels/worker_pool
go run ./lessons/06_concurrency/concurrency/channels/drop
go run ./lessons/06_concurrency/concurrency/channels/cancellation
go run ./lessons/06_concurrency/concurrency/channels/retry
```

3. Compare unbuffered, buffered, and bounded designs:

```bash
go run ./lessons/06_concurrency/concurrency/concurrency-guide/unbuffered_synchronization
go run ./lessons/06_concurrency/concurrency/concurrency-guide/buffered_decoupling
go run ./lessons/06_concurrency/concurrency/concurrency-guide/worker_pool
go run ./lessons/06_concurrency/concurrency/concurrency-guide/worker_pool_with_context
```

4. Study `concurrency/production/bounded_fetch.go`, which combines bounded work,
   ordered partial results, cancellation, and error joining.
5. Use [EXERCISES.md](EXERCISES.md) to change one constraint at a time.

## Focused verification

Every directory now contains one package instead of unrelated `main` functions:

```bash
go test ./lessons/06_concurrency/concurrency/channels/...
go test ./lessons/06_concurrency/concurrency/concurrency-guide/...
go test ./lessons/06_concurrency/concurrency/context/...
go test ./lessons/06_concurrency/concurrency/mutex/...
go test -race ./lessons/06_concurrency/concurrency/mutex/...
```

The HTTP request demonstrations make real network calls and are not part of the
deterministic test set.

## Production tradeoffs and common mistakes

- A goroutine blocked on an abandoned send or receive is a leak even if it consumes
  little CPU.
- Launching one goroutine per unbounded input shifts overload into memory, sockets,
  or a downstream service.
- Closing a channel broadcasts “no more values”; it is not a general request for
  one producer to stop.
- A larger buffer changes latency and overload policy. It does not repair unclear
  ownership.
- `RWMutex` can be slower than `Mutex`; choose it from contention measurements.
- Never copy a mutex, atomic value, or type that contains one after first use.
- Avoid `time.Sleep` as synchronization in tests. Wait for a state transition.

## Interview prompts

1. How would you prove a worker pool cannot leak after its caller cancels?
2. Who should close a results channel with multiple worker senders?
3. What does a channel buffer size mean operationally?
4. When would a mutex be clearer than a channel?
5. How would you preserve input ordering while processing concurrently?

## Confidence checklist

- [ ] Every goroutine has a named owner and stop path.
- [ ] Work and queues are bounded by an explicit constraint.
- [ ] Blocking sends and receives observe caller cancellation where necessary.
- [ ] One sender owns each channel close.
- [ ] Errors and partial results have a documented policy.
- [ ] Concurrent tests pass with the race detector.
