# Lesson 11: Package Design

## Learning goals and mental model

Design packages around one behavior that can be described without listing files.
Keep the exported surface small, direct dependencies inward, and let useful zero
values work when that does not hide invalid state. Package names describe what the
caller gets—`cache`, `auth`, or `invoice`—rather than vague buckets such as `common`.

The example exposes construction, lookup, and mutation while hiding map storage.
Both `NewCache` and the zero value are ready to use:

```bash
go test ./lessons/11_package_design
```

## Tradeoffs and common mistakes

- Every exported name becomes a compatibility commitment.
- `internal/` enforces repository boundaries; naming alone does not.
- A package named after a technical layer often accumulates unrelated behavior.
- Moving a shared type “down” can create a dependency cycle; move contracts toward
  the consumer or introduce a smaller neutral package only when it is cohesive.

## Interview prompts and practice

1. How do you decide whether a constructor is necessary?
2. Where should an interface live when two packages collaborate?
3. Refactor an oversized package into cohesive boundaries and draw the dependency
   direction before moving code.
