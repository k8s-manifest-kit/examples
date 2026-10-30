//go:build integration

// The Run smoke test renders live Helm charts from a public OCI registry and a
// public chart repository, so it runs only in the integration suite
// (make test/integration), never in the offline unit suite (make test).
package main_test

import (
	"context"
	"testing"
	"time"

	"github.com/k8s-manifest-kit/examples/internal/logger"
	"github.com/k8s-manifest-kit/examples/internal/testenv"
	example "github.com/k8s-manifest-kit/examples/multiple-sources"
)

func TestRun(t *testing.T) {
	testenv.IsolateHelmEnv(t)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	ctx = logger.WithLogger(ctx, t)

	if err := example.Run(ctx); err != nil {
		t.Fatalf("Run() failed: %v", err)
	}
}
