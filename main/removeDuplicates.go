package main

func removeDuplicates(nums []int) int {
	pos := 1
	i := 1
	for ; i < len(nums); i++ {
		if nums[i] != nums[pos] {
			nums[pos] = nums[i]
			pos++
		}
	}
	return pos + 1
}
