package main

import "fmt"

func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	count := len(nums2)
	for i := 0; i < count-1; i++ {
		for j := i + 1; j < count; j++ {
			if nums2[j] > nums2[i] {
				m[nums2[i]] = nums2[j]
				break
			}
		}
	}

	count = len(nums1)
	for i := 0; i < count; i++ {
		v, ok := m[nums1[i]]
		if ok {
			nums1[i] = v
		} else {
			nums1[i] = -1
		}
	}
	return nums1
}
