package main

func isMatch(s string, p string) bool {
	flag := make([][]bool, len(s)+1)
	for i := 0; i <= len(s); i++ {
		flag[i] = make([]bool, len(p)+1)
	}
	flag[0][0] = true
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}
	for i := 0; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] == '*' {
				flag[i][j] = flag[i][j-2]
				if matches(i, j-1) {
					flag[i][j] = flag[i][j] || flag[i-1][j]
				}
			}
			if matches(i, j) {
				flag[i][j] = flag[i-1][j-1]
			}
		}
	}
	return flag[len(s)][len(p)]
}
