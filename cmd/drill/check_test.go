package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// fixtureRepo builds a tiny repo: go.mod + one challenge with a test.
// good controls whether the test passes.
func fixtureRepo(t *testing.T, good bool) string {
	t.Helper()
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "go.mod"), []byte("module fixture\n\ngo 1.26.10\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	writeChallenge(t, root, "06-functions", "# 06-functions — Greet\n\n**15 minutes**\n")
	body := `package piscine

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Go")
	if got != "Hello, Go!" {
		t.Fatalf("Greet = %q", got)
	}
}
`
	impl := "package piscine\n\nfunc Greet(name string) string { return \"Hello, \" + name + \"!\" }\n"
	if !good {
		impl = "package piscine\n\nfunc Greet(name string) string { return \"nope\" }\n"
	}
	if err := os.WriteFile(filepath.Join(root, "06-functions", "greet_test.go"), []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "06-functions", "greet.go"), []byte(impl), 0o644); err != nil {
		t.Fatal(err)
	}
	return root
}

func TestRunGoTestPassAndFail(t *testing.T) {
	out, pass := runGoTest(fixtureRepo(t, true), "06-functions")
	if !pass {
		t.Fatalf("want pass, got:\n%s", out)
	}
	out, pass = runGoTest(fixtureRepo(t, false), "06-functions")
	if pass {
		t.Fatal("want fail")
	}
	if !strings.Contains(out, "Greet = ") {
		t.Fatalf("failure message lost:\n%s", out)
	}
}

func TestCmdCheckRecords(t *testing.T) {
	root := fixtureRepo(t, true)
	if err := cmdStart(root, []string{"6"}); err != nil {
		t.Fatal(err)
	}
	if err := cmdCheck(root, []string{"6"}); err != nil {
		t.Fatal(err)
	}
	prog, err := loadProgress(root)
	if err != nil {
		t.Fatal(err)
	}
	r := prog.record(6)
	if r.Passes != 1 || r.Stage != 1 {
		t.Fatalf("record = %+v", r)
	}
	if prog.Pending != nil {
		t.Fatal("pending should clear after check")
	}
	// A failing check records a fail without advancing the stage.
	root2 := fixtureRepo(t, false)
	if err := cmdCheck(root2, []string{"6"}); err != nil {
		t.Fatal(err)
	}
	prog2, _ := loadProgress(root2)
	r2 := prog2.record(6)
	if r2.Fails != 1 || r2.Stage != 0 {
		t.Fatalf("record = %+v", r2)
	}
}

func TestCmdSolutionHiddenUntilPass(t *testing.T) {
	root := fixtureRepo(t, true)
	solDir := filepath.Join(root, "solutions", "06-functions")
	if err := os.MkdirAll(solDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(solDir, "solution.go"), []byte("package piscine\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := cmdSolution(root, []string{"6"}); err == nil {
		t.Fatal("solution should be hidden before first pass")
	}
	if err := cmdCheck(root, []string{"6"}); err != nil {
		t.Fatal(err)
	}
	if err := cmdSolution(root, []string{"6"}); err != nil {
		t.Fatalf("solution should unlock after pass: %v", err)
	}
}

func TestHasTests(t *testing.T) {
	root := fixtureRepo(t, true)
	if !hasTests(root, "06-functions") {
		t.Fatal("want hasTests=true")
	}
	writeChallenge(t, root, "07-empty", "# x\n")
	if hasTests(root, "07-empty") {
		t.Fatal("want hasTests=false")
	}
}

func TestNestedModules(t *testing.T) {
	root := t.TempDir()
	writeChallenge(t, root, "156-api-lesson", "# 156\n")
	sub := filepath.Join(root, "156-api-lesson", "01-ex")
	if err := os.MkdirAll(sub, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(sub, "go.mod"), []byte("module ex\n\ngo 1.26.10\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(sub, "main.go"), []byte("package main\n\nfunc main() {}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(sub, "main_test.go"), []byte("package main\n\nimport \"testing\"\n\nfunc TestX(t *testing.T) {}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if !hasTests(root, "156-api-lesson") {
		t.Fatal("want hasTests=true via nested module")
	}
	if got := nestedModules(root, "156-api-lesson"); len(got) != 1 {
		t.Fatalf("nestedModules = %v", got)
	}
	out, pass := runGoTest(root, "156-api-lesson")
	if !pass {
		t.Fatalf("nested run failed:\n%s", out)
	}
	if !strings.Contains(out, "=== 156-api-lesson/01-ex ===") {
		t.Fatalf("nested output missing header:\n%s", out)
	}
}
