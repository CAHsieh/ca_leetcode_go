package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(allCellsDistOrder(1, 2, 0, 0))
	fmt.Println(allCellsDistOrder(2, 2, 0, 1))
	fmt.Println(allCellsDistOrder(2, 3, 1, 2))
}

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	r = r0
	c = c0
	nodes := nodeSlice([]Node{})
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			nodes = append(nodes, Node{i, j})
		}
	}

	sort.Sort(nodes)
	ans := [][]int{}
	for _, node := range nodes {
		ans = append(ans, []int{node.x, node.y})
	}
	return ans
}

var r int
var c int

type Node struct {
	x int
	y int
}

type nodeSlice []Node

// Len is the number of elements in the collection.
func (nodes nodeSlice) Len() int {
	return len(nodes)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (nodes nodeSlice) Less(i, j int) bool {
	di := abs(nodes[i].x-r) + abs(nodes[i].y-c)
	dj := abs(nodes[j].x-r) + abs(nodes[j].y-c)
	return di < dj
}

// Swap swaps the elements with indexes i and j.
func (nodes nodeSlice) Swap(i, j int) {
	nodes[i], nodes[j] = nodes[j], nodes[i]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
