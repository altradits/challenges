package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

// User represents a user in the system
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"` // Never serialize password
	Role     string `json:"role"`
}

// Claims represents JWT claims
type Claims struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type ctxKey string

const claimsKey ctxKey = "claims"

// InitAuthDB creates the users table if it does not exist.
func InitAuthDB(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		role TEXT NOT NULL DEFAULT 'user'
	)`)
	return err
}

// CreateAuthAPI creates an API with JWT authentication.
// Implement the following:
// - POST /api/v1/register -> Create user with hashed password (201)
// - POST /api/v1/login -> Return JWT token (200)
// - GET /api/v1/profile -> Protected route, 401 without valid token
// - GET /api/v1/admin -> Admin-only route, 403 for non-admins
// Hint: Use bcrypt for password hashing
// Hint: Use jwt.NewWithClaims for token creation
// Hint: Create middleware to validate tokens
// Hint: Use context to pass user info to handlers
// Hint: Return the handler; main() starts the server
func CreateAuthAPI(db *sql.DB, secret string) http.Handler {
	writeJSON := func(w http.ResponseWriter, code int, v any) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(v)
	}
	withAuth := func(requireAdmin bool, next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			tokenStr := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
			if tokenStr == "" || tokenStr == r.Header.Get("Authorization") {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}
			claims := &Claims{}
			token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
				return []byte(secret), nil
			})
			if err != nil || !token.Valid {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}
			if requireAdmin && claims.Role != "admin" {
				http.Error(w, "forbidden", http.StatusForbidden)
				return
			}
			next(w, r.WithContext(context.WithValue(r.Context(), claimsKey, claims)))
		}
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var in struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
			Role     string `json:"role"`
		}
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil || in.Email == "" || in.Password == "" {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		if in.Role == "" {
			in.Role = "user"
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
		res, err := db.Exec("INSERT INTO users (name, email, password, role) VALUES (?, ?, ?, ?)",
			in.Name, in.Email, string(hash), in.Role)
		if err != nil {
			http.Error(w, "cannot register", http.StatusConflict)
			return
		}
		id, _ := res.LastInsertId()
		writeJSON(w, http.StatusCreated, User{ID: int(id), Name: in.Name, Email: in.Email, Role: in.Role})
	})
	mux.HandleFunc("/api/v1/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var in struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		var u User
		var hash string
		err := db.QueryRow("SELECT id, name, email, password, role FROM users WHERE email = ?", in.Email).
			Scan(&u.ID, &u.Name, &u.Email, &hash, &u.Role)
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(in.Password)); err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
			UserID:           u.ID,
			Role:             u.Role,
			RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour))},
		})
		signed, err := token.SignedString([]byte(secret))
		if err != nil {
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{"token": signed})
	})
	mux.HandleFunc("/api/v1/profile", withAuth(false, func(w http.ResponseWriter, r *http.Request) {
		claims := r.Context().Value(claimsKey).(*Claims)
		var u User
		var hash string
		err := db.QueryRow("SELECT id, name, email, password, role FROM users WHERE id = ?", claims.UserID).
			Scan(&u.ID, &u.Name, &u.Email, &hash, &u.Role)
		if err != nil {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		writeJSON(w, http.StatusOK, u)
	}))
	mux.HandleFunc("/api/v1/admin", withAuth(true, func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"message": "welcome admin"})
	}))
	return mux
}

func main() {
	fmt.Println("Starting auth API on :8080...")
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		fmt.Println("db error:", err)
		return
	}
	defer db.Close()
	if err := InitAuthDB(db); err != nil {
		fmt.Println("init error:", err)
		return
	}
	if err := http.ListenAndServe(":8080", CreateAuthAPI(db, "dev-secret-change-me")); err != nil {
		fmt.Println("server error:", err)
	}
}
