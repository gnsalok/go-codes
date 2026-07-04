# Lesson 02: Composition

## Core Go concepts

- Building behavior through composition instead of inheritance.
- Grouping related data and behavior into cohesive types.
- Using type assertions and custom method sets.
- Evolving APIs from primitive values toward decoupled higher-level designs.

## Quick reading

The `composition/` examples show how Go programs are assembled from small concrete types and functions. The `grouping/` examples compare a problem version with a cleaner fix. The `api-design-decoupling/` series walks through multiple iterations of the same API design idea, making it useful to read the files in version order from `impl_v0` through `implV4`. The `type-assertion/` examples demonstrate how interface values can be inspected safely at runtime.

## Run the lesson

From the repository root, run individual examples or packages:

```bash
go run ./lessons/02_composition/composition/gptexample.go
```

Some directories contain multiple standalone examples. If a package has more than one `main` function, run a specific file instead of the whole directory.

## Run tests

This lesson currently has no `_test.go` files, but package compilation can be checked with:

```bash
go test ./lessons/02_composition/...
```
