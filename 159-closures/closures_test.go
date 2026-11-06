package piscine

import (
	"reflect"
	"testing"
)

func TestMakeCounter(t *testing.T) {
	c := MakeCounter()
	for want := 1; want <= 3; want++ {
		if got := c(); got != want {
			t.Fatalf("call %d = %d, want %d", want, got, want)
		}
	}
	other := MakeCounter()
	if got := other(); got != 1 {
		t.Errorf("fresh counter = %d, want 1 (independent state)", got)
	}
}

func TestMakeAdder(t *testing.T) {
	add5 := MakeAdder(5)
	if got := add5(3); got != 8 {
		t.Errorf("add5(3) = %d, want 8", got)
	}
	if got := add5(10); got != 15 {
		t.Errorf("add5(10) = %d, want 15", got)
	}
	addNeg := MakeAdder(-2)
	if got := addNeg(2); got != 0 {
		t.Errorf("addNeg(2) = %d, want 0", got)
	}
}

func TestFilterClosure(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }
	got := Filter([]int{1, 2, 3, 4, 5}, isEven)
	if want := []int{2, 4}; !reflect.DeepEqual(got, want) {
		t.Errorf("Filter = %v, want %v", got, want)
	}
	if got := Filter(nil, isEven); len(got) != 0 {
		t.Errorf("Filter nil = %v, want empty", got)
	}
}

func TestCompose(t *testing.T) {
	double := func(x int) int { return x * 2 }
	inc := func(x int) int { return x + 1 }
	if got := Compose(double, inc)(3); got != 8 {
		t.Errorf("Compose(double, inc)(3) = %d, want 8 (double first... no: inc first, then double)", got)
	}
	if got := Compose(inc, double)(3); got != 7 {
		t.Errorf("Compose(inc, double)(3) = %d, want 7", got)
	}
}
