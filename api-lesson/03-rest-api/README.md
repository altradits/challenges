# 03. REST API Design

## What You'll Learn

This exercise teaches you **REST API design principles**. By the end, you'll understand:
- RESTful resource design
- HTTP status codes
- Request/response patterns
- CRUD operations over HTTP
- API versioning

## Theory: REST API Design

### REST Principles

REST (Representational State Transfer) is an architectural style for APIs:

1. **Resources** - Everything is a resource (noun, not verb)
2. **HTTP Methods** - Use standard methods (GET, POST, PUT, DELETE, PATCH)
3. **Stateless** - Each request contains all necessary information
4. **Uniform Interface** - Consistent patterns across endpoints

### HTTP Status Codes

| Code | Meaning | When to Use |
|------|---------|-------------|
| 200 | OK | Successful GET, PUT |
| 201 | Created | Successful POST |
| 204 | No Content | Successful DELETE |
| 400 | Bad Request | Invalid input |
| 404 | Not Found | Resource doesn't exist |
| 500 | Internal Error | Server error |

### Resource Design

```
GET    /api/v1/users      - List all users
GET    /api/v1/users/123  - Get user 123
POST   /api/v1/users      - Create new user
PUT    /api/v1/users/123  - Update user 123
DELETE /api/v1/users/123  - Delete user 123
```

## Step-by-Step Approach

1. **Design resource structure** (users, tasks, etc.)
2. **Implement handlers** for each CRUD operation
3. **Return correct status codes** for each operation
4. **Add API versioning** to paths

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Using verbs in paths | Not RESTful | Use nouns: `/users` not `/getUsers` |
| Wrong status codes | Confuses clients | Use 201 for create, 204 for delete |
| No versioning | Breaking changes | Add `/api/v1/` prefix |
| Not handling errors | Poor UX | Return appropriate error codes |

## Practice Tips

- Use consistent response format
- Always return created resource on POST
- Use plural nouns for resources
- Keep endpoints shallow (max 2 levels)

## The Challenge

Create a RESTful API for managing tasks with proper status codes.

### Expected Function

```go
func CreateTaskAPI() {
    // Your code here
    // Implement:
    // GET /api/v1/tasks -> 200 + list of tasks
    // GET /api/v1/tasks/{id} -> 200 + task or 404
    // POST /api/v1/tasks -> 201 + created task
    // PUT /api/v1/tasks/{id} -> 200 + updated task or 404
    // DELETE /api/v1/tasks/{id} -> 204
}
```

### Test Cases

| Request | Expected Response |
|---------|-----------------|
| GET /api/v1/tasks | 200 + [] |
| POST /api/v1/tasks | 201 + created task |
| GET /api/v1/tasks/1 | 200 + task |
| GET /api/v1/tasks/999 | 404 |
| DELETE /api/v1/tasks/1 | 204 |

### Usage Example

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    fmt.Println("Starting Task API on :8080...")
    CreateTaskAPI()
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. What HTTP method should create a new resource?
2. What status code indicates a resource was created?
3. Why should you version your API?

## Next Steps

After completing this, you'll be ready for:
- **04-json**: Handling JSON request bodies
- **05-middleware**: Adding cross-cutting concerns