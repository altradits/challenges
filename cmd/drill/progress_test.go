package main

import (
	"testing"
	"time"
)

func TestRecordAttemptStages(t *testing.T) {
	p := &Progress{Records: map[string]*Record{}}
	at := time.Date(2026, 11, 6, 9, 0, 0, 0, time.UTC)

	r := p.recordAttempt(6, false, time.Minute, at)
	if r.Passes != 0 || r.Fails != 1 || r.Stage != 0 {
		t.Fatalf("after fail: %+v", r)
	}
	r = p.recordAttempt(6, true, 10*time.Minute, at)
	if r.Passes != 1 || r.Stage != 1 || r.BestMs != 10*60*1000 {
		t.Fatalf("after pass 1: %+v", r)
	}
	r = p.recordAttempt(6, true, 20*time.Minute, at.Add(time.Hour))
	if r.Stage != 2 || r.BestMs != 10*60*1000 {
		t.Fatalf("best should stay 10m: %+v", r)
	}
	for i := 0; i < 5; i++ {
		r = p.recordAttempt(6, true, time.Minute, at)
	}
	if r.Stage != len(intervals) {
		t.Fatalf("stage should cap at %d, got %d", len(intervals), r.Stage)
	}
}

func TestDueAt(t *testing.T) {
	at := time.Date(2026, 11, 6, 9, 0, 0, 0, time.UTC)
	if due, ok := dueAt(&Record{}); !due.IsZero() || !ok {
		t.Fatalf("fresh should be due now: %v %v", due, ok)
	}
	r := &Record{Passes: 1, Stage: 1, Last: at}
	if due, ok := dueAt(r); !ok || !due.Equal(at.Add(24*time.Hour)) {
		t.Fatalf("stage 1 due = %v %v, want +24h", due, ok)
	}
	r = &Record{Passes: 3, Stage: 3, Last: at}
	if due, ok := dueAt(r); !ok || !due.Equal(at.Add(7*24*time.Hour)) {
		t.Fatalf("stage 3 due = %v %v, want +7d", due, ok)
	}
	if _, ok := dueAt(&Record{Passes: 4, Stage: 4, Last: at}); ok {
		t.Fatal("stage 4 should be mastered")
	}
}

func TestTargetTimeShrinks(t *testing.T) {
	limit := 20 * time.Minute
	prev := limit + time.Minute
	for stage := 0; stage <= 3; stage++ {
		got := targetTime(limit, stage)
		if got >= prev {
			t.Fatalf("stage %d target %v not shrinking (prev %v)", stage, got, prev)
		}
		prev = got
	}
	if got := targetTime(limit, 4); got != targetTime(limit, 3) {
		t.Fatalf("mastered (stage 4) should keep the final target, got %v", got)
	}
	if got := targetTime(limit, 0); got != limit {
		t.Fatalf("stage 0 target = %v, want %v", got, limit)
	}
}

func TestProgressRoundTrip(t *testing.T) {
	root := t.TempDir()
	p := &Progress{Records: map[string]*Record{}}
	at := time.Date(2026, 11, 6, 9, 0, 0, 0, time.UTC)
	p.recordAttempt(6, true, 90*time.Second, at)
	p.Pending = &PendingStart{Num: 7, At: at}
	if err := saveProgress(root, p); err != nil {
		t.Fatal(err)
	}
	loaded, err := loadProgress(root)
	if err != nil {
		t.Fatal(err)
	}
	r := loaded.record(6)
	if r.Passes != 1 || r.BestMs != 90*1000 || !r.Last.Equal(at) {
		t.Fatalf("round trip lost data: %+v", r)
	}
	if loaded.Pending == nil || loaded.Pending.Num != 7 {
		t.Fatalf("round trip lost pending: %+v", loaded.Pending)
	}
}

func TestFormatDur(t *testing.T) {
	cases := map[time.Duration]string{
		45 * time.Second:            "45s",
		5 * time.Minute:             "5m",
		90 * time.Second:            "1m30s",
		2*time.Hour + 5*time.Minute: "2h5m",
	}
	for d, want := range cases {
		if got := formatDur(d); got != want {
			t.Errorf("formatDur(%v) = %q, want %q", d, got, want)
		}
	}
}
