# gin/05-auth-jwt — JWT Authentication

## Challenge

Add JWT-based authentication to your Gin API:

```
POST /auth/register   → create user, return JWT token
POST /auth/login      → verify credentials, return JWT token
GET  /api/me          → return current user (requires valid JWT)
POST /api/tasks       → protected endpoint (requires valid JWT)
```

## User Model

```go
type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Password string `json:"-"`  // never returned in JSON
}
```

## Requirements

- Hash passwords with `bcrypt`
- Sign JWTs with HS256 and a secret key from the environment
- JWT claims must include `user_id` and `exp` (expiry: 24 hours)
- Protected routes must return 401 if token is missing or invalid
- Extract the current user from the token in protected handlers

## Install

```bash
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt
```

## Test

```bash
# Register
curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"alice","password":"secret123"}'
# → {"token":"eyJ..."}

# Use token
curl -H "Authorization: Bearer eyJ..." http://localhost:8080/api/me
# → {"id":1,"username":"alice"}
```

Read `skills.md` before you start.
