package main

func removeElement(nums []int, val int) int {
	pos := 0
	i := 0
	for ; i < len(nums); i++ {
		if nums[i] != val {
			nums[pos] = nums[i]
			pos++
		}
	}
	return pos
}
