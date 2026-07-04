# Future Curriculum Roadmap

This repository is now organized as a sequential Go learning path. The next phase should add advanced core topics that help students move from syntax fluency to production-quality Go engineering.

## 1. Package design and module boundaries

- Package naming, exported APIs, and internal packages.
- Avoiding circular dependencies.
- Designing small packages around behavior rather than technical layers.
- Versioning modules and managing replace directives responsibly.

## 2. Context propagation in services

- Request-scoped cancellation and deadlines.
- Avoiding context misuse for optional parameters.
- Propagating trace IDs and request metadata safely.
- Testing timeout and cancellation behavior deterministically.

## 3. Advanced error design

- Creating stable domain error contracts.
- Separating user-facing messages from operational diagnostics.
- Joining multiple errors and reporting partial failures.
- Logging errors without losing wrapping context.

## 4. Production-grade HTTP services

- Routing, middleware, validation, and graceful shutdown.
- Structured JSON request/response handling.
- Health checks, readiness checks, and dependency checks.
- Integration testing with `httptest` and ephemeral dependencies.

## 5. Persistence and transactions

- `database/sql` fundamentals.
- Transaction boundaries and rollback patterns.
- Repository interfaces only where they clarify seams.
- Test strategies for persistence code.

## 6. Observability

- Structured logging.
- Metrics and service-level indicators.
- Distributed tracing concepts.
- Profiling CPU, heap, goroutine, and mutex contention in realistic programs.

## 7. Memory model and runtime behavior

- Escape analysis and allocation reduction.
- Garbage collector behavior and tuning tradeoffs.
- Data races, happens-before relationships, and the Go memory model.
- Scheduler behavior and goroutine lifecycle costs.

## 8. CLI application design

- Flags, environment configuration, and command structure.
- Exit codes and user-friendly terminal output.
- Testing command behavior without shelling out unnecessarily.

## 9. Generics beyond the basics

- Constraint design.
- Generic data structures.
- Comparing generics, interfaces, and code generation.
- Avoiding over-generalized APIs.

## 10. Capstone projects

Add small end-to-end projects that combine multiple lessons:

1. A concurrent URL checker with cancellation, tests, and benchmarks.
2. A JSON HTTP API with middleware, persistence, and graceful shutdown.
3. A CLI log analyzer using streaming I/O and profiling.
4. A typed in-memory cache demonstrating generics, synchronization, and observability hooks.
