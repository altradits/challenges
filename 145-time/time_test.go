package piscine

import (
	"testing"
	"time"
)

func TestAgeInYears(t *testing.T) {
	now := time.Now()
	thirty := now.AddDate(-30, 0, 0)
	if got := AgeInYears(thirty); got != 30 {
		t.Errorf("AgeInYears(30y ago) = %d, want 30", got)
	}
	almost := now.AddDate(-30, 0, 1) // birthday is tomorrow
	if got := AgeInYears(almost); got != 29 {
		t.Errorf("AgeInYears(birthday tomorrow) = %d, want 29", got)
	}
	if got := AgeInYears(now); got != 0 {
		t.Errorf("AgeInYears(today) = %d, want 0", got)
	}
}

func TestFormatLog(t *testing.T) {
	ts := time.Date(2026, 11, 6, 9, 30, 5, 0, time.UTC)
	if got, want := FormatLog(ts), "2026-11-06 09:30:05"; got != want {
		t.Errorf("FormatLog = %q, want %q", got, want)
	}
}

func TestParseDate(t *testing.T) {
	got, err := ParseDate("2026-11-06")
	if err != nil {
		t.Fatalf("ParseDate: %v", err)
	}
	if got.Year() != 2026 || got.Month() != 11 || got.Day() != 6 {
		t.Errorf("ParseDate = %v, want 2026-11-06", got)
	}
	if _, err := ParseDate("06/11/2026"); err == nil {
		t.Error("ParseDate of wrong layout should return an error")
	}
	if _, err := ParseDate(""); err == nil {
		t.Error("ParseDate of empty string should return an error")
	}
	if _, err := ParseDate("2026-13-01"); err == nil {
		t.Error("ParseDate of month 13 should return an error")
	}
}

func TestBusinessDays(t *testing.T) {
	mon := time.Date(2026, 11, 2, 0, 0, 0, 0, time.UTC)
	nextMon := time.Date(2026, 11, 9, 0, 0, 0, 0, time.UTC)
	if got := BusinessDays(mon, nextMon); got != 5 {
		t.Errorf("Mon->Mon BusinessDays = %d, want 5", got)
	}
	sat := time.Date(2026, 11, 7, 0, 0, 0, 0, time.UTC)
	if got := BusinessDays(sat, nextMon); got != 0 {
		t.Errorf("Sat->Mon BusinessDays = %d, want 0", got)
	}
	if got := BusinessDays(nextMon, mon); got != 0 {
		t.Errorf("reversed BusinessDays = %d, want 0", got)
	}
	if got := BusinessDays(mon, mon); got != 0 {
		t.Errorf("same-day BusinessDays = %d, want 0", got)
	}
}
