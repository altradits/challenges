package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

// User represents a user in the system
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}

// CreateJSONHandler builds handlers that properly handle JSON.
// Implement the following endpoints:
// - POST /api/v1/users -> Create user from JSON, return created user with 201
// - GET /api/v1/users/{id} -> Return user as JSON with 200, or 404
// - PUT /api/v1/users/{id} -> Update user from JSON, return updated user
// Hint: Use json.NewDecoder(r.Body).Decode() to read JSON
// Hint: Use json.Marshal() to write JSON response
// Hint: Set Content-Type header to application/json
// Hint: Return the handler; main() starts the server
func CreateJSONHandler() http.Handler {
	var mu sync.Mutex
	users := map[int]User{}
	nextID := 1
	writeJSON := func(w http.ResponseWriter, code int, v any) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(v)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		if r.URL.Path == "/api/v1/users" && r.Method == http.MethodPost {
			var in struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			}
			if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
				http.Error(w, "bad request", http.StatusBadRequest)
				return
			}
			u := User{ID: nextID, Name: in.Name, Email: in.Email}
			nextID++
			users[u.ID] = u
			writeJSON(w, http.StatusCreated, u)
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
		u, ok := users[id]
		if !ok {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		switch r.Method {
		case http.MethodGet:
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
			u.Name, u.Email = in.Name, in.Email
			users[id] = u
			writeJSON(w, http.StatusOK, u)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

func main() {
	fmt.Println("Starting JSON API on :8080...")
	if err := http.ListenAndServe(":8080", CreateJSONHandler()); err != nil {
		fmt.Println("server error:", err)
	}
}
