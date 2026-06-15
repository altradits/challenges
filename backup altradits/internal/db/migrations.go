package db

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func RunMigrations() error {
	// Create migrations tracking table
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS schema_migrations (
		version VARCHAR(255) PRIMARY KEY,
		applied_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	)`)
	if err != nil {
		return fmt.Errorf("create migrations table: %w", err)
	}

	// Find migration files
	migrationsDir := os.Getenv("MIGRATIONS_DIR")
	if migrationsDir == "" {
		migrationsDir = "docs/database/migrations"
	}

	files, err := filepath.Glob(filepath.Join(migrationsDir, "*.sql"))
	if err != nil {
		return fmt.Errorf("glob migrations: %w", err)
	}
	sort.Strings(files)

	for _, f := range files {
		version := filepath.Base(f)
		version = strings.TrimSuffix(version, ".sql")

		var exists string
		err := DB.QueryRow("SELECT version FROM schema_migrations WHERE version = $1", version).Scan(&exists)
		if err == nil {
			log.Printf("[migrations] already applied: %s", version)
			continue
		}

		content, err := os.ReadFile(f)
		if err != nil {
			return fmt.Errorf("read %s: %w", f, err)
		}

		log.Printf("[migrations] applying: %s", version)
		if _, err := DB.Exec(string(content)); err != nil {
			return fmt.Errorf("exec %s: %w", version, err)
		}

		if _, err := DB.Exec("INSERT INTO schema_migrations (version) VALUES ($1)", version); err != nil {
			return fmt.Errorf("record migration %s: %w", version, err)
		}
		log.Printf("[migrations] applied: %s", version)
	}
	return nil
}
