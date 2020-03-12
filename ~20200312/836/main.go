package main

import "fmt"

func main() {
	fmt.Println(isRectangleOverlap([]int{0, 0, 2, 2}, []int{1, 1, 3, 3}))
	fmt.Println(isRectangleOverlap([]int{0, 0, 1, 1}, []int{1, 0, 2, 1}))
	fmt.Println(isRectangleOverlap([]int{7, 8, 13, 15}, []int{10, 8, 12, 20}))
	fmt.Println(isRectangleOverlap([]int{4, 4, 14, 7}, []int{4, 3, 8, 8}))
}

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return rec1[0] < rec2[2] && rec2[0] < rec1[2] && rec1[1] < rec2[3] && rec2[1] < rec1[3]
}
