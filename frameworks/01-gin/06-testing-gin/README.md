# gin/06-testing-gin — Testing Gin HTTP Handlers

## Challenge

Write a complete test suite for the Task Manager API from challenge 04:

- Table-driven tests for every endpoint
- Test success cases, 400 errors, and 404 errors
- Test that creating a task returns 201 with the task body
- Test that an invalid JSON body returns 400
- Test PATCH partial update (only provided fields change)
- One benchmark for `GET /tasks`

## Test File

Create `main_test.go` in `package main` (or `package main_test`).

## Key Tools

```go
import (
    "net/http"
    "net/http/httptest"
    "testing"
    "encoding/json"
    "strings"
)

// Create a test request:
req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
w   := httptest.NewRecorder()
router.ServeHTTP(w, req)

// Assert:
if w.Code != http.StatusOK {
    t.Errorf("expected 200, got %d", w.Code)
}
```

Read `skills.md` before you start.
