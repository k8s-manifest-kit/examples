# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with examples in this repository.

## Project Overview

This repository contains runnable examples demonstrating the k8s-manifest-kit library capabilities. Each example is self-contained with its own main.go and test file showing real-world usage patterns.

## Repository Structure

```
examples/
├── quickstart/          # Basic Helm chart rendering
├── filtering-transformation/  # Applying filters and transformers
├── multiple-sources/    # Using multiple renderers (Helm + YAML + Kustomize)
├── production-features/ # Cache, metrics, error handling
├── real-world/          # Complete real-world scenario
└── internal/logger/     # Shared logger utility for examples
```

## Examples Overview

### 1. quickstart/ - Getting Started
The simplest way to get started with k8s-manifest-kit. Renders a Helm chart from OCI registry.

**Key concepts:**
- Creating an Engine with `engine.Helm()`
- Basic Helm chart rendering
- Inspecting rendered objects

**Run:**
```bash
go run ./quickstart
```

### 2. filtering-transformation/ - Filters and Transformers
Demonstrates applying filters and transformers to rendered manifests.

**Key concepts:**
- Using jq-based filters to select specific resources
- Applying label transformers to add common labels
- Renderer-level vs engine-level filtering/transformation

**Run:**
```bash
go run ./filtering-transformation
```

### 3. multiple-sources/ - Multiple Renderers
Shows how to combine manifests from multiple sources (Helm, YAML, Kustomize) in a single engine.

**Key concepts:**
- Creating an Engine with multiple renderers
- Using `engine.New()` with `engine.WithRenderer()` options
- Rendering from different source types
- Merging manifests from multiple sources

**Run:**
```bash
go run ./multiple-sources
```

### 4. production-features/ - Advanced Features
Demonstrates production-ready features like caching, metrics, and error handling.

**Key concepts:**
- Enabling caching for performance
- Using metrics to monitor render operations
- Proper error handling patterns
- Source annotations for tracking manifest origins

**Run:**
```bash
go run ./production-features
```

### 5. real-world/ - Complete Scenario
A comprehensive example showing a real-world scenario with multiple renderers, filters, transformers, and runtime value overrides.

**Key concepts:**
- Complex multi-source rendering
- Using multiple filters and transformers
- Runtime value overrides with `engine.WithValues()`
- Type-safe K8s object manipulation

**Run:**
```bash
go run ./real-world
```

## Key Concepts Across Examples

### Import Paths
All examples use the split repository structure:

- **Engine**: `github.com/k8s-manifest-kit/engine/pkg`
- **Renderers**: `github.com/k8s-manifest-kit/renderer-{helm,yaml,kustomize}/pkg`
- **Filters**: `github.com/k8s-manifest-kit/engine/pkg/filter`
- **Transformers**: `github.com/k8s-manifest-kit/engine/pkg/transformer`
- **Utilities**: `github.com/k8s-manifest-kit/pkg/util`

### Common Patterns

**1. Creating a single-renderer engine:**
```go
e, err := engine.Helm(helm.Source{...})
```

**2. Creating a multi-renderer engine:**
```go
e, err := engine.New(
    engine.WithRenderer(helmRenderer),
    engine.WithRenderer(yamlRenderer),
)
```

**3. Applying filters at renderer level:**
```go
renderer, err := helm.New([]helm.Source{...},
    helm.WithFilter(gvk.Filter(...)),
)
```

**4. Applying transformers at engine level:**
```go
e, err := engine.New(
    engine.WithRenderer(renderer),
    engine.WithTransformer(labels.Set(...)),
)
```

**5. Runtime value overrides:**
```go
objects, err := e.Render(ctx, engine.WithValues(map[string]any{
    "replicaCount": 5,
}))
```

## Development

### Running Examples
```bash
# Run all examples
go test -v ./...

# Run specific example
go run ./quickstart

# Test specific example
go test -v ./quickstart
```

### Testing
Each example includes a `main_test.go` that:
- Tests that the example runs without errors
- Validates basic output expectations
- Can be used as additional usage documentation

### Formatting and Linting
```bash
# Format code
make fmt

# Run linter
make lint

# Auto-fix linting issues
make lint/fix
```

### Adding New Examples

1. Create new directory under `examples/`
2. Create `main.go` with runnable example
3. Create `main_test.go` with test coverage
4. Update this CLAUDE.md with example description
5. Update README.md with example in the list

**Example structure:**
```go
package main

import (
    "context"
    "log"
    
    "github.com/k8s-manifest-kit/examples/internal/logger"
    engine "github.com/k8s-manifest-kit/engine/pkg"
    // ... other imports
)

func main() {
    ctx := logger.WithLogger(context.Background(), &logger.StdoutLogger{})
    if err := Run(ctx); err != nil {
        log.Fatalf("Error: %v", err)
    }
}

func Run(ctx context.Context) error {
    l := logger.FromContext(ctx)
    // ... example implementation
    return nil
}
```

## Testing Philosophy

Examples should:
- Be self-contained and runnable
- Demonstrate real-world usage patterns
- Include comments explaining key concepts
- Have test coverage validating basic functionality
- Use the shared logger for consistent output

## Common Issues

**Network Access:**
Some examples (Helm) require network access to fetch charts from OCI registries. Tests may fail in offline environments.

**Go Module Dependencies:**
This repository uses local replace directives in `go.mod` to reference sibling repositories during development. For released versions, these should point to proper version tags.

## Related Documentation

- [k8s-manifest-kit/engine](https://github.com/k8s-manifest-kit/engine) - Core engine documentation
- [k8s-manifest-kit/renderer-helm](https://github.com/k8s-manifest-kit/renderer-helm) - Helm renderer
- [k8s-manifest-kit/renderer-yaml](https://github.com/k8s-manifest-kit/renderer-yaml) - YAML renderer
- [k8s-manifest-kit/renderer-kustomize](https://github.com/k8s-manifest-kit/renderer-kustomize) - Kustomize renderer

