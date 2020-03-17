package main

import "fmt"

func main() {
	fmt.Println(shortestToChar("loveleetcode", 'e'))
}

func shortestToChar(S string, C byte) []int {
	count := len(S)
	result := make([]int, count)
	queue := []int{}
	for i := range S {
		result[i] = -1
		if S[i] == C {
			queue = append(queue, i)
			result[i] = 0
		}
	}

	for len(queue) > 0 {
		newQueue := []int{}

		for _, i := range queue {
			if i > 0 && (result[i-1] == -1 || result[i-1] > result[i]+1) {
				result[i-1] = result[i] + 1
				newQueue = append(newQueue, i-1)
			}
			if i < count-1 && (result[i+1] == -1 || result[i+1] > result[i]+1) {
				result[i+1] = result[i] + 1
				newQueue = append(newQueue, i+1)
			}
		}

		queue = newQueue
	}
	return result
}
