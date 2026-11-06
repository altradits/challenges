package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func doRequest(t *testing.T, h http.Handler, method, path, body string) (int, string) {
	t.Helper()
	var rdr *strings.Reader
	if body != "" {
		rdr = strings.NewReader(body)
	} else {
		rdr = strings.NewReader("")
	}
	req := httptest.NewRequest(method, path, rdr)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec.Code, rec.Body.String()
}

func TestTaskAPICRUD(t *testing.T) {
	h := CreateTaskAPI()
	if h == nil {
		t.Fatal("CreateTaskAPI returned nil")
	}
	code, body := doRequest(t, h, "POST", "/api/v1/tasks", `{"title":"write tests"}`)
	if code != http.StatusCreated {
		t.Fatalf("POST status = %d, want 201 (%s)", code, body)
	}
	var created Task
	if err := json.Unmarshal([]byte(body), &created); err != nil {
		t.Fatalf("POST body is not a task: %v", err)
	}
	if created.ID == 0 || created.Title != "write tests" {
		t.Fatalf("POST created %+v", created)
	}
	code, body = doRequest(t, h, "GET", "/api/v1/tasks", "")
	if code != 200 || !strings.Contains(body, "write tests") {
		t.Errorf("GET list = %d %q", code, body)
	}
	code, body = doRequest(t, h, "GET", "/api/v1/tasks/1", "")
	if code != 200 || !strings.Contains(body, "write tests") {
		t.Errorf("GET one = %d %q", code, body)
	}
	code, _ = doRequest(t, h, "GET", "/api/v1/tasks/999", "")
	if code != 404 {
		t.Errorf("GET missing = %d, want 404", code)
	}
	code, body = doRequest(t, h, "PUT", "/api/v1/tasks/1", `{"title":"write tests","done":true}`)
	if code != 200 || !strings.Contains(body, `"done":true`) {
		t.Errorf("PUT = %d %q", code, body)
	}
	code, _ = doRequest(t, h, "DELETE", "/api/v1/tasks/1", "")
	if code != 204 {
		t.Errorf("DELETE = %d, want 204", code)
	}
	code, _ = doRequest(t, h, "GET", "/api/v1/tasks/1", "")
	if code != 404 {
		t.Errorf("GET after delete = %d, want 404", code)
	}
}
