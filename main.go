package main

import (
	"fmt"
)

func main() {
	fmt.Println(climbStairs(1))
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}
