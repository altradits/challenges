package piscine

import (
	"flag"
	"os"
	"testing"
)

var configKeys = []string{"APP_HOST", "APP_PORT", "APP_DB_PATH", "APP_LOG_LEVEL", "APP_DEBUG"}

// clearEnv unsets all config keys (t.Setenv cannot unset).
func clearEnv(t *testing.T) {
	t.Helper()
	for _, k := range configKeys {
		if v, ok := os.LookupEnv(k); ok {
			t.Cleanup(func() { os.Setenv(k, v) })
		} else {
			t.Cleanup(func() { os.Unsetenv(k) })
		}
		os.Unsetenv(k)
	}
}

func TestLoadConfigDefaults(t *testing.T) {
	clearEnv(t)
	cfg := LoadConfig()
	want := Config{Host: "localhost", Port: 8080, DBPath: "app.db", LogLevel: "info", Debug: false}
	if cfg != want {
		t.Errorf("LoadConfig defaults = %+v, want %+v", cfg, want)
	}
}

func TestLoadConfigFromEnv(t *testing.T) {
	clearEnv(t)
	t.Setenv("APP_HOST", "db.internal")
	t.Setenv("APP_PORT", "5432")
	t.Setenv("APP_DB_PATH", "/data/app.db")
	t.Setenv("APP_LOG_LEVEL", "debug")
	t.Setenv("APP_DEBUG", "true")
	cfg := LoadConfig()
	want := Config{Host: "db.internal", Port: 5432, DBPath: "/data/app.db", LogLevel: "debug", Debug: true}
	if cfg != want {
		t.Errorf("LoadConfig = %+v, want %+v", cfg, want)
	}
}

func TestLoadConfigBadValues(t *testing.T) {
	clearEnv(t)
	t.Setenv("APP_PORT", "notaport")
	t.Setenv("APP_DEBUG", "notabool")
	cfg := LoadConfig()
	if cfg.Port != 8080 {
		t.Errorf("bad APP_PORT: got %d, want default 8080", cfg.Port)
	}
	if cfg.Debug != false {
		t.Errorf("bad APP_DEBUG: got true, want default false")
	}
}

func TestRegisterFlags(t *testing.T) {
	cfg := RegisterFlags()
	if cfg == nil {
		t.Fatal("RegisterFlags returned nil")
	}
	for _, name := range []string{"host", "port", "db-path", "log-level", "debug"} {
		if flag.Lookup(name) == nil {
			t.Errorf("flag -%s not registered", name)
		}
	}
	if cfg.Host != "localhost" || cfg.Port != 8080 || cfg.DBPath != "app.db" || cfg.LogLevel != "info" || cfg.Debug != false {
		t.Errorf("flag defaults = %+v", cfg)
	}
}
