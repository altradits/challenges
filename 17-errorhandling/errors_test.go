package piscine

import (
	"errors"
	"strconv"
	"testing"
)

func TestDivide(t *testing.T) {
	got, err := Divide(10, 2)
	if err != nil || got != 5 {
		t.Errorf("Divide(10, 2) = (%v, %v), want (5, nil)", got, err)
	}
	if _, err := Divide(1, 0); err == nil {
		t.Error("Divide(1, 0) should return a non-nil error")
	}
}

func TestParseAge(t *testing.T) {
	got, err := ParseAge("30")
	if err != nil || got != 30 {
		t.Errorf("ParseAge(\"30\") = (%d, %v), want (30, nil)", got, err)
	}
	if got, err := ParseAge("0"); err != nil || got != 0 {
		t.Errorf("ParseAge(\"0\") = (%d, %v), want (0, nil)", got, err)
	}
	_, err = ParseAge("-5")
	if !errors.Is(err, ErrNegativeAge) {
		t.Errorf("ParseAge(\"-5\") error = %v, want errors.Is ErrNegativeAge", err)
	}
	_, err = ParseAge("abc")
	if err == nil {
		t.Fatal("ParseAge(\"abc\") should return a non-nil error")
	}
	var numErr *strconv.NumError
	if !errors.As(err, &numErr) {
		t.Errorf("ParseAge(\"abc\") error = %v, want a wrapped strconv error", err)
	}
}

func TestValidationError(t *testing.T) {
	err := &ValidationError{Field: "age", Message: "must be >= 0"}
	if got, want := err.Error(), "age: must be >= 0"; got != want {
		t.Errorf("Error() = %q, want %q", got, want)
	}
	var _ error = err
}
