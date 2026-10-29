# Lesson 13: Stable Error Contracts

## Learning goals and mental model

Callers need stable facts, while operators need diagnostic context. Sentinels expose
identity, typed errors expose structured fields, and `%w` preserves those contracts
through higher-level operations. User-facing messages and internal diagnostics are
different outputs and should not share accidental string contracts.

```bash
go test ./lessons/13_advanced_errors
```

## Tradeoffs and common mistakes

- Exporting a sentinel or error type commits the package to its meaning.
- `errors.Join` preserves independent failures but callers still need a partial
  result policy.
- Logging at every layer duplicates one failure; enrich while returning and log at
  the boundary that decides the outcome.
- Validation fields should be stable identifiers, not presentation text.

## Interview prompts and practice

1. When would you choose a sentinel over a typed error?
2. How do you avoid leaking database or payment-provider details to an API client?
3. Extend `Charge` with a wrapped provider error and prove `errors.Is` still works.
