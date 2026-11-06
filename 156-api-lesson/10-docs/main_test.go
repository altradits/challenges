package main

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDocumentedAPI(t *testing.T) {
	h := CreateDocumentedAPI()
	if h == nil {
		t.Fatal("CreateDocumentedAPI returned nil")
	}
	serve := func(method, path, body string) (int, string) {
		t.Helper()
		req := httptest.NewRequest(method, path, strings.NewReader(body))
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, req)
		return rec.Code, rec.Body.String()
	}
	if code, body := serve("POST", "/api/v1/users", `{"name":"Ada"}`); code != 201 || !strings.Contains(body, "Ada") {
		t.Errorf("POST = %d %q", code, body)
	}
	if code, body := serve("GET", "/api/v1/users", ""); code != 200 || !strings.Contains(body, "Ada") {
		t.Errorf("GET list = %d %q", code, body)
	}
	if code, _ := serve("GET", "/api/v1/users/99", ""); code != 404 {
		t.Errorf("GET missing = %d, want 404", code)
	}
	code, body := serve("GET", "/swagger/doc.json", "")
	if code != 200 {
		t.Fatalf("GET /swagger/doc.json = %d", code)
	}
	var doc map[string]any
	if err := json.Unmarshal([]byte(body), &doc); err != nil {
		t.Fatalf("doc.json is not JSON: %v", err)
	}
	if doc["openapi"] != "3.0.0" {
		t.Errorf("openapi version = %v", doc["openapi"])
	}
	paths, _ := doc["paths"].(map[string]any)
	if _, ok := paths["/api/v1/users"]; !ok {
		t.Errorf("doc.json missing /api/v1/users paths: %v", paths)
	}
}
