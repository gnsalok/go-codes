# Lesson 07: Generics

## Learning goals

- Use type parameters when one algorithm applies to several concrete types.
- Write the smallest constraint that permits the required operations.
- Compare generics with interfaces and ordinary concrete functions.
- Explain type inference, approximation elements such as `~int`, and zero values.

## Mental model

Interfaces describe runtime behavior through method sets. Type parameters describe
a family of statically checked types used by one implementation. Choose generics
when preserving the caller's concrete type removes meaningful duplication; choose
an interface when the algorithm calls behavior supplied by different implementations.

```bash
go run ./lessons/07_generics/generics/generics.go
go test ./lessons/07_generics/...
```

## Production tradeoffs and common mistakes

- A broad constraint promises more than the implementation needs.
- A generic helper that only saves a few obvious lines can make call sites harder
  to understand.
- `~int` includes named types whose underlying type is `int`; `int` alone does not.
- Do not replace behavior-oriented interfaces with type lists merely because
  generics are available.

## Interview prompts

1. When does a type parameter provide a stronger contract than `any` plus a type
   assertion?
2. Why can a constraint interface be unusable as an ordinary value type?
3. How would you choose between an interface, a generic function, and code generation?

## Practice

Add a named integer type to the `Add` example, then narrow and widen the constraint
to observe which programs compile. Explain each change in the public contract.
