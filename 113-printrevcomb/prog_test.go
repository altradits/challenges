package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
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

func TestPrintRevComb(t *testing.T) {
	bin := buildBinary(t)
	cmd := exec.Command(bin)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		t.Fatalf("run: %v", err)
	}
	got := stdout.String()
	// Independent oracle: all distinct-digit triples, descending.
	var combos []string
	for a := 9; a >= 0; a-- {
		for b := a - 1; b >= 0; b-- {
			for c := b - 1; c >= 0; c-- {
				combos = append(combos, fmt.Sprintf("%d%d%d", a, b, c))
			}
		}
	}
	want := strings.Join(combos, ", ") + "\n"
	if got != want {
		t.Errorf("output mismatch: got %d bytes starting %q, want %d bytes starting %q",
			len(got), firstN(got, 40), len(want), firstN(want, 40))
	}
	if len(combos) != 120 {
		t.Fatalf("oracle built %d combos, want C(10,3)=120", len(combos))
	}
	if !strings.HasPrefix(got, "987, 986, 985") {
		t.Errorf("output should start with 987, 986, 985; got %q", firstN(got, 20))
	}
	if !strings.HasSuffix(got, "210\n") {
		t.Errorf("output should end with 210 + newline; got %q", lastN(got, 10))
	}
}

func firstN(s string, n int) string {
	if len(s) < n {
		return s
	}
	return s[:n]
}

func lastN(s string, n int) string {
	if len(s) < n {
		return s
	}
	return s[len(s)-n:]
}
