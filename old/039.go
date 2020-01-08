package main

func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	var tmp []int
	backtrack(&ans, tmp, candidates, target, 0)
	return ans
}

func backtrack(ans *[][]int, tmp []int, candidates []int, remain int, start int) {
	if remain < 0 {
		return
	} else if remain == 0 {
		in := make([]int, len(tmp))
		copy(in, tmp)
		*ans = append(*ans, in)
	} else {
		for i := start; i < len(candidates); i++ {
			tmp = append(tmp, candidates[i])
			backtrack(ans, tmp, candidates, remain-candidates[i], i)
			tmp = tmp[:len(tmp)-1]
		}
	}
}
