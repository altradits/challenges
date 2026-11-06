package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
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

func writeError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(&APIError{Code: code, Message: message})
}

// CreateErrorHandlingAPI builds an API with proper error handling.
// Implement the following:
// - GET /ok -> 200 {"status":"ok"}
// - GET /fail -> 500 + APIError JSON body
// - GET /panic -> panics; the recovery middleware must return 500 + APIError JSON
// - Consistent JSON error responses via a writeError helper
// - Panic recovery middleware
// Hint: Use log/slog for structured logging
// Hint: Create a writeError helper function
// Hint: Use defer/recover in middleware
// Hint: Return the handler; main() starts the server
func CreateErrorHandlingAPI() http.Handler {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	mux := http.NewServeMux()
	mux.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})
	mux.HandleFunc("/fail", func(w http.ResponseWriter, r *http.Request) {
		writeError(w, http.StatusInternalServerError, "something went wrong")
	})
	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("boom")
	})
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				logger.Error("panic", "path", r.URL.Path, "panic", rec)
				writeError(w, http.StatusInternalServerError, "internal server error")
			}
		}()
		logger.Info("request", "method", r.Method, "path", r.URL.Path)
		mux.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Println("Starting error handling API on :8080...")
	if err := http.ListenAndServe(":8080", CreateErrorHandlingAPI()); err != nil {
		fmt.Println("server error:", err)
	}
}
