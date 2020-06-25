package main

func romanToInt(s string) int {
	var m = make(map[int32]int)
	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000
	res := 0
	for i, i2 := range s {
		if i == len(s)-1 {
			res += m[i2]
		} else {
			if m[i2] < m[int32(s[i+1])] {
				res -= m[i2]
			} else {
				res += m[i2]
			}
		}
	}
	return res
}
