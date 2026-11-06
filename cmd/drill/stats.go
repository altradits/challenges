package main

import (
	"fmt"
	"sort"
	"time"
)

func cmdStats(root string, args []string) error {
	if len(args) != 0 {
		return fmt.Errorf("usage: drill stats")
	}
	cat, err := loadCatalog(root)
	if err != nil {
		return err
	}
	prog, err := loadProgress(root)
	if err != nil {
		return err
	}
	at := now()
	var passed, mastered, fails int
	for _, c := range cat {
		r := prog.record(c.Num)
		if r.Passes > 0 {
			passed++
		}
		if _, ok := dueAt(r); r.Passes > 0 && !ok {
			mastered++
		}
		fails += r.Fails
	}
	fmt.Printf("Challenges: %d passed, %d mastered, %d total\n", passed, mastered, len(cat))
	fmt.Printf("Day streak: %d\n\n", streakDays(prog, at))

	fmt.Println("Slowest recalls vs target (drill these first):")
	weak := weakChallenges(cat, prog, 5)
	if len(weak) == 0 {
		fmt.Println("  (no attempts yet)")
	}
	for _, w := range weak {
		fmt.Println("  " + w)
	}
	fmt.Println("\nWeakest phases by failure rate:")
	for _, line := range weakPhases(cat, prog) {
		fmt.Println("  " + line)
	}
	return nil
}

// streakDays counts consecutive local days with at least one pass, ending
// today or yesterday (a streak stays alive until the end of today).
func streakDays(prog *Progress, at time.Time) int {
	days := map[string]bool{}
	for _, r := range prog.Records {
		for _, a := range r.History {
			if a.Pass {
				y, m, d := a.At.Date()
				days[fmt.Sprintf("%04d-%02d-%02d", y, m, d)] = true
			}
		}
	}
	y, m, d := at.Date()
	today := time.Date(y, m, d, 0, 0, 0, 0, at.Location())
	streak := 0
	for day := today; ; day = day.AddDate(0, 0, -1) {
		yy, mm, dd := day.Date()
		if !days[fmt.Sprintf("%04d-%02d-%02d", yy, mm, dd)] {
			// Yesterday missing ends the streak; today missing is fine.
			if day.Equal(today) {
				continue
			}
			return streak
		}
		streak++
	}
}

// weakChallenges ranks attempted challenges by failures (then by best time
// over target) and formats the top n as one line each.
func weakChallenges(cat []Challenge, prog *Progress, n int) []string {
	type row struct {
		c     Challenge
		fails int
		over  float64 // best/target ratio, 0 when no pass yet
	}
	var rows []row
	for _, c := range cat {
		r := prog.record(c.Num)
		if r.Passes == 0 && r.Fails == 0 {
			continue
		}
		over := 0.0
		if r.BestMs > 0 && c.Limit > 0 {
			best := time.Duration(r.BestMs) * time.Millisecond
			over = float64(best) / float64(targetTime(c.Limit, r.Stage))
		}
		rows = append(rows, row{c, r.Fails, over})
	}
	sort.Slice(rows, func(i, j int) bool {
		if rows[i].fails != rows[j].fails {
			return rows[i].fails > rows[j].fails
		}
		return rows[i].over > rows[j].over
	})
	if len(rows) > n {
		rows = rows[:n]
	}
	var out []string
	for _, r := range rows {
		rec := prog.record(r.c.Num)
		out = append(out, fmt.Sprintf("%-22s %d fails, %d passes, best %s (target %s)",
			r.c.Dir, rec.Fails, rec.Passes, bestStr(rec),
			formatDur(targetTime(r.c.Limit, rec.Stage))))
	}
	return out
}

// phaseOf names the learning phase for stats rollups.
func phaseOf(num int) string {
	switch {
	case num <= 0:
		return "step 0 (keywords)"
	case num <= 5:
		return "phase 1 (hello)"
	case num <= 27:
		return "phase 2 (foundations)"
	case num <= 51:
		return "phase 3 (beginner)"
	case num <= 80:
		return "phase 4 (strings)"
	case num <= 144:
		return "phase 5 (problems)"
	case num <= 158:
		return "phase 6 (backend)"
	case num <= 169:
		return "phase 7 (depth)"
	case num <= 175:
		return "phase 9 (bitcoin I)"
	case num <= 185:
		return "phase 8 (algorithms)"
	default:
		return "phase 9 (bitcoin II)"
	}
}

func weakPhases(cat []Challenge, prog *Progress) []string {
	type agg struct {
		name          string
		passes, fails int
	}
	byPhase := map[string]*agg{}
	var order []string
	for _, c := range cat {
		name := phaseOf(c.Num)
		a, ok := byPhase[name]
		if !ok {
			a = &agg{name: name}
			byPhase[name] = a
			order = append(order, name)
		}
		r := prog.record(c.Num)
		a.passes += r.Passes
		a.fails += r.Fails
	}
	var rows []*agg
	for _, name := range order {
		a := byPhase[name]
		if a.passes+a.fails > 0 {
			rows = append(rows, a)
		}
	}
	sort.Slice(rows, func(i, j int) bool {
		ri := float64(rows[i].fails) / float64(rows[i].passes+rows[i].fails)
		rj := float64(rows[j].fails) / float64(rows[j].passes+rows[j].fails)
		return ri > rj
	})
	var out []string
	for _, a := range rows {
		rate := 100 * float64(a.fails) / float64(a.passes+a.fails)
		out = append(out, fmt.Sprintf("%-22s %d fails / %d passes (%.0f%% fail rate)",
			a.name, a.fails, a.passes, rate))
	}
	if len(out) == 0 {
		out = append(out, "(no attempts yet)")
	}
	return out
}
