package main

func intToRoman(num int) string {
	var item = [...]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var pay = [...]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	res := ""
	pos := 0
	for num > 0 {
		if num >= pay[pos] {
			res += item[pos]
			num -= pay[pos]
		} else {
			pos++
		}
	}
	return res
}
