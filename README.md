# Go Senior Interview Curriculum

![Go Version](https://img.shields.io/badge/go-v1.27.3-blue.svg)
[![License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](LICENSE)

This repository is a practical curriculum for senior Go interviews. It connects
language mechanics to the design decisions expected from engineers who build and
operate production services: API boundaries, cancellation, concurrency,
correctness, testing, performance, persistence, and observability.

The material is deliberately mixed:

- **Runnable examples** demonstrate one behavior and can be executed directly.
- **Tested packages** turn important contracts into executable specifications.
- **Design progressions** compare several versions of the same idea.
- **Intentional compiler failures** preserve common mistakes so you can learn to
  read the diagnostic. Their lesson README names the expected failure and the
  corresponding corrected example.

Because compiler failures are part of the curriculum, `go test ./...` is not the
repository's verification command. Use the focused command in each lesson.

## Prerequisites

- Go 1.27.3. The version is pinned in `mise.toml`; with
  [mise](https://mise.jdx.dev/), run `mise install`.
- Familiarity with Go syntax and the standard `go` tool.
- A willingness to explain tradeoffs aloud. Senior interviews evaluate reasoning,
  not only whether the final program compiles.

## How to study a lesson

1. Read the learning goals and mental model.
2. Predict the example's output or compiler diagnostic before running it.
3. Run the focused command shown in the lesson README.
4. Answer the interview prompts without looking at the notes.
5. Change one constraint, write a test, and explain what the new design trades away.

For code that may run concurrently, also use the race detector. For performance
work, record a benchmark before changing the implementation.

## Curriculum

| Lesson | Topic | Senior-level focus |
| --- | --- | --- |
| [01](lessons/01_basics/README.md) | Language basics | Zero values, slices, maps, runes, and value semantics |
| [02](lessons/02_composition/README.md) | Composition | Behavior over inheritance and API evolution |
| [03](lessons/03_interfaces_and_decoupling/README.md) | Interfaces | Method sets, consumer-owned interfaces, and decoupling |
| [04](lessons/04_errors/README.md) | Errors | Wrapping, inspection, typed nils, panic, and recovery |
| [05](lessons/05_testing/README.md) | Testing | Table tests, fakes, HTTP tests, and benchmarks |
| [06](lessons/06_concurrency/README.md) | Concurrency | Ownership, cancellation, backpressure, and bounded work |
| [07](lessons/07_generics/README.md) | Generics | Constraints and when concrete code is clearer |
| [08](lessons/08_time/README.md) | Time | Instants, locations, parsing, and testable clocks |
| [09](lessons/09_performance/README.md) | Performance | Measurement, allocation analysis, and profiling |
| [10](lessons/10_go_version_evolution/README.md) | Go version evolution | Reading release notes and adopting new APIs deliberately |
| [11](lessons/11_package_design/README.md) | Package design | Cohesion, exported surface area, and dependency direction |
| [12](lessons/12_context_services/README.md) | Context | Cancellation across service boundaries |
| [13](lessons/13_advanced_errors/README.md) | Error contracts | Stable domain errors and partial failure |
| [14](lessons/14_http_services/README.md) | HTTP services | Validation, middleware, health, and graceful shutdown |
| [15](lessons/15_persistence_transactions/README.md) | Persistence | Transaction boundaries and repository seams |
| [16](lessons/16_observability/README.md) | Observability | Logs, metrics, traces, and operational questions |
| [17](lessons/17_runtime_memory_model/README.md) | Runtime and memory | Happens-before, races, scheduling, GC, and escape analysis |
| [18](lessons/18_cli_design/README.md) | CLI design | Testable boundaries, streams, and exit codes |
| [19](lessons/19_advanced_generics/README.md) | Advanced generics | Generic data structures and API tradeoffs |
| [20](lessons/20_file_io_buffering/README.md) | File I/O | Streaming, buffering, lifecycle, and atomic replacement |
| [21](lessons/21_closures/README.md) | Closures | Captured state, middleware, and concurrency hazards |

## Interview practice loop

For each topic, be ready to do four things:

1. State the invariant: what must always remain true?
2. Identify ownership: who creates, mutates, closes, cancels, or releases it?
3. Name the failure modes: invalid input, partial work, timeout, cancellation,
   resource exhaustion, race, deadlock, or leak.
4. Describe the evidence: which test, benchmark, profile, log, or metric would
   prove the implementation works?

## Capstone exercises

Use these after the individual lessons. Each capstone should include a problem
statement, acceptance tests, a short design note, benchmarks where relevant, and
an explanation of failure handling.

1. **Concurrent URL checker:** bounded workers, per-request deadlines, ordered
   results, cancellation, deterministic HTTP tests, and benchmarks.
2. **JSON service:** middleware, validation, error contracts, transactions,
   readiness checks, structured logs, and graceful shutdown.
3. **Streaming log analyzer:** CLI boundaries, buffered I/O, cancellation,
   profiling, and atomic output replacement.
4. **Typed in-memory cache:** generics, synchronization, eviction policy,
   metrics, race tests, and package-level API design.

## Repository conventions

- Format Go source with `gofmt`.
- Prefer the standard library unless a dependency earns its cost.
- Keep interfaces small and declare them near the consumer.
- Comments explain intent, constraints, or surprising behavior; they do not
  narrate obvious syntax.
- Tests must be deterministic and must not depend on the public internet.
- Preserve attribution attached to examples derived from external teaching
  material.

See [CONTRIBUTING.md](CONTRIBUTING.md) for the change workflow.

## License

Repository-original material is available under the [MIT License](LICENSE).
Individual examples may carry an additional attribution or license notice; keep
those notices with the example when moving or adapting it.
