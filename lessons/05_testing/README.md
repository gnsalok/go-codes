# Lesson 05: Testing

## Core Go concepts

- Writing table-driven unit tests.
- Organizing subtests with `t.Run`.
- Testing HTTP handlers and mocked web servers.
- Writing benchmarks with Go's `testing` package.
- Isolating external behavior with simple mocks and fakes.

## Quick reading

The `testing/` directory introduces table-driven tests, benchmarks, and HTTP test servers. The `subtest/` example focuses on nested test cases. The `handlers/` package pairs production handler code with tests, making it a good starting point for learning the red-green-refactor loop. The `mocking/` examples show how to replace external collaborators with small local abstractions.

## Run the lesson

From the repository root, execute the test-oriented packages directly through `go test`:

```bash
go test ./lessons/05_testing/...
```

## Run benchmarks

```bash
go test -bench=. ./lessons/05_testing/...
```
