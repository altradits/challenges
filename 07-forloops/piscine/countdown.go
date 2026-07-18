package piscine

import (
	"fmt"
)

func Countdown(n int) {
	for i := n; i >= 0; i-- {
		fmt.Println(i)
	}
}
