## wordanatomy

### Instructions

Write a function `WordAnatomy` that takes a string and returns a struct containing analysis of the words in the string.

The struct should contain:
- `WordCount`: number of words (delimited by spaces)
- `VowelCount`: total number of vowels (a, e, i, o, u, case-insensitive)
- `ConsonantCount`: total number of consonants (letters that are not vowels)
- `LongestWord`: the longest word (first one if tie)
- `ShortestWord`: the shortest word (first one if tie)

### Expected function

```go
type WordStats struct {
	WordCount      int
	VowelCount     int
	ConsonantCount int
	LongestWord    string
	ShortestWord   string
}

func WordAnatomy(s string) WordStats {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
	"fmt"
)

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
```

And its output:

```console
$ go run .
Words: 2, Vowels: 3, Consonants: 7, Longest: Hello, Shortest: Hello
Words: 4, Vowels: 5, Consonants: 11, Longest: quick, Shortest: The
Words: 0, Vowels: 0, Consonants: 0, Longest: , Shortest: 