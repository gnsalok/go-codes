# Lesson 02: Composition and API Evolution

## Learning goals

- Build behavior by composing concrete types instead of modeling inheritance.
- Explain embedding as promotion and delegation, not subtyping.
- Design interfaces around behavior required by a consumer.
- Evolve an API from primitive coupling toward smaller capabilities.

## Mental model

Go types do not enter a subtype hierarchy because one struct embeds another.
Embedding promotes fields and methods; interfaces group types implicitly by the
methods they provide. Prefer concrete types until a caller needs to depend on a
smaller behavior.

## Examples and progression

```bash
go run ./lessons/02_composition/examples/behavior_grouping
go run ./lessons/02_composition/examples/custom_stringer
go run ./lessons/02_composition/examples/engine_composition
```

The `progression/v0` through `progression/v4` directories evolve one copy
operation from primitive values toward separate pulling and storing capabilities.
Read them in order and identify which dependency each version removes.

```bash
go run ./lessons/02_composition/progression/v0
go run ./lessons/02_composition/progression/v4
go test ./lessons/02_composition/examples/... ./lessons/02_composition/progression/...
```

## Intentional compiler failures

`compile_failures/state_hierarchy` attempts to store `Dog` and `Cat` values in a
`[]Animal`. Embedding `Animal` does not make either type an `Animal` subtype.

`compile_failures/interface_assignment` assigns a `Mover` interface value to a
`MoveLocker`. The static source interface does not guarantee the extra methods,
even when its current dynamic value happens to implement them.

Run each failure separately and read the diagnostic:

```bash
go run ./lessons/02_composition/compile_failures/state_hierarchy
go run ./lessons/02_composition/compile_failures/interface_assignment
```

## Production tradeoffs and common mistakes

- Embedding a large interface also embeds a large dependency and a nil-value risk.
- Returning interfaces often hides useful concrete behavior; accept interfaces at
  the consumer boundary when substitution is actually needed.
- Do not add an interface only to make a type “mockable.” A focused function or
  small consumer-owned interface is usually enough.
- Type assertions move a guarantee from compile time to runtime. Prefer ordinary
  method calls when the contract is known statically.

## Interview prompts

1. How does embedding differ from inheritance?
2. Why should the consumer usually own an interface?
3. When would you accept an interface but return a concrete type?
4. What information is stored inside a non-nil interface value?

## Practice

Add an in-memory source and destination to the final progression without changing
the copy orchestration. Explain whether each new type needs to be exported.
