package main

func maxSubArray(nums []int) int {
	res := nums[0]
	now := nums[0]
	for i := 1; i < len(nums); i++ {
		if now < 0 {
			now = 0
		}
		now += nums[i]
		res = max(res, now)
	}
	return res
}
