package main

import (
	"fmt"
)

func main() {
	// a := ListNode{Val: 1, Next: nil}
	// b := ListNode{Val: 2, Next: &a}
	// c := ListNode{Val: 2, Next: &b}
	// d := ListNode{Val: 1, Next: &c}
	//ta := TreeNode{Val: 2, Left: nil, Right: nil}
	//tb := TreeNode{Val: 3, Left: nil, Right: nil}
	//tc := TreeNode{Val: 1, Left: &ta, Right: &tb}

	// fmt.Println(isPalindrome(&d))
	// fmt.Println(countBits(5))
	// fmt.Println(reconstructQueue([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}))
	// fmt.Println(countSubstrings("aaa"))
	// fmt.Println(findCircleNum([][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}}))
	// fmt.Println(topKFrequent([]int{3, 0, 1, 0}, 1))
	// fmt.Println(generateParenthesis(3))
	// input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// rotate(input)
	// fmt.Println(input)
	// input = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	// rotate(input)
	// fmt.Println(input)
	// fmt.Println(isSameTree(&tc, &tc))

	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
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
