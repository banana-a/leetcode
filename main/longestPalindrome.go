package main

func longestPalindrome(s string) string{
	if len(s) == 0 {
		return s
	}
	p := 0
	rl,rr := 0,0
	for p < len(s) {
		left,right := p,p
		for left >= 0 && right < len(s) && string(s[left]) == string(s[right]){
			left--
			right++
		}
		if rr - rl < right - left - 2{
			rr = right - 1
			rl = left + 1
		}
		left,right = p,p+1
		for left >= 0 && right < len(s) && string(s[left]) == string(s[right]){
			left--
			right++
		}
		if rr - rl < right - left - 2{
			rr = right - 1
			rl = left + 1
		}
		p++
	}
	return s[rl:rr + 1]
}