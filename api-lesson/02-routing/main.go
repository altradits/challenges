package main

import (
	"fmt"
)

// CreateRouter creates a router with multiple endpoints.
// Handle the following routes:
// - GET / -> "Welcome to the API"
// - GET /users -> "List of users"
// - GET /users/{id} -> "User {id}"
// - POST /users -> "Create user"
// - PUT /users/{id} -> "Update user {id}"
// - DELETE /users/{id} -> "Delete user {id}"
// Hint: Use http.HandleFunc for each route
// Hint: Check r.Method to handle different HTTP methods
// Hint: Use strings.TrimPrefix to extract path parameters
func CreateRouter() {
	// TODO: Implement this function
	// 1. Register handlers for each route
	// 2. Handle different HTTP methods
	// 3. Extract user ID from paths like /users/123
	// 4. Start the server on port 8080
}

func main() {
	fmt.Println("Starting router on :8080...")
	CreateRouter()
}
