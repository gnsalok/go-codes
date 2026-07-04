# Lesson 17: Memory Model and Runtime Behavior

## Core Go concepts

- Reading escape analysis output and reducing unnecessary allocations.
- Understanding garbage collector behavior and tuning tradeoffs.
- Reasoning about data races and happens-before relationships.
- Knowing how the scheduler affects goroutine execution.
- Estimating goroutine lifecycle costs and avoiding leaks.

## Quick reading

The Go runtime handles scheduling, stacks, allocation, and garbage collection, but production engineers still need to understand the tradeoffs. Escape analysis explains why some values move from stack to heap. The garbage collector favors low pause times, but allocation-heavy code still creates work. The memory model explains when writes in one goroutine are guaranteed to be visible to another.

Use the race detector to catch unsynchronized access, but do not use it as a substitute for design. Establish clear ownership, synchronization, or channel communication. For performance investigations, measure before and after changes with benchmarks and profiles.

Pair this lesson with `lessons/06_concurrency` and `lessons/09_performance`.

## Suggested exercises

1. Run escape analysis on a small function and explain why a value escapes.
2. Benchmark two implementations with different allocation behavior.
3. Write a race, confirm `go test -race` detects it, then fix it.
4. Demonstrate a goroutine leak and add cancellation to stop it.
5. Compare mutex synchronization with channel ownership for shared data.

## Run the lesson

The `counter.go` example demonstrates mutex synchronization for concurrent writes. From the repository root:

```bash
go test ./lessons/17_runtime_memory_model/...
```

## Run tests

Run the example tests and package compilation check with:

```bash
go test ./lessons/17_runtime_memory_model/...
```

For race-focused examples, use:

```bash
go test -race ./lessons/17_runtime_memory_model/...
```
