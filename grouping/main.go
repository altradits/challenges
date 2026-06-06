package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	pattern := os.Args[1]
	text := os.Args[2]

	// Validate pattern format: must be (something)
	if len(pattern) < 2 || pattern[0] != '(' || pattern[len(pattern)-1] != ')' {
		return
	}

	// Extract inner pattern
	inner := pattern[1 : len(pattern)-1]
	if inner == "" {
		return
	}

	// Split by | for alternation
	parts := strings.Split(inner, "|")
	if len(parts) == 0 {
		return
	}

	// Build regex pattern for each alternative
	var patterns []*regexp.Regexp
	for _, part := range parts {
		if part == "" {
			continue
		}
		// Escape special regex chars except we want literal matching
		re, err := regexp.Compile(part)
		if err != nil {
			return
		}
		patterns = append(patterns, re)
	}

	if len(patterns) == 0 {
		return
	}

	// Split text into words (split by spaces)
	words := strings.Fields(text)
	if len(words) == 0 {
		return
	}

	// Find matching words
	count := 0
	for _, word := range words {
		for _, re := range patterns {
			if re.MatchString(word) {
				count++
				fmt.Printf("%d: %s\n", count, word)
				break
			}
		}
	}
}
