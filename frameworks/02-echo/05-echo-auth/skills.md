# Skills for echo/05-echo-auth

## What You'll Learn

**Previous:** [echo/04-echo-validation](../04-echo-validation/skills.md) | **Next:** [frameworks/03-fiber](../../03-fiber/README.md)

**Challenge:** Implement JWT auth using Echo's official JWT middleware.

## Echo JWT Middleware

```go
import (
    echojwt "github.com/labstack/echo-jwt/v4"
    "github.com/golang-jwt/jwt/v5"
)

// Apply to a group
api := e.Group("/api")
api.Use(echojwt.JWT([]byte("secret")))
```

The middleware validates the token and stores it in the context as `"user"`.

## Custom Claims

```go
type Claims struct {
    UserID   int    `json:"user_id"`
    Username string `json:"username"`
    Role     string `json:"role"`
    jwt.RegisteredClaims
}
```

Configure the middleware to use your claims type:

```go
api.Use(echojwt.WithConfig(echojwt.Config{
    SigningKey:  []byte("secret"),
    NewClaimsFunc: func(c echo.Context) jwt.Claims {
        return new(Claims)
    },
}))
```

## Extracting Claims

```go
func profile(c echo.Context) error {
    token := c.Get("user").(*jwt.Token)
    claims := token.Claims.(*Claims)
    return c.JSON(http.StatusOK, map[string]interface{}{
        "user_id":  claims.UserID,
        "username": claims.Username,
        "role":     claims.Role,
    })
}
```

## Role Authorization Middleware

```go
func RequireRole(role string) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            token := c.Get("user").(*jwt.Token)
            claims := token.Claims.(*Claims)
            if claims.Role != role {
                return echo.NewHTTPError(http.StatusForbidden,
                    "insufficient permissions")
            }
            return next(c)
        }
    }
}

// Usage:
api.GET("/admin", adminHandler, RequireRole("admin"))
```

## Documentation

- [echo-jwt middleware](https://github.com/labstack/echo-jwt)
- [golang-jwt](https://github.com/golang-jwt/jwt)
- [gin/05-auth-jwt](../../01-gin/05-auth-jwt/skills.md) — JWT concepts

**Next:** [frameworks/03-fiber](../../03-fiber/README.md)
