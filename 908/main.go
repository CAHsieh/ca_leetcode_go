package main

import "fmt"

func main() {
	fmt.Println(smallestRangeI([]int{1}, 0))
	fmt.Println(smallestRangeI([]int{0, 10}, 2))
	fmt.Println(smallestRangeI([]int{1, 3, 6}, 3))
}

func smallestRangeI(A []int, K int) int {
	min, max := A[0], A[0]
	for i := 1; i < len(A); i++ {
		if min > A[i] {
			min = A[i]
		}
		if max < A[i] {
			max = A[i]
		}
	}
	if max-min > 2*K {
		return max - min - 2*K
	}
	return 0
}
