package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2))
	fmt.Println(combine(5, 4))
}

var result [][]int

func combine(n int, k int) [][]int {
	result = [][]int{}
	for i := 0; i <= n-k; i++ {
		path := []int{i + 1}
		dfs(path, 1, n, k)
	}
	return result
}

func dfs(path []int, depth, limit, goal int) {
	if depth == goal {
		tmp := make([]int, depth)
		copy(tmp, path)
		result = append(result, tmp)
	}
	for i := path[depth-1] + 1; i <= limit; i++ {
		tmp := append(path, i)
		dfs(tmp, depth+1, limit, goal)
	}
}
