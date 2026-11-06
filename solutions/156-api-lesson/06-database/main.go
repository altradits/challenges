package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// User represents a user in the system
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

// InitDB creates the users table if it does not exist.
func InitDB(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`)
	return err
}

// CreateDatabaseAPI builds a RESTful API backed by the database.
// Implement the following endpoints:
// - POST /api/v1/users -> Create user, save to database, return 201
// - GET /api/v1/users -> Return all users from database
// - GET /api/v1/users/{id} -> Return user from database, or 404
// - PUT /api/v1/users/{id} -> Update user in database
// - DELETE /api/v1/users/{id} -> Delete user from database, return 204
// Hint: Use db.Query() for SELECT
// Hint: Use db.Exec() for INSERT/UPDATE/DELETE
// Hint: Use db.QueryRow() for single-row queries
// Hint: Return the handler; main() starts the server
func CreateDatabaseAPI(db *sql.DB) http.Handler {
	writeJSON := func(w http.ResponseWriter, code int, v any) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(v)
	}
	scanUser := func(row *sql.Row) (User, error) {
		var u User
		err := row.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)
		return u, err
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/users" {
			switch r.Method {
			case http.MethodGet:
				rows, err := db.Query("SELECT id, name, email, created_at FROM users ORDER BY id")
				if err != nil {
					http.Error(w, "db error", http.StatusInternalServerError)
					return
				}
				defer rows.Close()
				users := []User{}
				for rows.Next() {
					var u User
					if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt); err != nil {
						http.Error(w, "db error", http.StatusInternalServerError)
						return
					}
					users = append(users, u)
				}
				writeJSON(w, http.StatusOK, users)
			case http.MethodPost:
				var in struct {
					Name  string `json:"name"`
					Email string `json:"email"`
				}
				if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
					http.Error(w, "bad request", http.StatusBadRequest)
					return
				}
				res, err := db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", in.Name, in.Email)
				if err != nil {
					http.Error(w, "db error", http.StatusInternalServerError)
					return
				}
				id, _ := res.LastInsertId()
				u, err := scanUser(db.QueryRow("SELECT id, name, email, created_at FROM users WHERE id = ?", id))
				if err != nil {
					http.Error(w, "db error", http.StatusInternalServerError)
					return
				}
				writeJSON(w, http.StatusCreated, u)
			default:
				http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			}
			return
		}
		if !strings.HasPrefix(r.URL.Path, "/api/v1/users/") {
			http.NotFound(w, r)
			return
		}
		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/v1/users/"))
		if err != nil {
			http.Error(w, "bad id", http.StatusBadRequest)
			return
		}
		switch r.Method {
		case http.MethodGet:
			u, err := scanUser(db.QueryRow("SELECT id, name, email, created_at FROM users WHERE id = ?", id))
			if err == sql.ErrNoRows {
				http.Error(w, "not found", http.StatusNotFound)
				return
			}
			if err != nil {
				http.Error(w, "db error", http.StatusInternalServerError)
				return
			}
			writeJSON(w, http.StatusOK, u)
		case http.MethodPut:
			var in struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			}
			if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
				http.Error(w, "bad request", http.StatusBadRequest)
				return
			}
			res, err := db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", in.Name, in.Email, id)
			if err != nil {
				http.Error(w, "db error", http.StatusInternalServerError)
				return
			}
			if n, _ := res.RowsAffected(); n == 0 {
				http.Error(w, "not found", http.StatusNotFound)
				return
			}
			u, _ := scanUser(db.QueryRow("SELECT id, name, email, created_at FROM users WHERE id = ?", id))
			writeJSON(w, http.StatusOK, u)
		case http.MethodDelete:
			res, err := db.Exec("DELETE FROM users WHERE id = ?", id)
			if err != nil {
				http.Error(w, "db error", http.StatusInternalServerError)
				return
			}
			if n, _ := res.RowsAffected(); n == 0 {
				http.Error(w, "not found", http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

func main() {
	fmt.Println("Starting database API on :8080...")
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		fmt.Println("db error:", err)
		return
	}
	defer db.Close()
	if err := InitDB(db); err != nil {
		fmt.Println("init error:", err)
		return
	}
	if err := http.ListenAndServe(":8080", CreateDatabaseAPI(db)); err != nil {
		fmt.Println("server error:", err)
	}
}
