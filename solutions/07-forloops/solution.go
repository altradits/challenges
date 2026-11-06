package piscine

import "fmt"

func CountDown(n int) {
	for i := n; i >= 0; i-- {
		fmt.Println(i)
	}
}
