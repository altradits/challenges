# 02. Simple Routing

## What You'll Learn

This exercise teaches you **routing in Go HTTP servers**. By the end, you'll understand:
- Path-based routing with `http.HandleFunc`
- HTTP method handling (GET, POST, PUT, DELETE)
- URL path parameters
- Building a simple router

## Theory: Routing in Go

### DefaultServeMux

Go's `http.HandleFunc` uses a default multiplexer (router) that matches paths:

```go
http.HandleFunc("/users", usersHandler)
http.HandleFunc("/products", productsHandler)
```

### HTTP Methods

The `http.Request` struct contains the method:

```go
r.Method // "GET", "POST", "PUT", "DELETE", etc.
```

### Path Parameters

For dynamic paths like `/users/123`, you need to extract the ID:

```go
// Extract path after "/users/"
id := strings.TrimPrefix(r.URL.Path, "/users/")
```

## Step-by-Step Approach

1. **Create handlers** for different paths
2. **Check HTTP method** in each handler
3. **Extract path parameters** when needed
4. **Return appropriate responses** for each route

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Not checking HTTP method | All methods go to same handler | Check `r.Method` first |
| Forgetting to handle 404 | No response for unknown paths | Add a catch-all handler |
| Not parsing path params | Can't get dynamic values | Use `strings.TrimPrefix` or regex |

## Practice Tips

- Test each route with `curl -X METHOD /path`
- Use `http.NotFound` for unknown routes
- Log the method and path for debugging
- Return appropriate status codes

## The Challenge

Create a router that handles multiple paths and methods:

### Expected Function

```go
func CreateRouter() {
    // Your code here
    // Handle:
    // GET / -> "Welcome to the API"
    // GET /users -> "List of users"
    // GET /users/{id} -> "User {id}"
    // POST /users -> "Create user"
    // PUT /users/{id} -> "Update user {id}"
    // DELETE /users/{id} -> "Delete user {id}"
}
```

### Test Cases

| Request | Expected Response |
|---------|-----------------|
| GET / | "Welcome to the API" |
| GET /users | "List of users" |
| GET /users/123 | "User 123" |
| POST /users | "Create user" |
| PUT /users/456 | "Update user 456" |
| DELETE /users/789 | "Delete user 789" |

### Usage Example

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    fmt.Println("Starting router on :8080...")
    CreateRouter()
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. How do you register multiple handlers?
2. How do you check the HTTP method in a handler?
3. How do you extract a value from a path like `/users/123`?

## Next Steps

After completing this, you'll be ready for:
- **03-rest-api**: Building proper REST endpoints
- **04-json**: Handling JSON request/response bodies