package main

import "fmt"

func main() {
	fmt.Println(minFallingPathSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}

func minFallingPathSum(A [][]int) int {
	length := len(A)
	for i := 1; i < length; i++ {
		for j := 0; j < length; j++ {
			if j > 0 && j < length-1 {
				A[i][j] += min3(A[i-1][j-1], A[i-1][j], A[i-1][j+1])
			} else if j > 0 {
				A[i][j] += min2(A[i-1][j-1], A[i-1][j])
			} else {
				A[i][j] += min2(A[i-1][j], A[i-1][j+1])
			}
		}
	}

	min := A[length-1][0]
	for i := 1; i < length; i++ {
		if A[length-1][i] < min {
			min = A[length-1][i]
		}
	}
	return min
}

func min3(x, y, z int) int {
	return min2(min2(x, y), z)
}

func min2(x, y int) int {
	if x < y {
		return x
	}
	return y
}
