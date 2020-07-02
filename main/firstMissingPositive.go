package main

func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}
	for i, num := range nums {
		for num >= 0 && num < len(nums) && nums[num] != num {
			temp := nums[num]
			nums[num] = num
			nums[i] = temp
			num = temp
		}
	}

	res := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == res {
			res++
		}
	}
	if res == nums[0] {
		res++
	}
	return res
}
