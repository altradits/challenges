package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findPairs(arr []int, target int) [][]int {
	var pairs [][]int
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				pairs = append(pairs, []int{i, j})
			}
		}
	}
	return pairs
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid input.")
		return
	}
	arrStr := os.Args[1]
	targetStr := os.Args[2]
	if !strings.HasPrefix(arrStr, "[") || !strings.HasSuffix(arrStr, "]") {
		fmt.Println("Invalid input.")
		return
	}
	arrStr = strings.Trim(arrStr, "[]")
	if arrStr == "" {
		fmt.Println("Invalid input.")
		return
	}
	parts := strings.Split(arrStr, ", ")
	var arr []int
	for _, p := range parts {
		n, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			fmt.Printf("Invalid number: %s\n", p)
			return
		}
		arr = append(arr, n)
	}
	target, err := strconv.Atoi(targetStr)
	if err != nil {
		fmt.Println("Invalid target sum.")
		return
	}
	pairs := findPairs(arr, target)
	if len(pairs) == 0 {
		fmt.Println("No pairs found.")
	} else {
		fmt.Printf("Pairs with sum %d: %v\n", target, pairs)
	}
}
