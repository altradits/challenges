package piscine

import (
	"reflect"
	"sort"
	"testing"
)

func TestMapFilterReduce(t *testing.T) {
	got := Map([]int{1, 2, 3}, func(n int) int { return n * 10 })
	if want := []int{10, 20, 30}; !reflect.DeepEqual(got, want) {
		t.Errorf("Map = %v", got)
	}
	got2 := Filter([]int{1, 2, 3, 4}, func(n int) bool { return n%2 == 0 })
	if want := []int{2, 4}; !reflect.DeepEqual(got2, want) {
		t.Errorf("Filter = %v", got2)
	}
	sum := Reduce([]int{1, 2, 3}, 100, func(acc, n int) int { return acc + n })
	if sum != 106 {
		t.Errorf("Reduce = %d, want 106", sum)
	}
}

func TestContainsKeysMin(t *testing.T) {
	if !Contains([]string{"a"}, "a") || Contains([]int{1}, 2) {
		t.Error("Contains wrong")
	}
	keys := Keys(map[string]int{"x": 1, "y": 2})
	sort.Strings(keys)
	if !reflect.DeepEqual(keys, []string{"x", "y"}) {
		t.Errorf("Keys = %v", keys)
	}
	if got := Min([]int{5, 2, 8}); got != 2 {
		t.Errorf("Min ints = %d", got)
	}
	if got := Min([]string{"b", "a"}); got != "a" {
		t.Errorf("Min strings = %q", got)
	}
	if got := Min([]float64{1.5, 0.5}); got != 0.5 {
		t.Errorf("Min floats = %v", got)
	}
}
