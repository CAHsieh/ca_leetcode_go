package main

import (
	"fmt"
)

func main() {
	// a := ListNode{Val: 3, Next: nil}
	// b := ListNode{Val: 2, Next: &a}
	// c := ListNode{Val: 1, Next: &b}
	// d := ListNode{Val: 1, Next: &c}
	// ta := TreeNode{Val: 6, Left: nil, Right: nil}
	// tb := TreeNode{Val: 4, Left: nil, Right: nil}
	// tc := TreeNode{Val: 3, Left: nil, Right: nil}
	// td := TreeNode{Val: 5, Left: nil, Right: &ta}
	// te := TreeNode{Val: 2, Left: &tc, Right: &tb}
	// tf := TreeNode{Val: 1, Left: &te, Right: &td}

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

	// fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	// fmt.Println(canPartition([]int{1, 5, 11, 5}))
	// fmt.Println(canPartition([]int{1, 1}))
	// fmt.Println(canPartition([]int{1, 2, 5}))
	// fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	// flatten(&tf)
	// for tf.Right != nil {
	// 	fmt.Println(tf.Val)
	// 	tf = *tf.Right
	// }
	// fmt.Println(tf.Val)
	// fmt.Println("")

	// fmt.Println(numIslands([][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'},
	// 	{'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}))
	// fmt.Println(numIslands([][]byte{
	// 	{'1', '0', '1', '1', '1'},
	// 	{'1', '0', '1', '0', '1'},
	// 	{'1', '1', '1', '0', '1'}}))

	// fmt.Println(letterCombinations("23"))

	// fmt.Println(removeNthFromEnd(&c, 3))

	// input := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	// fmt.Println(minPathSum(input))

	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 1))
	fmt.Println(searchInsert(nums, 3))
	fmt.Println(searchInsert(nums, 5))
	fmt.Println(searchInsert(nums, 6))

	fmt.Println(searchInsert(nums, 0))
	fmt.Println(searchInsert(nums, 2))
	fmt.Println(searchInsert(nums, 7))
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

//Label used to record leader
type Label struct {
	root  *Label
	label int
	rank  int
}
