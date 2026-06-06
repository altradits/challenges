# 10. API Documentation

## What You'll Learn

This exercise teaches you **API documentation in Go**. By the end, you'll understand:
- OpenAPI/Swagger specification
- Generating documentation from code
- Using swaggo for automatic docs
- Documenting endpoints
- API versioning

## Theory: API Documentation

### OpenAPI Specification

OpenAPI (formerly Swagger) is a standard for API documentation:

```yaml
openapi: 3.0.0
info:
  title: Task API
  version: 1.0.0
paths:
  /users:
    get:
      summary: List all users
      responses:
        '200':
          description: Success
```

### Swaggo

Swaggo generates OpenAPI docs from code comments:

```go
// @Summary Create user
// @Description Create a new user
// @Accept json
// @Produce json
// @Param user body User true "User data"
// @Success 201 {object} User
// @Router /users [post]
```

### Documentation Best Practices

- Document all endpoints
- Include example requests/responses
- Specify error codes
- Keep docs in sync with code

## Step-by-Step Approach

1. **Add swaggo comments** to handlers
2. **Install swaggo CLI**
3. **Generate docs** with `swag init`
4. **Serve Swagger UI**
5. **Update docs** as you change endpoints

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| No documentation | Hard to use API | Document all endpoints |
| Outdated docs | Misleads users | Update with code changes |
| No examples | Unclear usage | Include example requests |
| Not specifying errors | Poor error handling | Document all error codes |

## Practice Tips

- Use consistent response format
- Document request bodies
- Include authentication requirements
- Use tags to group endpoints

## The Challenge

Add OpenAPI documentation to an API.

### Expected Function

```go
func CreateDocumentedAPI() {
    // Your code here
    // Add swaggo comments to all handlers
    // Generate OpenAPI spec
    // Serve Swagger UI
}
```

### Test Cases

| Request | Expected Response |
|---------|-----------------|
| GET /swagger/index.html | Swagger UI |
| GET /docs/doc.json | OpenAPI spec |

### Usage Example

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Starting documented API on :8080...")
    CreateDocumentedAPI()
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. What is OpenAPI and why use it?
2. How do you document a request body?
3. What's the difference between @Success and @Response?

## Next Steps

After completing this, you'll be ready for:
- **capstone-taskmanager**: Building a complete API
- **capstone-production**: Deploying to production