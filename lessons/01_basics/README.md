# Lesson 01: Language Basics

## Learning goals

- Explain zero values, typed and untyped constants, and composite literals.
- Distinguish a slice header from its backing array.
- Use the comma-ok form for map lookups and impose order only when required.
- Explain why a string is bytes while `range` over a string yields runes.

## Mental model

Go favors useful zero values and explicit data movement. Assigning an array copies
all its elements; assigning a slice copies a small descriptor that may still refer
to the same backing array. Map iteration order is unspecified. A string is an
immutable byte sequence and may contain UTF-8 text, so byte indexes and character
positions are different concepts.

## Examples

```bash
go run ./lessons/01_basics/examples/variables
go run ./lessons/01_basics/examples/maps
go run ./lessons/01_basics/examples/buffering
go test ./lessons/01_basics/examples/...
```

The runtime-failure example intentionally ends with an index-out-of-range panic. Predict the
printed values and the failing index before running it:

```bash
go run ./lessons/01_basics/runtime_failures/slice_shrink
```

Read [runes.md](runes.md) before implementing text-processing code.

## Production tradeoffs and common mistakes

- A slice passed to a function can mutate shared elements even though its header
  was copied.
- Appending may reuse or replace the backing array; do not infer ownership from
  the variable name.
- A missing map key returns the value type's zero value. Use comma-ok when absence
  has different meaning from a stored zero value.
- Never depend on map iteration order. Sort keys only when the output contract
  requires deterministic ordering.
- Do not index UTF-8 text and call the result a character.

## Interview prompts

1. What is copied when a slice is assigned or passed to a function?
2. When can appending through one slice affect another slice?
3. How would you distinguish a missing map entry from a stored zero value?
4. What do `len(text)`, `text[i]`, and `range text` count or return?

## Practice

Write a function that returns deterministically ordered map entries, then write a
UTF-8-safe function that reverses text by rune. State what each function allocates.
