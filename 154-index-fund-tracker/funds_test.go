package piscine

import (
	"strings"
	"testing"
)

func TestParseTicker(t *testing.T) {
	f, err := ParseTicker("VTI: $245.32 (+1.2%)")
	if err != nil {
		t.Fatalf("ParseTicker: %v", err)
	}
	if f.Ticker != "VTI" || f.Price != 245.32 || f.ChangePct != 1.2 {
		t.Errorf("ParseTicker = %+v", f)
	}
	f2, err := ParseTicker("VOO: $100.00 (-0.5%)")
	if err != nil {
		t.Fatal(err)
	}
	if f2.ChangePct != -0.5 {
		t.Errorf("negative change = %v", f2.ChangePct)
	}
	for _, bad := range []string{"no colon here", "VTI no price", "VTI: $abc (+1%)", "VTI: $10 (+x%)"} {
		if _, err := ParseTicker(bad); err == nil {
			t.Errorf("ParseTicker(%q) should fail", bad)
		}
	}
}

func TestFormatChange(t *testing.T) {
	if got := FormatChange(1.2); got != "+1.2%" {
		t.Errorf("FormatChange(1.2) = %q", got)
	}
	if got := FormatChange(-0.5); got != "-0.5%" {
		t.Errorf("FormatChange(-0.5) = %q", got)
	}
	if got := FormatChange(0); got != "+0.0%" {
		t.Errorf("FormatChange(0) = %q", got)
	}
}

func TestFormatComparison(t *testing.T) {
	vti := Fund{Ticker: "VTI", ChangePct: 3.5}
	voo := Fund{Ticker: "VOO", ChangePct: 1.2}
	if got, want := FormatComparison(vti, voo), "VTI outperforms VOO by 2.3%"; got != want {
		t.Errorf("comparison = %q, want %q", got, want)
	}
	if got, want := FormatComparison(voo, vti), "VOO underperforms VTI by 2.3%"; got != want {
		t.Errorf("comparison = %q, want %q", got, want)
	}
	if got := FormatComparison(vti, Fund{Ticker: "X", ChangePct: 3.5}); !strings.Contains(got, "matches") {
		t.Errorf("tie comparison = %q", got)
	}
}

func TestGeneratePerformanceReport(t *testing.T) {
	report := GeneratePerformanceReport([]Fund{
		{Ticker: "VTI", Price: 245.32, ChangePct: 1.2},
		{Ticker: "VOO", Price: 100, ChangePct: -0.5},
	})
	for _, want := range []string{"Fund Performance Report", "VTI: +1.2% @ $245.32", "VOO: -0.5% @ $100.00"} {
		if !strings.Contains(report, want) {
			t.Errorf("report missing %q:\n%s", want, report)
		}
	}
}
