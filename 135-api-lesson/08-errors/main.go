package main

import (
	"fmt"
)

// APIError represents an API error response
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error implements the error interface
func (e *APIError) Error() string {
	return e.Message
}

// CreateErrorHandlingAPI creates an API with proper error handling.
// Implement the following:
// - Custom APIError type with Code and Message
// - Structured logging for all requests
// - Consistent error responses (JSON format)
// - Panic recovery middleware
// Hint: Use log/slog for structured logging
// Hint: Create a writeError helper function
// Hint: Use defer/recover in middleware
// Hint: Log request method, path, and status
func CreateErrorHandlingAPI() {
	// TODO: Implement this function
	// 1. Create custom error type
	// 2. Implement logging middleware
	// 3. Implement error response helper
	// 4. Implement panic recovery
	// 5. Add test endpoints that trigger errors
	// 6. Start the server on port 8080
}

func main() {
	fmt.Println("Starting error handling API on :8080...")
	CreateErrorHandlingAPI()
}
