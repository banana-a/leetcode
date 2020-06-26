package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	pos := 0
	res := ""
	c := uint8('0')
	for true {
		for i, i2 := range strs {
			if pos >= len(i2) {
				return res
			}
			if i == 0 {
				c = i2[pos]
			}
			if c != i2[pos] {
				return res
			}
		}
		res += string(c)
		pos++
	}
	return res
}
