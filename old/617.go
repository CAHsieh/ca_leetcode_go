package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if nil == t1 && nil == t2 {
		return nil
	} else if nil == t1 {
		return t2
	} else if nil == t2 {
		return t1
	}

	t1.Val += t2.Val
	leftNode := mergeTrees(t1.Left, t2.Left)
	if nil == t1.Left {
		t1.Left = leftNode
	}

	rightNode := mergeTrees(t1.Right, t2.Right)
	if nil == t1.Right {
		t1.Right = rightNode
	}

	return t1
}
