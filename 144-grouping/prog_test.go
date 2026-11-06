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

func TestGrouping(t *testing.T) {
	bin := buildBinary(t)
	cases := []struct {
		name string
		args []string
		want string
	}{
		{"single group", []string{"(a)", "I'm heavy, jumpsuit is on steady, Lighter when I'm lower, higher when I'm heavy"}, "1: heavy\n2: steady\n3: heavy\n"},
		{"alternation", []string{"(e|n)", "I currently have 4 windows opened up… and I don't know why."}, "1: currently\n2: currently\n3: have\n4: windows\n5: opened\n6: opened\n7: and\n8: don't\n9: know\n"},
		{"hi", []string{"(hi)", "He swore he just saw his sushi move."}, "1: his\n2: sushi\n"},
		{"empty text", []string{"(s)", ""}, ""},
		{"invalid pattern", []string{"i", "Something in the air"}, ""},
		{"no args", nil, ""},
		{"no match", []string{"(z)", "abc def"}, ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := run(t, bin, tc.args...); got != tc.want {
				t.Errorf("args %q: got %q, want %q", tc.args, got, tc.want)
			}
		})
	}
}
