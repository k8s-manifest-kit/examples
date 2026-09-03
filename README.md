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

The examples target Go 1.26.8. Run one with:

```bash
go run ./quickstart
go run ./pipeline-hooks
```

Run the complete module test suite with `make test`. Formatting, linting, and
the combined validation command are available through `make fmt`, `make lint`,
and `make check`.

To add an example, follow [`AGENTS.md`](AGENTS.md) and keep the example
self-contained with a short README when it teaches a distinct workflow.

Related module documentation:
[`engine`](../engine), [`pkg`](../pkg), [`renderer-helm`](../renderer-helm),
[`renderer-kustomize`](../renderer-kustomize), [`renderer-mem`](../renderer-mem),
[`renderer-yaml`](../renderer-yaml), and
[`renderer-gotemplate`](../renderer-gotemplate).

## License

Apache License 2.0. See [LICENSE](LICENSE).
