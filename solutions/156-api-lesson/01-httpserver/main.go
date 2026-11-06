package main

import (
	"fmt"
	"net/http"
)

// CreateServer builds the handler for a basic HTTP server.
// The server should respond with "Hello, World!" on the root path.
// Hint: Use http.NewServeMux and handle "/"
// Hint: Return the mux; main() starts the server with ListenAndServe
func CreateServer() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})
	return mux
}

func main() {
	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", CreateServer()); err != nil {
		fmt.Println("server error:", err)
	}
}
