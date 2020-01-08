package main

func productExceptSelf(nums []int) []int {

	zeroCount := 0
	zeroIdx := -1
	sum := 1

	for i, n := range nums {
		if 0 != n {
			sum *= n
		} else {
			zeroCount++
			zeroIdx = i
		}
	}

	if 0 == zeroCount {
		for i, n := range nums {
			nums[i] = sum / n
		}
		return nums
	} else if 1 == zeroCount {
		ans := make([]int, len(nums))
		ans[zeroIdx] = sum
		return ans
	} else {
		return make([]int, len(nums))
	}
}
