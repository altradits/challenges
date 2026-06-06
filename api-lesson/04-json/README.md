# 04. JSON Handling

## What You'll Learn

This exercise teaches you **JSON handling in Go**. By the end, you'll understand:
- Struct tags for JSON field mapping
- Marshaling (encoding) Go structs to JSON
- Unmarshaling (decoding) JSON to Go structs
- JSON validation and error handling
- Working with nested JSON structures

## Theory: JSON in Go

### The encoding/json Package

Go's standard library provides `encoding/json` for JSON handling:

```go
import "encoding/json"
```

### Struct Tags

Use struct tags to control JSON field names:

```go
type User struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email,omitempty"` // omit if empty
}
```

### Marshaling

Convert Go struct to JSON:

```go
data, err := json.Marshal(user)
// or
data, err := json.MarshalIndent(user, "", "  ") // pretty print
```

### Unmarshaling

Convert JSON to Go struct:

```go
var user User
err := json.Unmarshal(data, &user)
```

### Reading JSON from Request

```go
var user User
err := json.NewDecoder(r.Body).Decode(&user)
```

## Step-by-Step Approach

1. **Define structs** with proper JSON tags
2. **Decode request body** to struct
3. **Validate input** data
4. **Encode response** to JSON
5. **Set Content-Type** header

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Forgetting `json:` tags | Fields not exported | Add struct tags |
| Not handling errors | Panics on invalid JSON | Always check errors |
| Not closing body | Resource leak | Use `defer r.Body.Close()` |
| Wrong Content-Type | Client confusion | Set `application/json` |

## Practice Tips

- Use `json.NewDecoder(r.Body).Decode()` for request bodies
- Use `json.Marshal()` for responses
- Validate required fields before processing
- Use `omitempty` for optional fields

## The Challenge

Create handlers that properly handle JSON request/response bodies.

### Expected Function

```go
func CreateJSONHandler() {
    // Your code here
    // Implement:
    // POST /api/v1/users -> Create user from JSON, return created user
    // GET /api/v1/users/{id} -> Return user as JSON
    // PUT /api/v1/users/{id} -> Update user from JSON
}
```

### Test Cases

| Request | Expected Response |
|---------|-----------------|
| POST /api/v1/users with `{"name":"John","email":"john@example.com"}` | 201 + created user with ID |
| GET /api/v1/users/1 | 200 + user JSON |
| PUT /api/v1/users/1 with `{"name":"Jane"}` | 200 + updated user |

### Usage Example

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    fmt.Println("Starting JSON API on :8080...")
    CreateJSONHandler()
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. How do you map a Go field to a different JSON name?
2. How do you decode JSON from an HTTP request body?
3. What does `omitempty` do in a struct tag?

## Next Steps

After completing this, you'll be ready for:
- **05-middleware**: Adding request processing layers
- **06-database**: Persisting data to a database