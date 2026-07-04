# Advanced Curriculum Plan

This repository now includes an advanced Go learning path that moves students from syntax fluency toward production-quality Go engineering. The former future roadmap has been converted into numbered lessons under `lessons/`, with capstone projects intentionally left as the next major expansion.

## Implemented advanced lessons

1. `lessons/11_package_design` — package naming, exported APIs, internal packages, circular dependency avoidance, and module boundary discipline.
2. `lessons/12_context_services` — request-scoped cancellation, deadlines, safe request metadata, and deterministic context testing.
3. `lessons/13_advanced_errors` — stable domain error contracts, user-facing versus diagnostic messages, error joining, and logging wrapped errors.
4. `lessons/14_http_services` — JSON HTTP APIs, routing, middleware, validation, health checks, readiness checks, graceful shutdown, and `httptest` integration tests.
5. `lessons/15_persistence_transactions` — `database/sql`, context-aware persistence, transaction boundaries, rollback patterns, repository seams, and persistence test strategies.
6. `lessons/16_observability` — structured logging, metrics, service-level indicators, distributed tracing concepts, and realistic profiling workflows.
7. `lessons/17_runtime_memory_model` — escape analysis, allocation reduction, garbage collector behavior, data races, happens-before relationships, scheduler behavior, and goroutine lifecycle costs.
8. `lessons/18_cli_design` — flags, environment configuration, command structure, exit codes, terminal output, and testable command behavior.
9. `lessons/19_advanced_generics` — constraint design, generic data structures, generics versus interfaces, generics versus code generation, and avoiding over-generalized APIs.

## Remaining planned expansion

Capstone projects should be added later as a separate lesson or project area after the advanced topic lessons have enough examples for students to combine.

Suggested capstones:

1. A concurrent URL checker with cancellation, tests, and benchmarks.
2. A JSON HTTP API with middleware, persistence, and graceful shutdown.
3. A CLI log analyzer using streaming I/O and profiling.
4. A typed in-memory cache demonstrating generics, synchronization, and observability hooks.

## Recommended capstone structure

When capstones are added, each project should include:

- A problem statement.
- Required features and acceptance criteria.
- Suggested package layout.
- Testing expectations.
- Benchmark or profiling guidance where relevant.
- Optional extension tasks for advanced students.
