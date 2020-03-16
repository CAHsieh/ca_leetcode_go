package main

import (
	"fmt"
	"math"
)

func main() {
	a := TreeNode{1, nil, nil}
	b := TreeNode{0, nil, nil}
	c := TreeNode{1, nil, nil}
	d := TreeNode{0, nil, nil}
	e := TreeNode{1, &b, &a}
	f := TreeNode{0, &d, &c}
	g := TreeNode{1, &f, &e}
	fmt.Println(sumRootToLeaf(&g))
	fmt.Println(sumRootToLeaf(&b))
}

var sum int

func sumRootToLeaf(root *TreeNode) int {
    sum=0
	dfs(root, int(math.Pow(2, 30)), 0, 1)
	return sum
}

func dfs(node *TreeNode, two, tmpSum, i int) {
	tmpSum += node.Val * two
	if node.Left == nil && node.Right == nil {
		sum += (tmpSum >> (31 - i))
		return
	}
    if node.Left != nil {
        dfs(node.Left, two/2, tmpSum, i+1)
    }
	if node.Right != nil {
	    dfs(node.Right, two/2, tmpSum, i+1)
    }
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
