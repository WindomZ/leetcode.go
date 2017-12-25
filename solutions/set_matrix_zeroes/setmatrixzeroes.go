package setmatrixzeroes

func setZeroes(matrix [][]int) {
	rows, cols := len(matrix), len(matrix[0])
	col0 := false

	for i := 0; i < rows; i++ {
		if matrix[i][0] == 0 {
			col0 = true // first col must be zero
		}
		for j := 1; j < cols; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0], matrix[0][j] = 0, 0 // sign zero row & col
			}
		}
	}

	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 1; j-- {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0 // if matrix[i][j] in zero row & col, be zero
			}
		}
		if col0 {
			matrix[i][0] = 0 // first col must be zero
		}
	}
}
