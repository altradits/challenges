package main

import "fmt"

func Split(s string, sep string) []string {
	if len(sep) == 0 {
		var result []string
		for _, c := range s {
			result = append(result, string(c))
		}
		return result
	}

	var result []string
	start := 0
	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i:i+len(sep)] == sep {
			result = append(result, s[start:i])
			start = i + len(sep)
			i += len(sep) - 1
		}
	}
	return result
}

func Join(arr []string, sep string) string {
	if len(arr) == 0 {
		return ""
	}

	result := arr[0]
	for i := 1; i < len(arr); i++ {
		result += sep + arr[i]
	}
	return result
}

func main() {
	// Test Split
	fmt.Printf("%q\n", Split("a,b,c", ","))
	fmt.Printf("%q\n", Split("Hello World", " "))
	fmt.Printf("%q\n", Split("abc", ""))
	fmt.Printf("%q\n", Split("", ","))

	// Test Join
	fmt.Println(Join([]string{"a", "b", "c"}, ","))
	fmt.Println(Join([]string{"Hello", "World"}, " "))
	fmt.Println(Join([]string{"a", "b", "c"}, ""))
	fmt.Println(Join([]string{""}, ","))
}
