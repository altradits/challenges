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

func getenv(key, def string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return def
}

func LoadConfig() Config {
	port := 8080
	if v, ok := os.LookupEnv("APP_PORT"); ok {
		if n, err := strconv.Atoi(v); err == nil {
			port = n
		}
	}
	debug := false
	if v, ok := os.LookupEnv("APP_DEBUG"); ok {
		if b, err := strconv.ParseBool(v); err == nil {
			debug = b
		}
	}
	return Config{
		Host:     getenv("APP_HOST", "localhost"),
		Port:     port,
		DBPath:   getenv("APP_DB_PATH", "app.db"),
		LogLevel: getenv("APP_LOG_LEVEL", "info"),
		Debug:    debug,
	}
}

func RegisterFlags() *Config {
	cfg := &Config{}
	flag.StringVar(&cfg.Host, "host", "localhost", "server host")
	flag.IntVar(&cfg.Port, "port", 8080, "server port")
	flag.StringVar(&cfg.DBPath, "db-path", "app.db", "database path")
	flag.StringVar(&cfg.LogLevel, "log-level", "info", "log level")
	flag.BoolVar(&cfg.Debug, "debug", false, "debug mode")
	return cfg
}
