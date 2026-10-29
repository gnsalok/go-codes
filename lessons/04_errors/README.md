# Lesson 04: Errors, Panics, and Recovery

## Learning goals

- Add operation-specific context while preserving an error chain.
- Choose `errors.Is`, `errors.As`, or direct comparison intentionally.
- Explain why an interface containing a typed nil pointer is non-nil.
- Use panic and recovery only at a boundary that can make a safe decision.

## Mental model

An error is a value and often part of an API contract. Wrap with `%w` when a
caller must retain access to the cause. Use `errors.Is` for identity or semantic
equivalence and `errors.As` to extract structured error data. Do not branch on
error strings.

Panic unwinds the current goroutine while running deferred calls. `recover` can
stop that unwind only when called by a deferred function in the same goroutine.
Recovery is appropriate at process or request boundaries that can log context,
restore invariants, and return a controlled failure; it is not ordinary error
handling.

## Runnable examples

```bash
go run ./lessons/04_errors/examples/wrapping
go run ./lessons/04_errors/examples/is_and_as
go run ./lessons/04_errors/examples/custom_type
go run ./lessons/04_errors/examples/switch_on_sentinel
go run ./lessons/04_errors/examples/recover_simple
go test ./lessons/04_errors/examples/...
```

## Typed-nil progression

Both programs compile. The problem exits through `log.Fatal` because assigning a
nil `*customError` to `error` creates an interface with a dynamic type. The solution
returns `error` directly and therefore returns a nil interface on success.

```bash
go run ./lessons/04_errors/progression/typed_nil/problem
go run ./lessons/04_errors/progression/typed_nil/solution
```

## Production tradeoffs and common mistakes

- Wrapping every layer with vague text creates noisy chains; add the operation and
  relevant identifiers instead.
- Sentinel errors are easy to inspect but form a long-lived public contract.
- Typed errors carry structured data but expose more API surface.
- Logging and returning the same error at every layer duplicates incidents. Add
  context while returning, and log once at the system boundary.
- Recovering without restoring a valid state can let corrupted work continue.

## Interview prompts

1. When should a caller use `errors.Is` rather than direct equality?
2. What exactly makes a typed-nil interface non-nil?
3. Where would you recover from a panic in an HTTP service, and what would you do
   before returning a response?
4. How do you expose a stable error contract without coupling callers to strings?

## Practice

Create a batch operation that returns successful results plus `errors.Join` for
independent failures. Write tests proving callers can still find each sentinel or
typed error through the joined chain.
