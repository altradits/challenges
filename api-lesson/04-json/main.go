package main

import (
	"fmt"
)

// User represents a user in the system
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}

// CreateJSONHandler creates handlers that properly handle JSON.
// Implement the following endpoints:
// - POST /api/v1/users -> Create user from JSON, return created user with 201
// - GET /api/v1/users/{id} -> Return user as JSON with 200, or 404
// - PUT /api/v1/users/{id} -> Update user from JSON, return updated user
// Hint: Use json.NewDecoder(r.Body).Decode() to read JSON
// Hint: Use json.Marshal() to write JSON response
// Hint: Set Content-Type header to application/json
func CreateJSONHandler() {
	// TODO: Implement this function
	// 1. Create in-memory storage for users
	// 2. Implement handlers for each endpoint
	// 3. Handle JSON encoding/decoding
	// 4. Set proper headers
	// 5. Start the server on port 8080
}

func main() {
	fmt.Println("Starting JSON API on :8080...")
	CreateJSONHandler()
}
