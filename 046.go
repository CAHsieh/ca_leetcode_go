package main

func permute(nums []int) [][]int {
	used := make([]bool, len(nums))
	ans := make([][]int, 0)

	for i, val := range nums {
		tmp := make([]int, 0)
		tmp = append(tmp, val)
		used[i] = true
		p(used, nums, &tmp, &ans)
		used[i] = false
	}

	return ans
}

func p(used []bool, nums []int, tmp *[]int, ans *[][]int) {
	if len(*tmp) == len(nums) {
		ttmp := make([]int, len(*tmp))
		copy(ttmp, *tmp)
		*ans = append(*ans, ttmp)
		return
	}
	for i, val := range nums {
		if !used[i] {
			*tmp = append(*tmp, val)
			used[i] = true
			p(used, nums, tmp, ans)
			used[i] = false
			*tmp = (*tmp)[:len(*tmp)-1]
		}
	}
}
