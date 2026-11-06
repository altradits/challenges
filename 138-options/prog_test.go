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

func TestOptions(t *testing.T) {
	bin := buildBinary(t)
	cases := []struct {
		name string
		args []string
		want string
	}{
		{"no args", nil, "options: abcdefghijklmnopqrstuvwxyz\n"},
		{"abc ijk", []string{"-abc", "-ijk"}, "00000000 00000000 00000111 00000111\n"},
		{"z", []string{"-z"}, "00000010 00000000 00000000 00000000\n"},
		{"h first wins", []string{"-abc", "-hijk"}, "options: abcdefghijklmnopqrstuvwxyz\n"},
		{"h alone", []string{"-h"}, "options: abcdefghijklmnopqrstuvwxyz\n"},
		{"h not first", []string{"-zh"}, "00000010 00000000 00000000 10000000\n"},
		{"h second arg", []string{"-z", "-h"}, "options: abcdefghijklmnopqrstuvwxyz\n"},
		{"many h", []string{"-hhhhhh"}, "options: abcdefghijklmnopqrstuvwxyz\n"},
		{"repeated e", []string{"-eeeeee"}, "00000000 00000000 00000000 00010000\n"},
		{"bad char", []string{"-%"}, "Invalid Option\n"},
		{"lone dash", []string{"-"}, "Invalid Option\n"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := run(t, bin, tc.args...); got != tc.want {
				t.Errorf("args %q: got %q, want %q", tc.args, got, tc.want)
			}
		})
	}
}
