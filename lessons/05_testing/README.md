# Lesson 05: Testing and Benchmarks

## Learning goals

- Write table-driven tests and descriptive subtests.
- Choose among real dependencies, fakes, stubs, and mocks.
- Test HTTP behavior without depending on the public internet.
- Write benchmarks that measure one decision and report allocations.

## Mental model

A test is an executable contract. Assert observable outcomes rather than internal
call sequences so implementation refactors do not invalidate correct behavior.
Keep most tests small and deterministic; use a real local boundary only when the
boundary itself is what you need to prove.

For HTTP code, `httptest.NewRequest` and `httptest.NewRecorder` exercise a handler
without opening a socket. A custom `http.RoundTripper` tests client behavior without
DNS or network access. `httptest.Server` is useful when transport behavior matters,
but it is a larger test.

## Run tests and benchmarks

```bash
go test ./lessons/05_testing/handlers
go test ./lessons/05_testing/subtest
go test ./lessons/05_testing/testing
go test ./lessons/05_testing/mocking/...
go test -bench=. -benchmem ./lessons/05_testing/testing/...
```

The tests use local recorders and transports; they do not call public websites.

## Production tradeoffs and common mistakes

- Tests that share global state or depend on execution order are hard to run in
  parallel and easy to make flaky.
- `time.Sleep` makes timing tests slow and unreliable. Synchronize on observable
  events or use a controllable clock.
- Mocking every collaborator proves the mocks agree with the test, not that the
  system works. Prefer real implementations, then fakes, stubs, and finally mocks.
- A benchmark without `-benchmem`, stable input, or before-and-after results rarely
  supports a performance decision.
- Do not call `t.Parallel` until the test owns its state and dependencies.

## Interview prompts

1. What belongs in a unit test versus an integration test?
2. How would you test retry or timeout behavior without sleeping?
3. When is a fake more useful than an interaction-based mock?
4. How do you prevent compiler optimization from invalidating a benchmark?

## Practice

Add cancellation and malformed-response cases to the HTTP examples. Make each
failure deterministic, name the contract in the test, and explain the resource
level of the test: small, medium, or large.
