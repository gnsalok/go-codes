# Lesson 08: Time

## Core Go concepts

- Working with Unix epoch timestamps.
- Converting and printing UTC times.
- Using the standard library `time` package for portable time handling.

## Quick reading

The `time/` examples are short demonstrations of common time operations. `epoch.go` focuses on Unix timestamp values, while `utc.go` highlights UTC-oriented formatting and conversion. These examples are intentionally small so students can experiment with layouts, zones, and parsing.

## Run the lesson

From the repository root:

```bash
go run ./lessons/08_time/time/epoch.go
```

## Run tests

This lesson currently has no `_test.go` files, but package compilation can be checked with:

```bash
go test ./lessons/08_time/...
```
