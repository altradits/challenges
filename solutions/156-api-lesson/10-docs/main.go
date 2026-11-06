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
	Email string `json:"email"`
}

const openAPIDoc = `{
  "openapi": "3.0.0",
  "info": {"title": "Users API", "version": "1.0.0"},
  "paths": {
    "/api/v1/users": {
      "get": {"summary": "List all users", "responses": {"200": {"description": "OK"}}},
      "post": {"summary": "Create a user", "responses": {"201": {"description": "Created"}}}
    },
    "/api/v1/users/{id}": {
      "get": {"summary": "Get a user", "responses": {"200": {"description": "OK"}, "404": {"description": "Not found"}}}
    }
  }
}`

// CreateDocumentedAPI builds an API with OpenAPI documentation.
// Add swaggo comments to all handlers.
// Implement:
// - GET /api/v1/users -> List users
// - POST /api/v1/users -> Create user
// - GET /api/v1/users/{id} -> Get user
// - GET /swagger/doc.json -> OpenAPI specification for the API above
// Hint: Use swaggo comments for documentation
// Hint: Add @Summary, @Description, @Param, @Success tags
// Hint: Serve the OpenAPI JSON at /swagger/doc.json
// Hint: Return the handler; main() starts the server
func CreateDocumentedAPI() http.Handler {
	var mu sync.Mutex
	users := map[int]User{}
	nextID := 1
	writeJSON := func(w http.ResponseWriter, code int, v any) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(v)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(openAPIDoc))
	})
	mux.HandleFunc("/api/v1/users", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		switch r.Method {
		case http.MethodGet:
			list := []User{}
			for _, u := range users {
				list = append(list, u)
			}
			writeJSON(w, http.StatusOK, list)
		case http.MethodPost:
			var u User
			if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
				writeJSON(w, http.StatusBadRequest, APIError{Code: 400, Message: "bad request"})
				return
			}
			u.ID = nextID
			nextID++
			users[u.ID] = u
			writeJSON(w, http.StatusCreated, u)
		default:
			writeJSON(w, http.StatusMethodNotAllowed, APIError{Code: 405, Message: "method not allowed"})
		}
	})
	mux.HandleFunc("/api/v1/users/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		if r.Method != http.MethodGet {
			writeJSON(w, http.StatusMethodNotAllowed, APIError{Code: 405, Message: "method not allowed"})
			return
		}
		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/v1/users/"))
		if err != nil {
			writeJSON(w, http.StatusBadRequest, APIError{Code: 400, Message: "bad id"})
			return
		}
		u, ok := users[id]
		if !ok {
			writeJSON(w, http.StatusNotFound, APIError{Code: 404, Message: "not found"})
			return
		}
		writeJSON(w, http.StatusOK, u)
	})
	return mux
}

// @Summary List all users
// @Description Get a list of all users
// @Accept json
// @Produce json
// @Success 200 {array} User
// @Router /api/v1/users [get]
func getUsers(w http.ResponseWriter, r *http.Request) {
	// Implemented inside CreateDocumentedAPI.
}

// @Summary Create a user
// @Description Create a new user
// @Accept json
// @Produce json
// @Param user body User true "User data"
// @Success 201 {object} User
// @Failure 400 {object} APIError
// @Router /api/v1/users [post]
func createUser(w http.ResponseWriter, r *http.Request) {
	// Implemented inside CreateDocumentedAPI.
}

// @Summary Get a user
// @Description Get a user by ID
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Failure 404 {object} APIError
// @Router /api/v1/users/{id} [get]
func getUser(w http.ResponseWriter, r *http.Request) {
	// Implemented inside CreateDocumentedAPI.
}

// APIError represents an error response
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {
	fmt.Println("Starting documented API on :8080...")
	if err := http.ListenAndServe(":8080", CreateDocumentedAPI()); err != nil {
		fmt.Println("server error:", err)
	}
}
