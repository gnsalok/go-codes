# Lesson 19: Generics Beyond the Basics

## Core Go concepts

- Designing constraints that describe required operations without overfitting.
- Building generic data structures with clear ownership and zero-value behavior.
- Comparing generics, interfaces, and code generation.
- Preserving readability while removing meaningful duplication.
- Avoiding over-generalized APIs that are harder to use than concrete code.

## Quick reading

Generics are most useful when the same algorithm or data structure works across multiple concrete types while preserving type safety. Constraints should be as small as possible: require only the operations the generic code actually performs. If behavior is best expressed through methods, an interface may be clearer. If performance or generated specialization is required, code generation may still be appropriate.

Start with `lessons/07_generics`, then use this lesson to evaluate API design tradeoffs. Ask whether a type parameter makes the caller's code simpler. If not, concrete functions or interfaces may be better.

## Suggested exercises

1. Write a generic stack or queue with useful zero-value behavior.
2. Design a constraint for ordered values and compare it with a callback-based API.
3. Rewrite an interface-based helper as a generic function, then compare readability.
4. Identify an over-generalized API and simplify it.
5. Add table-driven tests for multiple concrete instantiations.

## Run the lesson

The `stack.go` example implements a zero-value-friendly generic stack. From the repository root:

```bash
go test ./lessons/19_advanced_generics/...
```

## Run tests

Run the example tests and package compilation check with:

```bash
go test ./lessons/19_advanced_generics/...
```
