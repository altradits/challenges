# 150-logging — Logging in Go

## Challenge

Implement in `package piscine`:

```go
// NewLogger returns a *slog.Logger that writes JSON to w at the given level.
// level is one of: "debug", "info", "warn", "error"
func NewLogger(w io.Writer, level string) *slog.Logger

// RequestLogger returns an http.Handler that wraps next,
// logging method, path, and duration for every request.
func RequestLogger(logger *slog.Logger, next http.Handler) http.Handler
```

Read `skills.md` before you start.
