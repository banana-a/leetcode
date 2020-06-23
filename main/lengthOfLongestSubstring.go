package main

func lengthOfLongestSubstring(s string) int {
	left,right := 0,0                           //双指针的窗口
	maxl := 0
	bytes := []byte(s)
	m := make(map[byte]int)						//维护字符最后出现位置，作为移动左指针的依据
	for right < len(s) {
		c,ok := m[bytes[right]]
		if ok && c >= left{						//出现重复
			if right - left + 1 > maxl {
				 maxl = right - left
			}
			left = c + 1
		}
		m[bytes[right]] = right
		right++
	}
	if right == len(s) {
		if right - left + 1 > maxl {
			maxl = right - left
		}
	}
	return maxl
}