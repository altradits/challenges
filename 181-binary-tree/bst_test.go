package piscine

import (
	"reflect"
	"testing"
)

func buildBST(vals ...int) *BST {
	t := NewBST()
	for _, v := range vals {
		t.Insert(v)
	}
	return t
}

func TestBSTInsertContains(t *testing.T) {
	bst := buildBST(4, 2, 6, 1, 3, 5, 7)
	for _, v := range []int{1, 2, 3, 4, 5, 6, 7} {
		if !bst.Contains(v) {
			t.Errorf("Contains(%d) = false, want true", v)
		}
	}
	for _, v := range []int{0, 8, -1, 100} {
		if bst.Contains(v) {
			t.Errorf("Contains(%d) = true, want false", v)
		}
	}
	if bst.Size() != 7 {
		t.Errorf("Size = %d, want 7", bst.Size())
	}
}

func TestBSTTraversals(t *testing.T) {
	bst := buildBST(4, 2, 6, 1, 3, 5, 7)
	if got, want := bst.InOrder(), []int{1, 2, 3, 4, 5, 6, 7}; !reflect.DeepEqual(got, want) {
		t.Errorf("InOrder = %v, want %v", got, want)
	}
	if got, want := bst.PreOrder(), []int{4, 2, 1, 3, 6, 5, 7}; !reflect.DeepEqual(got, want) {
		t.Errorf("PreOrder = %v, want %v", got, want)
	}
	if got, want := bst.PostOrder(), []int{1, 3, 2, 5, 7, 6, 4}; !reflect.DeepEqual(got, want) {
		t.Errorf("PostOrder = %v, want %v", got, want)
	}
}

func TestBSTHeight(t *testing.T) {
	if got := NewBST().Height(); got != 0 {
		t.Errorf("empty Height = %d, want 0", got)
	}
	if got := buildBST(1).Height(); got != 1 {
		t.Errorf("single Height = %d, want 1", got)
	}
	if got := buildBST(4, 2, 6, 1, 3, 5, 7).Height(); got != 3 {
		t.Errorf("balanced Height = %d, want 3", got)
	}
	if got := buildBST(1, 2, 3, 4).Height(); got != 4 {
		t.Errorf("degenerate Height = %d, want 4", got)
	}
}
