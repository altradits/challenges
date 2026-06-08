// len(str) - returns total number of bytes or characters in a string
// str[index] - grabs the character at that specific position starting form 0
// text[0] - Go prints its raw numeric byte instead of the letter, to see the actual letter, use the verb %c in printf

package main

import "fmt"

func main() {
	text := "stano"

	fmt.Printf("The character is: %c\n", text[0])
	fmt.Println(text[0])

}
