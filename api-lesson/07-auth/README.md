# 07. Authentication

## What You'll Learn

This exercise teaches you **authentication in Go APIs**. By the end, you'll understand:
- Password hashing with bcrypt
- JWT (JSON Web Tokens)
- Session-based authentication
- Protected routes
- Role-based access control (RBAC)

## Theory: Authentication in Go

### Password Hashing

Use bcrypt for secure password storage:

```go
import "golang.org/x/crypto/bcrypt"

// Hash password
hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)

// Compare password
err := bcrypt.CompareHashAndPassword(hashed, []byte(password))
```

### JWT (JSON Web Tokens)

JWT provides stateless authentication:

```go
import "github.com/golang-jwt/jwt/v5"

// Create token
token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
signed, err := token.SignedString([]byte(secret))

// Verify token
token, err := jwt.Parse(signed, func(t *jwt.Token) (interface{}, error) {
    return []byte(secret), nil
})
```

### Protected Routes

Use middleware to protect routes:

```go
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Check Authorization header
        // Validate token
        // Call next or return 401
    })
}
```

## Step-by-Step Approach

1. **Hash passwords** before storing
2. **Create login endpoint** to issue tokens
3. **Create auth middleware** to validate tokens
4. **Protect routes** with middleware
5. **Implement RBAC** for role-based access

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Storing plain passwords | Security risk | Always hash with bcrypt |
| Not validating token | Anyone can access | Check token on every request |
| Using weak secret | Tokens can be forged | Use strong random secret |
| Not expiring tokens | Security risk | Set expiration time |

## Practice Tips

- Use environment variables for secrets
- Set short token expiration
- Use HTTPS in production
- Log authentication failures

## The Challenge

Create an API with JWT authentication and protected routes.

### Expected Function

```go
func CreateAuthAPI() {
    // Your code here
    // Implement:
    // - POST /api/v1/register -> Create user with hashed password
    // - POST /api/v1/login -> Return JWT token
    // - GET /api/v1/profile -> Protected route (requires token)
    // - GET /api/v1/admin -> Admin-only route
}
```

### Test Cases

| Request | Expected Response |
|---------|-----------------|
| POST /api/v1/register with valid data | 201 + user created |
| POST /api/v1/login with correct password | 200 + JWT token |
| GET /api/v1/profile with valid token | 200 + user profile |
| GET /api/v1/profile without token | 401 Unauthorized |
| GET /api/v1/admin as non-admin | 403 Forbidden |

### Usage Example

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Starting auth API on :8080...")
    CreateAuthAPI()
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. Why should you hash passwords before storing?
2. How do you validate a JWT token?
3. What's the difference between 401 and 403?

## Next Steps

After completing this, you'll be ready for:
- **08-errors**: Advanced error handling
- **09-testing**: Testing your authenticated endpoints