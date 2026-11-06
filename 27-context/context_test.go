package piscine

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestFetchWithTimeoutDone(t *testing.T) {
	got, err := FetchWithTimeout(context.Background(), 10*time.Millisecond)
	if err != nil || got != "done" {
		t.Errorf("FetchWithTimeout = (%q, %v), want (done, nil)", got, err)
	}
}

func TestFetchWithTimeoutCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := FetchWithTimeout(ctx, time.Hour)
	if !errors.Is(err, context.Canceled) {
		t.Errorf("cancelled FetchWithTimeout err = %v, want context.Canceled", err)
	}
}

func TestFetchWithTimeoutDeadline(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	_, err := FetchWithTimeout(ctx, time.Hour)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("expired FetchWithTimeout err = %v, want context.DeadlineExceeded", err)
	}
}

func TestProcessAll(t *testing.T) {
	var seen []string
	err := ProcessAll(context.Background(), []string{"a", "b", "c"}, func(s string) {
		seen = append(seen, s)
	})
	if err != nil {
		t.Fatalf("ProcessAll: %v", err)
	}
	if len(seen) != 3 {
		t.Errorf("ProcessAll visited %d items, want 3", len(seen))
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	called := 0
	err = ProcessAll(ctx, []string{"a", "b"}, func(s string) { called++ })
	if !errors.Is(err, context.Canceled) {
		t.Errorf("cancelled ProcessAll err = %v, want context.Canceled", err)
	}
	if called != 0 {
		t.Errorf("cancelled ProcessAll still ran fn %d times, want 0", called)
	}
}
