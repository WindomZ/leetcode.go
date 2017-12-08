package minimumpathsum

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	cur := make([]int, m)

	cur[0] = grid[0][0] // the first space
	for i := 1; i < m; i++ {
		cur[i] = cur[i-1] + grid[i][0] // count all column paths in the first row
	}

	for j := 1; j < n; j++ {
		cur[0] += grid[0][j] // move down in the first column
		for i := 1; i < m; i++ {
			// cur[i-1] -> count paths if move right
			// cur[i] -> count paths if move down
			// choose the min paths
			cur[i] = min(cur[i-1], cur[i]) + grid[i][j]
		}
	}
	return cur[m-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
