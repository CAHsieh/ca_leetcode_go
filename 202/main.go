package main

import "fmt"

func main() {
	fmt.Println(isHappy(7))
}

func isHappy(n int) bool {

	past := make(map[int]bool)

	for a := n; a != 1; {
		sum := 0
		for ; a > 0; a /= 10 {
			k := a % 10
			sum += k * k
		}
		a = sum
		_, ok := past[a]
		if ok {
			return false
		}
		past[a] = true
	}
	return true
}
