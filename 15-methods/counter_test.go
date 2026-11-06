package piscine

import "testing"

func TestCounter(t *testing.T) {
	c := Counter{}
	if got := c.Value(); got != 0 {
		t.Fatalf("fresh Value() = %d, want 0", got)
	}
	c.Increment()
	c.Increment()
	c.Add(4)
	if got := c.Value(); got != 6 {
		t.Errorf("after ++, ++, +4: Value() = %d, want 6", got)
	}
	c.Add(-2)
	if got := c.Value(); got != 4 {
		t.Errorf("after -2: Value() = %d, want 4", got)
	}
	c.Reset()
	if got := c.Value(); got != 0 {
		t.Errorf("after Reset: Value() = %d, want 0", got)
	}
}
