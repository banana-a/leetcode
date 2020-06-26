package main

func letterCombinations(digits string) []string {
	m := make(map[byte][]byte)
	m['2'] = []byte{'a', 'b', 'c'}
	m['3'] = []byte{'d', 'e', 'f'}
	m['4'] = []byte{'g', 'h', 'i'}
	m['5'] = []byte{'j', 'k', 'l'}
	m['6'] = []byte{'m', 'n', 'o'}
	m['7'] = []byte{'p', 'q', 'r', 's'}
	m['8'] = []byte{'t', 'u', 'v'}
	m['9'] = []byte{'w', 'x', 'y', 'z'}
	var res []string
	if digits == "" {
		return res
	}
	res = append(res, "")
	pos := 0
	for pos < len(digits) {
		start := len(res)
		for i := 0; i < start; i++ {
			for _, c := range m[digits[pos]] {
				res = append(res, res[i]+string(c))
			}
		}
		res = res[start:]
		pos++
	}
	return res
}
