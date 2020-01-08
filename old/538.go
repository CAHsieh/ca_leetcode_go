package main

func convertBST(root *TreeNode) *TreeNode {
	if nil == root {
		return nil
	}

	sum(root, 0)

	return root
}

func sum(node *TreeNode, num int) int {
	if nil == node {
		return num
	}

	right := sum(node.Right, num)
	left := sum(node.Left, node.Val+right)

	node.Val += right
	return left
}
