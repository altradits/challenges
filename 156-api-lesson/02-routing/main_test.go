package main

import (
	"net/http/httptest"
	"testing"
)

func TestCreateRouter(t *testing.T) {
	h := CreateRouter()
	if h == nil {
		t.Fatal("CreateRouter returned nil")
	}
	cases := []struct {
		method, path string
		code         int
		body         string
	}{
		{"GET", "/", 200, "Welcome to the API"},
		{"GET", "/users", 200, "List of users"},
		{"GET", "/users/123", 200, "User 123"},
		{"POST", "/users", 200, "Create user"},
		{"PUT", "/users/7", 200, "Update user 7"},
		{"DELETE", "/users/9", 200, "Delete user 9"},
		{"GET", "/nope", 404, ""},
		{"POST", "/", 404, ""},
	}
	for _, tc := range cases {
		req := httptest.NewRequest(tc.method, tc.path, nil)
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, req)
		if rec.Code != tc.code {
			t.Errorf("%s %s: status = %d, want %d", tc.method, tc.path, rec.Code, tc.code)
		}
		if tc.body != "" && rec.Body.String() != tc.body {
			t.Errorf("%s %s: body = %q, want %q", tc.method, tc.path, rec.Body.String(), tc.body)
		}
	}
}
