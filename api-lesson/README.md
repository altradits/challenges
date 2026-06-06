# Go API Development Mastery Course

Welcome to your complete API development course in Go! This collection of **10 progressive exercises** is designed to take you from **complete beginner** to **production-ready** API developer.

## How to Use

Each exercise has its own numbered folder containing:
- `README.md` - Comprehensive theory, examples, and instructions
- `main.go` - Starter code with TODO function for you to implement
- `go.mod` - Go module file (for exercises requiring dependencies)

**Your task:** Open each `main.go`, study the README to learn the concepts, then implement the function. Run `go run .` to test your solution.

## Course Structure

### Beginner Level (01-02) - HTTP Fundamentals

These exercises teach you the absolute fundamentals of HTTP servers in Go.

| # | Exercise | Folder | Key Concepts |
|---|----------|--------|--------------|
| 01 | [Basic HTTP Server](01-httpserver/README.md) | `01-httpserver/` | `net/http`, handlers, ListenAndServe, request/response cycle |
| 02 | [Simple Routing](02-routing/README.md) | `02-routing/` | Path matching, URL parameters, method routing, mux patterns |

**Beginner Tips:**
- Use `http.HandleFunc` for simple routing
- Always set `Content-Type` header for responses
- Use `http.ListenAndServe` to start the server
- Test with `curl` or a browser

### Intermediate Level (03-05) - REST API Design

These exercises build on basics and introduce REST principles and JSON handling.

| # | Exercise | Folder | Key Concepts |
|---|----------|--------|--------------|
| 03 | [REST API Design](03-rest-api/README.md) | `03-rest-api/` | HTTP methods, status codes, resource modeling, REST conventions |
| 04 | [JSON Handling](04-json/README.md) | `04-json/` | `encoding/json`, struct tags, marshaling/unmarshaling, validation |
| 05 | [Middleware](05-middleware/README.md) | `05-middleware/` | Middleware pattern, request chaining, logging, CORS, recovery |

**Intermediate Tips:**
- Use proper HTTP status codes (200, 201, 400, 404, 500)
- Design resources around nouns, not verbs
- Use struct tags for JSON field mapping
- Chain middleware in the correct order

### Advanced Level (06-10) - Production Features

These exercises cover production-ready features for real-world APIs.

| # | Exercise | Folder | Key Concepts |
|---|----------|--------|--------------|
| 06 | [Database Integration](06-database/README.md) | `06-database/` | `database/sql`, SQL queries, connection pooling, migrations |
| 07 | [Authentication](07-auth/README.md) | `07-auth/` | JWT, sessions, password hashing, middleware auth, RBAC |
| 08 | [Error Handling](08-errors/README.md) | `08-errors/` | Custom errors, error wrapping, structured logging, recovery |
| 09 | [Testing APIs](09-testing/README.md) | `09-testing/` | `httptest`, table tests, mocking, integration tests |
| 10 | [API Documentation](10-docs/README.md) | `10-docs/` | OpenAPI/Swagger, godoc, swagger UI, spec generation |

**Advanced Tips:**
- Use context for request cancellation and timeouts
- Implement proper error handling with custom types
- Write tests before or alongside your code
- Document your API as you build it

## Capstone Projects

Apply all your skills to build real-world API services.

### Capstone 1: Task Manager API

Build a complete REST API for task management with:
- User authentication and authorization
- CRUD operations for tasks
- Task filtering and pagination
- Input validation
- Comprehensive error handling
- Full test coverage
- OpenAPI documentation

**Skills Applied:**
- All concepts from exercises 01-10
- Real database design
- Security best practices
- Production deployment

### Capstone 2: Production Deployment

Deploy your Task Manager API to production with:
- Docker containerization
- Environment configuration
- Logging and monitoring
- Health checks
- Rate limiting
- HTTPS/TLS setup

**Skills Applied:**
- Docker and docker-compose
- Environment variables
- Structured logging
- Security headers
- Performance optimization

## Key Go API Concepts

### HTTP Handlers

Handlers are functions that process HTTP requests:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // Process request
    w.Write([]byte("Hello, World!"))
}
```

### The Request-Response Cycle

```
Client Request → Router → Middleware → Handler → Middleware → Response
```

### JSON Serialization

```go
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

// Encode to JSON
data, _ := json.Marshal(user)

// Decode from JSON
var user User
json.Unmarshal(data, &user)
```

## Progress Tracking

Work through the exercises in order. Each one builds on skills from the previous ones:

```
Beginner:        HTTP Server → Routing
                 ↓
Intermediate:    REST Design → JSON → Middleware
                 ↓
Advanced:        Database → Auth → Errors → Testing → Docs
                 ↓
Capstone:        Task Manager API → Production Deployment
```

## Getting Help

If you're stuck:
1. **Re-read the README** - It contains all the theory you need
2. **Look at the hints** in `main.go`
3. **Test with curl** - Verify your endpoints work
4. **Check similar exercises** - Patterns repeat across the course
5. **Use Go documentation** - `go doc net/http`

## What You'll Achieve

After completing all exercises, you will be able to:
- Build HTTP servers from scratch using Go's standard library
- Design and implement RESTful APIs
- Handle JSON requests and responses efficiently
- Create reusable middleware for cross-cutting concerns
- Integrate with SQL databases
- Implement secure authentication systems
- Write comprehensive tests for your APIs
- Document APIs with OpenAPI/Swagger
- Deploy APIs to production environments

## Next Steps

After completing this course:
1. Build your own API project from scratch
2. Explore popular Go frameworks (Gin, Echo, Fiber)
3. Learn about microservices patterns
4. Study distributed systems concepts
5. Contribute to open-source Go API projects

**Happy coding! Start with `01-httpserver` and work your way up.**