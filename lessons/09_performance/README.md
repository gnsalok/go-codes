# Lesson 09: Measurement, Benchmarks, and Profiling

## Learning goals

- Write benchmarks that isolate one decision and resist dead-code elimination.
- Interpret time, allocation, CPU, heap, goroutine, block, and mutex profiles.
- Use escape analysis as evidence, not as an optimization target by itself.
- Compare before and after results under representative input.

## Mental model

Performance work starts with a user-visible or operational problem, then narrows it
with measurements. Benchmarks compare controlled implementations; profiles show
where a running workload spends resources. Neither proves production impact unless
the workload represents the real system.

```bash
go test -bench=. -benchmem ./lessons/09_performance/benchmark
go test ./lessons/09_performance/...
go run ./lessons/09_performance/pprof
```

The pprof program intentionally retains per-request data in a global map. Generate
traffic against `/process`, capture multiple heap profiles, and compare retained
allocations over time. Do not expose pprof endpoints publicly without access controls.

## Production tradeoffs and common mistakes

- A microbenchmark can optimize code that is irrelevant to end-to-end latency.
- Moving an allocation from heap to stack is useful only when measurement shows
  allocation pressure matters.
- Benchmark setup belongs outside the timed region unless setup is the operation.
- Comparing one run is weak evidence; use repeated measurements and a statistical
  comparison tool for important decisions.
- A heap profile shows retained or allocated memory depending on the sample type;
  name which question you are answering.

## Interview prompts

1. How do you distinguish a memory leak from healthy cache growth?
2. Why can an interface appear allocation-free in one benchmark and allocate in
   another?
3. Which profile would you collect for lock contention, and what would it omit?

## Practice

Capture a baseline for the interface and concrete benchmarks, make one justified
change, and report both the result and the limitations of the experiment.
