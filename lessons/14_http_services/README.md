# Lesson 14: Production-Grade HTTP Services

## Core Go concepts

- Building JSON HTTP APIs with the standard library.
- Organizing routing, middleware, and handlers around clear dependencies.
- Validating requests and returning consistent error responses.
- Adding health, readiness, and dependency checks.
- Performing graceful shutdown and integration testing with `httptest`.

## Quick reading

A production HTTP service is more than a handler function. It needs explicit dependencies, request validation, predictable response shapes, observability hooks, and a shutdown path that lets in-flight requests finish. Keep handlers thin: decode input, call application behavior, map results to HTTP responses, and let lower layers own domain rules.

Use middleware for cross-cutting behavior such as request IDs, logging, authentication, panic recovery, and timeouts. Use `httptest` to test handlers without starting a real server, and prefer table-driven tests for request and response combinations.

This lesson builds on `lessons/05_testing`, `lessons/12_context_services`, and `lessons/13_advanced_errors`.

## Suggested exercises

1. Create a JSON handler with request validation and structured error responses.
2. Add middleware that attaches a request ID and logs status codes.
3. Implement `/healthz` and `/readyz` endpoints with different responsibilities.
4. Test a handler with `httptest.NewRecorder` and `httptest.NewRequest`.
5. Add graceful shutdown around an `http.Server`.

## Run the lesson

The `server.go` example provides a small JSON API with validation and a health endpoint. From the repository root:

```bash
go test ./lessons/14_http_services/...
```

## Run tests

Run the example tests and package compilation check with:

```bash
go test ./lessons/14_http_services/...
```
