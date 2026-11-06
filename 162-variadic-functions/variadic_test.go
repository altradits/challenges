package piscine

import (
	"reflect"
	"testing"
)

func TestSumVariadic(t *testing.T) {
	if got := Sum(1, 2, 3, 4); got != 10 {
		t.Errorf("Sum = %d, want 10", got)
	}
	if got := Sum(); got != 0 {
		t.Errorf("Sum() = %d, want 0", got)
	}
	if got := Sum(-1, 1); got != 0 {
		t.Errorf("Sum(-1, 1) = %d, want 0", got)
	}
}

func TestMaxVariadic(t *testing.T) {
	if got := Max(3, 1, 4, 1, 5); got != 5 {
		t.Errorf("Max = %d, want 5", got)
	}
	if got := Max(-3, -1); got != -1 {
		t.Errorf("Max negatives = %d, want -1", got)
	}
	defer func() {
		if recover() == nil {
			t.Error("Max() with no args should panic")
		}
	}()
	_ = Max()
}

func TestJoinVariadic(t *testing.T) {
	if got := Join("-", "a", "b", "c"); got != "a-b-c" {
		t.Errorf("Join = %q", got)
	}
	if got := Join(","); got != "" {
		t.Errorf("Join with no parts = %q, want empty", got)
	}
}

func TestWrap(t *testing.T) {
	double := func(n int) int { return n * 2 }
	got := Wrap(double, 1, 2, 3)
	if want := []int{2, 4, 6}; !reflect.DeepEqual(got, want) {
		t.Errorf("Wrap = %v, want %v", got, want)
	}
}
