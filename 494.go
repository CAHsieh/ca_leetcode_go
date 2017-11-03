package main

/*
 * S = sumP - sumN
 * S + sumP + sumN = 2 * sumP
 * S + SUM = 2 * sumP
 * sumP = (S+SUM)/2
 */

func findTargetSumWays(nums []int, S int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	if sum < S || (sum+S)%2 == 1 {
		return 0
	}
	return find(nums, (S+sum)/2)
}

func find(nums []int, sum int) int {
	dp := make([]int, sum+1)
	dp[0] = 1

	for _, val := range nums {
		for i := sum; i >= val; i-- {
			dp[i] += dp[i-val]
		}
	}

	return dp[sum]
}
