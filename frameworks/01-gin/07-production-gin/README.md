# gin/07-production-gin — Production-Ready Gin Server

## Challenge

Evolve your Task Manager into a production-ready server with:

1. **Graceful shutdown** — on SIGINT/SIGTERM, finish in-flight requests then stop
2. **Configuration from environment** — port, JWT secret, log level from env vars
3. **Structured logging** — use `log/slog` for JSON log output
4. **Health check** endpoint: `GET /health → { "status": "ok", "uptime": "3m12s" }`
5. **Request ID middleware** — add a unique `X-Request-ID` header to every response
6. **Timeout middleware** — cancel requests that take longer than 5 seconds
7. **Panic recovery** — convert panics to 500 JSON responses with the request ID

## Config

```go
type Config struct {
    Port       string        // PORT env var, default "8080"
    JWTSecret  string        // JWT_SECRET env var
    LogLevel   string        // LOG_LEVEL env var: "debug"|"info"|"warn"|"error"
    ShutdownTO time.Duration // SHUTDOWN_TIMEOUT env var, default 30s
}
```

## Test

```bash
PORT=9090 LOG_LEVEL=debug go run .

curl http://localhost:9090/health
# → {"status":"ok","uptime":"5s"}

# Server logs (JSON):
# {"time":"2026-01-01T00:00:05Z","level":"INFO","msg":"request","method":"GET","path":"/health","status":200,"latency":"0.5ms","request_id":"abc123"}
```

Read `skills.md` before you start.
