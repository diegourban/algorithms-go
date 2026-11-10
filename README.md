# algorithms-go
[![build](https://github.com/diegourban/algorithms-go/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/diegourban/algorithms-go/actions/workflows/build.yml)
[![codecov](https://codecov.io/gh/diegourban/algorithms-go/branch/main/graph/badge.svg)](https://codecov.io/gh/diegourban/algorithms-go)

A small Go library of algorithms and data structures, plus worked problems organised by topic. Go sibling of
[algorithms-java](https://github.com/diegourban/algorithms-java), which is the source of truth for the problem catalogue.

## Build

Requires Go 1.27 and [golangci-lint](https://golangci-lint.run/) v2.

```sh
make check   # lint (incl. gofmt/goimports and doc comments), tests with -race, coverage gate (90% of statements)
make fmt     # format the sources
make test    # tests only
```

## Layout

```
text/                          algorithms on strings (Java: string.StringUtils)
problems/<topic>/<name>/       worked problems (see below)
```

One concept per package, at the repo root, named so it never shadows the standard library (`text`, not `strings`;
`numbertheory`, not `math`). Every exported identifier has a doc comment; the linter enforces it.

## Code and test pattern

- Implementation in `<concept>.go`, tests in `<concept>_test.go` in the same package.
- Table-driven tests with the standard `testing` package; no assertion library.
- Property-style checks as Go fuzz tests (`func FuzzX(f *testing.F)`), the counterpart of the jqwik properties in
  algorithms-java. Inputs that once failed are kept under `testdata/fuzz/` as regression cases.

## Problems

Each problem is a package `problems/<topic>/<name>/` with a `Solve(...)` function and its tests:

- The package doc comment gives the title, the source (beecrowd/URI, Codewars, SPOJ, Interview Cake, HackerRank) and a
  link, plus a line on the approach and its complexity.
- A problem's id is `<topic>/<kebab-name>`, shared with algorithms-java. Go package names cannot contain `-`, so the
  directory drops the hyphens: `math/prime-generator` becomes `problems/math/primegenerator/`.
- Judge sample files (`.in`/`.out`) are copied from algorithms-java
  (`src/test/resources/problems/<topic>/<kebab-name>/`) into the problem's `testdata/` folder when it is ported.
