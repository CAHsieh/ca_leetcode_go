package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(calPoints([]string{"5", "2", "C", "D", "+"}))
	fmt.Println(calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"}))
}

func calPoints(ops []string) int {
	sum := 0
	count := 0
	stack := []int{}

	for _, v := range ops {

		switch v {
		case "+":
			if len(stack) > 1 {
				p := stack[count-1] + stack[count-2]
				sum += p
				stack = append(stack, p)
				count++
			}
		case "D":
			if len(stack) > 0 {
				p := stack[count-1] * 2
				sum += p
				stack = append(stack, p)
				count++
			}
		case "C":
			if len(stack) > 0 {
				sum -= stack[count-1]
				stack = stack[:count-1]
				count--
			}
		default:
			p, _ := strconv.Atoi(v)
			sum += p
			stack = append(stack, p)
			count++
		}
	}

	return sum
}
