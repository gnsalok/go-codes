# Lesson 18: Testable CLI Design

## Learning goals and mental model

Keep process concerns at the edge. Parse arguments into explicit configuration,
send normal output to stdout and diagnostics to stderr, return an exit code, and let
`main` be the only place that calls `os.Exit`. Application behavior should accept
ordinary values and `io.Reader`/`io.Writer` boundaries so tests need no subprocess.

```bash
go test ./lessons/18_cli_design
```

## Tradeoffs and common mistakes

- Exit codes are a compatibility contract for scripts.
- Writing diagnostics to stdout breaks pipelines.
- Global `flag.CommandLine` state makes tests and multiple command instances interfere.
- Environment variables need explicit precedence relative to flags and config files.
- Handle broken pipes and cancellation for commands that stream large output.

## Interview prompts and practice

1. When is a subprocess test worth its cost?
2. How would you design cancellation for a long-running command?
3. Add environment fallback for `-name` and test the documented precedence.
