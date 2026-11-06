package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// now is time.Now, replaceable in tests.
var now = time.Now

// taskSections are the README sections that make up "the task": the problem
// statement without any teaching material. Headers are matched case-insensitively;
// legacy headers (Instructions, Challenge, ...) cover READMEs not yet rewritten.
var taskSections = []string{
	"## Task",
	"## Input / Output",
	"## Input/Output",
	"## Examples",
	"### Usage",
	"## Constraints",
	"### Expected function",
	"## Time limit",
}

var legacyTaskHeaders = []string{
	"### Instructions",
	"## Challenge",
	"## The Challenge",
	"## Objective",
}

func cmdStart(root string, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: drill start NN")
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
	readme, err := os.ReadFile(filepath.Join(root, c.Dir, "README.md"))
	if err != nil {
		return err
	}
	rec := prog.record(c.Num)
	target := targetTime(c.Limit, rec.Stage)

	fmt.Printf("=== %s ===\n\n", c.Title)
	fmt.Printf("Round target: %s (stage %d/%d, best %s)\n\n",
		formatDur(target), rec.Stage, len(intervals), bestStr(rec))
	fmt.Println("BLIND: no docs, no search, no AI, no copying. Autocomplete off.")
	fmt.Println("Write every line from memory. Timer started — good luck.")
	fmt.Println(extractTask(string(readme)))
	fmt.Printf("\n---\nWhen done: drill check %d\n", c.Num)

	prog.Pending = &PendingStart{Num: c.Num, At: now()}
	return saveProgress(root, prog)
}

func bestStr(r *Record) string {
	if r.BestMs == 0 {
		return "—"
	}
	return formatDur(time.Duration(r.BestMs) * time.Millisecond)
}

// extractTask returns the problem statement: task-like sections plus the time
// limit, in README order. Anything teaching the solution (skills.md content,
// reference solutions) is never printed by drill.
func extractTask(readme string) string {
	lines := strings.Split(readme, "\n")
	want := map[string]bool{}
	for _, h := range taskSections {
		want[strings.ToLower(h)] = true
	}
	legacy := map[string]bool{}
	for _, h := range legacyTaskHeaders {
		legacy[strings.ToLower(h)] = true
	}
	var out []string
	inTask := false
	seenTask := false
	for _, line := range lines {
		lower := strings.ToLower(strings.TrimSpace(line))
		if strings.HasPrefix(lower, "## ") || strings.HasPrefix(lower, "### ") {
			header := lower
			// Normalize "## The Challenge" etc. to their canonical match.
			inTask = want[header] || legacy[header]
			if legacy[header] && !seenTask {
				seenTask = true
			}
			if inTask {
				out = append(out, line)
			}
			continue
		}
		if inTask {
			out = append(out, line)
		}
	}
	body := strings.TrimSpace(strings.Join(out, "\n"))
	if body == "" {
		return "(no task section found in README — read the file directly)"
	}
	return body
}
