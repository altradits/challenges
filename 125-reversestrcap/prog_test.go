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

func TestReverseStrCap(t *testing.T) {
	bin := buildBinary(t)
	cases := []struct {
		name string
		args []string
		want string
	}{
		{"basic", []string{"First SMALL TesT"}, "firsT smalL tesT\n"},
		{"three args", []string{"SEconD Test IS a LItTLE EasIEr", "bEwaRe IT'S NoT HARd WhEN ", " Go a dernier 0123456789 for the road e"},
			"seconD tesT iS A littlE easieR\nbewarE it'S noT harD wheN \n gO A dernieR 0123456789 foR thE roaD E\n"},
		{"no args", nil, ""},
		{"empty", []string{""}, "\n"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := run(t, bin, tc.args...); got != tc.want {
				t.Errorf("args %q: got %q, want %q", tc.args, got, tc.want)
			}
		})
	}
}
