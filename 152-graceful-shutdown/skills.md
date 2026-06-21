# Skills for 152-graceful-shutdown

## What You'll Learn

**Previous:** [151-generics](../151-generics/skills.md) | **Next:** [153-financial-freedom-api](../153-financial-freedom-api/skills.md)

**Challenge:** Start an HTTP server that shuts down cleanly when the OS sends a signal.

## Why Graceful Shutdown Matters

When Kubernetes stops a pod, or a user presses Ctrl+C, the OS sends `SIGTERM` (or `SIGINT`).
Without graceful shutdown, the server kills all in-flight requests instantly — transactions get half-committed, files get corrupted, clients get broken pipe errors.

Graceful shutdown means:
1. Stop accepting new connections
2. Wait for in-flight requests to finish
3. Close the database, flush logs, release resources
4. Exit cleanly

Every production Go service implements this.

## OS Signals

```go
import (
    "os"
    "os/signal"
    "syscall"
)

quit := make(chan os.Signal, 1)
signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

<-quit  // block here until a signal arrives
fmt.Println("shutting down...")
```

- `syscall.SIGINT` — Ctrl+C in the terminal
- `syscall.SIGTERM` — sent by `kill` command, Kubernetes, Docker
- Buffered channel `make(chan os.Signal, 1)` — prevents signal loss if we're not ready to receive

## `http.Server.Shutdown`

`http.ListenAndServe` doesn't support graceful shutdown. Use `http.Server` directly:

```go
srv := &http.Server{
    Addr:         ":8080",
    Handler:      mux,
    ReadTimeout:  5 * time.Second,
    WriteTimeout: 10 * time.Second,
    IdleTimeout:  60 * time.Second,
}

// Start in a goroutine so it doesn't block
go func() {
    if err := srv.ListenAndServe(); err != http.ErrServerClosed {
        log.Fatal("server error:", err)
    }
}()

// Wait for signal
quit := make(chan os.Signal, 1)
signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
<-quit

// Shut down with a deadline
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

if err := srv.Shutdown(ctx); err != nil {
    log.Fatal("forced shutdown:", err)
}
log.Println("server stopped cleanly")
```

**What `srv.Shutdown(ctx)` does:**
- Closes the listener (no new connections)
- Waits for active connections to finish
- Returns `nil` when all connections are done
- Returns `ctx.Err()` if the deadline expires before all connections finish

`srv.ListenAndServe()` returns `http.ErrServerClosed` after `Shutdown` is called — that is not an error, so we check for it explicitly.

## Timeouts — Always Set These

```go
srv := &http.Server{
    ReadTimeout:  5 * time.Second,   // max time to read the full request
    WriteTimeout: 10 * time.Second,  // max time to write the full response
    IdleTimeout:  60 * time.Second,  // max time to wait for the next request (keep-alive)
}
```

Without timeouts, a slow client can hold a connection open forever (a slow-loris attack).

## Cleanup on Shutdown

Combine with `defer` to release resources:

```go
func main() {
    db, _ := sql.Open("sqlite", "app.db")
    defer db.Close()

    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

    srv := &http.Server{Addr: ":8080", Handler: router(db, logger)}

    go func() {
        logger.Info("listening", "addr", srv.Addr)
        if err := srv.ListenAndServe(); err != http.ErrServerClosed {
            logger.Error("server error", "err", err)
            os.Exit(1)
        }
    }()

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    sig := <-quit
    logger.Info("signal received", "signal", sig)

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        logger.Error("shutdown error", "err", err)
    }
    logger.Info("shutdown complete")
    // defer db.Close() runs here
}
```

## Solving the Challenge

```go
package piscine

import (
    "context"
    "log/slog"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func NewServer(addr string, handler http.Handler) *http.Server {
    return &http.Server{
        Addr:         addr,
        Handler:      handler,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  60 * time.Second,
    }
}

func RunWithGracefulShutdown(srv *http.Server) error {
    logger := slog.Default()

    errCh := make(chan error, 1)
    go func() {
        logger.Info("server starting", "addr", srv.Addr)
        if err := srv.ListenAndServe(); err != http.ErrServerClosed {
            errCh <- err
        }
        close(errCh)
    }()

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

    select {
    case err := <-errCh:
        return err
    case sig := <-quit:
        logger.Info("signal received", "signal", sig)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        return err
    }
    logger.Info("shutdown complete")
    return <-errCh
}
```

**Next:** [153-financial-freedom-api](../153-financial-freedom-api/README.md)
