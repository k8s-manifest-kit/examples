# Examples

This module contains small runnable examples for the k8s-manifest-kit
renderers and shared pipeline.

| Example | Demonstrates |
| --- | --- |
| [`quickstart`](quickstart) | Basic renderer setup |
| [`multiple-sources`](multiple-sources) | Combining sources in one engine |
| [`filtering-transformation`](filtering-transformation) | Filters and transformers |
| [`pipeline-hooks`](pipeline-hooks) | Filters, transformers, and post-renderers |
| [`production-features`](production-features) | Production-oriented renderer options |
| [`real-world`](real-world) | A larger end-to-end example |
| [`pipeline-hooks`](pipeline-hooks) | Filters, transformers, and post-renderers |

The examples target Go 1.27.3. Run one with:

```bash
go run ./quickstart
go run ./pipeline-hooks
```

Run the complete module test suite with `make test`. Formatting, linting, and
the combined validation command are available through `make fmt`, `make lint`,
and `make check`.

## Testing

The module has two test suites with a strict boundary:

- `make test` and `make test/race` are offline: they compile every package and
  run the unit suite, which currently contains no behavioral tests. They never
  contact external services.
- `make test/integration` executes the six example `Run` smoke tests
  (`integration` build tag) against live chart sources, with the race detector
  and result caching disabled (`-count=1`) so a pass is always a fresh
  execution.

The integration suite needs network access to public registries and
repositories — Docker Hub (`bitnamicharts/nginx`, `daprio/dapr-shared-chart`,
both pinned to exact chart versions) and `https://dapr.github.io/helm-charts`
(`dapr`, pinned to an exact chart version). No credentials are required. CI
runs both suites on every push and pull request.

To add an example, follow [`AGENTS.md`](AGENTS.md) and keep the example
self-contained with a short README when it teaches a distinct workflow.

Related module documentation:
[`engine`](../engine), [`pkg`](../pkg), [`renderer-helm`](../renderer-helm),
[`renderer-kustomize`](../renderer-kustomize), [`renderer-mem`](../renderer-mem),
[`renderer-yaml`](../renderer-yaml), and
[`renderer-gotemplate`](../renderer-gotemplate).

## License

Apache License 2.0. See [LICENSE](LICENSE).
