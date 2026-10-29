# Lesson 16: Observability

## Learning goals and mental model

Logs explain events, metrics show aggregate behavior, traces connect work across
boundaries, and profiles explain resource use inside a process. Start with an
operational question—availability, latency, correctness, or saturation—then choose
the signal that answers it with controlled cardinality and cost.

```bash
go test ./lessons/16_observability
```

## Tradeoffs and common mistakes

- High-cardinality labels such as user IDs can make metrics expensive or unusable.
- Logs need stable field names and appropriate redaction; never emit secrets or
  unrestricted personal data.
- Sampling reduces trace cost but can hide rare failures unless policy accounts for
  error and latency signals.
- A 3xx response is not a 2xx success; the example names their combined classification
  explicitly as `IsSuccessOrRedirect`.

## Interview prompts and practice

1. Define RED metrics for the greeting handler in Lesson 14.
2. Where would you attach a request ID and how would it cross service boundaries?
3. Add a test `slog.Handler` and assert structured fields rather than formatted text.
