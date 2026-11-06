package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// goTestTimeout bounds one challenge check so a hanging solution cannot stall
// the drill. Unit-testable via runGoTest.
const goTestTimeout = 120 * time.Second

func cmdCheck(root string, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: drill check NN")
	}
	cat, err := loadCatalog(root)
	if err != nil {
		return err
	}
	c, err := findChallenge(cat, args[0])
	if err != nil {
		return err
	}
	prog, err := loadProgress(root)
	if err != nil {
		return err
	}
	at := now()
	elapsed := pendingElapsed(prog, c.Num, at)

	if c.Manual {
		return checkManual(root, prog, c, elapsed, at)
	}
	if !hasTests(root, c.Dir) {
		return fmt.Errorf("challenge %d has no tests yet and is not manual — nothing to check", c.Num)
	}
	out, pass := runGoTest(root, c.Dir)
	fmt.Print(out)
	rec := prog.recordAttempt(c.Num, pass, elapsed, at)
	if err := saveProgress(root, prog); err != nil {
		return err
	}
	if pass {
		stage := rec.Stage
		if stage >= len(intervals) {
			fmt.Printf("\nPASS in %s — MASTERED (4/4 spaced rounds done).\n", formatDur(elapsed))
		} else {
			due, _ := dueAt(rec)
			fmt.Printf("\nPASS in %s — next due %s (stage %d/%d).\n",
				formatDur(elapsed), due.Format("Mon 2006-01-02"), stage, len(intervals))
		}
		return nil
	}
	fmt.Printf("\nFAIL in %s — read the failure above, fix from memory, check again.\n", formatDur(elapsed))
	return nil
}

// pendingElapsed is the solve time measured from `drill start`.
func pendingElapsed(prog *Progress, num int, at time.Time) time.Duration {
	if prog.Pending != nil && prog.Pending.Num == num {
		return at.Sub(prog.Pending.At)
	}
	return 0
}

func hasTests(root, dir string) bool {
	matches, _ := filepath.Glob(filepath.Join(root, dir, "*_test.go"))
	if len(matches) > 0 {
		return true
	}
	return len(nestedModules(root, dir)) > 0
}

// nestedModules lists challenge subdirectories that are their own Go modules
// with tests (the 156-api-lesson sub-exercises).
func nestedModules(root, dir string) []string {
	var out []string
	mods, _ := filepath.Glob(filepath.Join(root, dir, "*", "go.mod"))
	for _, g := range mods {
		sub := filepath.Dir(g)
		if tests, _ := filepath.Glob(filepath.Join(sub, "*_test.go")); len(tests) > 0 {
			out = append(out, sub)
		}
	}
	return out
}

// runGoTest runs the hidden tests for one challenge and reports pass/fail.
// The full test output is returned so failure messages reach the student.
// Challenges with nested test modules (156) run each module in turn.
func runGoTest(root, dir string) (string, bool) {
	if subs := nestedModules(root, dir); len(subs) > 0 {
		var sb strings.Builder
		allPass := true
		for _, sub := range subs {
			rel, _ := filepath.Rel(root, sub)
			out, pass := runGoTestIn(sub, "./...")
			fmt.Fprintf(&sb, "=== %s ===\n%s", rel, out)
			allPass = allPass && pass
		}
		return sb.String(), allPass
	}
	return runGoTestIn(root, "./"+dir+"/")
}

// runGoTestIn runs go test for pkg with workdir as the module root.
func runGoTestIn(workdir, pkg string) (string, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), goTestTimeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, "go", "test", "-count=1", pkg)
	cmd.Dir = workdir
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	err := cmd.Run()
	out := buf.String()
	if ctx.Err() == context.DeadlineExceeded {
		return out + "\nTIMEOUT: tests took over " + formatDur(goTestTimeout) + " — check for deadlocks.\n", false
	}
	return out, err == nil
}

// checkManual attests a manual challenge (00, 22, 83, 158): no hidden tests
// exist, so the student's honest yes is the check.
func checkManual(root string, prog *Progress, c Challenge, elapsed time.Duration, at time.Time) error {
	fmt.Printf("=== %s ===\n\n", c.Title)
	fmt.Println("Manual challenge: there are no hidden tests to run.")
	fmt.Println("Answer honestly — you are only cheating your own memory.")
	fmt.Print("Did you complete every step in the Task section? [y/N] ")
	line, err := readLine()
	if err != nil {
		return err
	}
	pass := strings.ToLower(strings.TrimSpace(line)) == "y" || strings.ToLower(strings.TrimSpace(line)) == "yes"
	rec := prog.recordAttempt(c.Num, pass, elapsed, at)
	if err := saveProgress(root, prog); err != nil {
		return err
	}
	if pass {
		fmt.Printf("\nRecorded PASS (stage %d/%d). Well done.\n", rec.Stage, len(intervals))
		return nil
	}
	fmt.Println("\nRecorded as not done. Come back when it is really finished.")
	return nil
}

var stdinReader = bufio.NewReader(os.Stdin)

func readLine() (string, error) {
	return stdinReader.ReadString('\n')
}
