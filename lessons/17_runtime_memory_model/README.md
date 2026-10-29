# Lesson 17: Runtime and Memory Model

## Learning goals and mental model

A data race is unsynchronized concurrent access where at least one access writes.
The memory model defines which synchronization events establish a happens-before
relationship; timing and “it usually runs first” do not. The scheduler, allocator,
and garbage collector affect cost but do not repair missing synchronization.

```bash
go test ./lessons/17_runtime_memory_model
go test -race ./lessons/17_runtime_memory_model
```

## Tradeoffs and common mistakes

- The race detector finds executed races, not every possible race.
- Mutexes and atomics must not be copied after first use.
- Escape analysis explains compiler placement decisions; it is not a direct measure
  of application performance.
- Reducing allocations can increase retention or complexity; profile first.
- Goroutine stacks grow, but goroutines still consume memory and scheduler work.

| Shared-state requirement | Starting point |
| --- | --- |
| One invariant across several fields | `sync.Mutex` |
| Read-heavy state with measured reader contention | `sync.RWMutex` |
| Independent keys with write-once or disjoint-key workloads | Consider `sync.Map` |
| One numeric state transition | Typed `sync/atomic` operation |

`sync.Map` is specialized, not a general replacement for `map` plus a mutex. Prefer
the ordinary map when its type safety and compound invariants make ownership clearer.

## Interview prompts and practice

1. Which channel and mutex operations establish happens-before relationships?
2. How do you distinguish a race, deadlock, and goroutine leak?
3. Write a failing race-detector example, fix it, and explain the synchronization.
