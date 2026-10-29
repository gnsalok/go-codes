# Lesson 12: Context Across Service Boundaries

## Learning goals and mental model

`context.Context` carries cancellation, deadlines, and request-scoped metadata
across API boundaries. Pass it explicitly as the first parameter; do not store it in
a long-lived struct. A function that can block should define what cancellation means
and return promptly when the caller is no longer interested.

```bash
go test ./lessons/12_context_services
```

The example races a result channel against caller cancellation and distinguishes a
closed result stream from a canceled request.

## Tradeoffs and common mistakes

- Context values are for cross-boundary request metadata, not optional parameters
  or dependency injection.
- `context.Background()` inside request work severs the caller's lifetime.
- Always call the cancel function returned by `WithCancel`, `WithTimeout`, or
  `WithDeadline` so associated resources can be released early.
- Check cancellation around blocking work, not after an uninterruptible operation.

## Interview prompts and practice

1. Should cancellation be returned as `ctx.Err()` or wrapped with operation context?
2. When may cleanup continue after a request is canceled?
3. Add a timeout test using controlled channels rather than `time.Sleep`.
