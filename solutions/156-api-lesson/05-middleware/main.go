package main

import (
	"fmt"
	"log"
	"net/http"
)

// CreateMiddlewareChain wraps next with logging, CORS and recovery middleware.
// Implement the following middleware:
// - Logging: Log method, path, and status code
// - CORS: Add Access-Control-Allow-Origin: * header
// - Recovery: Recover from panics and return 500
// Hint: Use http.HandlerFunc to convert function to handler
// Hint: Use a custom ResponseWriter to capture status code
// Hint: Use defer/recover for panic recovery
func CreateMiddlewareChain(next http.Handler) http.Handler {
	return recoveryMiddleware(corsMiddleware(loggingMiddleware(next)))
}

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sw := &statusWriter{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(sw, r)
		log.Printf("%s %s %d", r.Method, r.URL.Path, sw.status)
	})
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Println("Starting middleware server on :8080...")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})
	if err := http.ListenAndServe(":8080", CreateMiddlewareChain(mux)); err != nil {
		fmt.Println("server error:", err)
	}
}
