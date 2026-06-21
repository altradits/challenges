# Skills for 149-environment-and-config

## What You'll Learn

**Previous:** [148-database-sql](../148-database-sql/skills.md) | **Next:** [150-logging](../150-logging/skills.md)

**Challenge:** Read environment variables and CLI flags into a config struct.

## Environment Variables

```go
import "os"

// Read — returns "" if not set
val := os.Getenv("DATABASE_URL")

// Read with ok — distinguish unset from empty
val, ok := os.LookupEnv("DATABASE_URL")
if !ok {
    val = "default.db"
}

// Set (rarely needed in production code — use for tests)
os.Setenv("APP_ENV", "test")
os.Unsetenv("APP_ENV")
```

### Pattern: Env with Default

```go
func getEnv(key, defaultVal string) string {
    if val, ok := os.LookupEnv(key); ok {
        return val
    }
    return defaultVal
}
```

### Parsing Typed Values from Env

```go
import "strconv"

portStr := os.Getenv("APP_PORT")
port, err := strconv.Atoi(portStr)
if err != nil {
    port = 8080  // default
}

debug := os.Getenv("APP_DEBUG") == "true"
```

## The `flag` Package

Go's built-in CLI flag parser:

```go
import "flag"

host := flag.String("host", "localhost", "server hostname")
port := flag.Int("port", 8080, "server port")
debug := flag.Bool("debug", false, "enable debug mode")

flag.Parse()  // MUST call before using flag values

fmt.Println(*host, *port, *debug)
```

Usage is automatically generated:
```
-host string
    server hostname (default "localhost")
-port int
    server port (default 8080)
```

Run: `./myapp -host 0.0.0.0 -port 9090 -debug`

### Binding Flags to Existing Variables

```go
var cfg Config
flag.StringVar(&cfg.Host, "host", "localhost", "server hostname")
flag.IntVar(&cfg.Port, "port", 8080, "server port")
flag.BoolVar(&cfg.Debug, "debug", false, "enable debug mode")
flag.Parse()
// cfg.Host is now set from -host flag
```

### Config Struct Pattern

The idiomatic way to manage configuration:

```go
type Config struct {
    Host     string
    Port     int
    DBPath   string
    LogLevel string
    Debug    bool
}

func LoadConfig() Config {
    return Config{
        Host:     getEnv("APP_HOST", "localhost"),
        Port:     getEnvInt("APP_PORT", 8080),
        DBPath:   getEnv("APP_DB_PATH", "app.db"),
        LogLevel: getEnv("APP_LOG_LEVEL", "info"),
        Debug:    os.Getenv("APP_DEBUG") == "true",
    }
}
```

### 12-Factor App Pattern

Configuration best practices for backend services:
- All config comes from environment variables (never hardcoded)
- Sensible defaults for local development
- Secrets (DB passwords, API keys) ONLY in env vars, never in code or files
- Use `.env` file loading (`github.com/joho/godotenv`) in development

```go
import "github.com/joho/godotenv"

func init() {
    // loads .env file if it exists (silently ignores if absent)
    godotenv.Load()
}
```

### Solving the Challenge

```go
package piscine

import (
    "flag"
    "os"
    "strconv"
)

type Config struct {
    Host     string
    Port     int
    DBPath   string
    LogLevel string
    Debug    bool
}

func getEnv(key, def string) string {
    if v, ok := os.LookupEnv(key); ok {
        return v
    }
    return def
}

func LoadConfig() Config {
    port, _ := strconv.Atoi(getEnv("APP_PORT", "8080"))
    return Config{
        Host:     getEnv("APP_HOST", "localhost"),
        Port:     port,
        DBPath:   getEnv("APP_DB_PATH", "app.db"),
        LogLevel: getEnv("APP_LOG_LEVEL", "info"),
        Debug:    os.Getenv("APP_DEBUG") == "true",
    }
}

func RegisterFlags() *Config {
    cfg := &Config{}
    flag.StringVar(&cfg.Host, "host", "localhost", "server hostname")
    flag.IntVar(&cfg.Port, "port", 8080, "server port")
    flag.StringVar(&cfg.DBPath, "db-path", "app.db", "SQLite database path")
    flag.StringVar(&cfg.LogLevel, "log-level", "info", "log level")
    flag.BoolVar(&cfg.Debug, "debug", false, "enable debug mode")
    return cfg
}
```

**Next:** [150-logging](../150-logging/README.md)
