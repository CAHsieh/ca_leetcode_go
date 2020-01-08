package main

func rob(nums []int) int {
	if nil == nums || len(nums) == 0 {
		return 0
	}
	Max := nums[0]
	for i := 1; i < len(nums); i++ {
		if i-3 >= 0 {
			nums[i] += max(nums[i-2], nums[i-3])
		} else if i-2 >= 0 {
			nums[i] += nums[i-2]
		}
		Max = max(Max, nums[i])
	}
	return Max
}
