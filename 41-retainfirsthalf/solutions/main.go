package main

import "fmt"

func RetainFirstHalf(s string) string {
	if len(s) <= 1 {
		return s
	}
	half := len(s) / 2
	return s[:half]
}

func main() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(RetainFirstHalf("A"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("Hello World"))
}
