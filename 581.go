package main

func findUnsortedSubarray(nums []int) int {

	e := -2
	s := -1
	Max := nums[0]
	Min := nums[len(nums)-1]

	for i := 1; i < len(nums); i++ {
		Max = max(nums[i], Max)
		Min = min(nums[len(nums)-i-1], Min)

		if Max > nums[i] {
			e = i
		}
		if Min < nums[len(nums)-i-1] {
			s = len(nums) - i - 1
		}
	}

	return e - s + 1
}
