package main

import "fmt"

func main() {
	fmt.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}}))
}

func transpose(A [][]int) [][]int {

	m := len(A)
	n := len(A[0])
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			result[i][j] = A[j][i]
		}
	}

	return result
}
