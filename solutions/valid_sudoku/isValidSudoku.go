package validsudoku

func isValidSudoku(board [][]byte) bool {
	var row, column, grid [9][9]bool
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[x]); y++ {
			n := board[x][y] - '0' - 1
			if n >= 0 && n < 9 { // n must in [0, 9]
				g := x/3 + y/3*3
				if row[x][n] || // check each row
					column[y][n] || // check each column
					grid[g][n] { // check each 3*3 grid
					return false
				}
				row[x][n] = true
				column[y][n] = true
				grid[g][n] = true
			}
		}
	}
	return true
}
