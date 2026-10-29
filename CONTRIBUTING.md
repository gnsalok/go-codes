# Contributing

Contributions should make the curriculum more accurate, focused, or useful for
senior Go interview preparation.

## Before changing an example

1. Identify whether it is runnable, tested, a design progression, or an
   intentional compiler failure.
2. Preserve its teaching objective. Do not turn a deliberately failing example
   into a passing one without moving the failure into an explicitly documented
   replacement.
3. Check neighboring examples and the lesson README for naming and structure.
4. Preserve license and attribution notices.

## Code changes

- Format source with `gofmt`.
- Use descriptive Go names and conventional snake-case filenames.
- Prefer clear control flow over compact or clever code.
- Add or update focused tests when behavior changes.
- Keep tests deterministic; use fakes, `httptest`, and temporary directories
  instead of public services or machine-specific paths.
- Run the commands documented by the affected lesson. Use `go test -race` when
  changing concurrent code.

## Documentation changes

Each lesson should answer:

- What should the learner be able to explain or implement?
- What mental model makes the behavior predictable?
- Which tradeoffs matter in production?
- Which mistakes are common in interviews and real systems?
- Which exact command runs or tests the example?

Remove duplicated prose instead of adding another summary. Keep commands, paths,
version requirements, caveats, and source links accurate.

## Pull requests

Keep changes focused and describe the learning outcome they improve. Separate
behavior changes from broad refactoring when practical. A pull request should
include:

- the affected lessons;
- the behavior intentionally preserved or changed;
- the commands used for verification; and
- any intentional compiler failure that remains outside normal package tests.

Use [Conventional Commits](https://www.conventionalcommits.org/) for commit
messages, for example `docs: clarify channel ownership` or
`refactor: isolate typed nil failure example`.
