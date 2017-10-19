package main

func maxProfit(prices []int) int {
	if nil == prices || 0 == len(prices) {
		return 0
	}

	maxP := 0
	minS := prices[0]
	for i := 1; i < len(prices); i++ {
		minS = min(minS, prices[i])
		maxP = max(maxP, prices[i]-minS)
	}
	return maxP
}

// func max(x int, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }

// func min(x int, y int) int {
// 	if x > y {
// 		return y
// 	}
// 	return x
// }
