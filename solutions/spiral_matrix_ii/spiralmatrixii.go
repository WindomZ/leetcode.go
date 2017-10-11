package spiralmatrixii

func generateMatrix(n int) [][]int {
	if n <= 0 {
		return [][]int{}
	}
	if n == 1 {
		return [][]int{{1}}
	}
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}
	for i, j, k := 0, 0, 1; k <= n*n; i++ {
		// top, horizontal, left to right
		for j = i; j < n-i; j, k = j+1, k+1 {
			ret[i][j] = k
		}
		// right, vertical, top to bottom
		for j = i + 1; j < n-i; j, k = j+1, k+1 {
			ret[j][n-i-1] = k
		}
		// bottom, horizontal, right to left
		for j = n - i - 2; j > i; j, k = j-1, k+1 {
			ret[n-i-1][j] = k
		}
		// left, vertical, bottom to top
		for j = n - i - 1; j > i; j, k = j-1, k+1 {
			ret[j][i] = k
		}
	}
	return ret
}
