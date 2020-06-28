package main

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if haystack == "" {
		return -1
	}
	next := make([][]int, len(needle))
	for i := 0; i < len(next); i++ {
		next[i] = make([]int, 256)
	}
	help := func(need string) {
		next[0][need[0]] = 1
		X := 0
		for i := 1; i < len(need); i++ {
			for j := 0; j < 256; j++ {
				next[i][j] = next[X][j]
				next[i][need[i]] = i + 1
			}
			X = next[X][need[i]]
		}
	}
	help(needle)
	pos := 0
	for i, value := range haystack {
		pos = next[pos][value]
		if pos == len(needle) {
			return i - len(needle) + 1
		}
	}
	return -1
}
