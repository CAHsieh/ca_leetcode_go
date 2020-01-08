package main

func findDuplicate(nums []int) int {
	n := len(nums)
	slow := n
	fast := n

	slow = nums[slow-1]
	fast = nums[nums[fast-1]-1]

	for slow != fast {
		slow = nums[slow-1]
		fast = nums[nums[fast-1]-1]
	}

	slow = n

	for slow != fast {
		slow = nums[slow-1]
		fast = nums[fast-1]
	}

	return slow
}
