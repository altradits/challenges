# Skills for echo/03-echo-middleware

## What You'll Learn

**Previous:** [echo/02-echo-routing](../02-echo-routing/skills.md) | **Next:** [echo/04-echo-validation](../04-echo-validation/skills.md)

**Challenge:** Add multiple middleware layers to your Echo app using built-in and custom middleware.

## Custom Echo Middleware

```go
func RequestID() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            id := c.Request().Header.Get("X-Request-ID")
            if id == "" {
                id = fmt.Sprintf("%d", time.Now().UnixNano())
            }
            c.Response().Header().Set("X-Request-ID", id)
            c.Set("request_id", id)
            return next(c)
        }
    }
}
```

Echo middleware has the type `func(echo.HandlerFunc) echo.HandlerFunc`. The pattern is identical to standard library middleware.

## Content-Type Guard

```go
func JSONOnly() echo.MiddlewareFunc {
    mutatingMethods := map[string]bool{"POST": true, "PUT": true, "PATCH": true}
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            if mutatingMethods[c.Request().Method] {
                ct := c.Request().Header.Get("Content-Type")
                if !strings.Contains(ct, "application/json") {
                    return echo.NewHTTPError(
                        http.StatusUnsupportedMediaType,
                        "Content-Type must be application/json")
                }
            }
            return next(c)
        }
    }
}
```

## Built-In Middleware

```go
// Logger with custom format
e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
    Format: `{"time":"${time_rfc3339}","method":"${method}","uri":"${uri}","status":${status},"latency":"${latency_human}"}` + "\n",
}))

// CORS with specific origins
e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    AllowOrigins: []string{"https://example.com"},
    AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
}))

// Secure headers
e.Use(middleware.Secure())

// Rate limiter (20 requests/second)
e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
```

## Documentation

- [Echo middleware](https://echo.labstack.com/docs/middleware)
- [Rate limiter](https://echo.labstack.com/docs/middleware/rate-limiter)

**Next:** [echo/04-echo-validation](../04-echo-validation/README.md)
