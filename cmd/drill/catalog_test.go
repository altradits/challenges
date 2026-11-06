package main

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestParseTitle(t *testing.T) {
	if got := parseTitle("# 06-functions — X\nbody\n"); got != "06-functions — X" {
		t.Fatalf("parseTitle = %q", got)
	}
	if got := parseTitle("no header\n"); got != "Untitled challenge" {
		t.Fatalf("parseTitle fallback = %q", got)
	}
}

func TestParseLimitFromReadme(t *testing.T) {
	got := parseLimit("## Time limit\n\n**25 minutes** (round 1).\n", 145)
	if got != 25*time.Minute {
		t.Fatalf("parseLimit = %v, want 25m", got)
	}
}

func TestParseLimitFallback(t *testing.T) {
	cases := []struct {
		num  int
		want time.Duration
	}{
		{1, 5 * time.Minute},
		{6, 15 * time.Minute},
		{28, 10 * time.Minute},
		{96, 20 * time.Minute},
		{142, 30 * time.Minute},
		{145, 25 * time.Minute},
		{153, 120 * time.Minute},
		{156, 45 * time.Minute},
		{157, 240 * time.Minute},
		{160, 20 * time.Minute},
		{180, 30 * time.Minute},
	}
	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			if got := parseLimit("no limit here", tc.num); got != tc.want {
				t.Fatalf("parseLimit(%d) = %v, want %v", tc.num, got, tc.want)
			}
		})
	}
}

func TestIsManualReadme(t *testing.T) {
	if !isManualReadme("This is a manual challenge — no tests.", "00-x") {
		t.Fatal("want manual=true")
	}
	if isManualReadme("## How to check\n\ngo test ./01-x/\n", "01-x") {
		t.Fatal("want manual=false")
	}
}

func writeChallenge(t *testing.T, root, dir, readme string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Join(root, dir), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, dir, "README.md"), []byte(readme), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestLoadCatalogAndFind(t *testing.T) {
	root := t.TempDir()
	writeChallenge(t, root, "06-functions", "# 06-functions — Greet\n\n**15 minutes**\n")
	writeChallenge(t, root, "158-bitcoin-oss", "# 158 — OSS\n\nmanual challenge\n")
	if err := os.MkdirAll(filepath.Join(root, "not-a-challenge"), 0o755); err != nil {
		t.Fatal(err)
	}

	cat, err := loadCatalog(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(cat) != 2 || cat[0].Num != 6 || cat[1].Num != 158 {
		t.Fatalf("catalog = %+v", cat)
	}
	if cat[1].Manual != true || cat[0].Manual != false {
		t.Fatalf("manual flags = %v %v", cat[0].Manual, cat[1].Manual)
	}

	for _, arg := range []string{"6", "06", "06-functions"} {
		c, err := findChallenge(cat, arg)
		if err != nil || c.Num != 6 {
			t.Fatalf("findChallenge(%q) = %+v, %v", arg, c, err)
		}
	}
	if _, err := findChallenge(cat, "nope"); err == nil {
		t.Fatal("findChallenge(nope) should fail")
	}
	if _, err := findChallenge(cat, "7"); err == nil {
		t.Fatal("findChallenge(7) should fail when missing")
	}
}
