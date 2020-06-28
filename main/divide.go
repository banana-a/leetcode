package main

import "math"

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	fuhao := 1
	if dividend*divisor < 0 {
		fuhao = -1
	} else if dividend == 0 {
		return 0
	}
	var div func(a int, b int) int
	div = func(a int, b int) int {
		if a < b {
			return 0
		}
		time := 1
		num := divisor
		for num+num <= dividend {
			if num+num == dividend {
				return time + time
			}
			num = num + num
			time = time + time
		}
		return time + divide(dividend-num, divisor)
	}
	dividend = int(math.Abs(float64(dividend)))
	divisor = int(math.Abs(float64(divisor)))
	return div(dividend, divisor) * fuhao
}
