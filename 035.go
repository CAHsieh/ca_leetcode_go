package main

func searchInsert(nums []int, target int) int {
	end := len(nums) - 1
	return binarySearch(nums, target, 0, end)
}

func binarySearch(nums []int, target, head, end int) int {

	if head > end {
		return head
	}

	index := head + (end-head+1)/2

	if nums[index] == target {
		return index
	} else if nums[index] < target {
		head = index + 1
	} else {
		end = index - 1
	}
	return binarySearch(nums, target, head, end)
}
