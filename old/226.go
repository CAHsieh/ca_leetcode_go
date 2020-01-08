package main

func invertTree(root *TreeNode) *TreeNode {

	if nil == root {
		return nil
	}

	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}
