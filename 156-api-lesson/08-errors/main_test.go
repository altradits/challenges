package main

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestErrorHandlingAPI(t *testing.T) {
	h := CreateErrorHandlingAPI()
	if h == nil {
		t.Fatal("CreateErrorHandlingAPI returned nil")
	}
	serve := func(path string) (int, string) {
		t.Helper()
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, httptest.NewRequest("GET", path, nil))
		return rec.Code, rec.Body.String()
	}
	if code, body := serve("/ok"); code != 200 || !strings.Contains(body, `"status":"ok"`) {
		t.Errorf("/ok = %d %q", code, body)
	}
	code, body := serve("/fail")
	if code != 500 {
		t.Fatalf("/fail = %d, want 500", code)
	}
	var apiErr APIError
	if err := json.Unmarshal([]byte(body), &apiErr); err != nil {
		t.Fatalf("/fail body is not APIError JSON: %v (%q)", err, body)
	}
	if apiErr.Code != 500 || apiErr.Message == "" {
		t.Errorf("/fail error = %+v", apiErr)
	}
	code, body = serve("/panic")
	if code != 500 {
		t.Fatalf("/panic = %d, want recovered 500", code)
	}
	if err := json.Unmarshal([]byte(body), &apiErr); err != nil {
		t.Fatalf("/panic body is not APIError JSON: %v", err)
	}
}
