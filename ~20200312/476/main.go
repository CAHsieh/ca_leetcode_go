package main

import (
	"fmt"
)

func main() {
	fmt.Println(findComplement(1))
	fmt.Println(findComplement(2))
	fmt.Println(findComplement(3))
	fmt.Println(findComplement(4))
	fmt.Println(findComplement(5))
	fmt.Println(findComplement(6))
	fmt.Println(findComplement(7))
	fmt.Println(findComplement(8))
}

func findComplement(num int) int {
	n := num
	p := 1
	for n > 0 {
		n /= 2
		p *= 2
	}
	return (num ^ (p - 1))
}
