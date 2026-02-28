package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	engine "github.com/k8s-manifest-kit/engine/pkg"
	"github.com/k8s-manifest-kit/engine/pkg/postrenderer"
	"github.com/k8s-manifest-kit/engine/pkg/render"
	"github.com/k8s-manifest-kit/engine/pkg/source"
	"github.com/k8s-manifest-kit/engine/pkg/transformer/meta/labels"
	"github.com/k8s-manifest-kit/engine/pkg/types"
	"github.com/k8s-manifest-kit/examples/internal/logger"
	helm "github.com/k8s-manifest-kit/renderer-helm/pkg"
)

func main() {
	ctx := logger.WithLogger(context.Background(), &logger.StdoutLogger{})
	if err := Run(ctx); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func Run(ctx context.Context) error {
	l := logger.FromContext(ctx)
	l.Log("=== Pipeline Hooks Example ===")
	l.Log("Demonstrates: SourceSelector, PostRenderer, render.WithValues, postrenderer.ApplyOrder")
	l.Log()

	// SourceSelector: skip internal-only charts in production.
	// source.Selector[helm.Source] provides type-safe access to Helm-specific fields;
	// non-Helm sources are automatically passed through.
	env := os.Getenv("DEPLOY_ENV")
	if env == "" {
		env = "staging"
	}

	chartSelector := source.Selector[helm.Source](
		func(_ context.Context, s helm.Source) (bool, error) {
			if env == "production" && s.ReleaseName == "internal-tools" {
				return false, nil
			}

			return true, nil
		},
	)

	// Source-level PostRenderer: add per-chart metadata
	addChartOriginLabel := func(
		_ context.Context, objects []unstructured.Unstructured,
	) ([]unstructured.Unstructured, error) {
		for i := range objects {
			lbls := objects[i].GetLabels()
			if lbls == nil {
				lbls = make(map[string]string)
			}

			lbls["pipeline-hooks-example/source-hook"] = "true"
			objects[i].SetLabels(lbls)
		}

		return objects, nil
	}

	helmRenderer, err := helm.New(
		[]helm.Source{
			{
				Chart:       "oci://registry-1.docker.io/bitnamicharts/nginx",
				ReleaseName: "web-frontend",
				Values: helm.Values(map[string]any{
					"replicaCount": 2,
				}),
				PostRenderers: []types.PostRenderer{addChartOriginLabel},
			},
			{
				Chart:       "oci://registry-1.docker.io/bitnamicharts/nginx",
				ReleaseName: "internal-tools",
				Values: helm.Values(map[string]any{
					"replicaCount": 1,
				}),
			},
		},
		helm.WithSourceSelector(chartSelector),
		helm.WithPostRenderer(func(
			_ context.Context, objects []unstructured.Unstructured,
		) ([]unstructured.Unstructured, error) {
			l.Logf("  [renderer PostRenderer] processing %d objects from all sources\n", len(objects))

			return objects, nil
		}),
	)
	if err != nil {
		return fmt.Errorf("failed to create helm renderer: %w", err)
	}

	// Engine: global PostRenderer sorts resources into apply-order
	e, err := engine.New(
		engine.WithRenderer(helmRenderer),
		engine.WithTransformer(labels.Set(map[string]string{
			"managed-by": "pipeline-hooks-example",
		})),
		engine.WithPostRenderer(postrenderer.ApplyOrder()),
	)
	if err != nil {
		return fmt.Errorf("failed to create engine: %w", err)
	}

	// render.WithValues: provide render-time values merged with Source-level values
	objects, err := e.Render(ctx,
		render.WithValues(types.Values{
			"commonAnnotations": map[string]any{
				"deploy-env": env,
			},
		}),
	)
	if err != nil {
		return fmt.Errorf("failed to render: %w", err)
	}

	l.Logf("\nRendered %d objects (DEPLOY_ENV=%s)\n\n", len(objects), env)

	l.Log("Pipeline hooks exercised:")
	l.Log("  1. source.Selector[helm.Source] — skips 'internal-tools' when DEPLOY_ENV=production")
	l.Log("  2. Source.PostRenderers          — adds 'source-hook=true' label per chart")
	l.Log("  3. helm.WithPostRenderer         — renderer-level batch logging")
	l.Log("  4. engine.WithTransformer        — global label injection")
	l.Log("  5. postrenderer.ApplyOrder()     — global resource sort (Namespaces first, Webhooks last)")
	l.Log("  6. render.WithValues             — render-time values merged with Source-level")
	l.Log()

	l.Log("Resources in apply-order:")

	for i, obj := range objects {
		sourceHook := ""
		if obj.GetLabels()["pipeline-hooks-example/source-hook"] == "true" {
			sourceHook = " [source-hook]"
		}

		l.Logf("  %2d. %-45s %s/%s%s\n",
			i+1,
			obj.GetKind(),
			obj.GetNamespace(),
			obj.GetName(),
			sourceHook,
		)
	}

	return nil
}
