package main

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left, right := 0, len(nums)-1
	res := make([]int, 2)
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target && (mid == 0 || nums[mid-1] != nums[mid]) {
			res[0] = mid
			break
		} else if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	left, right = 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target && (mid == len(nums)-1 || nums[mid+1] != nums[mid]) {
			res[1] = mid
			break
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left > right {
		res[0] = -1
		res[1] = -1
	}
	return res
}
