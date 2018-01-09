package main

func lengthOfLIS(nums []int) int {
	if nums == nil || len(nums) < 1 {
		return 0
	}

	length := make([]int, len(nums))

	for i := range nums {
		length[i] = 1
	}

	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				length[j] = max1(length[i]+1, length[j])
			}
		}
	}

	return max2(length)
}

func max1(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func max2(x []int) int {
	max := 0
	for _, v := range x {
		if max < v {
			max = v
		}
	}
	return max
}
