# Lesson 15: Persistence and Transactions

## Core Go concepts

- Using `database/sql` for connection pooling, queries, and scans.
- Passing context into database operations.
- Defining transaction boundaries close to the use case.
- Rolling back safely and committing only after all required work succeeds.
- Introducing repository interfaces only when they clarify seams.
- Choosing test strategies for persistence code.

## Quick reading

The `database/sql` package provides a common abstraction over SQL drivers. A `*sql.DB` is a concurrency-safe pool, not a single connection. Long-running calls should receive a context so request cancellation and deadlines can stop database work.

Transactions should express business boundaries. Begin the transaction near the application operation that requires atomicity, pass the transaction through the functions that need it, roll back on error, and commit at the end. Avoid creating repository interfaces by default; introduce them when they make tests or package boundaries clearer.

This lesson pairs naturally with `lessons/14_http_services`.

## Suggested exercises

1. Write a query function that accepts `context.Context` and scans rows carefully.
2. Implement a transaction helper that rolls back on error and commits on success.
3. Compare testing with a fake repository, a test database, and SQL-level integration tests.
4. Document where transaction boundaries belong in a sample service operation.

## Run the lesson

The `transactions.go` example shows a context-aware transaction boundary around an insert operation. From the repository root:

```bash
go test ./lessons/15_persistence_transactions/...
```

## Run tests

Run the example tests and package compilation check with:

```bash
go test ./lessons/15_persistence_transactions/...
```
