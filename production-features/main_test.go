package main_test

import (
	"context"
	"testing"
	"time"

	"github.com/k8s-manifest-kit/examples/internal/logger"
	"github.com/k8s-manifest-kit/examples/internal/testenv"
	example "github.com/k8s-manifest-kit/examples/production-features"
)

func TestRun(t *testing.T) {
	testenv.IsolateHelmEnv(t)

	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	ctx = logger.WithLogger(ctx, t)

	if err := example.Run(ctx); err != nil {
		t.Fatalf("Run() failed: %v", err)
	}
}
