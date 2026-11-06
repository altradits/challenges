package main

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"testing"
)

func buildBinary(t *testing.T) string {
	t.Helper()
	bin := filepath.Join(t.TempDir(), "prog")
	cmd := exec.Command("go", "build", "-o", bin, ".")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("go build failed: %v\n%s", err, out)
	}
	return bin
}

func run(t *testing.T, bin string, args ...string) string {
	t.Helper()
	cmd := exec.Command(bin, args...)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		t.Fatalf("run %v: %v", args, err)
	}
	return stdout.String()
}

func TestCleanStr(t *testing.T) {
	bin := buildBinary(t)
	cases := []struct {
		name string
		args []string
		want string
	}{
		{"plain", []string{"you see it's easy to display the same thing"}, "you see it's easy to display the same thing\n"},
		{"messy", []string{" only    it's  harder   "}, "only it's harder\n"},
		{"two args", []string{" how funny", "Did you   hear Mathilde ?"}, "\n"},
		{"empty", []string{""}, "\n"},
		{"no args", nil, "\n"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := run(t, bin, tc.args...); got != tc.want {
				t.Errorf("args %q: got %q, want %q", tc.args, got, tc.want)
			}
		})
	}
}
