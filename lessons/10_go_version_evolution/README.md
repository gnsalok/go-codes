# Lesson 10: Go Version Evolution

## Learning goals

- Read release notes for language, toolchain, runtime, and library changes.
- Separate a language-version requirement from a toolchain preference.
- Adopt new APIs only when they improve the code's contract or readability.
- Identify compatibility, rollout, and benchmark work required by an upgrade.

## Mental model

An upgrade is a behavior change you did not write. The Go 1 compatibility promise
greatly reduces source breakage, but runtime defaults, vet checks, performance,
platform support, and exact diagnostics can still change. Read every intervening
release note, run tests and benchmarks, and review `GODEBUG` compatibility controls.

The repository targets Go 1.27.3. [release_notes.md](release_notes.md) summarizes a
small interview-relevant subset; the official notes remain authoritative.

## Example

Go 1.25 added `sync.WaitGroup.Go`, which combines starting and accounting for a
goroutine. The example keeps ownership visible while removing manual `Add` and
`Done` bookkeeping:

```bash
go run ./lessons/10_go_version_evolution/examples/wait_group_go
go test ./lessons/10_go_version_evolution/...
```

## Production tradeoffs and common mistakes

- Do not adopt a new API without updating the module's minimum Go version.
- Compiler and runtime speedups can change benchmark baselines without a source
  change; record toolchain versions with results.
- Experimental packages and `GOEXPERIMENT` behavior are not stable contracts.
- Exact error text, compressed bytes, profiles, and symbol names are often poor
  compatibility assertions.

## Interview prompts

1. How would you roll out a new Go toolchain across a fleet and detect regressions?
2. What does the `go` directive control compared with a `toolchain` directive?
3. Which tests would you run before adopting a new JSON implementation or GC?

## Practice

Choose one change from each release in `release_notes.md`. Explain whether it is a
language, API, tool, or runtime change and what evidence you would require before
using it in production.
