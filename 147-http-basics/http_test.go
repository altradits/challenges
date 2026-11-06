package piscine

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestJSONHandler(t *testing.T) {
	rec := httptest.NewRecorder()
	JSONHandler(map[string]string{"a": "b"})(rec, httptest.NewRequest("GET", "/", nil))
	if ct := rec.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("Content-Type = %q, want application/json", ct)
	}
	if !strings.Contains(rec.Body.String(), `"a":"b"`) {
		t.Errorf("body = %q", rec.Body.String())
	}
}

func TestRouter(t *testing.T) {
	router := Router()
	serve := func(method, target, body string) (int, string) {
		t.Helper()
		var rdr io.Reader
		if body != "" {
			rdr = strings.NewReader(body)
		}
		req := httptest.NewRequest(method, target, rdr)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)
		return rec.Code, rec.Body.String()
	}
	if code, body := serve("GET", "/health", ""); code != 200 || !strings.Contains(body, `"status":"ok"`) {
		t.Errorf("GET /health = %d %q", code, body)
	}
	if code, body := serve("GET", "/echo?msg=hi", ""); code != 200 || !strings.Contains(body, `"message":"hi"`) {
		t.Errorf("GET /echo = %d %q", code, body)
	}
	if code, body := serve("POST", "/reverse", `{"text":"abc"}`); code != 200 || !strings.Contains(body, `"result":"cba"`) {
		t.Errorf("POST /reverse = %d %q", code, body)
	}
	if code, _ := serve("POST", "/reverse", `nope`); code != http.StatusBadRequest {
		t.Errorf("POST /reverse invalid = %d, want 400", code)
	}
	if code, _ := serve("POST", "/health", ""); code != http.StatusMethodNotAllowed {
		t.Errorf("POST /health = %d, want 405", code)
	}
}

func TestFetchJSON(t *testing.T) {
	srv := httptest.NewServer(JSONHandler(map[string]int{"n": 42}))
	defer srv.Close()
	var v map[string]int
	if err := FetchJSON(srv.URL, &v); err != nil {
		t.Fatalf("FetchJSON: %v", err)
	}
	if v["n"] != 42 {
		t.Errorf("FetchJSON decoded %v", v)
	}
	if err := FetchJSON("http://127.0.0.1:1/none", &v); err == nil {
		t.Error("FetchJSON of dead server should return an error")
	}
}
