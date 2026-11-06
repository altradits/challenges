package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Challenge is one numbered drill discovered from the repo layout.
type Challenge struct {
	Num    int
	Dir    string // e.g. "06-functions"
	Title  string // first "# " line of the challenge README
	Limit  time.Duration
	Manual bool // no hidden tests: check() attests instead of running go test
}

// manualNums are challenges with no runnable tests: a preface, terminal
// tasks, a reference doc, and an external merged PR.
var manualNums = map[int]bool{0: true, 22: true, 83: true, 158: true}

var (
	dirRe   = regexp.MustCompile(`^(\d+)-.+`)
	limitRe = regexp.MustCompile(`\*\*(\d+) minutes\*\*`)
)

// loadCatalog scans root for NN-name directories and parses each README.
func loadCatalog(root string) ([]Challenge, error) {
	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}
	var out []Challenge
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		m := dirRe.FindStringSubmatch(e.Name())
		if m == nil {
			continue
		}
		num, _ := strconv.Atoi(m[1])
		readme, err := os.ReadFile(filepath.Join(root, e.Name(), "README.md"))
		if err != nil {
			continue // not a challenge dir
		}
		out = append(out, Challenge{
			Num:    num,
			Dir:    e.Name(),
			Title:  parseTitle(string(readme)),
			Limit:  parseLimit(string(readme), num),
			Manual: manualNums[num] || isManualReadme(string(readme), e.Name()),
		})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Num < out[j].Num })
	if len(out) == 0 {
		return nil, fmt.Errorf("no challenges found under %s", root)
	}
	return out, nil
}

// findChallenge resolves "6", "06" or "06-functions" to a catalog entry.
func findChallenge(cat []Challenge, arg string) (Challenge, error) {
	arg = strings.TrimSpace(arg)
	for _, c := range cat {
		if arg == c.Dir {
			return c, nil
		}
	}
	n, err := strconv.Atoi(strings.TrimLeft(arg, "0"))
	if arg == "0" || arg == "00" {
		n, err = 0, nil
	}
	if err != nil {
		return Challenge{}, fmt.Errorf("not a challenge: %q (use NN like 6, 06, or 06-functions)", arg)
	}
	for _, c := range cat {
		if c.Num == n {
			return c, nil
		}
	}
	return Challenge{}, fmt.Errorf("no challenge %d", n)
}

func parseTitle(readme string) string {
	for _, line := range strings.Split(readme, "\n") {
		if strings.HasPrefix(line, "# ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "# "))
		}
	}
	return "Untitled challenge"
}

// parseLimit reads the round-1 limit from "**M minutes**" in the README,
// falling back to the phase table when the README has no limit yet.
func parseLimit(readme string, num int) time.Duration {
	if m := limitRe.FindStringSubmatch(readme); m != nil {
		if mins, err := strconv.Atoi(m[1]); err == nil {
			return time.Duration(mins) * time.Minute
		}
	}
	return fallbackLimit(num)
}

func fallbackLimit(num int) time.Duration {
	switch {
	case num <= 0:
		return 15 * time.Minute
	case num <= 5:
		return 5 * time.Minute
	case num <= 27:
		return 15 * time.Minute
	case num <= 80:
		return 10 * time.Minute
	case num <= 140:
		return 20 * time.Minute
	case num <= 144:
		return 30 * time.Minute
	case num <= 152:
		return 25 * time.Minute
	case num <= 155:
		return 120 * time.Minute
	case num == 156:
		return 45 * time.Minute
	case num == 157:
		return 240 * time.Minute
	case num <= 169:
		return 20 * time.Minute
	default:
		return 30 * time.Minute
	}
}

// isManualReadme reports whether the README declares a manual challenge.
func isManualReadme(readme, dir string) bool {
	_ = dir
	return strings.Contains(strings.ToLower(readme), "manual challenge")
}
