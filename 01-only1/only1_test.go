package main

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"testing"
)

// buildBinary compiles the student's program in this directory.
func buildBinary(t *testing.T) string {
	t.Helper()
	bin := filepath.Join(t.TempDir(), "only1")
	cmd := exec.Command("go", "build", "-o", bin, ".")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("go build failed: %v\n%s", err, out)
	}
	return bin
}

func TestOnly1(t *testing.T) {
	bin := buildBinary(t)
	argSets := [][]string{
		{},
		{"a", "b"},
		{"a", "b", "c"},
	}
	for _, args := range argSets {
		cmd := exec.Command(bin, args...)
		var stdout bytes.Buffer
		cmd.Stdout = &stdout
		if err := cmd.Run(); err != nil {
			t.Fatalf("run %v: program exited with error: %v", args, err)
		}
		if got, want := stdout.String(), "1\n"; got != want {
			t.Errorf("args %v: output = %q, want %q (exactly one '1' plus newline)", args, got, want)
		}
	}
}
