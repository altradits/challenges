package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// intervals is the spaced-repetition schedule: after pass N the challenge is
// due again after intervals[N-1]. Four passes (1/3/7/21 days) master it.
var intervals = []time.Duration{
	24 * time.Hour,
	3 * 24 * time.Hour,
	7 * 24 * time.Hour,
	21 * 24 * time.Hour,
}

// targetFactor shrinks the target time as the stage rises: the deeper the
// memory, the faster the recall must be.
var targetFactor = []float64{1, 0.75, 0.6, 0.5}

// Attempt is one recorded check.
type Attempt struct {
	At        time.Time `json:"at"`
	Pass      bool      `json:"pass"`
	ElapsedMs int64     `json:"elapsed_ms"`
}

// Record is the full history of one challenge.
type Record struct {
	Passes  int       `json:"passes"`
	Fails   int       `json:"fails"`
	Stage   int       `json:"stage"` // completed passes, 0..len(intervals)
	BestMs  int64     `json:"best_ms"`
	Last    time.Time `json:"last"`
	History []Attempt `json:"history,omitempty"`
}

// PendingStart is an open timer from `drill start`.
type PendingStart struct {
	Num int       `json:"num"`
	At  time.Time `json:"at"`
}

// Progress is the whole .drill/progress.json file.
type Progress struct {
	Records map[string]*Record `json:"records"`
	Pending *PendingStart      `json:"pending,omitempty"`
}

func progressPath(root string) string {
	return filepath.Join(root, ".drill", "progress.json")
}

func loadProgress(root string) (*Progress, error) {
	p := &Progress{Records: map[string]*Record{}}
	raw, err := os.ReadFile(progressPath(root))
	if os.IsNotExist(err) {
		return p, nil
	}
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(raw, p); err != nil {
		return nil, fmt.Errorf("corrupt %s: %w (delete it to restart)", progressPath(root), err)
	}
	if p.Records == nil {
		p.Records = map[string]*Record{}
	}
	return p, nil
}

func saveProgress(root string, p *Progress) error {
	if err := os.MkdirAll(filepath.Join(root, ".drill"), 0o755); err != nil {
		return err
	}
	raw, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(progressPath(root), append(raw, '\n'), 0o644)
}

func key(num int) string { return fmt.Sprintf("%03d", num) }

func (p *Progress) record(num int) *Record {
	k := key(num)
	r, ok := p.Records[k]
	if !ok {
		r = &Record{}
		p.Records[k] = r
	}
	return r
}

// recordAttempt stores one check result. A pass advances the stage (capped at
// mastered); a fail leaves the stage but counts toward weak-topic stats.
func (p *Progress) recordAttempt(num int, pass bool, elapsed time.Duration, at time.Time) *Record {
	r := p.record(num)
	r.Last = at
	ms := elapsed.Milliseconds()
	if ms < 0 {
		ms = 0
	}
	if pass {
		r.Passes++
		if r.Stage < len(intervals) {
			r.Stage++
		}
		if r.BestMs == 0 || ms < r.BestMs {
			r.BestMs = ms
		}
	} else {
		r.Fails++
	}
	r.History = append(r.History, Attempt{At: at, Pass: pass, ElapsedMs: ms})
	if len(r.History) > 50 {
		r.History = r.History[len(r.History)-50:]
	}
	if p.Pending != nil && p.Pending.Num == num {
		p.Pending = nil
	}
	return r
}

// dueAt returns when the challenge is next due. The zero time means due now;
// ok=false means mastered (never due).
func dueAt(r *Record) (due time.Time, ok bool) {
	if r.Passes == 0 {
		return time.Time{}, true
	}
	if r.Stage >= len(intervals) {
		return time.Time{}, false
	}
	return r.Last.Add(intervals[r.Stage-1]), true
}

// targetTime is the time to beat for the current stage.
func targetTime(limit time.Duration, stage int) time.Duration {
	if stage < 0 {
		stage = 0
	}
	if stage >= len(targetFactor) {
		stage = len(targetFactor) - 1
	}
	return time.Duration(float64(limit) * targetFactor[stage])
}

func formatDur(d time.Duration) string {
	d = d.Round(time.Second)
	if d < time.Minute {
		return fmt.Sprintf("%ds", int(d.Seconds()))
	}
	if d < time.Hour {
		m, s := int(d.Minutes()), int(d.Seconds())%60
		if s == 0 {
			return fmt.Sprintf("%dm", m)
		}
		return fmt.Sprintf("%dm%ds", m, s)
	}
	return fmt.Sprintf("%dh%dm", int(d.Hours()), int(d.Minutes())%60)
}
