package main

func countBits(num int) []int {
	ans := make([]int, num+1)
	threshold := 2
	offset := 1
	for i := 1; i <= num; i++ {
		if i == threshold {
			threshold *= 2
			offset *= 2
		}
		k := i ^ offset
		ans[i] = ans[k] + 1
	}

	return ans
}
