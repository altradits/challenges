package piscine

import "testing"

func TestStack(t *testing.T) {
	s := NewStack()
	if !s.IsEmpty() || s.Size() != 0 {
		t.Fatal("fresh stack should be empty")
	}
	if _, ok := s.Pop(); ok {
		t.Error("Pop on empty should return false")
	}
	if _, ok := s.Peek(); ok {
		t.Error("Peek on empty should return false")
	}
	s.Push(1)
	s.Push(2)
	if s.Size() != 2 {
		t.Errorf("Size = %d, want 2", s.Size())
	}
	if v, ok := s.Peek(); !ok || v != 2 {
		t.Errorf("Peek = (%d, %v), want (2, true)", v, ok)
	}
	if v, ok := s.Pop(); !ok || v != 2 {
		t.Errorf("Pop = (%d, %v), want (2, true)", v, ok)
	}
	if v, ok := s.Pop(); !ok || v != 1 {
		t.Errorf("Pop = (%d, %v), want (1, true)", v, ok)
	}
	if !s.IsEmpty() {
		t.Error("stack should be empty again")
	}
}

func TestIsBalanced(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		{"({[]})", true},
		{"([)]", false},
		{"(((", false},
		{"", true},
		{"hello", true},
		{"(a[b{c}d]e)", true},
		{"(]", false},
		{")(", false},
	}
	for _, tc := range cases {
		if got := IsBalanced(tc.input); got != tc.want {
			t.Errorf("IsBalanced(%q) = %v, want %v", tc.input, got, tc.want)
		}
	}
}
