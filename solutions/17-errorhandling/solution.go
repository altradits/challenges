package piscine

import (
	"errors"
	"fmt"
	"strconv"
)

var ErrNegativeAge = errors.New("age cannot be negative")

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func ParseAge(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid age %q: %w", s, err)
	}
	if n < 0 {
		return 0, ErrNegativeAge
	}
	return n, nil
}

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Field + ": " + e.Message
}
