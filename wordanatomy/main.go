package main

import "fmt"

type WordStats struct {
	WordCount      int
	VowelCount     int
	ConsonantCount int
	LongestWord    string
	ShortestWord   string
}

func WordAnatomy(s string) WordStats {
	words := fields(s)
	stats := WordStats{
		WordCount: len(words),
	}

	if len(words) == 0 {
		return stats
	}

	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true}

	longest := words[0]
	shortest := words[0]

	for _, word := range words {
		if len(word) > len(longest) {
			longest = word
		}
		if len(word) < len(shortest) {
			shortest = word
		}

		for i := 0; i < len(word); i++ {
			c := word[i]
			if isLetter(c) {
				if vowels[c] {
					stats.VowelCount++
				} else {
					stats.ConsonantCount++
				}
			}
		}
	}

	stats.LongestWord = longest
	stats.ShortestWord = shortest

	return stats
}

func fields(s string) []string {
	var words []string
	var current []byte
	inWord := false

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == ' ' || c == '\t' || c == '\n' || c == '\r' {
			if inWord {
				words = append(words, string(current))
				current = nil
				inWord = false
			}
		} else {
			current = append(current, c)
			inWord = true
		}
	}

	if inWord {
		words = append(words, string(current))
	}

	return words
}

func isLetter(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
}

func main() {
	stats := WordAnatomy("Hello World")
	fmt.Printf("Words: %d, Vowels: %d, Consonants: %d, Longest: %s, Shortest: %s\n",
		stats.WordCount, stats.VowelCount, stats.ConsonantCount, stats.LongestWord, stats.ShortestWord)

	stats = WordAnatomy("The quick brown fox")
	fmt.Printf("Words: %d, Vowels: %d, Consonants: %d, Longest: %s, Shortest: %s\n",
		stats.WordCount, stats.VowelCount, stats.ConsonantCount, stats.LongestWord, stats.ShortestWord)

	stats = WordAnatomy("")
	fmt.Printf("Words: %d, Vowels: %d, Consonants: %d, Longest: %s, Shortest: %s\n",
		stats.WordCount, stats.VowelCount, stats.ConsonantCount, stats.LongestWord, stats.ShortestWord)
}
