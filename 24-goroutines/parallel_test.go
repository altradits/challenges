package piscine

import (
	"reflect"
	"testing"
)

func TestConcurrentMap(t *testing.T) {
	double := func(n int) int { return n * 2 }
	got := ConcurrentMap([]int{1, 2, 3, 4, 5}, double)
	if want := []int{2, 4, 6, 8, 10}; !reflect.DeepEqual(got, want) {
		t.Errorf("ConcurrentMap = %v, want %v (order must match input)", got, want)
	}
	if got := ConcurrentMap(nil, double); len(got) != 0 {
		t.Errorf("ConcurrentMap(nil) = %v, want empty", got)
	}
	// Bigger input to shake out races (also run with -race).
	big := make([]int, 500)
	for i := range big {
		big[i] = i
	}
	got = ConcurrentMap(big, func(n int) int { return -n })
	for i, v := range got {
		if v != -i {
			t.Fatalf("ConcurrentMap[%d] = %d, want %d", i, v, -i)
		}
	}
}

func TestParallelSum(t *testing.T) {
	cases := []struct {
		name    string
		data    []int
		workers int
		want    int
	}{
		{"even split", []int{1, 2, 3, 4}, 2, 10},
		{"uneven split", []int{1, 2, 3, 4, 5}, 2, 15},
		{"one worker", []int{1, 2, 3}, 1, 6},
		{"more workers than data", []int{1, 2}, 8, 3},
		{"empty", []int{}, 4, 0},
		{"nil", nil, 4, 0},
		{"zero workers", []int{1, 2, 3}, 0, 6},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ParallelSum(tc.data, tc.workers); got != tc.want {
				t.Errorf("ParallelSum(%v, %d) = %d, want %d", tc.data, tc.workers, got, tc.want)
			}
		})
	}
}
