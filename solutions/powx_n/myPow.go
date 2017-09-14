package powxn

import "math"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == math.MinInt32 {
		if x == 1 || x == -1 {
			return 1
		}
		return 0
	}
	if n == 2 {
		return x * x
	}

	var r float64 = 1
	for i := int64(math.Abs(float64(n))); i > 0; {
		if (i & 1) == 1 {
			r *= x
		}
		i >>= 1
		x *= x
	}
	if n < 0 {
		return 1 / r
	}
	return r
}
