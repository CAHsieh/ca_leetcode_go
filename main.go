package main

import (
	"fmt"
)

func main() {
	a := ListNode{Val: 1, Next: nil}
	b := ListNode{Val: 2, Next: nil}
	fmt.Println(mergeTwoLists(&a, &b).Val)
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
