package sqrtx

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}
