package main

func canPartition(nums []int) bool {

	sum := 0

	for _, v := range nums {
		sum += v
	}

	if sum%2 == 1 {
		return false
	}

	sum /= 2

	dp := make([]bool, sum+1)

	dp[0] = true

	for _, v := range nums {
		for k := sum; k > 0; k-- {
			if k >= v {
				dp[k] = dp[k] || dp[k-v]
			}
		}
	}

	return dp[sum]
}
