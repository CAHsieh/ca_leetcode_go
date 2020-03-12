package main

import "fmt"

func main() {
	fmt.Println(findTheLongestSubstring("eleetminicoworoep"))
	// fmt.Println(findTheLongestSubstring("leetcodeisgreat"))
	// fmt.Println(findTheLongestSubstring("bcbcbc"))
}

func findTheLongestSubstring(s string) int {
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	result := 0
	status := 0
	m := make(map[int]int)

	for i, c := range s {
		
		if contains(vowels, c) {
			status ^= (1 << (c - 'a'))
		}

		if status == 0 {
			result = i + 1
		} else {
			lastI, ok := m[status]
			if !ok {
				m[status] = i
			} else {
				result = max(result, i-lastI)
			}
		}
	}

	return result
}

func contains(vowels []rune, c rune) bool {
	for _, v := range vowels {
		if v == c {
			return true
		}
	}
	return false
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
