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
