# 01. Basic HTTP Server

## What You'll Learn

This exercise teaches you the **fundamentals of HTTP servers in Go**. By the end, you'll understand:
- How Go's `net/http` package works
- Creating HTTP handlers
- Starting a server with `ListenAndServe`
- The request-response cycle
- Writing responses to clients

## Theory: HTTP in Go

### The net/http Package

Go's standard library includes a powerful `net/http` package for building HTTP servers. It provides everything you need to handle HTTP requests without external dependencies.

```go
import "net/http"
```

### HTTP Handlers

A handler is a function that processes HTTP requests. It has the signature:

```go
func(w http.ResponseWriter, r *http.Request)
```

- `w` - ResponseWriter: used to write the response
- `r` - Request: contains all request information

### Starting a Server

```go
http.ListenAndServe(":8080", nil)
```

This starts a server on port 8080. The second parameter is a handler (nil uses DefaultServeMux).

## Step-by-Step Approach

1. **Import the net/http package**
2. **Create a handler function** that writes a response
3. **Register the handler** with `http.HandleFunc`
4. **Start the server** with `http.ListenAndServe`

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Forgetting to import `net/http` | Code won't compile | Add `"net/http"` to imports |
| Not calling `ListenAndServe` | Server never starts | Call it in `main()` |
| Forgetting to write response | Client gets nothing | Use `w.Write()` or `w.WriteHeader()` |
| Wrong port format | Server won't start | Use `":8080"` format |

## Practice Tips

- Test your server with `curl http://localhost:8080`
- Use `w.WriteHeader(http.StatusOK)` before writing
- Set `Content-Type` header for proper responses
- Check the terminal for server logs

## The Challenge

Create a basic HTTP server that responds with "Hello, World!" on the root path.

### Expected Function

```go
func CreateServer() {
    // Your code here
    // Start server on port 8080
}
```

### Test Cases

| Request | Expected Response |
|---------|-----------------|
| GET / | "Hello, World!" |
| Server running | Should be accessible on localhost:8080 |

### Usage Example

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Your server will run here
    // Test with: curl http://localhost:8080
    CreateServer()
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. What is the signature of an HTTP handler function?
2. How do you write a response to the client?
3. What does `ListenAndServe` return when the server fails?

## Next Steps

After completing this, you'll be ready for:
- **02-routing**: Handling different paths and methods
- **03-rest-api**: Building RESTful endpoints