package main

import "fmt"

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	if !isValidCamelCase(s) {
		return s
	}

	var result []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' && i > 0 {
			result = append(result, '_')
		}
		result = append(result, c)
	}
	return string(result)
}

func isValidCamelCase(s string) bool {
	if len(s) == 0 {
		return false
	}

	// Check for numbers or punctuation (only letters allowed)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !((c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')) {
			return false
		}
	}

	// Check if ends with uppercase
	last := s[len(s)-1]
	if last >= 'A' && last <= 'Z' {
		return false
	}

	// Check for consecutive uppercase letters
	prevUpper := false
	for i := 0; i < len(s); i++ {
		c := s[i]
		isUpper := c >= 'A' && c <= 'Z'
		if isUpper && prevUpper {
			return false
		}
		prevUpper = isUpper
	}

	// Must have at least one uppercase letter to be camelCase
	hasUpper := false
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			hasUpper = true
			break
		}
	}
	if !hasUpper {
		return false
	}

	return true
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
