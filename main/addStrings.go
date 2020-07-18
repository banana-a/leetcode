package main

func addStrings(num1 string, num2 string) string {
	help := func(a string) string {
		temp := []rune(a)
		i, j := 0, len(a)-1
		for i < j {
			temp[i], temp[j] = temp[j], temp[i]
			i++
			j--
		}
		return string(temp)
	}
	res := ""
	up := 0
	p, q := len(num1)-1, len(num2)-1
	for p >= 0 || q >= 0 || up > 0 {
		item := 0
		if p >= 0 {
			item += int(num1[p] - '0')
			p--
		}
		if q >= 0 {
			item += int(num2[q] - '0')
			q--
		}
		item += up
		res += string(int(item)%10 + '0')
		up = int(item) / 10
	}
	return help(res)
}
