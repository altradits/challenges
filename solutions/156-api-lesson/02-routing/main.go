package main

import (
	"fmt"
	"net/http"
	"strings"
)

// CreateRouter builds a router with multiple endpoints.
// Handle the following routes:
// - GET / -> "Welcome to the API"
// - GET /users -> "List of users"
// - GET /users/{id} -> "User {id}"
// - POST /users -> "Create user"
// - PUT /users/{id} -> "Update user {id}"
// - DELETE /users/{id} -> "Delete user {id}"
// Hint: Check r.Method to handle different HTTP methods
// Hint: Use strings.TrimPrefix to extract path parameters
// Hint: Return the handler; main() starts the server
func CreateRouter() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		switch {
		case path == "/" && r.Method == http.MethodGet:
			fmt.Fprint(w, "Welcome to the API")
		case path == "/users" && r.Method == http.MethodGet:
			fmt.Fprint(w, "List of users")
		case path == "/users" && r.Method == http.MethodPost:
			fmt.Fprint(w, "Create user")
		case strings.HasPrefix(path, "/users/"):
			id := strings.TrimPrefix(path, "/users/")
			switch r.Method {
			case http.MethodGet:
				fmt.Fprintf(w, "User %s", id)
			case http.MethodPut:
				fmt.Fprintf(w, "Update user %s", id)
			case http.MethodDelete:
				fmt.Fprintf(w, "Delete user %s", id)
			default:
				http.NotFound(w, r)
			}
		default:
			http.NotFound(w, r)
		}
	})
}

func main() {
	fmt.Println("Starting router on :8080...")
	if err := http.ListenAndServe(":8080", CreateRouter()); err != nil {
		fmt.Println("server error:", err)
	}
}
