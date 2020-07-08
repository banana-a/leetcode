package main

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i, _ := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 0; i <= amount; i++ {
		for _, value := range coins {
			if i+value <= amount {
				dp[i+value] = min(dp[i+value], dp[i]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
