package main

import (
	"fmt"
)

// Task represents a task in the system
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// CreateTaskAPI creates a RESTful API for task management.
// Implement the following endpoints:
// - GET /api/v1/tasks -> 200 + list of tasks
// - GET /api/v1/tasks/{id} -> 200 + task or 404
// - POST /api/v1/tasks -> 201 + created task
// - PUT /api/v1/tasks/{id} -> 200 + updated task or 404
// - DELETE /api/v1/tasks/{id} -> 204
// Hint: Use an in-memory slice to store tasks
// Hint: Use http.StatusXxx constants for status codes
// Hint: Use encoding/json for JSON handling
func CreateTaskAPI() {
	// TODO: Implement this function
	// 1. Create in-memory storage for tasks
	// 2. Implement handlers for each endpoint
	// 3. Return appropriate status codes
	// 4. Start the server on port 8080
}

func main() {
	fmt.Println("Starting Task API on :8080...")
	CreateTaskAPI()
}
