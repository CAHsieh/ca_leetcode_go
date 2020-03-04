package main

import (
	"fmt"
)

func main() {
	fmt.Println(distributeCandies([]int{1, 1, 2, 2, 3, 3}))
	fmt.Println(distributeCandies([]int{1, 1, 2, 3}))
}

func distributeCandies(candies []int) int {

	m := make(map[int]int)
	for _, candy := range candies {
		m[candy] = m[candy] + 1
	}

	return min(len(m), len(candies)/2)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
