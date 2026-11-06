package piscine

import (
	"fmt"
	"strconv"
	"strings"
)

type Expense struct {
	Category string
	Amount   float64
	Date     string
}

// ParseExpense parses "Groceries: $150.50" (Date left empty).
func ParseExpense(s string) (Expense, error) {
	parts := strings.SplitN(s, ":", 2)
	if len(parts) != 2 {
		return Expense{}, fmt.Errorf("missing colon in %q", s)
	}
	amount := strings.TrimSpace(parts[1])
	amount = strings.ReplaceAll(amount, "$", "")
	amount = strings.ReplaceAll(amount, ",", "")
	n, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return Expense{}, fmt.Errorf("bad amount in %q: %w", s, err)
	}
	return Expense{Category: strings.TrimSpace(parts[0]), Amount: n}, nil
}

// CalculateRunway converts savings into (years, months). Non-positive
// expenses mean the savings never drain: (0, 0).
func CalculateRunway(savings, monthlyExpenses float64) (years, months int) {
	if monthlyExpenses <= 0 || savings <= 0 {
		return 0, 0
	}
	total := int(savings / monthlyExpenses)
	return total / 12, total % 12
}

func FormatRunway(years, months int) string {
	yWord, mWord := "years", "months"
	if years == 1 {
		yWord = "year"
	}
	if months == 1 {
		mWord = "month"
	}
	return fmt.Sprintf("%d %s %d %s", years, yWord, months, mWord)
}

func GenerateMilestone(months int) string {
	if months < 12 {
		return fmt.Sprintf("Milestone: %d months of savings!", months)
	}
	years := months / 12
	if years == 1 {
		return "Milestone: 1 year of savings!"
	}
	return fmt.Sprintf("Milestone: %d years of savings!", years)
}
