package piscine

import (
	"fmt"
	"math"
	"strings"
)

type FinancialData struct {
	MonthlyIncome   float64
	MonthlyExpenses float64
	CurrentSavings  float64
	TargetSavings   float64
}

// CalculateRunway returns whole months until TargetSavings is reached, 0 if
// already there, or -1 if expenses meet or exceed income.
func CalculateRunway(data FinancialData) int {
	remaining := data.TargetSavings - data.CurrentSavings
	if remaining <= 0 {
		return 0
	}
	monthly := data.MonthlyIncome - data.MonthlyExpenses
	if monthly <= 0 {
		return -1
	}
	return int(math.Ceil(remaining / monthly))
}

// ValidateAccountNumber reports whether s is exactly 10 ASCII digits.
func ValidateAccountNumber(s string) bool {
	if len(s) != 10 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

// NormalizeCurrency strips "$", commas and surrounding spaces.
func NormalizeCurrency(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "$", "")
	s = strings.ReplaceAll(s, ",", "")
	return s
}

func GenerateReport(data FinancialData) string {
	months := CalculateRunway(data)
	if months < 0 {
		return "Savings Runway: never (expenses exceed income)"
	}
	return fmt.Sprintf("Savings Runway: %d months (%d years)", months, months/12)
}
