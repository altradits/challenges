package middleware

import (
	"net/http"
	"strings"

	"github.com/altradits/altradits/internal/db"
)

type User struct {
	ID         string
	Identifier string
	Role       string
	FullName   string
}

func GetUser(r *http.Request) *User {
	token := extractToken(r)
	if token == "" {
		return nil
	}

	var u User
	err := db.DB.QueryRow(`
		SELECT u.id, u.identifier, u.role, COALESCE(u.full_name, '')
		FROM sessions s
		JOIN users u ON u.id = s.user_id
		WHERE s.token = $1 AND s.expires_at > NOW() AND u.is_blocked = FALSE`,
		token).Scan(&u.ID, &u.Identifier, &u.Role, &u.FullName)
	if err != nil {
		return nil
	}

	// Update last_seen
	db.DB.Exec("UPDATE users SET last_seen = NOW() WHERE id = $1", u.ID)
	return &u
}

func extractToken(r *http.Request) string {
	// From cookie first
	c, err := r.Cookie("altradits_session")
	if err == nil && c.Value != "" {
		return c.Value
	}
	// From Authorization header
	h := r.Header.Get("Authorization")
	if strings.HasPrefix(h, "Bearer ") {
		return strings.TrimPrefix(h, "Bearer ")
	}
	return ""
}

func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := GetUser(r)
		if u == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}

func RequireRole(role string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := GetUser(r)
		if u == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		if u.Role != role && u.Role != "admin" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next(w, r)
	}
}
