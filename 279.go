package main

func numSquares(n int) int {
	dp := make([]int, n+1)

	dp[0] = 0

	for i := 1; i <= n; i++ {
		minV := n
		for j := 1; i-j*j >= 0; j++ {
			minV = min(minV, dp[i-j*j]+1)
		}
		dp[i] = minV
	}

	return dp[n]
}
