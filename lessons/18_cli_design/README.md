# Lesson 18: CLI Application Design

## Core Go concepts

- Designing command-line flags and environment configuration.
- Separating command parsing from application behavior.
- Returning meaningful exit codes.
- Producing user-friendly terminal output.
- Testing command behavior without shelling out unnecessarily.

## Quick reading

A good CLI is predictable, scriptable, and easy to test. Keep parsing and presentation at the edges, and move business logic into functions that accept explicit inputs and outputs. This makes commands easier to reuse from tests and other packages.

Use flags for command-specific options and environment variables for deployment or user-level configuration. Print normal output to standard output, diagnostics to standard error, and map known failures to clear exit codes. Tests should call command functions directly whenever possible instead of spawning subprocesses.

## Suggested exercises

1. Build a command function that accepts `io.Reader`, `io.Writer`, and `io.Writer` for stderr.
2. Parse flags into a configuration struct and pass it to application logic.
3. Define exit codes for usage errors, validation failures, and runtime failures.
4. Write tests that verify stdout, stderr, and exit code behavior.
5. Add environment variable fallback for one optional setting.

## Run the lesson

The `command.go` example keeps parsing, output, and exit-code behavior testable without shelling out. From the repository root:

```bash
go test ./lessons/18_cli_design/...
```

## Run tests

Run the example tests and package compilation check with:

```bash
go test ./lessons/18_cli_design/...
```
