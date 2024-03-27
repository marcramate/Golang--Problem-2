package main

import (
	"fmt"
	"sort"
)

func maxProChickens(n, k int, poschickens []int) int {
	sort.Ints(poschickens)

	left, right := 0, 0
	maxChickens := 0

	for right < n {
		end := poschickens[right] + k
		InRange := 0
		
		for left < n && poschickens[left] < end {
			InRange++
			left++
		}

		maxChickens = max(maxChickens, InRange)
		right++
	}

	return maxChickens
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	n1, k1 := 5, 5
	poschickens1 := []int{2, 5, 10, 12, 15}

	maxChickens1 := maxProChickens(n1, k1, poschickens1)
	fmt.Printf("Maximum chickens : %d\n", maxChickens1)

	n2, k2 := 6, 10
	poschickens2 := []int{1, 11, 30, 34, 35, 37}

	maxChickens2 := maxProChickens(n2, k2, poschickens2)
	fmt.Printf("Maximum chickens : %d\n", maxChickens2)
}
