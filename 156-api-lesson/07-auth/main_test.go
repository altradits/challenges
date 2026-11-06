package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
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
	if err := InitAuthDB(db); err != nil {
		t.Fatalf("InitAuthDB: %v", err)
	}
	return db
}

func serve(t *testing.T, h http.Handler, method, path, body, token string) (int, string) {
	t.Helper()
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec.Code, rec.Body.String()
}

func register(t *testing.T, h http.Handler, name, email, password, role string) {
	t.Helper()
	code, body := serve(t, h, "POST", "/api/v1/register",
		`{"name":"`+name+`","email":"`+email+`","password":"`+password+`","role":"`+role+`"}`, "")
	if code != 201 {
		t.Fatalf("register %s = %d, want 201 (%s)", email, code, body)
	}
	if strings.Contains(body, password) {
		t.Fatalf("register response leaks password: %s", body)
	}
}

func login(t *testing.T, h http.Handler, email, password string) string {
	t.Helper()
	code, body := serve(t, h, "POST", "/api/v1/login",
		`{"email":"`+email+`","password":"`+password+`"}`, "")
	if code != 200 {
		t.Fatalf("login %s = %d, want 200 (%s)", email, code, body)
	}
	var out struct {
		Token string `json:"token"`
	}
	if err := json.Unmarshal([]byte(body), &out); err != nil || out.Token == "" {
		t.Fatalf("login response has no token: %q", body)
	}
	return out.Token
}

func TestAuthFlow(t *testing.T) {
	h := CreateAuthAPI(testDB(t), "test-secret")
	if h == nil {
		t.Fatal("CreateAuthAPI returned nil")
	}
	register(t, h, "Root", "root@x.io", "s3cret!", "admin")
	register(t, h, "Ada", "ada@x.io", "password1", "user")

	if code, _ := serve(t, h, "POST", "/api/v1/login", `{"email":"ada@x.io","password":"wrong"}`, ""); code != 401 {
		t.Errorf("wrong password = %d, want 401", code)
	}
	userToken := login(t, h, "ada@x.io", "password1")
	adminToken := login(t, h, "root@x.io", "s3cret!")

	if code, _ := serve(t, h, "GET", "/api/v1/profile", "", ""); code != 401 {
		t.Errorf("profile without token = %d, want 401", code)
	}
	if code, _ := serve(t, h, "GET", "/api/v1/profile", "", "bogus"); code != 401 {
		t.Errorf("profile with bad token = %d, want 401", code)
	}
	code, body := serve(t, h, "GET", "/api/v1/profile", "", userToken)
	if code != 200 || !strings.Contains(body, "ada@x.io") {
		t.Errorf("profile = %d %q", code, body)
	}
	if strings.Contains(body, "password1") {
		t.Errorf("profile leaks password hash material: %q", body)
	}
	if code, _ := serve(t, h, "GET", "/api/v1/admin", "", userToken); code != 403 {
		t.Errorf("admin as user = %d, want 403", code)
	}
	if code, body := serve(t, h, "GET", "/api/v1/admin", "", adminToken); code != 200 || !strings.Contains(body, "admin") {
		t.Errorf("admin as admin = %d %q", code, body)
	}
}
