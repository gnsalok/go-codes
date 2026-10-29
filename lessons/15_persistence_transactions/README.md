# Lesson 15: Persistence and Transactions

## Learning goals and mental model

`*sql.DB` is a concurrency-safe pool, not one connection. Put transaction boundaries
around business invariants, begin close to the use case, pass the transaction to the
operations that participate, roll back on every early return, and commit once at the
end. Context bounds acquisition and database calls.

```bash
go test ./lessons/15_persistence_transactions
```

The example intentionally depends only on `database/sql`; exercising it requires a
driver and integration database chosen by the application.

## Tradeoffs and common mistakes

- A repository interface should clarify an application seam, not mirror every SQL
  method.
- Long transactions increase lock time and reduce concurrency.
- Driver placeholder syntax and `LastInsertId` support are database-specific.
- Always decide how commit errors and rollback errors are reported.

## Interview prompts and practice

1. Which layer owns the transaction for a multi-repository use case?
2. How do you test rollback and commit failure paths?
3. Implement a transfer operation whose balance invariant is atomic.
