package piscine

import (
	"strings"
)

func CleanList(inputs []string) []string {
	var result []string
	for _, s := range inputs {
		trimmed := strings.TrimSpace(s)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}
