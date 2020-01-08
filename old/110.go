package main

func isBalanced(root *TreeNode) bool {
	if nil == root {
		return true
	}

	if abs(getDepth(root.Left)-getDepth(root.Right)) >= 2 {
		return false
	}

	return (isBalanced(root.Left) && isBalanced(root.Right))
}

// func getDepth(node *TreeNode) int {
// 	if nil == node {
// 		return 0
// 	} else {
// 		return max(getDepth(node.Left)+1, getDepth(node.Right)+1)
// 	}
// }
