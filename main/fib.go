package main

func fib(N int) int {
	if N == 0 {
		return 0
	}
	a, res := 1, 1
	if N == 1 || N == 2 {
		return res
	}
	for i := 3; i <= N; i++ {
		temp := res
		res = a + res
		a = temp
	}
	return res
}
