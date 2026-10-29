# Interview-Relevant Changes in Go 1.25–1.27

This is a reading guide, not a replacement for the official release notes.

## Go 1.25

- `sync.WaitGroup.Go` expresses the common “start and account for this goroutine”
  operation directly.
- `testing/synctest` became a supported tool for deterministic tests involving
  goroutines and time.
- The runtime gained container-aware `GOMAXPROCS` behavior and the Green Tea garbage
  collector was available experimentally.

Source: [Go 1.25 release notes](https://go.dev/doc/go1.25).

## Go 1.26

- `new(expression)` can allocate a variable initialized from an expression.
- Green Tea became the default garbage collector after its Go 1.25 experiment.
- The runtime added an experimental goroutine-leak profile and reduced baseline
  cgo call overhead.

Source: [Go 1.26 release notes](https://go.dev/doc/go1.26).

## Go 1.27

- Concrete methods may declare type parameters; interface methods still may not.
- Type inference for generic functions applies in more assignment contexts.
- `go test` enables the `stdversion` vet check by default.
- The goroutine-leak profile became generally available.
- `encoding/json/v2`, `encoding/json/jsontext`, and `uuid` joined the standard
  library. The v1 JSON API remains supported.

Source: [Go 1.27 release notes](https://go.dev/doc/go1.27).

## Upgrade discussion checklist

1. Minimum supported operating systems and architectures.
2. Module `go` version, toolchain availability, and developer/CI parity.
3. New vet diagnostics and changed tests.
4. Runtime or standard-library behavior controlled by `GODEBUG`.
5. Benchmark, profile, binary-size, and memory changes under representative load.
6. Staged rollout and rollback path.
