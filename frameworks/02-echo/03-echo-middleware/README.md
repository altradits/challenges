# echo/03-echo-middleware — Echo Middleware

## Challenge

Add the following to your Blog API:

1. Use Echo's built-in **Logger**, **Recover**, **CORS**, and **Secure** middleware globally
2. Write a custom **RequestID** middleware that adds `X-Request-ID` to every response
3. Write a custom **JSON-only** middleware that returns 415 if the `Content-Type` is not `application/json` for POST/PUT/PATCH requests
4. Add **rate limiting** using `middleware.RateLimiter` from the echo/v4/middleware package

## Echo Built-In Rate Limiter

```go
import "github.com/labstack/echo/v4/middleware"

e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
```

Read `skills.md` before you start.
