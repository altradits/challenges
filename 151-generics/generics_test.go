package piscine

import (
	"reflect"
	"sort"
	"testing"
)

func TestMapGeneric(t *testing.T) {
	got := Map([]int{1, 2, 3}, func(n int) string {
		return string(rune('a' + n))
	})
	if want := []string{"b", "c", "d"}; !reflect.DeepEqual(got, want) {
		t.Errorf("Map = %v, want %v", got, want)
	}
	if got := Map([]int{}, func(n int) int { return n }); len(got) != 0 {
		t.Errorf("Map empty = %v, want empty", got)
	}
}

func TestFilterGeneric(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }
	got := Filter([]int{1, 2, 3, 4}, isEven)
	if want := []int{2, 4}; !reflect.DeepEqual(got, want) {
		t.Errorf("Filter = %v, want %v", got, want)
	}
}

func TestReduceGeneric(t *testing.T) {
	sum := Reduce([]int{1, 2, 3, 4}, 0, func(acc, n int) int { return acc + n })
	if sum != 10 {
		t.Errorf("Reduce sum = %d, want 10", sum)
	}
	concat := Reduce([]string{"a", "b"}, "seed:", func(acc, s string) string { return acc + s })
	if concat != "seed:ab" {
		t.Errorf("Reduce concat = %q", concat)
	}
}

func TestKeys(t *testing.T) {
	got := Keys(map[string]int{"b": 2, "a": 1})
	sort.Strings(got)
	if want := []string{"a", "b"}; !reflect.DeepEqual(got, want) {
		t.Errorf("Keys = %v, want %v", got, want)
	}
	if got := Keys(map[string]int{}); len(got) != 0 {
		t.Errorf("Keys empty = %v", got)
	}
}

func TestContainsGeneric(t *testing.T) {
	if !Contains([]string{"a", "b"}, "b") {
		t.Error("Contains should find b")
	}
	if Contains([]string{"a", "b"}, "z") {
		t.Error("Contains should not find z")
	}
	if Contains([]int{}, 1) {
		t.Error("Contains on empty should be false")
	}
}
