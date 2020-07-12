package main

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := amount; j >= 0; j-- {
			num := 1
			for j-coins[i]*num >= 0 {
				dp[j] += dp[j-num*coins[i]]
				num++
			}
		}
	}
	return dp[amount]
}
