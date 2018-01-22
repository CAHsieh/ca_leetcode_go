package main

func numIslands(grid [][]byte) int {
	count := 0

	for i, col := range grid {
		for j, p := range col {
			if p == '1' {
				mark(grid, i, j)
				count++
			}
		}
	}

	return count
}

func mark(grid [][]byte, i int, j int) {
	if grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'

	if i > 0 {
		mark(grid, i-1, j)
	}
	if j > 0 {
		mark(grid, i, j-1)
	}
	if i < len(grid)-1 {
		mark(grid, i+1, j)
	}
	if j < len(grid[0])-1 {
		mark(grid, i, j+1)
	}
}
