package main

import (
	"strings"
	"testing"
	"time"
)

func passAt(at time.Time) Attempt { return Attempt{At: at, Pass: true} }

func TestStreakDays(t *testing.T) {
	// Friday 2026-11-06 12:00 UTC.
	friday := time.Date(2026, 11, 6, 12, 0, 0, 0, time.UTC)
	p := &Progress{Records: map[string]*Record{
		"006": {History: []Attempt{passAt(friday), passAt(friday.AddDate(0, 0, -1)), passAt(friday.AddDate(0, 0, -2))}},
	}}
	if got := streakDays(p, friday); got != 3 {
		t.Fatalf("streak = %d, want 3", got)
	}
	// Missing today keeps the streak alive.
	if got := streakDays(p, friday.AddDate(0, 0, 1)); got != 3 {
		t.Fatalf("streak next day = %d, want 3", got)
	}
	// A gap ends it.
	p.Records["006"].History = []Attempt{passAt(friday), passAt(friday.AddDate(0, 0, -2))}
	if got := streakDays(p, friday); got != 1 {
		t.Fatalf("streak with gap = %d, want 1", got)
	}
	if got := streakDays(&Progress{Records: map[string]*Record{}}, friday); got != 0 {
		t.Fatalf("empty streak = %d, want 0", got)
	}
}

func TestWeakChallengesOrdering(t *testing.T) {
	cat := []Challenge{
		{Num: 6, Dir: "06-a", Limit: 10 * time.Minute},
		{Num: 7, Dir: "07-b", Limit: 10 * time.Minute},
		{Num: 8, Dir: "08-c", Limit: 10 * time.Minute},
	}
	p := &Progress{Records: map[string]*Record{}}
	p.recordAttempt(6, true, time.Minute, now())
	p.recordAttempt(7, false, time.Minute, now())
	p.recordAttempt(7, false, time.Minute, now())
	weak := weakChallenges(cat, p, 5)
	if len(weak) != 2 {
		t.Fatalf("weak = %v", weak)
	}
	if !strings.Contains(weak[0], "07-b") || !strings.Contains(weak[0], "2 fails") {
		t.Fatalf("most-failed should rank first: %v", weak)
	}
}

func TestPhaseOf(t *testing.T) {
	cases := map[int]string{
		0: "step 0 (keywords)", 3: "phase 1 (hello)", 27: "phase 2 (foundations)",
		51: "phase 3 (beginner)", 80: "phase 4 (strings)", 144: "phase 5 (problems)",
		158: "phase 6 (backend)", 169: "phase 7 (depth)", 171: "phase 9 (bitcoin I)",
		185: "phase 8 (algorithms)", 190: "phase 9 (bitcoin II)",
	}
	for num, want := range cases {
		if got := phaseOf(num); got != want {
			t.Errorf("phaseOf(%d) = %q, want %q", num, got, want)
		}
	}
}
