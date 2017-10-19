package main

import "sort"

func majorityElement(nums []int) int {
	sort.Ints(nums)

	if nums[0] == nums[len(nums)/2] {
		return nums[0]
	} else if nums[len(nums)-1] == nums[len(nums)/2+1] {
		return nums[len(nums)-1]
	} else {
		return nums[len(nums)/2]
	}
}
