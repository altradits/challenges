package piscine

import (
	"sync"
	"testing"
)

func TestCounterBasic(t *testing.T) {
	c := NewCounter()
	if v := c.Value(); v != 0 {
		t.Errorf("fresh Value = %d, want 0", v)
	}
	if v := c.Add(5); v != 5 {
		t.Errorf("Add(5) = %d, want 5", v)
	}
	if v := c.Add(-2); v != 3 {
		t.Errorf("Add(-2) = %d, want 3", v)
	}
	if v := c.Value(); v != 3 {
		t.Errorf("Value = %d, want 3", v)
	}
}

func TestCounterConcurrent(t *testing.T) {
	c := NewCounter()
	var wg sync.WaitGroup
	for g := 0; g < 50; g++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				c.Add(1)
			}
		}()
	}
	wg.Wait()
	if v := c.Value(); v != 5000 {
		t.Errorf("Value = %d, want 5000 (lost updates?)", v)
	}
}

func TestCASAdd(t *testing.T) {
	var x int64
	if v := CASAdd(&x, 10); v != 10 || x != 10 {
		t.Errorf("CASAdd = %d (x=%d), want 10", v, x)
	}
	if v := CASAdd(&x, -4); v != 6 || x != 6 {
		t.Errorf("CASAdd = %d (x=%d), want 6", v, x)
	}
	var wg sync.WaitGroup
	for g := 0; g < 20; g++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				CASAdd(&x, 1)
			}
		}()
	}
	wg.Wait()
	if x != 2006 {
		t.Errorf("x = %d, want 2006 (lost updates?)", x)
	}
}
