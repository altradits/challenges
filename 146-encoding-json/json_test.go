package piscine

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestMarshalUser(t *testing.T) {
	u := User{ID: 1, Name: "Ada", Email: "ada@example.com",
		CreatedAt: time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC), Password: "secret"}
	raw, err := MarshalUser(u)
	if err != nil {
		t.Fatalf("MarshalUser: %v", err)
	}
	s := string(raw)
	for _, want := range []string{`"id":1`, `"name":"Ada"`, `"email":"ada@example.com"`, `"created_at":"2026-01-02T03:04:05Z"`} {
		if !strings.Contains(s, want) {
			t.Errorf("marshaled %s: missing %s", s, want)
		}
	}
	if strings.Contains(s, "secret") || strings.Contains(s, "password") {
		t.Errorf("marshaled %s: password must never be serialized", s)
	}
}

func TestUnmarshalUser(t *testing.T) {
	u, err := UnmarshalUser([]byte(`{"id":2,"name":"Bo","email":"bo@x.io","created_at":"2026-05-06T00:00:00Z"}`))
	if err != nil {
		t.Fatalf("UnmarshalUser: %v", err)
	}
	if u.ID != 2 || u.Name != "Bo" || u.Email != "bo@x.io" {
		t.Errorf("UnmarshalUser = %+v", u)
	}
	if u.CreatedAt.Year() != 2026 || u.CreatedAt.Month() != 5 {
		t.Errorf("UnmarshalUser time = %v", u.CreatedAt)
	}
	if _, err := UnmarshalUser([]byte(`{oops`)); err == nil {
		t.Error("UnmarshalUser of invalid JSON should return an error")
	}
}

func TestUserRoundTrip(t *testing.T) {
	in := User{ID: 7, Name: "Cy", Email: "cy@x.io", CreatedAt: time.Now().UTC().Truncate(time.Second)}
	raw, err := MarshalUser(in)
	if err != nil {
		t.Fatal(err)
	}
	out, err := UnmarshalUser(raw)
	if err != nil {
		t.Fatal(err)
	}
	if out != in {
		t.Errorf("round trip = %+v, want %+v", out, in)
	}
}

func TestPrettyPrint(t *testing.T) {
	var buf bytes.Buffer
	if err := PrettyPrint(&buf, map[string]int{"b": 2}); err != nil {
		t.Fatalf("PrettyPrint: %v", err)
	}
	if !strings.Contains(buf.String(), "\n  ") {
		t.Errorf("PrettyPrint should indent, got %q", buf.String())
	}
}
