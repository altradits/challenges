package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}

	str := args[0]
	oldChar := args[1]
	newChar := args[2]

	if len(oldChar) != 1 || len(newChar) != 1 {
		return
	}

	if !contains(str, oldChar) {
		fmt.Println(str)
		return
	}

	result := replaceAll(str, oldChar, newChar)
	fmt.Println(result)
}

func contains(s, substr string) bool {
	if len(substr) == 0 {
		return true
	}
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func replaceAll(s, old, new string) string {
	if len(old) == 0 {
		return s
	}
	var result []byte
	for i := 0; i < len(s); {
		if i <= len(s)-len(old) && s[i:i+len(old)] == old {
			result = append(result, new...)
			i += len(old)
		} else {
			result = append(result, s[i])
			i++
		}
	}
	return string(result)
}
