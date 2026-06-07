package main

import "fmt"

func RemoveSpaces(s string) string {
	result := ""
	for _, c := range s {
		if c != ' ' {
			result += string(c)
		}
	}
	return result
}

func main() {
	fmt.Println(RemoveSpaces("Hello World"))
	fmt.Println(RemoveSpaces("Go is fun!"))
	fmt.Println(RemoveSpaces(""))
	fmt.Println(RemoveSpaces("NoSpacesHere"))
}
