package piscine

import "testing"

func TestCalculateRunway153(t *testing.T) {
	cases := []struct {
		name string
		data FinancialData
		want int
	}{
		{"exact", FinancialData{MonthlyIncome: 5000, MonthlyExpenses: 3000, CurrentSavings: 0, TargetSavings: 24000}, 12},
		{"rounds up", FinancialData{MonthlyIncome: 4000, MonthlyExpenses: 3000, CurrentSavings: 0, TargetSavings: 2500}, 3},
		{"already there", FinancialData{MonthlyIncome: 1000, MonthlyExpenses: 500, CurrentSavings: 10000, TargetSavings: 8000}, 0},
		{"spending too much", FinancialData{MonthlyIncome: 2000, MonthlyExpenses: 2500, CurrentSavings: 0, TargetSavings: 1000}, -1},
		{"break even", FinancialData{MonthlyIncome: 2000, MonthlyExpenses: 2000, CurrentSavings: 0, TargetSavings: 1000}, -1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := CalculateRunway(tc.data); got != tc.want {
				t.Errorf("CalculateRunway(%+v) = %d, want %d", tc.data, got, tc.want)
			}
		})
	}
}

func TestValidateAccountNumber(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  bool
	}{
		{"valid", "1234567890", true},
		{"short", "12345", false},
		{"long", "12345678901", false},
		{"letters", "12345abcde", false},
		{"empty", "", false},
		{"spaces", "12345 7890", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ValidateAccountNumber(tc.input); got != tc.want {
				t.Errorf("ValidateAccountNumber(%q) = %v, want %v (exactly 10 digits)", tc.input, got, tc.want)
			}
		})
	}
}

func TestNormalizeCurrency(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"dollar comma", "$1,234.56", "1234.56"},
		{"plain", "99.9", "99.9"},
		{"spaces", "  $5  ", "5"},
		{"empty", "", ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := NormalizeCurrency(tc.input); got != tc.want {
				t.Errorf("NormalizeCurrency(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}

func TestGenerateReport(t *testing.T) {
	data := FinancialData{MonthlyIncome: 5000, MonthlyExpenses: 3000, CurrentSavings: 0, TargetSavings: 48000}
	if got, want := GenerateReport(data), "Savings Runway: 24 months (2 years)"; got != want {
		t.Errorf("GenerateReport = %q, want %q", got, want)
	}
	poor := FinancialData{MonthlyIncome: 1000, MonthlyExpenses: 2000, CurrentSavings: 0, TargetSavings: 5000}
	if got := GenerateReport(poor); got != "Savings Runway: never (expenses exceed income)" {
		t.Errorf("GenerateReport broke = %q", got)
	}
}
