package main

func canPartition(nums []int) bool {
	target := 0
	for _, i2 := range nums {
		target += i2
	}
	if target%2 != 0 {
		return false
	}
	target /= 2
	dp := make([]bool, target+1)
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		for j := target; j >= 0; j-- {
			if j-nums[i] >= 0 {
				dp[j] = dp[j] || dp[j-nums[i]]
			}
		}
	}
	return dp[target]
}
