package main

import "fmt"

func RetainFirstHalf(str string) string {
	if len(str) <= 1 {
		return str
	}
	return str[:len(str)/2]
}

func main() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(RetainFirstHalf("A"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("Hello World"))
}
