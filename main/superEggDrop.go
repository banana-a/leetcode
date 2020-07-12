package main

func superEggDrop(K int, N int) int {
	memo := make([][]int, K+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, N+1)
	}
	var dp func(k int, n int) int
	dp = func(k int, n int) int {
		if k == 1 {
			return n
		}
		if n == 0 {
			return 0
		}
		if memo[k][n] != 0 {
			return memo[k][n]
		}
		res := 0
		for i := 1; i <= n; i++ {
			res = min(res, max(dp(k, n-i), dp(k-1, i-1))) + 1
		}
		memo[k][n] = res
		return res
	}
	return dp(K, N)
}
