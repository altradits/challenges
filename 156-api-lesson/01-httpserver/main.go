package main

import (
	"fmt"
)

// CreateServer starts a basic HTTP server on port 8080.
// The server should respond with "Hello, World!" on the root path.
// Hint: Use http.HandleFunc to register a handler for "/"
// Hint: Use http.ListenAndServe to start the server
func CreateServer() {
	// TODO: Implement this function
	// 1. Register a handler for the root path "/"
	// 2. The handler should write "Hello, World!" as the response
	// 3. Start the server on port 8080
	// 4. Handle any errors from ListenAndServe
}

func main() {
	fmt.Println("Starting server on :8080...")
	CreateServer()
}
