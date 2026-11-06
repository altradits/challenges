package piscine

import "testing"

func TestParseExpense(t *testing.T) {
	e, err := ParseExpense("Groceries: $150.50")
	if err != nil {
		t.Fatalf("ParseExpense: %v", err)
	}
	if e.Category != "Groceries" || e.Amount != 150.50 {
		t.Errorf("ParseExpense = %+v", e)
	}
	if _, err := ParseExpense("no colon"); err == nil {
		t.Error("ParseExpense without colon should fail")
	}
	if _, err := ParseExpense("Food: $abc"); err == nil {
		t.Error("ParseExpense with bad amount should fail")
	}
}

func TestCalculateRunway155(t *testing.T) {
	y, m := CalculateRunway(3000, 100)
	if y != 2 || m != 6 {
		t.Errorf("CalculateRunway(3000, 100) = (%d, %d), want (2, 6)", y, m)
	}
	y, m = CalculateRunway(1200, 100)
	if y != 1 || m != 0 {
		t.Errorf("CalculateRunway(1200, 100) = (%d, %d), want (1, 0)", y, m)
	}
	y, m = CalculateRunway(50, 100)
	if y != 0 || m != 0 {
		t.Errorf("CalculateRunway(50, 100) = (%d, %d), want (0, 0)", y, m)
	}
	y, m = CalculateRunway(1000, 0)
	if y != 0 || m != 0 {
		t.Errorf("CalculateRunway zero expenses = (%d, %d), want (0, 0)", y, m)
	}
}

func TestFormatRunway(t *testing.T) {
	if got := FormatRunway(2, 6); got != "2 years 6 months" {
		t.Errorf("FormatRunway = %q", got)
	}
	if got := FormatRunway(1, 1); got != "1 year 1 month" {
		t.Errorf("FormatRunway singular = %q", got)
	}
}

func TestGenerateMilestone(t *testing.T) {
	if got := GenerateMilestone(12); got != "Milestone: 1 year of savings!" {
		t.Errorf("GenerateMilestone(12) = %q", got)
	}
	if got := GenerateMilestone(30); got != "Milestone: 2 years of savings!" {
		t.Errorf("GenerateMilestone(30) = %q", got)
	}
	if got := GenerateMilestone(6); got != "Milestone: 6 months of savings!" {
		t.Errorf("GenerateMilestone(6) = %q", got)
	}
}
