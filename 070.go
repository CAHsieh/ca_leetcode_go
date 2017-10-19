package main

func climbStairs(n int) int {
	if n < 2 {
		return 1
	}
	step := make([]int, n)
	step[0] = 1
	step[1] = 2
	for i := 2; i < n; i++ {
		step[i] = step[i-1] + step[i-2]
	}
	return step[n-1]
}
