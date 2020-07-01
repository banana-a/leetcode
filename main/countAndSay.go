package main

import "strconv"

func countAndSay(n int) string {
	res := make([]string, 31)
	res[1] = "1"
	for i := 2; i <= n; i++ {
		res[i] = help(res[i-1])
	}
	return res[n]
}

func help(s string) string {
	var num int = 0
	var res string = ""
	for i := 0; i <= len(s); i++ {
		if i == 0 {
			num = 1
			continue
		}
		if i == len(s) || s[i] != s[i-1] {
			res += strconv.Itoa(num) + string(s[i-1])
			num = 1
			continue
		} else {
			num++
		}
	}
	return res
}
