package main

import (
	"fmt"
)

func main() {
	a := ListNode{Val: 1, Next: nil}
	b := ListNode{Val: 2, Next: &a}
	c := ListNode{Val: 2, Next: &b}
	d := ListNode{Val: 1, Next: &c}
	fmt.Println(isPalindrome(&d))
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

func abs(x int) int {
	if 0 <= x {
		return x
	}
	return -x
}

func getDepth(node *TreeNode) int {
	if nil == node {
		return 0
	}
	return max(getDepth(node.Left)+1, getDepth(node.Right)+1)
}

func reverse(cur *ListNode, next *ListNode) {
	if nil == next {
		return
	}
	reverse(next, next.Next)
	next.Next = cur
}

//TreeNode used for tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//ListNode used for node
type ListNode struct {
	Val  int
	Next *ListNode
}
