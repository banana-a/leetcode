package main

import "math"

func myAtoi(str string) int {
	if str == "" {
		return 0
	}
	left := 0
	temp := str[left]
	for !(temp == '+' || temp == '-' || (temp >= '0' && temp <= '9')) {
		if temp != ' ' {
			return 0
		}
		left++
		if left >= len(str) {
			return 0
		}
		temp = str[left]
	}
	right := left + 1
	for right < len(str) && str[right] >= '0' && str[right] <= '9' {
		right++
	}
	sub := str[left:right]
	if sub == "-" || sub == "+" {
		return 0
	}
	res := int64(0)
	start := 0
	fuhao := int64(1)
	if sub[0] == '+' {
		start++
	} else if sub[0] == '-' {
		fuhao = -1
		start++
	}
	for start < len(sub) {
		res = res*10 + int64(sub[start]) - int64('0')
		if res*fuhao > math.MaxInt32 {
			return math.MaxInt32
		} else if res*fuhao < math.MinInt32 {
			return math.MinInt32
		}
		start++
	}
	return int(res * fuhao)
}
