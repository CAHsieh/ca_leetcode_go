package main

import "fmt"

func main() {
	fmt.Println(shiftGrid([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1))
}

func shiftGrid(grid [][]int, k int) [][]int {
	y := len(grid)
	x := len(grid[0])
	nGrid := make([][]int, y)
	for i := 0; i < y; i++ {
		nGrid[i] = make([]int, x)
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			ny := (k / x) + i
			if (j + (k % x)) >= x {
				ny++
			}
			nGrid[ny%y][(j+(k%x))%x] = grid[i][j]
		}
	}
	return nGrid
}
