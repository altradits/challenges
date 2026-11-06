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

func TestAddPrimeSum(t *testing.T) {
	bin := buildBinary(t)
	cases := []struct {
		name string
		args []string
		want string
	}{
		{"5", []string{"5"}, "10\n"},
		{"7", []string{"7"}, "17\n"},
		{"negative", []string{"-2"}, "0\n"},
		{"zero", []string{"0"}, "0\n"},
		{"no args", nil, "0\n"},
		{"two args", []string{"5", "7"}, "0\n"},
		{"non-numeric", []string{"a"}, "0\n"},
		{"two", []string{"2"}, "2\n"},
		{"ten", []string{"10"}, "17\n"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := run(t, bin, tc.args...); got != tc.want {
				t.Errorf("args %q: got %q, want %q", tc.args, got, tc.want)
			}
		})
	}
}
