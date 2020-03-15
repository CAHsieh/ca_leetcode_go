package main

import "fmt"

func main() {
	fmt.Println(minCostToMoveChips([]int{1, 2, 3}))
	fmt.Println(minCostToMoveChips([]int{2, 2, 2, 3, 3}))
}

func minCostToMoveChips(chips []int) int {
	odd, even := 0, 0
	for _, v := range chips {
		if v%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	if even < odd {
		return even
	}
	return odd
}
