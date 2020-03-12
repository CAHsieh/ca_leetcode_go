package main

import (
	"fmt"
)

func main() {
	fmt.Println(diStringMatch("IDID"))
	fmt.Println(diStringMatch("III"))
	fmt.Println(diStringMatch("DDI"))
}

func diStringMatch(S string) []int {
	var arr []int
	e := len(S)
	s := 0

	for _, c := range S {
		if c == 'I' {
			arr = append(arr, s)
			s++
		} else if c == 'D' {
			arr = append(arr, e)
			e--
		}
	}
	arr = append(arr, s)

	return arr
}
