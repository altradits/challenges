# Skills for 150-logging

## What You'll Learn

**Previous:** [149-environment-and-config](../149-environment-and-config/skills.md) | **Next:** [151-generics](../151-generics/skills.md)

**Challenge:** Log structured messages with the `log` and `slog` packages.

## The `log` Package

The simple, traditional logger:

```go
import "log"

log.Print("server starting")           // 2024/03/15 09:30:00 server starting
log.Printf("listening on :%d", port)   // with format
log.Println("done")

log.Fatal("database error: ", err)  // logs + os.Exit(1)
log.Panic("unexpected nil")         // logs + panic
```

### Configuring the Default Logger

```go
log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
// 2024/03/15 09:30:00 main.go:42: message

log.SetPrefix("[API] ")
// [API] 2024/03/15 09:30:00 message
```

### Custom Logger

```go
import (
    "log"
    "os"
)

logger := log.New(os.Stderr, "[ERROR] ", log.LstdFlags)
logger.Printf("query failed: %v", err)
```

## The `slog` Package (Go 1.21+)

`slog` provides structured logging — key-value pairs that are machine-parseable:

```go
import "log/slog"

slog.Info("user created", "id", 42, "name", "Alice")
// {"time":"2024-03-15T09:30:00Z","level":"INFO","msg":"user created","id":42,"name":"Alice"}

slog.Error("query failed", "error", err, "query", "SELECT ...")
slog.Debug("cache miss", "key", key)
slog.Warn("slow query", "duration_ms", ms)
```

### Log Levels

```
DEBUG < INFO < WARN < ERROR
```

Production typically uses INFO. Development uses DEBUG.

### Creating a Custom slog Logger

```go
// JSON output (machine-readable, for production)
logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
    Level: slog.LevelDebug,
}))

// Text output (human-readable, for development)
logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
    Level: slog.LevelInfo,
}))

// Use it
logger.Info("starting", "port", 8080)
logger.Error("failed", "err", err)
```

### Adding Context with `With`

```go
// Create a child logger with fixed fields (request ID, user ID, etc.)
reqLogger := logger.With("request_id", requestID, "user_id", userID)
reqLogger.Info("handler called")  // always includes request_id and user_id
reqLogger.Info("handler done", "status", 200)
```

### Request Logger Middleware Pattern

```go
func RequestLogger(logger *slog.Logger, next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        logger.Info("request",
            "method", r.Method,
            "path", r.URL.Path,
            "duration_ms", time.Since(start).Milliseconds(),
        )
    })
}
```

### Setting the Default slog Logger

```go
slog.SetDefault(logger)  // now slog.Info, slog.Error etc. use this logger
```

### Solving the Challenge

```go
package piscine

import (
    "io"
    "log/slog"
    "net/http"
    "strings"
    "time"
)

func NewLogger(w io.Writer, level string) *slog.Logger {
    var lvl slog.Level
    switch strings.ToLower(level) {
    case "debug":
        lvl = slog.LevelDebug
    case "warn":
        lvl = slog.LevelWarn
    case "error":
        lvl = slog.LevelError
    default:
        lvl = slog.LevelInfo
    }
    return slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{Level: lvl}))
}

func RequestLogger(logger *slog.Logger, next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        logger.Info("request",
            "method", r.Method,
            "path", r.URL.Path,
            "duration_ms", time.Since(start).Milliseconds(),
        )
    })
}
```

**Next:** [151-generics](../151-generics/README.md)
