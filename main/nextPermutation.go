package main

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 {
		if nums[i] < nums[i+1] {
			break
		}
		i--
	}
	if i == -1 {
		left, right := 0, len(nums)-1
		for left < right {
			temp := nums[left]
			nums[left] = nums[right]
			nums[right] = temp
			left++
			right--
		}
		return
	}
	j := i + 1
	for i < len(nums) {
		if j == len(nums)-1 || (nums[j] > nums[i] && nums[j+1] <= nums[i]) {
			start, end := i+1, len(nums)-1
			temp := nums[i]
			nums[i] = nums[j]
			nums[j] = temp
			for start < end {
				tm := nums[start]
				nums[start] = nums[end]
				nums[end] = tm
				start++
				end--
			}
			return
		}
		j++
	}
	return
}
