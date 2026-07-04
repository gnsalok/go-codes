# Lesson 04: Errors

## Core Go concepts

- Returning and checking explicit `error` values.
- Creating custom error types and sentinel error variables.
- Wrapping errors and inspecting chains with `errors.Is` and `errors.As`.
- Recovering from panics with `defer` and `recover`.
- Understanding typed-nil error traps.

## Quick reading

The files in `errors/` progress from basic error handling to more advanced error chains and recovery. `error_wrap.go`, `error_is_and_as.go`, and the custom error examples are best read together because they show how callers preserve context while still allowing programmatic checks. The `tricky_problem/` directory demonstrates the typed-nil problem and its fix, which is an important Go interview and production debugging topic.

## Run the lesson

From the repository root:

```bash
go run ./lessons/04_errors/errors/error_wrap.go
```

Many files are independent demonstrations, so run specific files when a directory contains more than one `main` function.

## Run tests

This lesson currently has no `_test.go` files, but package compilation can be checked with:

```bash
go test ./lessons/04_errors/...
```
