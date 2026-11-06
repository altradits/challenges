package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// User represents a user in the system
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserStore defines the interface for user storage
type UserStore interface {
	GetAll() ([]User, error)
	Get(id int) (*User, error)
	Create(user *User) error
}

// CreateUserHandler creates a handler for user endpoints.
// Implement:
// - GET /users -> return all users
// - GET /users/{id} -> return user by ID
// - POST /users -> create new user
// Hint: Use the UserStore interface for data access
// Hint: Return appropriate status codes
// Hint: Handle errors gracefully
func CreateUserHandler(store UserStore) http.Handler {
	if store == nil {
		return nil
	}
	writeJSON := func(w http.ResponseWriter, code int, v any) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(v)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/users" {
			switch r.Method {
			case http.MethodGet:
				users, err := store.GetAll()
				if err != nil {
					http.Error(w, "store error", http.StatusInternalServerError)
					return
				}
				if users == nil {
					users = []User{}
				}
				writeJSON(w, http.StatusOK, users)
			case http.MethodPost:
				var u User
				if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
					http.Error(w, "bad request", http.StatusBadRequest)
					return
				}
				if err := store.Create(&u); err != nil {
					http.Error(w, "store error", http.StatusInternalServerError)
					return
				}
				writeJSON(w, http.StatusCreated, u)
			default:
				http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			}
			return
		}
		if !strings.HasPrefix(r.URL.Path, "/users/") || r.Method != http.MethodGet {
			http.NotFound(w, r)
			return
		}
		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/users/"))
		if err != nil {
			http.Error(w, "bad id", http.StatusBadRequest)
			return
		}
		u, err := store.Get(id)
		if err != nil {
			http.Error(w, "store error", http.StatusInternalServerError)
			return
		}
		if u == nil {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		writeJSON(w, http.StatusOK, u)
	})
}

// In-memory implementation of UserStore
type MemoryStore struct {
	users  []User
	nextID int
}

func (s *MemoryStore) GetAll() ([]User, error) {
	return s.users, nil
}

func (s *MemoryStore) Get(id int) (*User, error) {
	for i := range s.users {
		if s.users[i].ID == id {
			return &s.users[i], nil
		}
	}
	return nil, nil
}

func (s *MemoryStore) Create(user *User) error {
	user.ID = s.nextID
	s.nextID++
	s.users = append(s.users, *user)
	return nil
}

func main() {
	fmt.Println("Starting testable API on :8080...")
	_ = http.ListenAndServe(":8080", CreateUserHandler(&MemoryStore{}))
}
