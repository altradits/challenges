package main

import "fmt"

func WordAnatomy(word string) (string, string, string) {
	prefixes := []string{"un", "re", "pre", "mis", "dis", "over", "under", "anti", "inter", "sub"}
	suffixes := []string{"ing", "ed", "er", "est", "ly", "ness", "ment", "tion", "able", "ful"}

	prefix := ""
	root := word
	suffix := ""

	// Find longest matching prefix
	for _, p := range prefixes {
		if len(word) >= len(p) && word[:len(p)] == p && len(p) > len(prefix) {
			prefix = p
		}
	}

	// Find longest matching suffix
	for _, s := range suffixes {
		if len(word) >= len(s) && word[len(word)-len(s):] == s && len(s) > len(suffix) {
			suffix = s
		}
	}

	// Calculate root word
	if len(prefix) > 0 && len(suffix) > 0 {
		root = word[len(prefix) : len(word)-len(suffix)]
	} else if len(prefix) > 0 {
		root = word[len(prefix):]
	} else if len(suffix) > 0 {
		root = word[:len(word)-len(suffix)]
	}

	return prefix, root, suffix
}

func main() {
	tests := []string{
		"unhappy",
		"runner",
		"redoing",
		"kindness",
		"planet",
		"careful",
		"misunderstanding",
		"greatest",
		"underpaid",
	}

	for _, word := range tests {
		prefix, root, suffix := WordAnatomy(word)

		fmt.Printf(
			"%s -> Prefix:%q Root:%q Suffix:%q\n",
			word,
			prefix,
			root,
			suffix,
		)
	}
}
