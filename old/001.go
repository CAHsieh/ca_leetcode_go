package main

func twoSum(nums []int, target int) []int {
	if nil == nums {
		return nil
	}
	maps := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		mTarget := target - nums[i]
		idx, hasKey := maps[mTarget]
		if hasKey {
			return []int{idx, i}
		}
		maps[nums[i]] = i
	}
	return nil
}
