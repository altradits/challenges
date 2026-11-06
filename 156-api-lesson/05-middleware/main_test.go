package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddlewareChain(t *testing.T) {
	inner := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/panic" {
			panic("boom")
		}
		w.Write([]byte("ok"))
	})
	h := CreateMiddlewareChain(inner)
	if h == nil {
		t.Fatal("CreateMiddlewareChain returned nil")
	}
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("GET", "/", nil))
	if rec.Code != 200 || rec.Body.String() != "ok" {
		t.Errorf("passthrough = %d %q, want 200 ok", rec.Code, rec.Body.String())
	}
	if got := rec.Header().Get("Access-Control-Allow-Origin"); got != "*" {
		t.Errorf("CORS header = %q, want *", got)
	}
	rec = httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("GET", "/panic", nil))
	if rec.Code != 500 {
		t.Errorf("panic recovered with %d, want 500", rec.Code)
	}
}
