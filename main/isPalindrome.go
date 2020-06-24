package main

import "strconv"

func isPalindrome(x int) bool {
	num := strconv.Itoa(x)
	left, right := 0, len(num)-1
	for left < right {
		if num[left] != num[right] {
			return false
		}
		left++
		right--
	}
	return true
}
