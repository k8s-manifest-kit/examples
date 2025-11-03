package main_test

import (
	"context"
	"testing"
	"time"

	example "github.com/k8s-manifest-kit/examples/filtering-transformation"
	"github.com/k8s-manifest-kit/examples/internal/logger"
)

func TestRun(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	ctx = logger.WithLogger(ctx, t)

	if err := example.Run(ctx); err != nil {
		t.Fatalf("Run() failed: %v", err)
	}
}
