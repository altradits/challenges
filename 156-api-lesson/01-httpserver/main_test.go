package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateServer(t *testing.T) {
	h := CreateServer()
	if h == nil {
		t.Fatal("CreateServer returned nil")
	}
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("GET / status = %d, want 200", rec.Code)
	}
	if rec.Body.String() != "Hello, World!" {
		t.Errorf("GET / body = %q, want %q", rec.Body.String(), "Hello, World!")
	}
}
