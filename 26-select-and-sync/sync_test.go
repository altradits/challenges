package piscine

import (
	"errors"
	"sync"
	"testing"
	"time"
)

func TestSafeCounterConcurrent(t *testing.T) {
	c := NewSafeCounter()
	var wg sync.WaitGroup
	for g := 0; g < 10; g++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				c.Inc()
			}
		}()
	}
	wg.Wait()
	if got := c.Value(); got != 1000 {
		t.Errorf("Value() = %d, want 1000 (lost updates = missing mutex)", got)
	}
}

func TestTimeout(t *testing.T) {
	got, err := Timeout(func() string { return "fast" }, time.Second)
	if err != nil || got != "fast" {
		t.Errorf("fast Timeout = (%q, %v), want (fast, nil)", got, err)
	}
	got, err = Timeout(func() string {
		time.Sleep(300 * time.Millisecond)
		return "slow"
	}, 20*time.Millisecond)
	if !errors.Is(err, ErrTimeout) {
		t.Errorf("slow Timeout err = %v, want errors.Is ErrTimeout", err)
	}
	if got != "" {
		t.Errorf("slow Timeout value = %q, want empty", got)
	}
}

func TestNewOnce(t *testing.T) {
	calls := 0
	once := NewOnce(func() string {
		calls++
		return "result"
	})
	var wg sync.WaitGroup
	results := make([]string, 20)
	for i := range results {
		wg.Add(1)
		go func() {
			defer wg.Done()
			results[i] = once()
		}()
	}
	wg.Wait()
	for i, r := range results {
		if r != "result" {
			t.Fatalf("call %d = %q, want \"result\"", i, r)
		}
	}
	if calls != 1 {
		t.Errorf("fn ran %d times, want exactly 1", calls)
	}
}
