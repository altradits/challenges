package piscine

import (
	"reflect"
	"testing"
)

func sampleGraph() *Graph {
	g := NewGraph()
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	return g
}

func TestGraphBFS(t *testing.T) {
	if got, want := sampleGraph().BFS(1), []int{1, 2, 3, 4}; !reflect.DeepEqual(got, want) {
		t.Errorf("BFS(1) = %v, want %v", got, want)
	}
	if got := NewGraph().BFS(9); !reflect.DeepEqual(got, []int{9}) {
		t.Errorf("BFS on empty graph = %v, want [9]", got)
	}
}

func TestGraphDFS(t *testing.T) {
	if got, want := sampleGraph().DFS(1), []int{1, 2, 4, 3}; !reflect.DeepEqual(got, want) {
		t.Errorf("DFS(1) = %v, want %v", got, want)
	}
	if got := NewGraph().DFS(9); !reflect.DeepEqual(got, []int{9}) {
		t.Errorf("DFS on empty graph = %v, want [9]", got)
	}
}

func TestGraphHasPath(t *testing.T) {
	g := sampleGraph()
	if !g.HasPath(1, 4) {
		t.Error("HasPath(1,4) = false, want true")
	}
	if !g.HasPath(3, 4) {
		t.Error("HasPath(3,4) = false, want true")
	}
	if g.HasPath(1, 99) {
		t.Error("HasPath(1,99) = true, want false")
	}
	if !g.HasPath(2, 2) {
		t.Error("HasPath(2,2) = false, want true")
	}
}

func TestGraphConnectedComponents(t *testing.T) {
	g := NewGraph()
	g.AddEdge(1, 2)
	g.AddEdge(3, 4)
	if got := g.ConnectedComponents(); got != 2 {
		t.Errorf("ConnectedComponents = %d, want 2", got)
	}
	if got := NewGraph().ConnectedComponents(); got != 0 {
		t.Errorf("empty ConnectedComponents = %d, want 0", got)
	}
}
