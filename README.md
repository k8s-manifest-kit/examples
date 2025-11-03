# k8s-manifest-kit Examples

Comprehensive examples demonstrating the capabilities of the k8s-manifest-kit library for rendering, filtering, and transforming Kubernetes manifests.

Part of the [k8s-manifest-kit](https://github.com/k8s-manifest-kit) organization.

## 📚 Examples

### 1. [Quickstart](./quickstart) - Getting Started

The simplest way to get started with k8s-manifest-kit. Renders a Helm chart from an OCI registry.

```bash
go run ./quickstart
```

**Key concepts:**
- Creating an Engine with a single Helm renderer
- Basic manifest rendering
- Inspecting rendered objects

---

### 2. [Filtering & Transformation](./filtering-transformation) - Filters and Transformers

Demonstrates applying filters and transformers to rendered manifests.

```bash
go run ./filtering-transformation
```

**Key concepts:**
- Using jq-based filters to select specific resources
- Applying label transformers to add common labels  
- Renderer-level vs engine-level filtering/transformation

---

### 3. [Multiple Sources](./multiple-sources) - Combining Renderers

Shows how to combine manifests from multiple sources (Helm, YAML, Kustomize) in a single engine.

```bash
go run ./multiple-sources
```

**Key concepts:**
- Creating an Engine with multiple renderers
- Rendering from different source types (Helm, YAML, Kustomize)
- Merging manifests from multiple sources

---

### 4. [Production Features](./production-features) - Advanced Features

Demonstrates production-ready features like caching, metrics, and error handling.

```bash
go run ./production-features
```

**Key concepts:**
- Enabling caching for performance optimization
- Using metrics to monitor render operations
- Proper error handling patterns
- Source annotations for tracking manifest origins

---

### 5. [Real World](./real-world) - Complete Scenario

A comprehensive example showing a real-world scenario with multiple renderers, filters, transformers, and runtime value overrides.

```bash
go run ./real-world
```

**Key concepts:**
- Complex multi-source rendering
- Using multiple filters and transformers
- Runtime value overrides
- Type-safe Kubernetes object manipulation

---

## 🚀 Quick Start

### Prerequisites

- Go 1.23.8 or later
- Network access (for Helm OCI registry examples)

### Installation

Clone this repository (assuming you have the k8s-manifest-kit repositories):

```bash
git clone https://github.com/k8s-manifest-kit/examples.git
cd examples
```

### Running Examples

Run any example directly:

```bash
# Run a specific example
go run ./quickstart
go run ./filtering-transformation
go run ./multiple-sources
go run ./production-features
go run ./real-world

# Run tests for all examples
go test -v ./...

# Run tests for a specific example
go test -v ./quickstart
```

## 📖 Documentation

### Core Concepts

The examples demonstrate three key concepts from k8s-manifest-kit:

1. **Renderers** - Convert manifest sources (Helm, Kustomize, YAML) to Kubernetes objects
2. **Filters** - Select which objects to include in the output
3. **Transformers** - Modify objects before they're returned

### Import Structure

The k8s-manifest-kit project is split into multiple repositories:

```go
import (
    // Core engine
    engine "github.com/k8s-manifest-kit/engine/pkg"
    
    // Renderers
    helm "github.com/k8s-manifest-kit/renderer-helm/pkg"
    yaml "github.com/k8s-manifest-kit/renderer-yaml/pkg"
    kustomize "github.com/k8s-manifest-kit/renderer-kustomize/pkg"
    
    // Filters and transformers
    "github.com/k8s-manifest-kit/engine/pkg/filter/jq"
    "github.com/k8s-manifest-kit/engine/pkg/transformer/meta/labels"
    
    // Utilities
    "github.com/k8s-manifest-kit/pkg/util/cache"
    "github.com/k8s-manifest-kit/pkg/util/metrics"
)
```

### Example Structure

Each example follows a consistent pattern:

- **main.go** - Runnable example demonstrating specific features
- **main_test.go** - Test coverage validating the example works correctly
- **Comments** - Inline explanations of key concepts

## 🔧 Development

### Building

```bash
# Build all examples
go build ./...
```

### Testing

```bash
# Test all examples
make test

# Test specific example
go test -v ./quickstart
```

### Linting

```bash
# Run linter
make lint

# Auto-fix linting issues
make lint/fix
```

### Formatting

```bash
# Format code
make fmt
```

## 📚 Related Documentation

- **[k8s-manifest-kit/engine](https://github.com/k8s-manifest-kit/engine)** - Core engine and types
- **[k8s-manifest-kit/renderer-helm](https://github.com/k8s-manifest-kit/renderer-helm)** - Helm renderer
- **[k8s-manifest-kit/renderer-yaml](https://github.com/k8s-manifest-kit/renderer-yaml)** - YAML renderer
- **[k8s-manifest-kit/renderer-kustomize](https://github.com/k8s-manifest-kit/renderer-kustomize)** - Kustomize renderer
- **[k8s-manifest-kit/pkg](https://github.com/k8s-manifest-kit/pkg)** - Shared utilities

## 📋 Example Features Matrix

| Example | Helm | YAML | Kustomize | Filters | Transformers | Cache | Metrics |
|---------|------|------|-----------|---------|--------------|-------|---------|
| quickstart | ✅ | - | - | - | - | - | - |
| filtering-transformation | ✅ | - | - | ✅ | ✅ | - | - |
| multiple-sources | ✅ | ✅ | ✅ | - | - | - | - |
| production-features | - | ✅ | - | - | - | ✅ | ✅ |
| real-world | ✅ | ✅ | ✅ | ✅ | ✅ | - | - |

## 🤝 Contributing

Contributions are welcome! When adding new examples:

1. Create a new directory under `examples/`
2. Add `main.go` with a runnable example
3. Add `main_test.go` with test coverage
4. Update this README with the example description
5. Update `CLAUDE.md` with technical details

## 📝 License

Apache License 2.0 - See [LICENSE](LICENSE) for details.

## 🔗 Links

- [k8s-manifest-kit Organization](https://github.com/k8s-manifest-kit)
- [Report Issues](https://github.com/k8s-manifest-kit/examples/issues)
- [Contributing Guidelines](https://github.com/k8s-manifest-kit/docs/blob/main/CONTRIBUTING.md)
