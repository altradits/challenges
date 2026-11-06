package main

import (
	"database/sql"
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func testDB(t *testing.T) *sql.DB {
	t.Helper()
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { db.Close() })
	if err := InitDB(db); err != nil {
		t.Fatalf("InitDB: %v", err)
	}
	return db
}

func TestDatabaseAPI(t *testing.T) {
	h := CreateDatabaseAPI(testDB(t))
	if h == nil {
		t.Fatal("CreateDatabaseAPI returned nil")
	}
	serve := func(method, path, body string) (int, string) {
		t.Helper()
		req := httptest.NewRequest(method, path, strings.NewReader(body))
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, req)
		return rec.Code, rec.Body.String()
	}
	code, body := serve("POST", "/api/v1/users", `{"name":"Ada","email":"ada@x.io"}`)
	if code != 201 {
		t.Fatalf("POST = %d, want 201 (%s)", code, body)
	}
	var created User
	if err := json.Unmarshal([]byte(body), &created); err != nil || created.Name != "Ada" || created.ID == 0 {
		t.Fatalf("POST created %q: %v", body, err)
	}
	if created.CreatedAt == "" {
		t.Error("CreatedAt should default to CURRENT_TIMESTAMP")
	}
	if code, body := serve("GET", "/api/v1/users", ""); code != 200 || !strings.Contains(body, "Ada") {
		t.Errorf("GET list = %d %q", code, body)
	}
	if code, _ := serve("GET", "/api/v1/users/999", ""); code != 404 {
		t.Errorf("GET missing = %d, want 404", code)
	}
	if code, body := serve("PUT", "/api/v1/users/1", `{"name":"Bo","email":"bo@x.io"}`); code != 200 || !strings.Contains(body, "Bo") {
		t.Errorf("PUT = %d %q", code, body)
	}
	if code, _ := serve("DELETE", "/api/v1/users/1", ""); code != 204 {
		t.Errorf("DELETE = %d, want 204", code)
	}
	if code, _ := serve("GET", "/api/v1/users/1", ""); code != 404 {
		t.Errorf("GET after delete = %d, want 404", code)
	}
}
