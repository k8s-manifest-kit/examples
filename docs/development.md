# Examples Development

This guide covers the contributor workflow for the `examples` module: the
runnable examples for the k8s-manifest-kit renderers and shared pipeline. For
the example catalog see [`README.md`](../README.md); for agent navigation see
[`AGENTS.md`](../AGENTS.md).

## Prerequisites

Build with the Go version required by [`go.mod`](../go.mod). Every example is
a `main` package that exposes `Run(context.Context) error`, so it can be run
directly with `go run ./<example>` and exercised by a test.

## Commands

Run all commands from the module root:

```bash
make test
make test/race
make test/integration
make fmt
make lint
make lint/fix
make lint/integration
make check
```

Coverage boundaries — keep these descriptions accurate when editing them:

- `make test` and `make test/race` are offline. They compile every package
  and run the unit suite, which currently contains no behavioral tests: a
  pass is a compile and vet check, not evidence that the examples render
  correctly.
- `make test/integration` is the live suite. It first verifies that `TestRun`
  is selected by the `integration` build tag for every entry in the
  Makefile's `INTEGRATION_EXAMPLES` list, then runs the tagged suite with the
  race detector and `-count=1`. `-count=1` disables result caching so a pass
  is always a fresh live execution; a previously cached success is never
  evidence of live verification.
- `make lint` runs the pinned golangci-lint (see `GOLANGCI_VERSION` in the
  Makefile) on the default build. `make lint/integration` runs the same
  linter with `--build-tags integration` so the tagged smoke tests are
  covered by the same configuration.
- `make check` runs `lint` and `vulncheck` only. It does not run integration
  lint, any tests, or formatting. CI runs `make lint/integration` and a
  separate formatting verification alongside `make check`.
- `make fmt` is the formatting target and is separate from `check`.

## Live sources and chart pins

The integration suite contacts public chart sources without credentials:
Docker Hub OCI registries (`bitnamicharts/nginx`,
`daprio/dapr-shared-chart`) and the Dapr chart repository
(`https://dapr.github.io/helm-charts`). Live Helm sources must set an exact
published `ReleaseVersion` pin:

- Verify a pin against the chart's published versions before using it; do
  not guess version availability and do not replace pins with
  latest-version lookups.
- Bumping a pin is a deliberate change: confirm the new version is
  published, then run `make test/integration` for a fresh live execution.

## Adding an example

1. Create `<example>/main.go` with a `main()` that calls `Run(ctx)` and an
   exported `Run(context.Context) error` holding the demonstrated behavior.
2. Add `<example>/main_test.go` behind `//go:build integration` in the
   external `main_test` package, following the existing smoke tests:
   isolate the Helm environment with `testenv.IsolateHelmEnv(t)`, attach the
   test logger with `logger.WithLogger`, bound the context with a timeout,
   and fail the test on any `Run` error. Pin every live chart to an exact
   published `ReleaseVersion`.
3. Add the example directory to `INTEGRATION_EXAMPLES` in the Makefile. The
   `test/integration` target fails if `TestRun` is not selected by the
   integration tag, so the list and the smoke tests must stay in sync —
   adding a file alone is not enough.
4. Run `make lint/integration` so the new tagged test is linted.
5. Update the catalog in [`README.md`](../README.md) and the example list in
   [`AGENTS.md`](../AGENTS.md).

## Code conventions

- Give each parameter its own type: `name string, namespace string`, not
  `name, namespace string`.
- Keep examples readable and self-contained. Shared helpers live in
  `internal/` (for example `internal/logger` and `internal/testenv`) rather
  than being duplicated per example.
- Wrap errors with operation context using `%w` where a cause exists, so a
  failure identifies the rendering or transformation step that produced it.
- Follow the repository's formatter and import ordering through `make fmt`;
  the linter configuration is `.golangci.yml`.
- Keep the offline suites offline: anything that needs network access
  belongs behind the `integration` build tag.

### Multiline strings

- Use one descriptive, standalone, package-level `const name =` declaration
  with a raw string literal for every multiline string: fixtures, manifests,
  templates, and expected output.
- Never embed multiline string literals in function or method bodies, and
  never declare them as function-local constants.
- Never group multiline constants in a `const (...)` block. Related
  single-line scalar constants can still be grouped, and single-line inline
  strings are unaffected.
- Keep large literal content separate from execution logic for readability.
  Preserve meaningful content indentation; the declaration layout must not
  introduce unintended bytes.
- A function may format the named template into a local variable. The
  prohibition concerns where the literal is declared, not where its string
  value is used.

```go
const fixtureYAML = `apiVersion: v1
kind: ConfigMap
`
```
