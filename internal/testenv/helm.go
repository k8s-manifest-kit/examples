package testenv

import (
	"path/filepath"
	"testing"
)

// IsolateHelmEnv gives each test its own Helm cache and config paths
// so OCI-backed example runs do not race through shared global state.
func IsolateHelmEnv(t *testing.T) {
	t.Helper()

	base := t.TempDir()

	t.Setenv("HELM_CACHE_HOME", filepath.Join(base, "cache"))
	t.Setenv("HELM_CONFIG_HOME", filepath.Join(base, "config"))
	t.Setenv("HELM_DATA_HOME", filepath.Join(base, "data"))
	t.Setenv("HELM_REPOSITORY_CACHE", filepath.Join(base, "repository"))
	t.Setenv("HELM_REPOSITORY_CONFIG", filepath.Join(base, "config", "repositories.yaml"))
	t.Setenv("HELM_REGISTRY_CONFIG", filepath.Join(base, "config", "registry.json"))
}
