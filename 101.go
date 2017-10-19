package main

func isSymmetric(root *TreeNode) bool {
	if nil == root {
		return true
	}
	return isSame(root.Left, root.Right)
}

func isSame(node1 *TreeNode, node2 *TreeNode) bool {
	if nil == node1 && nil == node2 {
		return true
	} else if nil == node1 || nil == node2 {
		return false
	} else {
		if node1.Val == node2.Val {
			return (isSame(node1.Left, node2.Right) && isSame(node1.Right, node2.Left))
		} else {
			return false
		}
	}
}
