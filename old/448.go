package main

func findDisappearedNumbers(nums []int) []int {
	if nil == nums {
		return nil
	}

	count := make([]int, len(nums)+1)
	length := len(nums)

	for i := 0; i < len(nums); i++ {
		if 0 == count[nums[i]] {
			length--
		}
		count[nums[i]]++
	}

	ans := make([]int, length)
	s := 0

	for i := 1; i <= len(nums); i++ {
		if 0 == count[i] {
			ans[s] = i
			s++
		}
	}

	return ans
}
