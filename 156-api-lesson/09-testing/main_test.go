package main

import (
	"encoding/json"
	"errors"
	"net/http/httptest"
	"strings"
	"testing"
)

type stubStore struct {
	users  []User
	nextID int
	fail   bool
}

func (s *stubStore) GetAll() ([]User, error) {
	if s.fail {
		return nil, errors.New("boom")
	}
	return s.users, nil
}

func (s *stubStore) Get(id int) (*User, error) {
	if s.fail {
		return nil, errors.New("boom")
	}
	for i := range s.users {
		if s.users[i].ID == id {
			return &s.users[i], nil
		}
	}
	return nil, nil
}

func (s *stubStore) Create(user *User) error {
	if s.fail {
		return errors.New("boom")
	}
	user.ID = s.nextID
	s.nextID++
	s.users = append(s.users, *user)
	return nil
}

func TestCreateUserHandler(t *testing.T) {
	h := CreateUserHandler(&stubStore{})
	if h == nil {
		t.Fatal("CreateUserHandler returned nil")
	}
	serve := func(method, path, body string) (int, string) {
		t.Helper()
		req := httptest.NewRequest(method, path, strings.NewReader(body))
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, req)
		return rec.Code, rec.Body.String()
	}
	code, body := serve("POST", "/users", `{"name":"Ada","email":"ada@x.io"}`)
	if code != 201 {
		t.Fatalf("POST = %d, want 201 (%s)", code, body)
	}
	var created User
	if err := json.Unmarshal([]byte(body), &created); err != nil || created.Name != "Ada" {
		t.Fatalf("POST created %q: %v", body, err)
	}
	code, body = serve("GET", "/users", "")
	if code != 200 || !strings.Contains(body, "Ada") {
		t.Errorf("GET list = %d %q", code, body)
	}
	code, body = serve("GET", "/users/0", "")
	if code != 200 || !strings.Contains(body, "Ada") {
		t.Errorf("GET one = %d %q", code, body)
	}
	if code, _ := serve("GET", "/users/99", ""); code != 404 {
		t.Errorf("GET missing = %d, want 404", code)
	}
	if code, _ := serve("GET", "/users/abc", ""); code != 400 {
		t.Errorf("GET bad id = %d, want 400", code)
	}
	if code, _ := serve("POST", "/users", `nope`); code != 400 {
		t.Errorf("POST invalid = %d, want 400", code)
	}
}

func TestCreateUserHandlerStoreError(t *testing.T) {
	h := CreateUserHandler(&stubStore{fail: true})
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("GET", "/users", nil))
	if rec.Code != 500 {
		t.Errorf("store failure = %d, want 500", rec.Code)
	}
}
