package dividetwointegers

import "math"

func divide(dividend int, divisor int) int {
	if divisor == 0 ||
		(dividend == math.MinInt32 && divisor == -1) { // -INT_MIN = INT_MAX + 1
		return math.MaxInt32
	} else if dividend == 0 {
		return 0
	}
	result := 0
	for den, sor := abs(dividend), abs(divisor); den >= sor; {
		temp, quadratics := sor, 1 // temp = abs(divisor)
		for den >= (temp << 1) {   // ensure den >= 0
			temp <<= 1       // temp *= 2
			quadratics <<= 1 // quadratics *=2
		}
		den -= temp
		result += quadratics
	}
	if (dividend >= 0) != (divisor >= 0) {
		return -result
	}
	return result
}

func abs(x int) int64 {
	if x < 0 {
		return int64(-x)
	}
	return int64(x)
}
