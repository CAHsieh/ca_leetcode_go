package main

import (
	"fmt"
)

func main() {
	// fmt.Println(numDecodings("12"))
	// fmt.Println(numDecodings("226"))
	fmt.Println(numDecodings("110"))
}

func isValid(a, b byte) bool {
	return ((a == '1' && b <= '9') || (a == '2' && b <= '6'))
}

func numDecodings(s string) int {

	n := len(s)
	d := make([]int, n+1)

	d[0] = 1
	if s[0] != '0' {
		d[1] = 1
	} else {
		d[1] = 0
	}

	for i := 2; i <= n; i++ {
		if s[i-1] != '0' {
			d[i] += d[i-1]
		}
		if isValid(s[i-2], s[i-1]) {
			d[i] += d[i-2]
		}
	}

	return d[n]
}
