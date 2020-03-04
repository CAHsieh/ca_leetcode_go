package main

import "fmt"

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
	fmt.Println(uniqueOccurrences([]int{1, 2}))
	fmt.Println(uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}))
}

func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)

	for _, val := range arr {
		count, ok := m[val]
		if ok == false {
			m[val] = 1
		} else {
			m[val] = count + 1
		}
	}

	counts := make([]int, 2000)

	for _, v := range m {
		counts[v]++
		if counts[v] > 1 {
			return false
		}
	}

	return true
}
