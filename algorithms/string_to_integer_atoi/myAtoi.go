package string_to_integer_atoi

import "math"

func myAtoi(str string) int {
	var l, idx int = len(str), 0
	for ; idx < l && str[idx] == ' '; idx++ {
	}
	if idx == l {
		return 0
	}
	var b byte
	var sign int = 1
	if b = str[idx]; b == '+' {
		idx++
	} else if b == '-' {
		sign = -1
		idx++
	}
	var n, max uint = 0, math.MaxInt32 / 10
	for ; idx < l; idx++ {
		if b = str[idx]; b < '0' || b > '9' {
			return sign * int(n)
		} else if n > max || (n == max && b > '7') {
			if sign == 1 {
				return math.MaxInt32
			}
			return -math.MaxInt32 - 1
		}
		n = 10*n + uint(b-'0')
	}

	return sign * int(n)
}
