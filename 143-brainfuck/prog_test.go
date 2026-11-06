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

func TestBrainfuck(t *testing.T) {
	bin := buildBinary(t)
	cases := []struct {
		name string
		args []string
		want string
	}{
		{"hello", []string{"++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>."}, "Hello World!\n"},
		{"hi", []string{"+++++[>++++[>++++H>+++++i<<-]>>>++\\n<<<<-]>>--------.>+++++.>."}, "Hi\n"},
		{"abc", []string{"++++++++++[>++++++++++>++++++++++>++++++++++<<<-]>---.>--.>-.>++++++++++."}, "abc\n"},
		{"no args", nil, ""},
		{"empty program", []string{""}, ""},
		{"comments ignored", []string{"this is all comments, no ops"}, ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := run(t, bin, tc.args...); got != tc.want {
				t.Errorf("program %q: got %q, want %q", tc.args, got, tc.want)
			}
		})
	}
}
