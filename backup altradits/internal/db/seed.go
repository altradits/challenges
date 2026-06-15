package db

import (
	"crypto/sha256"
	"fmt"
	"log"
)

func Seed() error {
	// Check if admin exists
	var count int
	DB.QueryRow("SELECT COUNT(*) FROM users WHERE role='admin'").Scan(&count)
	if count == 0 {
		// Create default admin
		hash := sha256.Sum256([]byte("admin123"))
		passwordHash := fmt.Sprintf("%x", hash)

		_, err := DB.Exec(`INSERT INTO users (identifier, password, role, full_name)
			VALUES ($1, $2, 'admin', 'Altradits Admin')
			ON CONFLICT (identifier) DO NOTHING`,
			"admin@altradits.com", passwordHash)
		if err != nil {
			return fmt.Errorf("seed admin: %w", err)
		}

		var adminID string
		DB.QueryRow("SELECT id FROM users WHERE identifier = 'admin@altradits.com'").Scan(&adminID)
		if adminID != "" {
			DB.Exec(`INSERT INTO wallets (user_id, lightning_addr) VALUES ($1, 'admin@altradits.com') ON CONFLICT DO NOTHING`, adminID)
		}

		// Create demo trader
		hash2 := sha256.Sum256([]byte("trader123"))
		DB.Exec(`INSERT INTO users (identifier, password, role, full_name)
			VALUES ('trader@altradits.com', $1, 'trader', 'Altradits Trader')
			ON CONFLICT (identifier) DO NOTHING`, fmt.Sprintf("%x", hash2))

		var traderID string
		DB.QueryRow("SELECT id FROM users WHERE identifier = 'trader@altradits.com'").Scan(&traderID)
		if traderID != "" {
			DB.Exec(`INSERT INTO wallets (user_id, lightning_addr) VALUES ($1, 'trader@altradits.com') ON CONFLICT DO NOTHING`, traderID)
		}

		log.Println("[seed] default admin: admin@altradits.com / admin123")
		log.Println("[seed] default trader: trader@altradits.com / trader123")
	}

	// Create demo agent (independent of admin, so it backfills on existing installs)
	var agentCount int
	DB.QueryRow("SELECT COUNT(*) FROM users WHERE identifier='agent@altradits.com'").Scan(&agentCount)
	if agentCount == 0 {
		hash3 := sha256.Sum256([]byte("agent123"))
		DB.Exec(`INSERT INTO users (identifier, password, role, full_name)
			VALUES ('agent@altradits.com', $1, 'agent', 'Wanjiku — Agent')
			ON CONFLICT (identifier) DO NOTHING`, fmt.Sprintf("%x", hash3))

		var agentID string
		DB.QueryRow("SELECT id FROM users WHERE identifier = 'agent@altradits.com'").Scan(&agentID)
		if agentID != "" {
			DB.Exec(`INSERT INTO wallets (user_id, lightning_addr, balance_sats) VALUES ($1, 'agent@altradits.com', 250000) ON CONFLICT DO NOTHING`, agentID)
		}

		log.Println("[seed] default agent: agent@altradits.com / agent123")
	}

	return nil
}
