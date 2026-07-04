# Lesson 07: Generics

## Core Go concepts

- Writing functions and types with type parameters.
- Defining constraints for reusable algorithms.
- Choosing generics only when they reduce duplication without hiding simple code.

## Quick reading

The `generics/` example introduces Go's type parameter syntax in a compact form. Read it as a contrast to interface-based polymorphism from the previous lessons: generics preserve static type information, while interfaces focus on behavior through method sets.

## Run the lesson

From the repository root:

```bash
go run ./lessons/07_generics/generics/generics.go
```

## Run tests

This lesson currently has no `_test.go` files, but package compilation can be checked with:

```bash
go test ./lessons/07_generics/...
```
