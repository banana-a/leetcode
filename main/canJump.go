package main

func canJump(nums []int) bool {
	m := 0
	for i := 0; i <= m; i++ {
		m = max(m, i+nums[i])
		if m >= len(nums)-1 {
			return true
		}
	}
	return false
}
