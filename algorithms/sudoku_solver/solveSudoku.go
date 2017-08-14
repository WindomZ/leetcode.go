package sudokusolver

func solveSudoku(board [][]byte) {
	solve(board)
}

func solve(board [][]byte) bool {
	var v byte
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if board[x][y] == '.' {
				for v = '1'; v <= '9'; v++ {
					i := 0
					for ; i < 9; i++ {
						if i != x && board[i][y] != '.' &&
							board[i][y] == v { // check each row
							break
						}
						if i != y && board[x][i] != '.' &&
							board[x][i] == v { // check each column
							break
						}
						row, col := x/3*3+i/3, y/3*3+i%3
						if row != x && col != y &&
							board[row][col] != '.' &&
							board[row][col] == v { // check each 3*3 grid
							break
						}
					}
					if i == 9 {
						board[x][y] = v
						if solve(board) {
							return true
						} else {
							board[x][y] = '.' // restore
						}
					}
				}
				return false
			}
		}
	}
	return true
}
