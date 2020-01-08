package main

func islandPerimeter(grid [][]int) int {

	total := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if 0 == grid[i][j] {
				continue
			}

			num := 4
			if i > 0 && 1 == grid[i-1][j] {
				num--
			}
			if i < len(grid)-1 && 1 == grid[i+1][j] {
				num--
			}
			if j > 0 && 1 == grid[i][j-1] {
				num--
			}
			if j < len(grid[i])-1 && 1 == grid[i][j+1] {
				num--
			}

			total += num
		}
	}

	return total
}
