package piscine

// Count counts overlapping occurrences of substr in s.
func Count(s, substr string) int {
	if substr == "" {
		return 0
	}
	count := 0
	for i := 0; i+len(substr) <= len(s); i++ {
		if s[i:i+len(substr)] == substr {
			count++
		}
	}
	return count
}
