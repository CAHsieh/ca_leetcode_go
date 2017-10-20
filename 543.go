package main

func diameterOfBinaryTree(root *TreeNode) int {
	if nil == root {
		return 0
	} else if nil == root.Left && nil == root.Right {
		return 0
	}
	rootDepth := getDepth(root.Left) + getDepth(root.Right) - 1
	subDepth := max(diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right))
	return max(rootDepth, subDepth)
}

// func getDepth(node *TreeNode) int {
// 	if nil == node {
// 		return 0
// 	} else {
// 		return max(getDepth(node.Left)+1, getDepth(node.Right)+1)
// 	}
// }

// func max(x int, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }
