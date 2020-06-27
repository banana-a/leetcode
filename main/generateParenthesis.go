package main

func generateParenthesis(n int) []string {
	var help func(left int, right int, item string)
	var res []string
	help = func(left int, right int, item string) {
		if left == 0 && right == 0 {
			res = append(res, item)
		}
		if left > 0 {
			help(left-1, right+1, item+"(")
		}
		if right > 0 {
			help(left, right-1, item+")")
		}
	}
	help(n, 0, "")
	return res
}
