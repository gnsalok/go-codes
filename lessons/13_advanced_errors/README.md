# Lesson 13: Advanced Error Design

## Core Go concepts

- Creating stable domain error contracts for callers.
- Separating user-facing messages from operational diagnostics.
- Wrapping errors while preserving programmatic inspection.
- Joining multiple errors and reporting partial failures.
- Logging errors without losing chain context.

## Quick reading

Start with `lessons/04_errors`, then use this lesson to think about errors as API contracts. A caller should be able to make safe decisions with `errors.Is` or `errors.As` without depending on fragile string comparisons. Error messages can change; error contracts should remain stable.

Keep user-facing messages separate from internal diagnostics. For example, an HTTP handler may return a short validation message to a client while logging the wrapped database or network error for operators. When multiple independent operations fail, `errors.Join` can preserve all failures while still allowing callers to inspect the combined error.

## Suggested exercises

1. Define sentinel domain errors and wrap them with operation-specific context.
2. Create a typed validation error that exposes fields for callers but keeps formatting flexible.
3. Use `errors.Join` to report partial failures from a batch operation.
4. Write a logging helper that records the full error chain without replacing the original error.

## Run the lesson

The `errors.go` example demonstrates stable domain errors, typed validation errors, and wrapped operational context. From the repository root:

```bash
go test ./lessons/13_advanced_errors/...
```

## Run tests

Run the example tests and package compilation check with:

```bash
go test ./lessons/13_advanced_errors/...
```
