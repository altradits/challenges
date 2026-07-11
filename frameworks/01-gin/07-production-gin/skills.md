# Skills for gin/07-production-gin

## What You'll Learn

**Previous:** [gin/06-testing-gin](../06-testing-gin/skills.md) | **Next:** [frameworks/02-echo](../../02-echo/README.md)

**Challenge:** Wire together graceful shutdown, config, structured logging, request IDs, and timeout middleware.

## Graceful Shutdown

When a signal arrives, stop accepting new connections but finish existing ones:

```go
import (
    "context"
    "os/signal"
    "syscall"
    "net/http"
)

srv := &http.Server{
    Addr:    ":" + cfg.Port,
    Handler: router,
}

go func() {
    if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        log.Fatal("listen:", err)
    }
}()

quit := make(chan os.Signal, 1)
signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
<-quit

log.Println("shutting down...")
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

if err := srv.Shutdown(ctx); err != nil {
    log.Fatal("forced shutdown:", err)
}
```

`srv.Shutdown(ctx)` waits for active requests to complete or the context to expire.

## Structured Logging with `log/slog`

```go
import "log/slog"

// JSON output
logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
    Level: slog.LevelInfo,
}))
slog.SetDefault(logger)

// Logging
slog.Info("request",
    "method", c.Request.Method,
    "path", c.Request.URL.Path,
    "status", c.Writer.Status(),
    "latency", time.Since(start).String(),
    "request_id", c.GetString("request_id"),
)
```

## Request ID Middleware

```go
import "github.com/google/uuid"

func RequestID() gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.GetHeader("X-Request-ID")
        if id == "" {
            id = uuid.NewString()
        }
        c.Set("request_id", id)
        c.Header("X-Request-ID", id)
        c.Next()
    }
}
```

## Timeout Middleware

```go
func Timeout(d time.Duration) gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(c.Request.Context(), d)
        defer cancel()
        c.Request = c.Request.WithContext(ctx)

        finished := make(chan struct{})
        go func() {
            c.Next()
            close(finished)
        }()

        select {
        case <-finished:
        case <-ctx.Done():
            c.AbortWithStatusJSON(http.StatusGatewayTimeout,
                gin.H{"error": "request timeout"})
        }
    }
}
```

## Config from Environment

```go
import (
    "os"
    "strconv"
)

type Config struct {
    Port      string
    JWTSecret string
}

func LoadConfig() Config {
    return Config{
        Port:      getEnvOrDefault("PORT", "8080"),
        JWTSecret: getEnvOrDefault("JWT_SECRET", "dev-secret"),
    }
}

func getEnvOrDefault(key, def string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return def
}
```

## Documentation

- [net/http.Server.Shutdown](https://pkg.go.dev/net/http#Server.Shutdown)
- [log/slog](https://pkg.go.dev/log/slog) (Go 1.21+)
- [152-graceful-shutdown](../../../152-graceful-shutdown/README.md)
- [os/signal](https://pkg.go.dev/os/signal)

**Next:** [frameworks/02-echo](../../02-echo/README.md)
