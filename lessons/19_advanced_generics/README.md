# Lesson 19: Generic Data Structures and API Tradeoffs

## Learning goals and mental model

Generic containers should preserve static types, have useful zero values, and avoid
retaining removed values in their backing storage. Constraints expose only the
operations an algorithm requires. A type parameter should make the caller's code
simpler; otherwise concrete code or a behavior-oriented interface is often clearer.

```bash
go test ./lessons/19_advanced_generics
```

The stack clears a removed slot before shortening its slice so pointer-containing
values can be reclaimed even while the backing array remains allocated.

## Tradeoffs and common mistakes

- Zero-value-friendly containers reduce constructor ceremony.
- Returning `(T, bool)` avoids inventing a sentinel for arbitrary `T`.
- Generic APIs can increase binary code and documentation surface.
- Do not expose the backing slice unless callers are meant to share mutation and
  retention behavior.

## Interview prompts and practice

1. Why does `Pop` clear an element that is already outside the slice length?
2. When would an iterator be safer than returning a slice?
3. Implement a generic queue, then document its zero value and memory-retention policy.
