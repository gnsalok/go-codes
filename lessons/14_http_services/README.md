# Lesson 14: HTTP Service Boundaries

## Learning goals and mental model

An HTTP handler is an adapter: decode and validate transport input, call application
behavior, then map the result to a stable response. Keep domain decisions out of the
transport layer. Middleware owns cross-cutting request behavior such as identity,
logging, authentication, recovery, and timeouts.

```bash
go test ./lessons/14_http_services
```

The tests exercise success, method rejection, malformed JSON, and validation with
`httptest` and no network listener.

## Tradeoffs and common mistakes

- Health and readiness answer different questions: process liveness versus ability
  to serve traffic with required dependencies.
- Limit request bodies, define unknown-field policy, and return consistent JSON
  errors before accepting untrusted input in production.
- Write headers only after all fallible response preparation that can still change
  the status code.
- Graceful shutdown needs a deadline and must stop accepting new work.

## Interview prompts and practice

1. Where should authentication and domain authorization live?
2. How would you test graceful shutdown with an in-flight request?
3. Add a request-size limit and test the exact transport contract.
