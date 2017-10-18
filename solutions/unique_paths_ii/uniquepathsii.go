package uniquepathsii

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}
	width := len(obstacleGrid[0])
	dp := make([]int, width)
	dp[0] = 1
	for _, row := range obstacleGrid {
		for col := 0; col < width; col++ {
			if row[col] == 1 {
				dp[col] = 0 // reset current paths to zero if in obstacle space
			} else if col > 0 {
				dp[col] += dp[col-1] // add up last paths if in empty space
			}
		}
	}
	return dp[width-1]
}
