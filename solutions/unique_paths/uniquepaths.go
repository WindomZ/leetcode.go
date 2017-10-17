package uniquepaths

import "math"

func uniquePaths(m int, n int) int {
	var f float64 = 1
	// make sure m > n
	if m <= n {
		n, m = m, n
	}
	// f = ( m * (m+1) * (m+2) * ... * (m+n-2) ) / n!
	for i := 1; i < n; i++ {
		f *= (float64)(m+i-1) / (float64)(i)
	}
	return (int)(math.Floor(f))
}
