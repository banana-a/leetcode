package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	res := ""
	pos := 0
	for k := 0; k < numRows; k++ {
		pos = 0
		for ; pos < len(s); pos += 2*numRows - 2 {
			if k == 0 {
				res += string(s[pos])
			} else if k == numRows-1 {
				if pos+k < len(s) {
					res += string(s[pos+k])
				}
			} else {
				if pos+k < len(s) {
					res += string(s[pos+k])
				}
				if pos+2*numRows-2-k < len(s) {
					res += string(s[pos+2*numRows-2-k])
				}
			}
		}
	}
	return res
}
