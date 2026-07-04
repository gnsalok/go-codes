# go-codes

![Go Version](https://img.shields.io/badge/go-v1.25.6-blue.svg) [![License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://github.com/gnsalok/go-mastermind/blob/main/LICENSE)

GoCodes is a collection of various Go language programming exercises and solutions. It serves as a learning resource for Go beginners and a practice ground for experienced developers to hone their Go skills. Each exercise includes a problem statement and a corresponding solution implemented in Go.

## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Curriculum Layout](#curriculum-layout)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Features

- Diverse Collection: **go-codes** covers a wide range of programming exercises, including algorithms, data structures, and common coding challenges, API, Testing, Benchmarking with extensive notes.
- Detailed Concurrency and Channel patterns are documented under `lessons/06_concurrency/concurrency/channels/*` directory.
- Clear Problem Statements: Each exercise comes with a well-defined problem statement to guide your implementation.
- Structured Solutions: Solutions are provided with comments to help you understand the thought process and best practices used.
- Unit Tests: The repository includes unit tests for each exercise, ensuring code correctness.
- Continuous Updates: The repository is regularly updated with new exercises and improvements.

## Requirements

To run the exercises and solutions in this repository, you need to have the following installed on your system:

- Go (v1.25.6 or above)

## Installation

To use this repository, simply clone it to your local machine:

```bash
git clone https://github.com/gnsalok/go-codes.git
```

## Curriculum Layout

Exercises are organized into numbered lesson directories under `lessons/` so students can progress from fundamentals to advanced topics:

1. `lessons/01_basics` — syntax, variables, constants, structs, maps, and slices.
2. `lessons/02_composition` — composition, grouping, API design, and type assertions.
3. `lessons/03_interfaces_and_decoupling` — interfaces, method semantics, and decoupling.
4. `lessons/04_errors` — error values, wrapping, inspection, panic recovery, and typed-nil pitfalls.
5. `lessons/05_testing` — unit tests, subtests, benchmarks, HTTP tests, and mocks.
6. `lessons/06_concurrency` — goroutines, channels, context, mutexes, atomics, and worker pools.
7. `lessons/07_generics` — type parameters and constraints.
8. `lessons/08_time` — epoch and UTC time handling.
9. `lessons/09_performance` — benchmarking and profiling.
10. `lessons/10_go_1_25_features` — Go 1.25-focused examples and notes.
11. `lessons/11_package_design` — package naming, exported APIs, internal packages, and module boundaries.
12. `lessons/12_context_services` — request-scoped cancellation, deadlines, metadata, and deterministic context tests.
13. `lessons/13_advanced_errors` — stable error contracts, diagnostics, joined errors, and logging context.
14. `lessons/14_http_services` — JSON APIs, middleware, validation, health checks, graceful shutdown, and `httptest`.
15. `lessons/15_persistence_transactions` — `database/sql`, transaction boundaries, rollback patterns, and persistence tests.
16. `lessons/16_observability` — structured logging, metrics, tracing concepts, and production profiling.
17. `lessons/17_runtime_memory_model` — escape analysis, garbage collection, races, happens-before rules, and scheduler behavior.
18. `lessons/18_cli_design` — flags, environment configuration, command structure, exit codes, and CLI tests.
19. `lessons/19_advanced_generics` — constraint design, generic data structures, and API tradeoffs.

Each lesson includes a `README.md` with concepts, a quick reading guide, and commands for running examples and tests. See `ADVANCED_CURRICULUM_PLAN.md` for the implemented advanced roadmap and remaining capstone plan.

## Usage

Navigate to the numbered lesson directory you want to explore under `lessons/`. Each exercise folder contains a Go file with the problem statement, a solution Go file, and a corresponding unit test Go file.

To run the unit tests for a specific exercise, use the following command:

```bash
go test -v
```

Feel free to experiment with the code, modify the solutions, and create new exercises. Contributions and suggestions for improvements are highly welcome!

## Contributing

Contributions to **go-codes** are encouraged! Whether it's adding new exercises, improving existing solutions, or enhancing the documentation, every contribution is valuable. Please read our [Contribution Guidelines](CONTRIBUTING.md) for more details on how to get involved.

## Lead Maintainer
**Alok Tripathi** - *Github* - [Alok Tripathi](https://github.com/gnsalok)

---

Explore, learn, and have fun with **go-codes**! We hope you find this repository useful on your journey to mastering Go programming. Happy coding! 🚀💻

