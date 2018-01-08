package main

func rotate(matrix [][]int) {
	// sol 1 , too stupid
	// length := len(matrix)
	// if length%2 == 0 {
	// 	rotateEven(matrix, length)
	// } else {
	// 	rotateOdd(matrix, length)
	// }

	//sol 2 common

	len := len(matrix)

	for i := 0; i < len/2; i++ {
		for j := 0; j < len; j++ {
			swap(&matrix[i][j], &matrix[len-i-1][j])
		}
	}

	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			swap(&matrix[i][j], &matrix[j][i])
		}
	}
}

func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func rotateOdd(matrix [][]int, len int) {

	offset := len / 2

	newMatrix := make([][]int, len)
	for i := range newMatrix {
		newMatrix[i] = make([]int, len)
	}

	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {

			transI := -(j - offset)
			transJ := (i - offset)

			newMatrix[i][j] = matrix[transI+offset][transJ+offset]
		}
	}

	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {

			matrix[i][j] = newMatrix[i][j]
		}
	}
}

func rotateEven(matrix [][]int, len int) {

	offset := len / 2

	newMatrix := make([][]int, len)
	for i := range newMatrix {
		newMatrix[i] = make([]int, len)
	}

	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {
			offseti := offset
			offsetj := offset

			if i >= offset {
				offseti--
			}

			if j >= offset {
				offsetj--
			}

			transI := -(j - offsetj) + offsetj
			transJ := (i - offseti) + offseti

			if transI > offset {
				transI--
			} else {
				transI++
			}

			newMatrix[i][j] = matrix[transI][transJ]
		}
	}

	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {

			matrix[i][j] = newMatrix[i][j]
		}
	}
}
