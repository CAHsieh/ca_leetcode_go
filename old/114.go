package main

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	if root.Left != nil {
		findLast(root.Left).Right = root.Right
	}
	if root.Left != nil {
		root.Right = root.Left
		root.Left = nil
	}
}

func findLast(node *TreeNode) *TreeNode {
	if node.Right != nil {
		return findLast(node.Right)
	} else if node.Left != nil {
		return findLast(node.Left)
	} else {
		return node
	}
}
