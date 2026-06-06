package main

import (
	"testing"
)

// TestUserHandler tests the user handler
func TestUserHandler(t *testing.T) {
	// TODO: Implement this function
	// 1. Create a MemoryStore
	// 2. Create the handler with the store
	// 3. Write table-driven tests for:
	//    - GET /users returns 200
	//    - POST /users with valid data returns 201
	//    - POST /users with invalid data returns 400
	//    - GET /users/999 returns 404
}

// Test cases to implement:
// 1. GET /users - should return 200 and empty array
// 2. POST /users - should return 201 and created user
// 3. POST /users with invalid JSON - should return 400
// 4. GET /users/1 - should return 200 and user
// 5. GET /users/999 - should return 404
