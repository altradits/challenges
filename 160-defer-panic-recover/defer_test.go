package piscine

import (
	"reflect"
	"testing"
)

func TestSafeDivide(t *testing.T) {
	got, err := SafeDivide(10, 2)
	if err != nil || got != 5 {
		t.Errorf("SafeDivide(10, 2) = (%d, %v), want (5, nil)", got, err)
	}
	got, err = SafeDivide(7, 0)
	if err == nil || got != 0 {
		t.Errorf("SafeDivide(7, 0) = (%d, %v), want (0, non-nil error)", got, err)
	}
}

func TestWithCleanup(t *testing.T) {
	cleaned := false
	err := WithCleanup(func() {}, func() { cleaned = true })
	if err != nil {
		t.Errorf("no panic: err = %v, want nil", err)
	}
	if !cleaned {
		t.Error("cleanup did not run on success")
	}
	cleaned = false
	err = WithCleanup(func() { panic("boom") }, func() { cleaned = true })
	if err == nil {
		t.Error("panic: err = nil, want non-nil")
	}
	if !cleaned {
		t.Error("cleanup did not run on panic")
	}
}

func TestDeferOrder(t *testing.T) {
	got := DeferOrder()
	want := []string{"third", "second", "first"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("DeferOrder = %v, want %v (LIFO)", got, want)
	}
}
