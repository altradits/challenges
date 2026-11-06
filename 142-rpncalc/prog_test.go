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

func TestRPNCalc(t *testing.T) {
	bin := buildBinary(t)
	cases := []struct {
		name string
		args []string
		want string
	}{
		{"basic", []string{"1 2 * 3 * 4 +"}, "10\n"},
		{"too many operands", []string{"1 2 3 4 +"}, "Error\n"},
		{"no args", nil, "Error\n"},
		{"extra spaces", []string{"     1      3 * 2 -"}, "1\n"},
		{"bad token", []string{"     1      3 * ksd 2 -"}, "Error\n"},
		{"div", []string{"8 3 /"}, "2\n"},
		{"mod", []string{"8 3 %"}, "2\n"},
		{"sub negative", []string{"3 8 -"}, "-5\n"},
		{"empty", []string{""}, "Error\n"},
		{"two args", []string{"1 2 +", "3"}, "Error\n"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := run(t, bin, tc.args...); got != tc.want {
				t.Errorf("args %q: got %q, want %q", tc.args, got, tc.want)
			}
		})
	}
}
