package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(kthSmallest([][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 3))
	fmt.Println(kthSmallest([][]int{{-5}}, 1))
}

func kthSmallest(matrix [][]int, k int) int {

	oneD := []int{}

	for _, v := range matrix {
		oneD = append(oneD, v...)
	}
	sort.Ints(oneD)
	return oneD[k-1]
}

//binary search will beat 100%
