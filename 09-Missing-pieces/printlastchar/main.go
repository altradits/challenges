// Challenge: Print the last character of any string using len() — without hardcoding the index.
// text := "stano"
// print the last character using len(text)
// Expected output: o
// The twist: len() gives you the count, but indexes start at 0 — so what index is the last character actually at?

package main

import (
	"fmt"
)

func main() {
	text := "stano"

	fmt.Printf("The last character is: %c\n", text[4])

	// Second Attempt

	// index of the last element
	// len(text)-1

	// the byte
	// text[index]
	
	// string(rune)

	last := string(text[len(text)-1])

	fmt.Println(last)
}
