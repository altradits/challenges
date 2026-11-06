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

func TestFPrime(t *testing.T) {
	bin := buildBinary(t)
	cases := []struct {
		name string
		args []string
		want string
	}{
		{"225225", []string{"225225"}, "3*3*5*5*7*11*13\n"},
		{"8333325", []string{"8333325"}, "3*3*5*5*7*11*13*37\n"},
		{"prime", []string{"9539"}, "9539\n"},
		{"prime 2", []string{"804577"}, "804577\n"},
		{"42", []string{"42"}, "2*3*7\n"},
		{"non-numeric", []string{"a"}, ""},
		{"zero", []string{"0"}, ""},
		{"one", []string{"1"}, ""},
		{"no args", nil, ""},
		{"two args", []string{"4", "6"}, ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := run(t, bin, tc.args...); got != tc.want {
				t.Errorf("args %q: got %q, want %q", tc.args, got, tc.want)
			}
		})
	}
}
