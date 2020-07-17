package main

func jump(nums []int) int {
	left := 0
	right := 0
	step := 0
	for right < len(nums)-1 {
		temp := left + nums[left]
		for i := left; i <= right; i++ {
			temp = max(temp, i+nums[i])
		}
		left = right + 1
		right = temp
		step++
	}
	return step
}
