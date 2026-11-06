package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func doRequest(t *testing.T, h http.Handler, method, path, body string) (int, http.Header, string) {
	t.Helper()
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec.Code, rec.Header(), rec.Body.String()
}

func TestJSONHandlerCRUD(t *testing.T) {
	h := CreateJSONHandler()
	if h == nil {
		t.Fatal("CreateJSONHandler returned nil")
	}
	code, hdr, body := doRequest(t, h, "POST", "/api/v1/users", `{"name":"Ada","email":"ada@x.io"}`)
	if code != 201 {
		t.Fatalf("POST = %d, want 201 (%s)", code, body)
	}
	if ct := hdr.Get("Content-Type"); ct != "application/json" {
		t.Errorf("Content-Type = %q, want application/json", ct)
	}
	var created User
	if err := json.Unmarshal([]byte(body), &created); err != nil {
		t.Fatalf("POST body: %v", err)
	}
	if created.ID == 0 || created.Name != "Ada" {
		t.Fatalf("POST created %+v", created)
	}
	code, _, body = doRequest(t, h, "GET", "/api/v1/users/1", "")
	if code != 200 || !strings.Contains(body, "Ada") {
		t.Errorf("GET = %d %q", code, body)
	}
	code, _, _ = doRequest(t, h, "GET", "/api/v1/users/999", "")
	if code != 404 {
		t.Errorf("GET missing = %d, want 404", code)
	}
	code, _, body = doRequest(t, h, "PUT", "/api/v1/users/1", `{"name":"Bo"}`)
	if code != 200 || !strings.Contains(body, `"name":"Bo"`) {
		t.Errorf("PUT = %d %q", code, body)
	}
	code, _, _ = doRequest(t, h, "POST", "/api/v1/users", `not json`)
	if code != 400 {
		t.Errorf("POST invalid = %d, want 400", code)
	}
}
