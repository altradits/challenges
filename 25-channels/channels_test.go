package piscine

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

func collect(t *testing.T, ch <-chan int) []int {
	t.Helper()
	var out []int
	timeout := time.After(5 * time.Second)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return out
			}
			out = append(out, v)
		case <-timeout:
			t.Fatal("timed out waiting for channel to close (leaked goroutine or missing close?)")
		}
	}
}

func TestGenerator(t *testing.T) {
	if got := collect(t, Generator(5)); !reflect.DeepEqual(got, []int{0, 1, 2, 3, 4}) {
		t.Errorf("Generator(5) = %v, want [0 1 2 3 4] in order", got)
	}
	if got := collect(t, Generator(0)); len(got) != 0 {
		t.Errorf("Generator(0) = %v, want empty (closed at once)", got)
	}
}

func TestMerge(t *testing.T) {
	a := make(chan int, 3)
	b := make(chan int, 2)
	a <- 1
	a <- 3
	a <- 5
	close(a)
	b <- 2
	b <- 4
	close(b)
	got := collect(t, Merge(a, b))
	sort.Ints(got)
	if want := []int{1, 2, 3, 4, 5}; !reflect.DeepEqual(got, want) {
		t.Errorf("Merge = %v, want %v (any order, then closed)", got, want)
	}
	empty := make(chan int)
	close(empty)
	if got := collect(t, Merge(empty, empty)); len(got) != 0 {
		t.Errorf("Merge of closed channels = %v, want empty", got)
	}
}

func TestWorkerPool(t *testing.T) {
	jobs := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		jobs <- i
	}
	close(jobs)
	got := collect(t, WorkerPool(jobs, 3, func(n int) int { return n * n }))
	sort.Ints(got)
	want := []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("WorkerPool squares = %v, want %v", got, want)
	}
}
