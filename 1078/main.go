package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findOcurrences("alice is a good girl she is a good student", "a", "good"))
	fmt.Println(findOcurrences("we will we will rock you", "we", "will"))
}

func findOcurrences(text string, first string, second string) []string {
	words := strings.Split(text, " ")
	result := []string{}
	count := len(words)
	for i := 0; i < count-1; i++ {
		if words[i] == first && words[i+1] == second && i+2 < count {
			result = append(result, words[i+2])
			i++
		}
	}
	return result
}
