package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

// Task represents a task in the system
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// CreateTaskAPI builds a RESTful API for task management.
// Implement the following endpoints:
// - GET /api/v1/tasks -> 200 + list of tasks
// - GET /api/v1/tasks/{id} -> 200 + task or 404
// - POST /api/v1/tasks -> 201 + created task
// - PUT /api/v1/tasks/{id} -> 200 + updated task or 404
// - DELETE /api/v1/tasks/{id} -> 204
// Hint: Use an in-memory slice to store tasks
// Hint: Use http.StatusXxx constants for status codes
// Hint: Use encoding/json for JSON handling
// Hint: Return the handler; main() starts the server
func CreateTaskAPI() http.Handler {
	var mu sync.Mutex
	tasks := map[int]Task{}
	nextID := 1
	writeJSON := func(w http.ResponseWriter, code int, v any) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(v)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		if r.URL.Path == "/api/v1/tasks" {
			switch r.Method {
			case http.MethodGet:
				list := []Task{}
				for _, tsk := range tasks {
					list = append(list, tsk)
				}
				writeJSON(w, http.StatusOK, list)
			case http.MethodPost:
				var in struct {
					Title string `json:"title"`
					Done  bool   `json:"done"`
				}
				if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
					http.Error(w, "bad request", http.StatusBadRequest)
					return
				}
				tsk := Task{ID: nextID, Title: in.Title, Done: in.Done}
				nextID++
				tasks[tsk.ID] = tsk
				writeJSON(w, http.StatusCreated, tsk)
			default:
				http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			}
			return
		}
		if !strings.HasPrefix(r.URL.Path, "/api/v1/tasks/") {
			http.NotFound(w, r)
			return
		}
		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/v1/tasks/"))
		if err != nil {
			http.Error(w, "bad id", http.StatusBadRequest)
			return
		}
		tsk, ok := tasks[id]
		if !ok {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, http.StatusOK, tsk)
		case http.MethodPut:
			var in struct {
				Title string `json:"title"`
				Done  bool   `json:"done"`
			}
			if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
				http.Error(w, "bad request", http.StatusBadRequest)
				return
			}
			tsk.Title, tsk.Done = in.Title, in.Done
			tasks[id] = tsk
			writeJSON(w, http.StatusOK, tsk)
		case http.MethodDelete:
			delete(tasks, id)
			w.WriteHeader(http.StatusNoContent)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

func main() {
	fmt.Println("Starting Task API on :8080...")
	if err := http.ListenAndServe(":8080", CreateTaskAPI()); err != nil {
		fmt.Println("server error:", err)
	}
}
