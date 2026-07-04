# Lesson 06: Concurrency

## Core Go concepts

- Launching goroutines and coordinating them with channels.
- Using `sync.WaitGroup`, mutexes, and atomics safely.
- Applying worker-pool and fan-out/fan-in patterns.
- Propagating cancellation and deadlines with `context`.
- Reasoning about synchronization, buffering, scarcity, and correctness.

## Quick reading

Start with `concurrency/concurrency-guide/0-concurrency-guide.md`, then move through the guide examples in the same directory. The `channels/` examples demonstrate orchestration, buffering, and channel-based coordination. The `context/` examples show timeout, value, and API-call patterns. The `mutex/` directory contrasts locking with atomic operations, while `Interview/` contains practice patterns such as worker pools and parallel URL fetching.

## Run the lesson

From the repository root:

```bash
go run ./lessons/06_concurrency/concurrency/concurrency-guide/worker-pool-pattern.go
```

For directories with multiple standalone `main` functions, run the specific file you want to study.

## Run tests

This lesson currently has no `_test.go` files, but package compilation can be checked with:

```bash
go test ./lessons/06_concurrency/...
```
