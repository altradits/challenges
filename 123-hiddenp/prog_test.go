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

func TestHiddenP(t *testing.T) {
	bin := buildBinary(t)
	cases := []struct {
		name string
		args []string
		want string
	}{
		{"hidden", []string{"fgex.;", "tyf34gdgf;'ektufjhgdgex.;.;rtjynur6"}, "1\n"},
		{"hidden abc", []string{"abc", "2altrb53c.sse"}, "1\n"},
		{"not hidden", []string{"abc", "btarc"}, "0\n"},
		{"single D", []string{"DD", "DABC"}, "0\n"},
		{"no args", nil, ""},
		{"one arg", []string{"a"}, ""},
		{"empty needle", []string{"", "abc"}, "1\n"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := run(t, bin, tc.args...); got != tc.want {
				t.Errorf("args %q: got %q, want %q (1 = subsequence of s2)", tc.args, got, tc.want)
			}
		})
	}
}
