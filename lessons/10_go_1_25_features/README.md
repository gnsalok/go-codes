# Lesson 10: Go 1.25 Features

## Core Go concepts

- Reviewing new language and standard-library behavior introduced around Go 1.25.
- Practicing modern `sync.WaitGroup` helper patterns where appropriate.
- Keeping examples compatible with the repository's `go 1.25.6` module directive.

## Quick reading

The `go1.25-features/` directory contains the current Go 1.25 feature notes and a small WaitGroup-related example. Use this lesson after the concurrency lesson so students already understand goroutines, synchronization, and why helper APIs improve readability.

## Run the lesson

From the repository root:

```bash
go run ./lessons/10_go_1_25_features/go1.25-features/wg.gofeat.go
```

## Run tests

This lesson currently has no `_test.go` files, but package compilation can be checked with:

```bash
go test ./lessons/10_go_1_25_features/...
```
