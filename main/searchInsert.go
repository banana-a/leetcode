package main

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	if target <= nums[0] {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target || (target < nums[mid] && target > nums[mid-1]) {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return 0
}
