package main

import "fmt"

func main() {
	a := TreeNode{1, nil, nil}
	b := TreeNode{2, &a, nil}
	c := TreeNode{4, nil, nil}
	d := TreeNode{0, nil, &b}
	e := TreeNode{3, &d, &c}
	r := trimBST(&e, 1, 3)
	fmt.Println(r.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < L {
		return trimBST(root.Right, L, R)
	} else if root.Val > R {
		return trimBST(root.Left, L, R)
	}
	root.Left = trimBST(root.Left, L, R)
	root.Right = trimBST(root.Right, L, R)
	return root
}
