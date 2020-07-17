package main

func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	res0 := make([]int, len(nums))
	res1 := make([]int, len(nums))
	res0[0] = nums[0]
	res0[1] = nums[0]
	res1[0] = 0
	res1[1] = nums[1]
	for i := 2; i < len(nums); i++ {
		res0[i] = max(res0[i-1], res0[i-2]+nums[i])
		res1[i] = max(res1[i-1], res1[i-2]+nums[i])
	}
	return max(res0[len(nums)-2], res1[len(nums)-1])
}
