package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mrob(root *TreeNode) int {

	res := robSub(root)
	return max(res[0], res[1])
}

func robSub(root *TreeNode) []int {
	res := make([]int, 2)
	if nil == root {
		return res
	}

	left := robSub(root.Left)
	right := robSub(root.Right)

	res[0] = root.Val + left[1] + right[1]
	res[1] = max(left[0], right[1]) + max(left[1], right[0])

	return res
}
