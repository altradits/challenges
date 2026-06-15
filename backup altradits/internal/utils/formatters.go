package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// HashPassword uses SHA-256 (upgrade to bcrypt when go modules are available)
func HashPassword(password string) string {
	h := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", h)
}

func CheckPassword(password, hash string) bool {
	return HashPassword(password) == hash
}

// GenerateToken creates a secure random session token
func GenerateToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// GenerateQRSecret creates a short secret for QR check-ins
func GenerateQRSecret() (string, error) {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// GenerateCertHash creates a unique certificate identifier
func GenerateCertHash(userID, hackathonID string) string {
	n, _ := rand.Int(rand.Reader, big.NewInt(999999))
	raw := fmt.Sprintf("%s-%s-%d-%d", userID, hackathonID, n.Int64(), time.Now().UnixNano())
	h := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(h[:])
}

// FormatSats formats a satoshi amount with thousand separators
func FormatSats(sats int64) string {
	s := strconv.FormatInt(sats, 10)
	// Insert commas
	result := ""
	for i, c := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			result += ","
		}
		result += string(c)
	}
	return result + " sats"
}

// FormatKES formats Kenyan shillings
func FormatKES(kes float64) string {
	return fmt.Sprintf("KES %.2f", kes)
}

// FormatKESAmount formats a KES amount with 2 decimals, no currency prefix
func FormatKESAmount(kes float64) string {
	return fmt.Sprintf("%.2f", kes)
}

// SatsToKES converts sats to KES using stored rate
func SatsToKES(sats int64, kesPerSat float64) float64 {
	return float64(sats) * kesPerSat
}

// FormatBTC formats a satoshi amount as BTC, trimming trailing zeros
func FormatBTC(sats int64) string {
	btc := float64(sats) / 100000000.0
	s := strconv.FormatFloat(btc, 'f', 8, 64)
	s = strings.TrimRight(s, "0")
	s = strings.TrimRight(s, ".")
	if s == "" || s == "-" {
		s = "0"
	}
	return "₿" + s
}

// KESToSats converts KES to sats
func KESToSats(kes float64, kesPerSat float64) int64 {
	if kesPerSat == 0 {
		return 0
	}
	return int64(kes / kesPerSat)
}

// ValidateIdentifier checks email or phone
func ValidateIdentifier(id string) bool {
	id = strings.TrimSpace(id)
	emailRe := regexp.MustCompile(`^[^@]+@[^@]+\.[^@]+$`)
	phoneRe := regexp.MustCompile(`^\+?[\d\s\-]{9,15}$`)
	return emailRe.MatchString(id) || phoneRe.MatchString(id)
}

// SessionExpiry returns session expiry time (72 hours)
func SessionExpiry() time.Time {
	return time.Now().Add(72 * time.Hour)
}

// TimeAgo returns human-readable time difference
func TimeAgo(t time.Time) string {
	d := time.Since(t)
	switch {
	case d < time.Minute:
		return "just now"
	case d < time.Hour:
		m := int(d.Minutes())
		if m == 1 {
			return "1 minute ago"
		}
		return fmt.Sprintf("%d minutes ago", m)
	case d < 24*time.Hour:
		h := int(d.Hours())
		if h == 1 {
			return "1 hour ago"
		}
		return fmt.Sprintf("%d hours ago", h)
	case d < 30*24*time.Hour:
		days := int(d.Hours() / 24)
		if days == 1 {
			return "yesterday"
		}
		return fmt.Sprintf("%d days ago", days)
	default:
		return t.Format("Jan 2, 2006")
	}
}

// DaysUntil returns days until a future time
func DaysUntil(t time.Time) int {
	return int(time.Until(t).Hours() / 24)
}

// EarlyWithdrawalPenalty returns penalty percentage based on years elapsed
func EarlyWithdrawalPenalty(lockedAt time.Time) int {
	years := time.Since(lockedAt).Hours() / 8760
	switch {
	case years < 1:
		return 100
	case years < 2:
		return 75
	case years < 3:
		return 50
	case years < 4:
		return 25
	default:
		return 10
	}
}
