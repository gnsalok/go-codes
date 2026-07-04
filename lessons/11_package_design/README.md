# Lesson 11: Package Design

## Core Go concepts

- Naming packages after the behavior they provide rather than broad technical layers.
- Designing small exported APIs and keeping implementation details unexported.
- Using `internal/` packages to enforce module boundaries.
- Avoiding circular dependencies by moving shared behavior behind smaller abstractions.
- Managing module versions and `replace` directives responsibly.

## Quick reading

Start by reviewing the public surface of a package: exported names, constructor functions, and the interfaces accepted by callers. Good Go packages are usually small, cohesive, and easy to describe in one sentence. Prefer package names such as `auth`, `cache`, or `invoice` over names such as `helpers` or `common` because behavior-oriented names make dependencies easier to understand.

Use `internal/` when a package is useful inside the repository but should not become part of the public module contract. If two packages import each other, split the shared contract into a smaller package or move behavior toward the caller that owns the dependency.

Pair this lesson with `lessons/03_interfaces_and_decoupling` before designing public APIs.

## Suggested exercises

1. Take a package with too many exported names and make implementation details unexported.
2. Move repository-only helpers into an `internal/` package and confirm external packages cannot import them.
3. Refactor a circular dependency sketch into one-directional dependencies.
4. Write a short package comment that explains the package behavior without describing its files.

## Run the lesson

The `cache.go` example shows a small behavior-focused package with a narrow exported API and hidden storage details:

```bash
go test ./lessons/11_package_design/...
```

## Run tests

Run the example tests and package compilation check with:

```bash
go test ./lessons/11_package_design/...
```
