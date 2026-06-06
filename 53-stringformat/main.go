package main

import (
	"fmt"
)

func Format(template string, args ...interface{}) string {
	result := template
	argIndex := 0
	for i := 0; i < len(result)-1 && argIndex < len(args); i++ {
		if result[i] == '%' {
			switch result[i+1] {
			case 's':
				result = result[:i] + fmt.Sprint(args[argIndex]) + result[i+2:]
				argIndex++
				i = -1 // restart scan
			case 'd':
				result = result[:i] + fmt.Sprint(args[argIndex]) + result[i+2:]
				argIndex++
				i = -1
			case 'f':
				result = result[:i] + fmt.Sprint(args[argIndex]) + result[i+2:]
				argIndex++
				i = -1
			}
		}
	}
	return result
}

func main() {
	fmt.Println(Format("Hello %s", "World"))
	fmt.Println(Format("%d + %d = %d", 1, 2, 3))
	fmt.Println(Format("%s is %d years old", "Alice", 30))
	fmt.Println(Format("No placeholders"))
}
