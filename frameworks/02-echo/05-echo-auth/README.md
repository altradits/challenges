# echo/05-echo-auth — JWT Auth with Echo

## Challenge

Use Echo's built-in JWT middleware:

```
POST /auth/login    → returns JWT token
GET  /api/profile   → protected by JWT middleware
GET  /api/dashboard → protected by JWT middleware
```

Use `github.com/labstack/echo-jwt/v4` — Echo's official JWT middleware.

```bash
go get github.com/labstack/echo-jwt/v4
go get github.com/golang-jwt/jwt/v5
```

## Requirements

- Use `echojwt.JWT([]byte(secret))` as middleware
- Extract user from token context: `c.Get("user").(*jwt.Token)`
- Return the user's claims in `/api/profile`
- Add a `/api/admin` endpoint that additionally checks the user's `role` claim is `"admin"`

Read `skills.md` before you start.
