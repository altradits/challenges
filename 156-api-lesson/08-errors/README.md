# 08. Error Handling and Logging

## What You'll Learn

This exercise teaches you **error handling and logging in Go APIs**. By the end, you'll understand:
- Custom error types
- Error wrapping and unwrapping
- Structured logging
- HTTP error responses
- Panic recovery

## Theory: Error Handling in Go

### Custom Error Types

Create custom error types for better error handling:

```go
type APIError struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

func (e *APIError) Error() string {
    return e.Message
}
```

### Error Wrapping

Use `fmt.Errorf` to wrap errors:

```go
import "fmt"

if err != nil {
    return fmt.Errorf("failed to create user: %w", err)
}
```

### Structured Logging

Use structured logging for better debugging:

```go
import "log/slog"

slog.Info("User created", "user_id", user.ID, "name", user.Name)
slog.Error("Database error", "error", err)
```

### HTTP Error Responses

Return consistent error responses:

```go
func writeError(w http.ResponseWriter, code int, message string) {
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(APIError{Code: code, Message: message})
}
```

## Step-by-Step Approach

1. **Create custom error types**
2. **Implement error response helper**
3. **Add structured logging**
4. **Handle errors consistently**
5. **Add panic recovery**

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Not wrapping errors | Lose context | Use `%w` format verb |
| Returning raw errors | Security risk | Return generic messages |
| Not logging errors | Hard to debug | Log all errors |
| Not using structured logs | Hard to search | Use key-value pairs |

## Practice Tips

- Log at appropriate levels (Info, Error, Debug)
- Include request ID in logs
- Don't expose internal errors to clients
- Use error codes for client handling

## The Challenge

Create an API with proper error handling and structured logging.

### Expected Function

```go
func CreateErrorHandlingAPI() {
    // Your code here
    // Implement:
    // - Custom APIError type
    // - Structured logging for requests
    // - Consistent error responses
    // - Panic recovery middleware
}
```

### Test Cases

| Request | Expected Response |
|---------|-----------------|
| Valid request | 200 + logged |
| Invalid JSON | 400 + error logged |
| Resource not found | 404 + error logged |
| Handler panic | 500 + recovered |

### Usage Example

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Starting error handling API on :8080...")
    CreateErrorHandlingAPI()
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. How do you create a custom error type?
2. What's the difference between `%v` and `%w` in error formatting?
3. Why should you use structured logging?

## Next Steps

After completing this, you'll be ready for:
- **09-testing**: Testing your error handling
- **10-docs**: Documenting your API
---

## Time limit

**45 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close the exercise README, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 156
```

or run this exercise's tests directly:

```bash
(cd 156-api-lesson/08-errors && go test ./...)
```

Green means done. Red means read the failure message, fix from memory, check again.

## Repeat schedule

Redo this exercise from scratch after **1 day, 3 days, 7 days, 21 days**. `go run ./cmd/drill due` tells you what is due today.
