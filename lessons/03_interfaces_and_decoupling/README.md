# Lesson 03: Interfaces and Decoupling

## Core Go concepts

- Defining behavior with small interfaces.
- Understanding method-set semantics and mixed receiver choices.
- Avoiding interface pollution by accepting concrete types until abstraction is needed.
- Using embedding and functions to decouple call sites from implementation details.

## Quick reading

The `interface/` examples focus on method semantics and common pitfalls. The `interface/inteface_pollution/` examples contrast over-abstracted code with a cleaner implementation. The `decoupling/` examples show several ways to separate behavior: embedded types, function values, interface contracts, and method receiver decisions. Read the accompanying markdown note in `decoupling/` before changing these examples.

## Run the lesson

From the repository root:

```bash
go run ./lessons/03_interfaces_and_decoupling/decoupling/embedding.go
```

Use file-specific `go run` commands for directories that contain multiple standalone `main` examples.

## Run tests

This lesson currently has no `_test.go` files, but package compilation can be checked with:

```bash
go test ./lessons/03_interfaces_and_decoupling/...
```
