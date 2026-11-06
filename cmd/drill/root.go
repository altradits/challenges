package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// findRoot locates the repo root: the nearest parent (or override) holding
// go.mod and cmd/drill. DRILL_ROOT overrides discovery (used by tests).
func findRoot() (string, error) {
	if override := os.Getenv("DRILL_ROOT"); override != "" {
		return override, nil
	}
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if isRoot(dir) {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("not inside the challenges repo (no go.mod + cmd/drill found above %s)", dir)
		}
		dir = parent
	}
}

func isRoot(dir string) bool {
	for _, name := range []string{"go.mod", filepath.Join("cmd", "drill")} {
		if _, err := os.Stat(filepath.Join(dir, name)); err != nil {
			return false
		}
	}
	return true
}
