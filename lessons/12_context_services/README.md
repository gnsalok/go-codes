# Lesson 12: Context Propagation in Services

## Core Go concepts

- Passing `context.Context` as the first parameter to request-scoped work.
- Propagating cancellation and deadlines through service boundaries.
- Avoiding context values for optional function parameters.
- Carrying request metadata, such as trace IDs, safely and sparingly.
- Testing timeout and cancellation behavior deterministically.

## Quick reading

Context connects the lifetime of a request to the goroutines, network calls, database queries, and background work that serve it. Functions that may block should accept a `context.Context`, check `ctx.Done()` where appropriate, and return promptly when cancellation is observed.

Do not store long-lived state in a context. Use context values only for request-scoped metadata that crosses API boundaries, such as trace IDs or authentication claims. Configuration, loggers, repositories, and optional parameters should be passed explicitly through constructors or function arguments.

Review `lessons/06_concurrency` before this lesson, especially the context examples.

## Suggested exercises

1. Add cancellation support to a function that performs slow work.
2. Write a service function that respects `context.WithTimeout`.
3. Replace a context value used as an optional parameter with an explicit function argument.
4. Test cancellation by using controlled channels instead of sleeping for real time.

## Run the lesson

The `service.go` example waits for either service output or context cancellation. From the repository root:

```bash
go test ./lessons/12_context_services/...
```

## Run tests

Run the example tests and package compilation check with:

```bash
go test ./lessons/12_context_services/...
```
