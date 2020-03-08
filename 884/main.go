package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(uncommonFromSentences("this apple is sweet", "this apple is sour"))
	fmt.Println(uncommonFromSentences("apple apple", "banana"))
}

func uncommonFromSentences(A string, B string) []string {
	m := make(map[string]int)

	arr := strings.Split(A, " ")
	arr = append(arr, strings.Split(B, " ")...)

	for _, s := range arr {
		m[s]++
	}

	result := []string{}
	for k, v := range m {
		if v == 1 {
			result = append(result, k)
		}
	}

	return result
}
