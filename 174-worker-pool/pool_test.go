package piscine

import (
	"reflect"
	"sync/atomic"
	"testing"
	"time"
)

func TestRunPoolOrder(t *testing.T) {
	jobs := make([]func() int, 20)
	for i := range jobs {
		i := i
		jobs[i] = func() int {
			time.Sleep(time.Duration(20-i) * time.Millisecond)
			return i * i
		}
	}
	got := RunPool(jobs, 4)
	want := make([]int, 20)
	for i := range want {
		want[i] = i * i
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("RunPool order wrong:\n%v\nwant\n%v", got, want)
	}
}

func TestRunPoolConcurrency(t *testing.T) {
	var current, maxSeen atomic.Int64
	jobs := make([]func() int, 8)
	for i := range jobs {
		jobs[i] = func() int {
			n := current.Add(1)
			for {
				m := maxSeen.Load()
				if n <= m || maxSeen.CompareAndSwap(m, n) {
					break
				}
			}
			time.Sleep(50 * time.Millisecond)
			current.Add(-1)
			return 1
		}
	}
	got := RunPool(jobs, 4)
	if len(got) != 8 {
		t.Fatalf("RunPool returned %d results, want 8", len(got))
	}
	if m := maxSeen.Load(); m < 2 {
		t.Errorf("max concurrent jobs = %d, want >= 2 (jobs ran sequentially?)", m)
	} else if m > 4 {
		t.Errorf("max concurrent jobs = %d, want <= 4 workers", m)
	}
}

func TestRunPoolEdges(t *testing.T) {
	if got := RunPool(nil, 4); got == nil || len(got) != 0 {
		t.Errorf("RunPool(nil) = %v, want empty non-nil slice", got)
	}
	got := RunPool([]func() int{func() int { return 7 }}, 0)
	if !reflect.DeepEqual(got, []int{7}) {
		t.Errorf("RunPool(workers=0) = %v, want [7]", got)
	}
}
