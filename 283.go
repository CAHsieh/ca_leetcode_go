package main

func moveZeroes(nums []int) {
	count := 0
	for i := 0; i < len(nums); i++ {
		if 0 == nums[i] {
			count++
		} else if 0 != count {
			nums[i-count] = nums[i]
			nums[i] = 0
		}
	}
}
