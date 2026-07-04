# Go Concurrency: Goroutines, Channels, and Production Confidence

Concurrency in Go is not just “add `go` before a function.” Production concurrency requires a lifecycle: start work intentionally, communicate clearly, bound scarce resources, report errors, and stop every goroutine.

## 1. Unbuffered channels: synchronization and hand-off

An unbuffered channel has no storage. A send and receive must meet at the same time. This is called a rendezvous.

```go
ch := make(chan Result)
```

Use unbuffered channels when:

- The sender needs to know a receiver accepted the value.
- Ownership of a value is handed from one goroutine to another.
- You want backpressure immediately when the receiver is not ready.

Analogy: a relay baton. The runner cannot drop the baton on a table and leave; another runner must be present to take it.

## 2. Buffered channels: queues, burst absorption, and semaphores

A buffered channel has a fixed-size queue. A send blocks only when the buffer is full. A receive blocks only when the buffer is empty.

```go
jobs := make(chan Job, 100)
sem := make(chan struct{}, 5)
```

Use buffered channels when:

- You intentionally want to decouple producer and consumer speed.
- You need a bounded queue for bursts.
- You need a semaphore to limit concurrent access to a scarce resource.

A buffer size is a design decision. Choose it from a real constraint such as worker count, expected burst size, memory budget, or external service limit. Do not add a buffer just to make a deadlock disappear.

## 3. Goroutine ownership

Every goroutine should have a clear owner. The owner is responsible for knowing:

- Why the goroutine exists.
- Which signal tells it to stop.
- Where its error goes.
- Which goroutine closes any channels it writes to.

If you cannot answer those questions, the goroutine is likely to leak under timeout, shutdown, or error conditions.

## 4. Combining goroutines and channels: worker pool

The most common production pattern is bounded fan-out/fan-in:

- **Jobs channel**: a finite stream of work.
- **Workers**: a fixed number of goroutines reading from the jobs channel.
- **Results channel**: workers send output or errors.
- **WaitGroup/errgroup**: waits for workers and coordinates shutdown.
- **Context**: cancels work when the caller gives up or one worker fails.

Channel closing ownership matters:

- The job producer closes `jobs` after sending the final job.
- Workers never close `jobs`; they only receive from it.
- The result channel is closed only after every worker has stopped sending.
- Receivers range over `results` until it is closed.

## 5. Common production failures

| Failure | Symptom | Fix |
| --- | --- | --- |
| Goroutine leak | A goroutine waits forever on send, receive, timer, or I/O | Add context cancellation and select around blocking channel operations |
| Deadlock | Runtime reports all goroutines are asleep | Check sender/receiver pairing and channel closing ownership |
| Send on closed channel | Panic | Only the sending owner closes the channel after all sends are done |
| Unbounded fan-out | Memory, CPU, or connection spikes | Use a fixed worker pool or semaphore |
| Data race | Incorrect or inconsistent shared state | Transfer ownership through channels, or use mutex/atomic |
| Swallowed error | Main flow succeeds while background work failed | Return errors through results or use `errgroup` |

## 6. Production checklist

Before using a goroutine/channel pattern, ask:

1. Is the work finite or cancellable?
2. Is concurrency bounded?
3. Is the channel closed by the sender?
4. Can a sender block forever if the receiver exits?
5. Can a receiver block forever if the sender exits?
6. Are errors propagated?
7. Is shared state protected or avoided?
