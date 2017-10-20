package main

func pathSum(root *TreeNode, sum int) int {
	if nil == root {
		return 0
	}

	count := trace(root, sum)
	count += pathSum(root.Left, sum)
	count += pathSum(root.Right, sum)

	return count
}

func trace(node *TreeNode, sum int) int {
	if nil == node {
		return 0
	}
	sum -= node.Val
	if 0 == sum {
		return 1 + (trace(node.Left, sum) + trace(node.Right, sum))
	}
	return (trace(node.Left, sum) + trace(node.Right, sum))
}
