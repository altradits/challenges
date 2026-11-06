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

func TestFindPairs(t *testing.T) {
	bin := buildBinary(t)
	cases := []struct {
		name string
		args []string
		want string
	}{
		{"pairs", []string{"[1, 2, 3, 4, 5]", "6"}, "Pairs with sum 6: [[0 4] [1 3]]\n"},
		{"negatives", []string{"[-1, 2, -3, 4, -5]", "1"}, "Pairs with sum 1: [[0 1] [2 3]]\n"},
		{"none", []string{"[1, 2, 3, 4, 5]", "10"}, "No pairs found.\n"},
		{"neg target", []string{"[-1, -2, -3, -4, -5]", "-5"}, "Pairs with sum -5: [[0 3] [1 2]]\n"},
		{"bad target", []string{"[1, 2, 3, 4, 20, -4, 5]", "2 5"}, "Invalid target sum.\n"},
		{"bad number", []string{"[1, 2, 3, 4, 20, p, 5]", "5"}, "Invalid number: p\n"},
		{"missing bracket", []string{"[1, 2, 3, 4", "5"}, "Invalid input.\n"},
		{"no brackets", []string{"1, 2, 3, 4", "5"}, "Invalid input.\n"},
		{"no args", nil, "Invalid input.\n"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := run(t, bin, tc.args...); got != tc.want {
				t.Errorf("args %q: got %q, want %q", tc.args, got, tc.want)
			}
		})
	}
}
