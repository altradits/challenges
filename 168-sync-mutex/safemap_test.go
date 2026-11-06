package piscine

import (
	"sync"
	"testing"
)

func TestSafeCounter168(t *testing.T) {
	c := NewSafeCounter()
	var wg sync.WaitGroup
	for g := 0; g < 8; g++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 125; i++ {
				c.Inc()
			}
		}()
	}
	wg.Wait()
	if got := c.Value(); got != 1000 {
		t.Errorf("Value = %d, want 1000", got)
	}
	c.Add(5)
	if got := c.Value(); got != 1005 {
		t.Errorf("after Add(5): %d", got)
	}
	c.Reset()
	if got := c.Value(); got != 0 {
		t.Errorf("after Reset: %d", got)
	}
}

func TestSafeMap168(t *testing.T) {
	m := NewSafeMap()
	var wg sync.WaitGroup
	for g := 0; g < 8; g++ {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				m.Set(string(rune('a'+g))+string(rune('0'+i%10)), i)
			}
		}(g)
	}
	wg.Wait()
	if got := m.Len(); got != 80 {
		t.Errorf("Len = %d, want 80", got)
	}
	v, ok := m.Get("a0")
	if !ok {
		t.Error("Get(a0) missing")
	}
	_ = v
	m.Delete("a0")
	if _, ok := m.Get("a0"); ok {
		t.Error("Delete did not remove a0")
	}
}
