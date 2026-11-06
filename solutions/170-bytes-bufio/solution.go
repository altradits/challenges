package piscine

import (
	"bufio"
	"bytes"
	"strings"
)

func CountLines(s string) int {
	n := 0
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		n++
	}
	return n
}

func WordFreq(s string) map[string]int {
	freq := make(map[string]int)
	for _, w := range bytes.Fields([]byte(s)) {
		freq[string(w)]++
	}
	return freq
}

func JoinLines(lines []string) string {
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	for i, l := range lines {
		if i > 0 {
			w.WriteByte('\n')
		}
		w.WriteString(l)
	}
	w.Flush()
	return buf.String()
}

func ReverseWords(s string) string {
	words := bytes.Fields([]byte(s))
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return string(bytes.Join(words, []byte(" ")))
}
