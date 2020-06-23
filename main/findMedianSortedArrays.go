package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2,nums1)
	}
	left,right := 0, len(nums1)
	halfLen := (len(nums1) + len(nums2) + 1) / 2
	for left <= right {
		i := (left + right) / 2
		j := halfLen - i
		if i < right && nums1[i] < nums2[j - 1] {
			left = i + 1
		} else if i > left && nums1[i - 1] > nums2[j] {
			right = i - 1
		} else {
			maxLeft := 0.0
			if (i == 0) {
				maxLeft = float64(nums2[j-1] )
			} else if (j == 0) {
				maxLeft = float64(nums1[i-1] )
			} else {
				maxLeft = math.Max(float64(nums1[i-1]), float64(nums2[j-1]))
			}
			if (len(nums1) + len(nums2)) % 2 == 1{
				return maxLeft
			}
			minRight := 0.0
			if i == len(nums1) {
				minRight = float64(nums2[j])
			}else if j == len(nums2) {
				minRight = float64(nums1[i])
			}else {
				minRight = math.Min(float64(nums1[i]),float64(nums2[j]))
			}
			fmt.Println(left,right)
			return (maxLeft + minRight) / 2.0
		}
	}
	return 0.0
}