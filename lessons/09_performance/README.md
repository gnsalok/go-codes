# Lesson 09: Performance

## Core Go concepts

- Measuring behavior with benchmarks.
- Observing allocations and interface-related memory costs.
- Starting CPU/profile investigations with `pprof` examples.

## Quick reading

The `benchmark/` example demonstrates how interface choices can affect memory and benchmark results. The `pprof/` example gives students a small entry point for learning profiling workflows. Pair this lesson with the notes under `lessons/notes/pprof.md` for a deeper reading path.

## Run the lesson

From the repository root:

```bash
go run ./lessons/09_performance/pprof/main.go
```

## Run benchmarks and tests

```bash
go test -bench=. ./lessons/09_performance/...
```
