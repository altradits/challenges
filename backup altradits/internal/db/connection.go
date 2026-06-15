package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/altradits/altradits/internal/pgdrv"
)

var DB *sql.DB

func Connect() error {
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		dsn = "postgres://altradits:password@localhost:5432/altradits?sslmode=disable"
	}

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("sql.Open: %w", err)
	}

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(5 * time.Minute)

	for i := 0; i < 10; i++ {
		if err = DB.Ping(); err == nil {
			log.Println("[db] connected to PostgreSQL")
			return nil
		}
		log.Printf("[db] waiting for postgres (%d/10): %v", i+1, err)
		time.Sleep(2 * time.Second)
	}
	return fmt.Errorf("db.Ping: %w", err)
}
