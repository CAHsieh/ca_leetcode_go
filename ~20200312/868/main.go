package main

import "fmt"

func main() {
	fmt.Println(binaryGap(22))
	fmt.Println(binaryGap(5))
	fmt.Println(binaryGap(6))
	fmt.Println(binaryGap(8))
}

func binaryGap(N int) int {
	d := 0
	past := -1
	index := 0
	for N > 0 {
		if N%2 == 1 {
			if past != -1 {
				curD := index - past
				if curD > d {
					d = curD
				}
			}
			past = index
		}
		index++
		N /= 2
	}
	return d
}
