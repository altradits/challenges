package main

import "fmt"

// CamelToSnakeCase converts a camelCase string to snake_case.
// If the string is not valid camelCase, return it unchanged.
func CamelToSnakeCase(s string) string {
	// TODO: Implement this function
	// Hint: First validate the string is camelCase.
	// Then iterate and add '_' before uppercase letters (except first).
	// Convert all to lowercase.
	return ""
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
