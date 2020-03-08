package main

import "fmt"

func main() {
	a := TreeNode{7, nil, nil}
	b := TreeNode{15, nil, nil}
	c := TreeNode{20, &b, &a}
	d := TreeNode{9, nil, nil}
	e := TreeNode{3, &d, &c}
	fmt.Println(averageOfLevels(&e))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var count map[int]int

func averageOfLevels(root *TreeNode) []float64 {

	if root.Left == nil && root.Right == nil {
		return []float64{float64(root.Val)}
	}

	count = make(map[int]int)
	levelSum := sum(root, 0)
	depth := len(levelSum)
	avg := make([]float64, depth)
	avg[0] = float64(root.Val)
	for i := 1; i < depth; i++ {
		avg[i] = float64(levelSum[i]) / float64(count[i])
	}
	return avg
}

func sum(root *TreeNode, depth int) []int {
	curr := []int{root.Val}
	if root.Left == nil && root.Right == nil {
		return curr
	}
	var dls []int
	llen := 0
	var drs []int
	rlen := 0
	if root.Left != nil {
		count[depth+1]++
		dls = sum(root.Left, depth+1)
		llen = len(dls)
	}
	if root.Right != nil {
		count[depth+1]++
		drs = sum(root.Right, depth+1)
		rlen = len(drs)
	}
	if llen < rlen {
		for i, v := range dls {
			drs[i] += v
		}
		return append(curr, drs...)
	} else {
		for i, v := range drs {
			dls[i] += v
		}
		return append(curr, dls...)
	}
}

// fastest solution
// func averageOfLevels(root *TreeNode) []float64 {
// 	result := []float64{}
// 	if root == nil {
// 		return result
// 	}
// 	queue := []*TreeNode{root}
// 	for len(queue) > 0 {
// 		newQueue := []*TreeNode{}
// 		level := len(result)
// 		result = append(result, 0)
// 		for _, node := range queue {
// 			result[level] += float64(node.Val)
// 			if node.Left != nil {
// 				newQueue = append(newQueue, node.Left)
// 			}
// 			if node.Right != nil {
// 				newQueue = append(newQueue, node.Right)
// 			}
// 		}
// 		result[level] /= float64(len(queue))
// 		queue = newQueue
// 	}
// 	return result
// }
