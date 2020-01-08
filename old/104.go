package main

func maxDepth(root *TreeNode) int {
	if nil == root {
		return 0
	}

	return int(max(maxDepth(root.Left)+1, maxDepth(root.Right)+1))
}

// func max(x int, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }
