package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
}

func findWords(words []string) []string {
	row := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	result := []string{}
	for _, w := range words {
		s := strings.Split(strings.ToLower(w), "")
		row0 := 0
		row1 := 0
		row2 := 0
		for _, a := range s {
			if strings.Contains(row[0], a) {
				row0 = 1
			} else if strings.Contains(row[1], a) {
				row1 = 1
			} else {
				row2 = 1
			}
		}
		if row0+row1+row2 == 1 {
			result = append(result, w)
		}
	}
	return result
}
