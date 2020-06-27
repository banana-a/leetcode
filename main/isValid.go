package main

func isValid(s string) bool {
	var stack []int
	for _, c := range s {
		switch c {
		case '(':
			stack = append(stack, 1)
		case '[':
			stack = append(stack, 2)
		case '{':
			stack = append(stack, 3)
		case ')':
			if len(stack) > 0 && stack[len(stack)-1] == 1 {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case ']':
			if len(stack) > 0 && stack[len(stack)-1] == 2 {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case '}':
			if len(stack) > 0 && stack[len(stack)-1] == 3 {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}

		}
	}
	return len(stack) == 0
}
