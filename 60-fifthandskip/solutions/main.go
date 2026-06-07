package main

import (
	"fmt"
)

func FifthAndSkip(str string) string {
	if len(str) == 0 {
		return "\n"
	}

	// Remove spaces
	noSpaces := ""
	for _, c := range str {
		if c != ' ' {
			noSpaces += string(c)
		}
	}

	if len(noSpaces) < 5 {
		return "Invalid Input\n"
	}

	result := ""
	for i := 0; i < len(noSpaces); {
		// Save 5 characters
		end := i + 5
		if end > len(noSpaces) {
			end = len(noSpaces)
		}
		if i > 0 {
			result += " "
		}
		result += noSpaces[i:end]
		i += 5
		// Skip 1 character
		if i < len(noSpaces) {
			i++
		}
	}

	return result + "\n"
}

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
}
