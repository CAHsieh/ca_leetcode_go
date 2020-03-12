package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(countPrimeSetBits(6, 10))
	fmt.Println(countPrimeSetBits(10, 15))
}

func countPrimeSetBits(L int, R int) int {
	count := 0

	for ; L <= R; L++ {
		num := bits.OnesCount(uint(L))
		if isPrime(num) {
			count++
		}
	}

	return count
}

// func isPrime(num int) bool {
// 	if num < 2 {
// 		return false
// 	}
// 	mid := int(math.Sqrt(float64(num))) + 1
// 	for i := 2; i < mid; i++ {
// 		if num%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

//faster:
//因應題目規定的範圍有可能的奇數值
func isPrime(x int) bool {
	return (x == 2 || x == 3 || x == 5 || x == 7 ||
		x == 11 || x == 13 || x == 17 || x == 19)
}
