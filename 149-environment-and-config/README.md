# 149-environment-and-config — Environment Variables and Configuration

## Challenge

Implement in `package piscine`:

```go
type Config struct {
    Host     string
    Port     int
    DBPath   string
    LogLevel string
    Debug    bool
}

// LoadConfig builds a Config from environment variables.
// Env vars: APP_HOST (default "localhost"), APP_PORT (default 8080),
//           APP_DB_PATH (default "app.db"), APP_LOG_LEVEL (default "info"),
//           APP_DEBUG (default false).
func LoadConfig() Config

// RegisterFlags registers config fields as CLI flags and returns a pointer to Config.
// Flags: -host, -port, -db-path, -log-level, -debug
func RegisterFlags() *Config
```

Read `skills.md` before you start.

---

## Time limit

**25 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 149
```

or run the tests directly:

```bash
go test ./149-environment-and-config/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 149` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
