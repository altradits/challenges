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

// CreateDocumentedAPI creates an API with OpenAPI documentation.
// Add swaggo comments to all handlers.
// Implement:
// - GET /api/v1/users -> List users
// - POST /api/v1/users -> Create user
// - GET /api/v1/users/{id} -> Get user
// Hint: Use swaggo comments for documentation
// Hint: Add @Summary, @Description, @Param, @Success tags
// Hint: Serve Swagger UI at /swagger/index.html
func CreateDocumentedAPI() {
	// TODO: Implement this function
	// 1. Add swaggo comments to handlers
	// 2. Implement user handlers
	// 3. Generate and serve Swagger UI
	// 4. Start the server on port 8080
}

// @Summary List all users
// @Description Get a list of all users
// @Accept json
// @Produce json
// @Success 200 {array} User
// @Router /api/v1/users [get]
func getUsers(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement
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
	// TODO: Implement
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
	// TODO: Implement
}

// APIError represents an error response
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
