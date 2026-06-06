package main

import (
	"fmt"
)

// CreateMiddlewareChain creates a server with middleware.
// Implement the following middleware:
// - Logging: Log method, path, and status code
// - CORS: Add Access-Control-Allow-Origin: * header
// - Recovery: Recover from panics and return 500
// Hint: Use http.HandlerFunc to convert function to handler
// Hint: Use a custom ResponseWriter to capture status code
// Hint: Use defer/recover for panic recovery
func CreateMiddlewareChain() {
	// TODO: Implement this function
	// 1. Create logging middleware
	// 2. Create CORS middleware
	// 3. Create recovery middleware
	// 4. Apply middleware to all routes
	// 5. Start the server on port 8080
}

func main() {
	fmt.Println("Starting middleware server on :8080...")
	CreateMiddlewareChain()
}
