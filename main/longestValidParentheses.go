package main

import (
	"fmt"
	"math"
)

func longestValidParentheses(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 0
	max := 0
	for i, c := range s {
		if i == 0 {
			dp[i+1] = 0
			continue
		}
		if c == '(' {
			dp[i+1] = 0
		} else {
			if s[i-1] == '(' {
				dp[i+1] = dp[i-1] + 2
			}
			if s[i-1] == ')' {
				if i-1-dp[i] >= 0 && s[i-1-dp[i]] == '(' {
					dp[i+1] = dp[i] + dp[i-1-dp[i]] + 2
				}
			}
		}
		max = int(math.Max(float64(max), float64(dp[i+1])))
	}
	fmt.Println(dp)
	return max
}
