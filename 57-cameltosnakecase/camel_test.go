package piscine

import "testing"

func TestCamelToSnakeCase(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"upper camel", "HelloWorld", "Hello_World"},
		{"lower camel", "helloWorld", "hello_World"},
		{"basic", "camelCase", "camel_Case"},
		{"consecutive caps invalid", "CAMELtoSnackCASE", "CAMELtoSnackCASE"},
		{"multi", "camelToSnakeCase", "camel_To_Snake_Case"},
		{"digit invalid", "hey2", "hey2"},
		{"empty", "", ""},
		{"single lower", "a", "a"},
		{"single upper", "A", "A"},
		{"space invalid", "a b", "a b"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := CamelToSnakeCase(tc.input); got != tc.want {
				t.Errorf("CamelToSnakeCase(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
