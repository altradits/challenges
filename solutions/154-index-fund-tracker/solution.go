package piscine

import (
	"fmt"
	"strconv"
	"strings"
)

type Fund struct {
	Ticker    string
	Name      string
	Price     float64
	Change    float64
	ChangePct float64
}

// ParseTicker parses "VTI: $245.32 (+1.2%)" into a Fund (Name left empty).
func ParseTicker(s string) (Fund, error) {
	var f Fund
	parts := strings.SplitN(s, ":", 2)
	if len(parts) != 2 {
		return Fund{}, fmt.Errorf("missing ticker colon in %q", s)
	}
	f.Ticker = strings.TrimSpace(parts[0])
	rest := strings.TrimSpace(parts[1])
	dollar := strings.Index(rest, "$")
	lparen := strings.Index(rest, "(")
	rparen := strings.Index(rest, ")")
	if dollar < 0 || lparen < 0 || rparen < lparen {
		return Fund{}, fmt.Errorf("bad fund quote %q", s)
	}
	price, err := strconv.ParseFloat(strings.TrimSpace(rest[dollar+1:lparen]), 64)
	if err != nil {
		return Fund{}, fmt.Errorf("bad price in %q: %w", s, err)
	}
	pct := strings.TrimSuffix(strings.TrimSpace(rest[lparen+1:rparen]), "%")
	changePct, err := strconv.ParseFloat(pct, 64)
	if err != nil {
		return Fund{}, fmt.Errorf("bad change in %q: %w", s, err)
	}
	f.Price = price
	f.ChangePct = changePct
	return f, nil
}

func FormatChange(change float64) string {
	return fmt.Sprintf("%+.1f%%", change)
}

func FormatComparison(f1, f2 Fund) string {
	diff := f1.ChangePct - f2.ChangePct
	switch {
	case diff > 0:
		return fmt.Sprintf("%s outperforms %s by %.1f%%", f1.Ticker, f2.Ticker, diff)
	case diff < 0:
		return fmt.Sprintf("%s underperforms %s by %.1f%%", f1.Ticker, f2.Ticker, -diff)
	default:
		return fmt.Sprintf("%s matches %s", f1.Ticker, f2.Ticker)
	}
}

func GeneratePerformanceReport(funds []Fund) string {
	var b strings.Builder
	b.WriteString("Fund Performance Report\n")
	for _, f := range funds {
		fmt.Fprintf(&b, "%s: %s @ $%.2f\n", f.Ticker, FormatChange(f.ChangePct), f.Price)
	}
	return b.String()
}
