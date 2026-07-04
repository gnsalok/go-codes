# Lesson 01: Basics

## Core Go concepts

- Declaring constants and variables with explicit and inferred types.
- Defining structs and constructing values.
- Working with maps, slices, and `range` loops.
- Reading small `main` package examples as standalone programs.

## Quick reading

The exercises in `syntax/` are intentionally small and focused. Each file demonstrates one language building block: `var.go` introduces variable declarations, `const.go` introduces constants, `struct.go` models data with structs, `maps.go` shows key/value collections, and `slice_range.go` demonstrates iteration over slices. The top-level `test.go` remains a simple starter program for experimenting with package-level code.

## Run the lesson

From the repository root:

```bash
go run ./lessons/01_basics/syntax
```

You can also run an individual file when it has a `main` function:

```bash
go run ./lessons/01_basics/syntax/var.go
```

## Run tests

This lesson currently has no `_test.go` files, but it can still be compiled with:

```bash
go test ./lessons/01_basics/...
```
