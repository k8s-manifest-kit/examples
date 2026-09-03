# Agent Guide: examples

This module contains runnable examples and tests for k8s-manifest-kit. Each example exposes `Run(context.Context) error` and has a `main_test.go` that exercises it.

## Examples

- `quickstart`: one Helm chart through `helm.NewEngine`.
- `filtering-transformation`: engine-level JQ filtering and label transformation.
- `multiple-sources`: multiple Helm sources in one renderer.
- `pipeline-hooks`: source selectors, source/renderer/engine post-renderers, render-time values, and apply ordering.
- `production-features`: Helm/YAML rendering with cache, metrics, and source annotations.
- `real-world`: composed filters and environment-specific transformations.

All examples use split module imports such as `github.com/k8s-manifest-kit/engine/pkg`, `renderer-helm/pkg`, and `pkg/util`. Helm examples require network access to public OCI registries or repositories.

## Development

Run commands from this directory:

```bash
make test
make fmt
make lint
make lint/fix
make check
```

Run an individual example with `go run ./<example>`. Keep examples runnable, update the README and this guide when adding one, and keep tests aligned with the behavior demonstrated.

