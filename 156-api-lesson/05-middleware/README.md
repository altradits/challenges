# 05. Middleware

## What You'll Learn

This exercise teaches you **middleware patterns in Go**. By the end, you'll understand:
- Middleware function signature
- Request/response chaining
- Common middleware use cases (logging, CORS, recovery)
- Building a middleware chain

## Theory: Middleware in Go

### Middleware Pattern

Middleware is a function that wraps an HTTP handler:

```go
func Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Before request
        next.ServeHTTP(w, r)
        // After request
    })
}
```

### Common Middleware Types

1. **Logging** - Log each request
2. **CORS** - Handle cross-origin requests
3. **Recovery** - Recover from panics
4. **Authentication** - Validate requests
5. **Rate Limiting** - Limit request frequency

### Chaining Middleware

```go
handler := http.Handler(yourHandler)
handler = loggingMiddleware(handler)
handler = corsMiddleware(handler)
http.Handle("/path", handler)
```

## Step-by-Step Approach

1. **Create middleware function** with proper signature
2. **Process request before** calling next handler
3. **Call next.ServeHTTP** to continue chain
4. **Process response after** if needed

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Not calling next handler | Request hangs | Always call `next.ServeHTTP` |
| Modifying response after | Can't change | Modify before or use wrapper |
| Not handling errors | Panics crash server | Use recovery middleware |

## Practice Tips

- Order matters: logging should be first
- Use `httptest` for testing middleware
- Keep middleware focused on one concern
- Use context for passing data between middleware

## The Challenge

Create middleware for logging, CORS, and recovery.

### Expected Function

```go
func CreateMiddlewareChain() {
    // Your code here
    // Implement:
    // - Logging middleware (log method, path, status)
    // - CORS middleware (allow all origins)
    // - Recovery middleware (recover from panics)
    // - Apply to all routes
}
```

### Test Cases

| Request | Expected Behavior |
|---------|-----------------|
| Any request | Should be logged |
| OPTIONS request | Should return CORS headers |
| Handler panic | Should recover and return 500 |

### Usage Example

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    fmt.Println("Starting middleware server on :8080...")
    CreateMiddlewareChain()
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. What is the signature of a middleware function?
2. How do you pass data from middleware to handlers?
3. Why should you use recovery middleware?

## Next Steps

After completing this, you'll be ready for:
- **06-database**: Adding persistence
- **07-auth**: Implementing authentication
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
(cd 156-api-lesson/05-middleware && go test ./...)
```

Green means done. Red means read the failure message, fix from memory, check again.

## Repeat schedule

Redo this exercise from scratch after **1 day, 3 days, 7 days, 21 days**. `go run ./cmd/drill due` tells you what is due today.
