package main

func inorderTraversal(root *TreeNode) []int {
	s := make([]int, 0)
	if nil == root {
		return s
	}

	if nil != root.Left {
		s = append(s, inorderTraversal(root.Left)...)
	}
	s = append(s, root.Val)

	if nil != root.Right {
		s = append(s, inorderTraversal(root.Right)...)
	}

	return s
}
