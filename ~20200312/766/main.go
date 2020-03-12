package main

import "fmt"

func main() {
	fmt.Println(isToeplitzMatrix([][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}))
	fmt.Println(isToeplitzMatrix([][]int{{1, 2}, {2, 2}}))
}

func isToeplitzMatrix(matrix [][]int) bool {

	for i := 0; i < len(matrix[0]); i++ {

		current := matrix[0][i]
		r := 1
		c := i + 1

		for r < len(matrix) && c < len(matrix[0]) {
			if matrix[r][c] != current {
				return false
			}
			r++
			c++
		}

	}

	for i := 1; i < len(matrix); i++ {
		//i,0
		r := i + 1
		c := 1
		current := matrix[i][0]

		for r < len(matrix) && c < len(matrix[i]) {
			if matrix[r][c] != current {
				return false
			}
			r++
			c++
		}

	}

	return true
}
