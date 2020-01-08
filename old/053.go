package main

func maxSubArray(nums []int) int {
	if nil == nums || len(nums) == 0 {
		return 0
	}

	sum := nums[0]
	Max := nums[0]
	for i := 1; i < len(nums); i++ {
		sum = max(nums[i], sum+nums[i])

		if Max < sum {
			Max = sum
		}
	}

	return Max
}
