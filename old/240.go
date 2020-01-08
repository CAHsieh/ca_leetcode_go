package main

func searchMatrix(matrix [][]int, target int) bool {
	if nil == matrix || len(matrix) < 1 || len(matrix[0]) < 1 {
		return false
	}

	col := len(matrix[0]) - 1
	row := 0

	for col >= 0 && row < len(matrix) {

		if target == matrix[row][col] {
			return true
		} else if target < matrix[row][col] {
			col--
		} else if target > matrix[row][col] {
			row++
		}
	}

	return false
}
