package piscine

import "testing"

func TestEmbedding(t *testing.T) {
	s := NewServer("example.com", 8080)
	if s.Prefix != "SERVER" {
		t.Errorf("Prefix = %q, want SERVER", s.Prefix)
	}
	if got := s.Address(); got != "example.com:8080" {
		t.Errorf("Address = %q", got)
	}
	if got := s.Log("hi"); got != "[SERVER] hi" {
		t.Errorf("promoted Log = %q", got)
	}
	if got := s.Start(); got != "[SERVER] Starting example.com:8080" {
		t.Errorf("Start = %q", got)
	}
}

func TestCache(t *testing.T) {
	c := NewCache()
	c.Set("k", "v")
	got, ok := c.Get("k")
	if !ok || got != "v" {
		t.Errorf("Get = (%q, %v)", got, ok)
	}
	if _, ok := c.Get("missing"); ok {
		t.Error("Get missing should return ok=false")
	}
}

func TestCachedServer(t *testing.T) {
	cs := NewCachedServer("h", 1)
	if got := cs.Address(); got != "h:1" {
		t.Errorf("promoted Address = %q", got)
	}
	cs.Store("k", "v")
	got, ok := cs.Fetch("k")
	if !ok || got != "v" {
		t.Errorf("Fetch = (%q, %v)", got, ok)
	}
}
