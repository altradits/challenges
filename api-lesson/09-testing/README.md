# 09. Testing APIs

## What You'll Learn

This exercise teaches you **testing Go APIs**. By the end, you'll understand:
- Using `httptest` for testing HTTP handlers
- Table-driven tests
- Mocking database and external services
- Integration testing
- Test coverage

## Theory: Testing in Go

### The httptest Package

Go's `net/http/httptest` package provides tools for testing HTTP:

```go
import "net/http/httptest"

// Create test request
req := httptest.NewRequest("GET", "/users", nil)

// Create response recorder
rr := httptest.NewRecorder()

// Call handler
handler(rr, req)

// Check response
if rr.Code != http.StatusOK {
    t.Errorf("expected 200, got %d", rr.Code)
}
```

### Table-Driven Tests

Organize tests in a table format:

```go
tests := []struct {
    name       string
    method     string
    path       string
    wantStatus int
}{
    {"get users", "GET", "/users", 200},
    {"create user", "POST", "/users", 201},
}

for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        // Test code
    })
}
```

### Mocking

Create interfaces for mocking:

```go
type UserStore interface {
    Get(id int) (*User, error)
    Create(user *User) error
}

// Mock implementation
type MockStore struct{}
```

## Step-by-Step Approach

1. **Create test file** with `_test.go` suffix
2. **Write table-driven tests** for each endpoint
3. **Use httptest** for request/response testing
4. **Mock dependencies** for isolated tests
5. **Run tests** with `go test -v`

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Testing with real DB | Slow, flaky tests | Use mocks or test DB |
| Not checking response body | Incomplete tests | Check both status and body |
| Not using t.Run | Poor test output | Use subtests for clarity |
| Not cleaning up | Test pollution | Use t.Cleanup or defer |

## Practice Tips

- Test error cases, not just happy paths
- Use `httptest.NewServer` for integration tests
- Mock external HTTP calls
- Aim for 80%+ test coverage

## The Challenge

Write comprehensive tests for an API.

### Expected Function

```go
func TestUserHandler(t *testing.T) {
    // Your code here
    // Test:
    // - GET /users returns 200
    // - POST /users with valid data returns 201
    // - POST /users with invalid data returns 400
    // - GET /users/999 returns 404
}
```

### Test Cases

| Test | Expected |
|------|----------|
| GET /users | 200 + JSON array |
| POST /users valid | 201 + created user |
| POST /users invalid | 400 + error |
| GET /users/999 | 404 |

### Usage Example

```go
package main

import (
    "testing"
)

func TestUserHandler(t *testing.T) {
    // Your tests here
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. How do you create a test HTTP request?
2. How do you check the response status code?
3. What's the benefit of table-driven tests?

## Next Steps

After completing this, you'll be ready for:
- **10-docs**: Documenting your API
- **capstone-taskmanager**: Building a complete API