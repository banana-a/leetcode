package main

func findTargetSumWays(nums []int, S int) int {
	target := 0
	for _, i2 := range nums {
		target += i2
	}
	target += S
	if target%2 != 0 {
		return 0
	}
	target = target / 2
	if target < 0 {
		return 0
	}
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := target; j >= 0; j-- {
			dp[j] = dp[j]
			if j-nums[i] >= 0 {
				dp[j] += dp[j-nums[i]]
			}
		}
	}
	return dp[len(dp)-1]
}
