package main

import (
	"net/http"
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
	// TODO: Implement this function
	// 1. Create handler function
	// 2. Handle different HTTP methods
	// 3. Use store for data access
	// 4. Return JSON responses
	return nil
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
