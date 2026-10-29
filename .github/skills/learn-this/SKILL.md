---
name: learn-this
description: Learn the structure, conventions, and verification model of the go-codes senior interview curriculum.
argument-hint: Optional lesson number or topic, such as concurrency, errors, or generics
---

# Learn This Codebase

## Start here

1. Read the root `README.md` for the curriculum map and repository conventions.
2. Read the target lesson's README before its Go files.
3. Classify each relevant directory as a runnable example, tested package, design
   progression, or intentional compiler failure.
4. Use the focused commands in the lesson README. Do not use `go test ./...` as a
   success gate because compiler failures are retained as teaching material.

## Repository shape

- `lessons/01_basics` through `lessons/21_closures` form the ordered curriculum.
- `examples/` contains runnable demonstrations.
- `progression/` compares versions of a design.
- `compile_failures/` contains source that must fail for a documented reason.
- Tested production-design packages are concentrated in Lessons 11–21.
- `REFERENCES.md` contains primary sources and attribution.

## Change discipline

- Preserve runnable behavior unless the lesson contains a technical error.
- Keep intentional failures isolated so unrelated redeclarations do not obscure the
  expected diagnostic.
- Prefer lesson-sized changes and focused verification.
- Run `gofmt`, `go vet`, and the lesson's tests; use `go test -race` for shared-memory
  and concurrency changes.
- Keep tests offline and deterministic.
