package main

import "fmt"

func CleanList(lst []string) []string {
	if len(lst) == 0 {
		return []string{}
	}

	cleaned := make([]string, 0, len(lst)+1)
	hasMilk := false

	for i, item := range lst {
		// Trim spaces manually
		trimmed := trimSpace(item)

		// Capitalize first letter, rest lowercase
		capitalized := capitalize(trimmed)

		// Check for milk (case-insensitive)
		if equalFold(trimmed, "milk") {
			hasMilk = true
		}

		// Add index prefix
		cleaned = append(cleaned, fmt.Sprintf("%d/ %s", i+1, capitalized))
	}

	// Add milk if missing
	if !hasMilk {
		cleaned = append(cleaned, fmt.Sprintf("%d/ Milk", len(cleaned)+1))
	}

	return cleaned
}

func trimSpace(s string) string {
	start := 0
	end := len(s)
	for start < end && (s[start] == ' ' || s[start] == '\t' || s[start] == '\n' || s[start] == '\r') {
		start++
	}
	for end > start && (s[end-1] == ' ' || s[end-1] == '\t' || s[end-1] == '\n' || s[end-1] == '\r') {
		end--
	}
	return s[start:end]
}

func capitalize(s string) string {
	if len(s) == 0 {
		return ""
	}
	result := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if i == 0 {
			if c >= 'a' && c <= 'z' {
				result[i] = c - 'a' + 'A'
			} else {
				result[i] = c
			}
		} else {
			if c >= 'A' && c <= 'Z' {
				result[i] = c - 'A' + 'a'
			} else {
				result[i] = c
			}
		}
	}
	return string(result)
}

func equalFold(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		c1 := s1[i]
		c2 := s2[i]
		// Convert both to lowercase for comparison
		if c1 >= 'A' && c1 <= 'Z' {
			c1 = c1 - 'A' + 'a'
		}
		if c2 >= 'A' && c2 <= 'Z' {
			c2 = c2 - 'A' + 'a'
		}
		if c1 != c2 {
			return false
		}
	}
	return true
}

func main() {
	// Test case from README
	shoppingList := []string{"tomatoes", "pastas", "milk", "salt"}
	fmt.Println("Standard:", CleanList(shoppingList))

	// Test without milk
	shoppingList2 := []string{"tomatoes", "pastas", "salt"}
	fmt.Println("Without milk:", CleanList(shoppingList2))

	// Test empty list
	fmt.Println("Empty list:", CleanList([]string{}))

	// Test with spaces and uppercase MILK
	shoppingList3 := []string{"  tomatoes  ", "  MILK  ", "salt"}
	fmt.Println("With spaces and MILK:", CleanList(shoppingList3))
}
