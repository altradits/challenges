# 152-graceful-shutdown — Graceful HTTP Shutdown

## Challenge

Implement in `package piscine`:

```go
// NewServer builds an http.Server with the given handler and addr.
func NewServer(addr string, handler http.Handler) *http.Server

// RunWithGracefulShutdown starts srv and blocks until SIGINT or SIGTERM is received.
// It then calls srv.Shutdown with a 10-second deadline, waiting for in-flight requests
// to complete before returning.
func RunWithGracefulShutdown(srv *http.Server) error
```

Read `skills.md` before you start.
