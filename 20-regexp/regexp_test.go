package piscine

import (
	"reflect"
	"testing"
)

func TestValidateEmail(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  bool
	}{
		{"simple", "user@example.com", true},
		{"dots and tags", "first.last+tag@mail.example.co", true},
		{"empty", "", false},
		{"no at", "not-an-email", false},
		{"no domain", "user@", false},
		{"no tld", "user@localhost", false},
		{"spaces", "user @example.com", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ValidateEmail(tc.input); got != tc.want {
				t.Errorf("ValidateEmail(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestExtractNumbers(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  []string
	}{
		{"two groups", "foo123bar456", []string{"123", "456"}},
		{"none", "no digits", []string{}},
		{"empty", "", []string{}},
		{"single", "abc7", []string{"7"}},
		{"adjacent spaces", "1 22 333", []string{"1", "22", "333"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := ExtractNumbers(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ExtractNumbers(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestMaskPhone(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"ten digits", "0712345678", "******5678"},
		{"short", "123", "123"},
		{"four", "1234", "1234"},
		{"empty", "", ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := MaskPhone(tc.input); got != tc.want {
				t.Errorf("MaskPhone(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
