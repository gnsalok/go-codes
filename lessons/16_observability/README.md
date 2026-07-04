# Lesson 16: Observability

## Core Go concepts

- Writing structured logs that help operators answer questions quickly.
- Defining metrics, service-level indicators, and service-level objectives.
- Understanding distributed tracing concepts and request correlation.
- Profiling CPU, heap, goroutine, and mutex contention.
- Connecting observability signals without overwhelming the codebase.

## Quick reading

Observability helps you understand a running system from the outside. Logs explain discrete events, metrics show trends and health, traces connect work across service boundaries, and profiles reveal resource costs inside a process. Good observability starts with clear questions: is the service available, is it fast enough, and what changed when it became unhealthy?

Prefer structured logs over free-form strings for operational events. Use metrics for counts, durations, queue lengths, and resource usage. Use traces to follow request paths through multiple components. Use profiling when you need evidence about CPU, memory, goroutines, blocking, or mutex contention.

Review `lessons/09_performance` and `lessons/notes/pprof.md` before this lesson.

## Suggested exercises

1. Replace ad-hoc log messages with structured key-value logging.
2. Define service-level indicators for a simple HTTP API.
3. Add request correlation fields to logs and handler responses.
4. Capture a CPU or heap profile from a small benchmark.
5. Explain when a metric is better than a log line.

## Run the lesson

The `observability.go` example records structured request fields and separates status classification from logging. From the repository root:

```bash
go test ./lessons/16_observability/...
```

## Run tests

Run the example tests and package compilation check with:

```bash
go test ./lessons/16_observability/...
```
