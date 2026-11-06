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

func TestRomanNumbers(t *testing.T) {
	bin := buildBinary(t)
	cases := []struct {
		name string
		args []string
		want string
	}{
		{"non-numeric", []string{"hello"}, "ERROR: cannot convert to roman digit\n"},
		{"123", []string{"123"}, "C+X+X+I+I+I\nCXXIII\n"},
		{"999", []string{"999"}, "(M-C)+(C-X)+(X-I)\nCMXCIX\n"},
		{"3999", []string{"3999"}, "M+M+M+(M-C)+(C-X)+(X-I)\nMMMCMXCIX\n"},
		{"too big", []string{"4000"}, "ERROR: cannot convert to roman digit\n"},
		{"zero", []string{"0"}, "ERROR: cannot convert to roman digit\n"},
		{"no args", nil, "ERROR: cannot convert to roman digit\n"},
		{"one", []string{"1"}, "I\nI\n"},
		{"four", []string{"4"}, "(V-I)\nIV\n"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := run(t, bin, tc.args...); got != tc.want {
				t.Errorf("args %q: got %q, want %q", tc.args, got, tc.want)
			}
		})
	}
}
