package main

import (
	"fmt"
	"strconv"
)

func main() {
	node3 := TreeNode{3, nil, nil}
	node2 := TreeNode{2, nil, &node3}
	node1 := TreeNode{1, nil, &node2}

	node5 := TreeNode{3, nil, nil}
	node4 := TreeNode{1, nil, &node5}
	fmt.Println(leafSimilar(&node1, &node4))
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	seq1 := ""
	seq2 := ""
	travel(root1, &seq1)
	travel(root2, &seq2)
	return (seq1 == seq2)
}

func travel(root *TreeNode, seq *string) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*seq += strconv.Itoa(root.Val)
		return
	}

	travel(root.Left, seq)
	travel(root.Right, seq)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
