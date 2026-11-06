package piscine

import (
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	if !q.IsEmpty() || q.Size() != 0 {
		t.Fatal("fresh queue should be empty")
	}
	if _, ok := q.Dequeue(); ok {
		t.Error("Dequeue on empty should return false")
	}
	if _, ok := q.Front(); ok {
		t.Error("Front on empty should return false")
	}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	if q.Size() != 3 {
		t.Errorf("Size = %d, want 3", q.Size())
	}
	if v, ok := q.Front(); !ok || v != 1 {
		t.Errorf("Front = (%d, %v), want (1, true)", v, ok)
	}
	for _, want := range []int{1, 2, 3} {
		v, ok := q.Dequeue()
		if !ok || v != want {
			t.Fatalf("Dequeue = (%d, %v), want (%d, true)", v, ok, want)
		}
	}
	if !q.IsEmpty() {
		t.Error("queue should be empty again")
	}
}

func TestBFS178(t *testing.T) {
	graph := map[int][]int{
		1: {2, 3},
		2: {4},
		3: {4},
		4: {},
	}
	got := BFS(graph, 1)
	if want := []int{1, 2, 3, 4}; !reflect.DeepEqual(got, want) {
		t.Errorf("BFS = %v, want %v", got, want)
	}
	if got := BFS(graph, 9); !reflect.DeepEqual(got, []int{9}) {
		t.Errorf("BFS from unknown = %v, want [9]", got)
	}
}
