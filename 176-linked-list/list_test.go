package piscine

import (
	"reflect"
	"testing"
)

func TestLinkedListBasics(t *testing.T) {
	l := NewLinkedList()
	if l.Len() != 0 {
		t.Fatalf("Len = %d, want 0", l.Len())
	}
	l.Push(1)
	l.Push(2)
	l.Append(3)
	if got, want := l.ToSlice(), []int{2, 1, 3}; !reflect.DeepEqual(got, want) {
		t.Errorf("ToSlice = %v, want %v", got, want)
	}
	if l.Len() != 3 {
		t.Errorf("Len = %d, want 3", l.Len())
	}
	if !l.Contains(1) || l.Contains(99) {
		t.Error("Contains wrong")
	}
	v, ok := l.Pop()
	if !ok || v != 2 {
		t.Errorf("Pop = (%d, %v), want (2, true)", v, ok)
	}
	if _, ok := NewLinkedList().Pop(); ok {
		t.Error("Pop on empty should return ok=false")
	}
}

func TestLinkedListReverse(t *testing.T) {
	l := NewLinkedList()
	for _, v := range []int{1, 2, 3} {
		l.Append(v)
	}
	l.Reverse()
	if got, want := l.ToSlice(), []int{3, 2, 1}; !reflect.DeepEqual(got, want) {
		t.Errorf("after Reverse: %v, want %v", got, want)
	}
}

func TestLinkedListRemoveAt(t *testing.T) {
	l := NewLinkedList()
	for _, v := range []int{10, 20, 30} {
		l.Append(v)
	}
	if !l.RemoveAt(1) {
		t.Fatal("RemoveAt(1) = false")
	}
	if got, want := l.ToSlice(), []int{10, 30}; !reflect.DeepEqual(got, want) {
		t.Errorf("after remove: %v, want %v", got, want)
	}
	if l.RemoveAt(5) || l.RemoveAt(-1) {
		t.Error("out-of-range RemoveAt should return false")
	}
	if !l.RemoveAt(0) {
		t.Error("RemoveAt(0) should work")
	}
}
