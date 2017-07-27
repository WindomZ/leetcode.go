package reverse_integer

import "math"

func reverse(x int) int {
	if x >= -9 && x <= 9 {
		return x
	}

	var r, t, n int
	for neg := x < 0; x != 0; x = x / 10 {
		t = x % 10
		n = r*10 + t
		if n > math.MaxInt32 || n < -math.MaxInt32 {
			return 0
		} else if neg {
			if n > 0 {
				return 0
			}
		} else if n < 0 {
			return 0
		}
		r = n
	}

	return r
}
