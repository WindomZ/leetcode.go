package reverseinteger

import "math"

func reverse(x int) int {
	if x >= -9 && x <= 9 {
		return x
	}

	var r, t, n int
	for ; x != 0; x = x / 10 {
		t = x % 10
		n = r*10 + t
		if n > math.MaxInt32 || n < -math.MaxInt32 {
			return 0
		}
		r = n
	}

	return r
}
